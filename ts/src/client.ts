import { Connection, Client } from '@temporalio/client';
import { countWords } from './workflows';
import { nanoid } from 'nanoid';

const args = [
  "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
  "Aenean feugiat felis sed turpis scelerisque, at imperdiet ante viverra.",
  "Aenean nec dui nec tellus dapibus ultricies sit amet a nulla.",
  "Integer eget dolor quis dolor luctus vestibulum.",
  "Nullam et turpis ac diam pellentesque feugiat.",
  "Maecenas scelerisque lorem at diam dictum, sit amet bibendum quam sollicitudin.",
  "Sed iaculis felis vitae dui elementum rhoncus ac vitae nisi.",
  "Etiam suscipit nulla sit amet semper efficitur.",
  "Cras pulvinar dui sit amet lacus pharetra congue.",
  "Duis tristique ante a lectus venenatis, ac congue nibh euismod.",
  "Aenean accumsan nibh eu dolor gravida condimentum.",
  "Maecenas laoreet turpis in erat fermentum, nec rutrum erat facilisis.",
  "Morbi malesuada turpis sit amet fermentum volutpat.",
  "Aliquam in ligula porttitor, molestie mi sit amet, tincidunt urna.",
  "Fusce at leo sed arcu fringilla eleifend id nec libero.",
  "Proin non lectus fringilla, varius ipsum eget, vulputate dui.",
];

async function run() {
  // Connect to the default Server location
  const connection = await Connection.connect({ address: 'localhost:7233' });
  // In production, pass options to configure TLS and other settings:
  // {
  //   address: 'foo.bar.tmprl.cloud',
  //   tls: {}
  // }

  const client = new Client({
    connection,
    // namespace: 'foo.bar', // connects to 'default' namespace if not specified
  });

  for (const arg of args) {
    const handle = await client.workflow.start(countWords, {
      taskQueue: 'count-words-ts-task-queue',
      args: [arg],
      workflowId: 'workflow-' + nanoid(),
    });
    console.log(`Started workflow ${handle.workflowId}`);

    console.log(await handle.result());
  }
}

run().catch((err) => {
  console.error(err);
  process.exit(1);
});
