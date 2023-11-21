import { TestWorkflowEnvironment } from '@temporalio/testing';
import { after, before, it } from 'mocha';
import { Worker } from '@temporalio/worker';
import { countWords } from '../workflows';
import assert from 'assert';

describe('Count Words workflow with mocks', () => {
  let testEnv: TestWorkflowEnvironment;

  before(async () => {
    testEnv = await TestWorkflowEnvironment.createLocal();
  });

  after(async () => {
    await testEnv?.teardown();
  });

  it('successfully completes the Workflow with a mocked Activity', async () => {
    const { client, nativeConnection } = testEnv;
    const taskQueue = 'test';

    const text = 'Temporal or Kafka that is the question in temporal or kafka';
    let mapped = {words: text.split(' ')};
    const worker = await Worker.create({
      connection: nativeConnection,
      taskQueue,
      workflowsPath: require.resolve('../workflows'),
      activities: {
        map: async () => mapped,
        reduce: async () => ({}),
      },
    });

    const result = await worker.runUntil(
      client.workflow.execute(countWords, {
        args: ['Temporal'],
        workflowId: 'test',
        taskQueue,
      })
    );
    assert.ok(result);
  });
});
