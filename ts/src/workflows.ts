import * as workflow from '@temporalio/workflow';
// Only import the activity types
import type * as mapActivity from './activity/map';
import {Mapped} from "./activity/map";
import type * as reduceActivity from './activity/reduce';
import {Reduced} from "./activity/reduce";


const { map } = workflow.proxyActivities<typeof mapActivity>({
  startToCloseTimeout: '1 minute',
});

const { reduce } = workflow.proxyActivities<typeof reduceActivity>({
  startToCloseTimeout: '1 minute',
});



/** A workflow that simply calls an activity */
export async function countWords(text: string): Promise<Reduced> {
  const mapped:Mapped = await map(text);
  return reduce(mapped);
}
