import { MockActivityEnvironment } from '@temporalio/testing';
import { describe, it } from 'mocha';
import {map, Mapped} from '../activity/map';
import {reduce, Reduced} from '../activity/reduce';
import assert from 'assert';

describe('reduce activity', async () => {
  it('successfully reduces', async () => {
    const env = new MockActivityEnvironment();
    const text = 'Temporal or Kafka that is the question in temporal or kafka';
    const mapped:Mapped = await env.run(map, text);
    const result:Reduced = await env.run(reduce, mapped);
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
