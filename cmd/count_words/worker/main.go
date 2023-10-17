package main

import (
	"github.com/juliocnsouzadev/temporal-io-labs/internal/count_words/tracing"
	prom "github.com/prometheus/client_golang/prometheus"
	"github.com/uber-go/tally/v4"
	"github.com/uber-go/tally/v4/prometheus"
	"go.temporal.io/sdk/interceptor"
	"log"
	"time"

	"github.com/juliocnsouzadev/temporal-io-labs/internal/count_words/activity"
	"github.com/juliocnsouzadev/temporal-io-labs/internal/count_words/workflow"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/contrib/opentracing"
	sdktally "go.temporal.io/sdk/contrib/tally"
	"go.temporal.io/sdk/worker"
	temporalWorkflow "go.temporal.io/sdk/workflow"
)

func main() {

	// Set tracer which will be returned by opentracing.GlobalTracer().
	closer := tracing.SetJaegerGlobalTracer("word-count")
	defer func() { _ = closer.Close() }()

	// Create interceptor
	tracingInterceptor, err := opentracing.NewInterceptor(opentracing.TracerOptions{})
	if err != nil {
		log.Fatalf("Failed creating interceptor: %v", err)
	}

	options := client.Options{
		MetricsHandler: sdktally.NewMetricsHandler(newPrometheusScope(prometheus.Configuration{
			ListenAddress: "0.0.0.0:9090",
			TimerType:     "histogram",
		})),
		ContextPropagators: []temporalWorkflow.ContextPropagator{tracing.NewContextPropagator()},
		Interceptors:       []interceptor.ClientInterceptor{tracingInterceptor},
	}
	c, err := client.Dial(options)
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	w := worker.New(c, workflow.CountWordsTaskQueue, worker.Options{
		EnableLoggingInReplay: true,
	})

	w.RegisterWorkflow(workflow.CountWords)
	w.RegisterActivity(activity.Map)
	w.RegisterActivity(activity.Reduce)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}

func newPrometheusScope(c prometheus.Configuration) tally.Scope {
	reporter, err := c.NewReporter(
		prometheus.ConfigurationOptions{
			Registry: prom.NewRegistry(),
			OnError: func(err error) {
				log.Println("error in prometheus reporter", err)
			},
		},
	)
	if err != nil {
		log.Fatalln("error creating prometheus reporter", err)
	}
	scopeOpts := tally.ScopeOptions{
		CachedReporter:  reporter,
		Separator:       prometheus.DefaultSeparator,
		SanitizeOptions: &sdktally.PrometheusSanitizeOptions,
		Prefix:          "temporal_samples",
	}
	scope, _ := tally.NewRootScope(scopeOpts, time.Second)
	scope = sdktally.NewPrometheusNamingScope(scope)

	log.Println("prometheus metrics scope created")
	return scope
}
