import { TestWorkflowEnvironment } from '@temporalio/testing';
import { before, describe, it } from 'mocha';
import { Worker } from '@temporalio/worker';
import { countWords } from '../workflows';
import assert from 'assert';
import {map} from "../activity/map";
import {reduce, Reduced} from "../activity/reduce";

describe('Count Words workflow', () => {
  let testEnv: TestWorkflowEnvironment;

  before(async () => {
    testEnv = await TestWorkflowEnvironment.createLocal();
  });

  after(async () => {
    await testEnv?.teardown();
  });

  it('successfully completes the Workflow', async () => {
    const { client, nativeConnection } = testEnv;
    const taskQueue = 'test';

    const worker = await Worker.create({
      connection: nativeConnection,
      taskQueue,
      workflowsPath: require.resolve('../workflows'),
      activities: {
        map,
        reduce
      },
    });

    const text = 'Temporal or Kafka that is the question in temporal or kafka';
    const result:Reduced = await worker.runUntil(
      client.workflow.execute(countWords, {
        args: [text],
        workflowId: 'test',
        taskQueue,
      })
    );
    assert.equal(result.wordCounts['temporal'], 2);
    assert.equal(result.wordCounts['or'], 2);
    assert.equal(result.wordCounts['kafka'], 2);
    assert.equal(result.wordCounts['that'], 1);
    assert.equal(result.wordCounts['is'], 1);
    assert.equal(result.wordCounts['the'], 1);
    assert.equal(result.wordCounts['question'], 1);
    assert.equal(result.wordCounts['in'], 1);
  });
});
