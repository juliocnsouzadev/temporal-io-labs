import { NativeConnection, Worker } from '@temporalio/worker';
import {map} from "./activity/map";
import {reduce} from "./activity/reduce";

async function run() {
  // Step 1: Establish a connection with Temporal server.
  //
  // Worker code uses `@temporalio/worker.NativeConnection`.
  // (But in your application code it's `@temporalio/client.Connection`.)
  const connection = await NativeConnection.connect({
    address: 'localhost:7233',
    // TLS and gRPC metadata configuration goes here.
  });
  // Step 2: Register Workflows and Activities with the Worker.
  const worker = await Worker.create({
    connection,
    namespace: 'default',
    taskQueue: 'count-words-ts-task-queue',
    // Workflows are registered using a path as they run in a separate JS context.
    workflowsPath: require.resolve('./workflows'),
    activities: {
      map,
      reduce,
    },
  });
  await worker.run();
}

run().catch((err) => {
  console.error(err);
  process.exit(1);
});
