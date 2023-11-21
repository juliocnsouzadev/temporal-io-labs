import { MockActivityEnvironment } from '@temporalio/testing';
import { describe, it } from 'mocha';
import {map, Mapped} from '../activity/map';
import assert from 'assert';

describe('map activity', async () => {
  it('successfully maps', async () => {
    const env = new MockActivityEnvironment();
    const text = 'Temporal or Kafka that is the question';
    const result:Mapped = await env.run(map, text);
    assert.equal(result.words.length, 7);
  });
});
