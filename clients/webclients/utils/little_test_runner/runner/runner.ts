// little test runner
// brian taylor vann

import {
  Assertions,
  Test,
  TestParams,
  TestCollection,
} from "../runner/test_types/test_types";
import {
  startTestRun,
  startTestCollection,
  cancelRun,
  endTestCollection,
  endTestRun,
} from "./relay_results/relay_results";
import {
  getTimestamp,
  updateTimestamp,
} from "./timestamp_sieve/timestamp_sieve";
import { runTestsInOrder, runTestsAllAtOnce } from "./run_tests/run_tests";

interface StartLtrTestCollectionRunParams {
  testCollection: TestCollection;
  startTime: number;
}
type StartLtrTestCollectionRun = (
  params: StartLtrTestCollectionRunParams
) => Promise<void>;

interface StartLtrTestRunParams {
  testCollection: TestCollection;
}
type StartLtrTestRun = (params: StartLtrTestRunParams) => Promise<void>;

// create a test collection

const startLtrTestCollectionRun: StartLtrTestCollectionRun = async ({
  testCollection,
  startTime,
}) => {
  startTestRun(startTime);

  let collectionID = 0;
  for (const collection of testCollection) {
    if (startTime < getTimestamp()) {
      return;
    }

    const { tests, runTestsAsynchronously, timeoutInterval } = collection;
    const runParams = {
      collectionID,
      tests,
      startTime,
      timeoutInterval,
    };
    startTestCollection({
      collectionID,
      startTime,
    });

    if (runTestsAsynchronously) {
      await runTestsInOrder(runParams);
    } else {
      await runTestsAllAtOnce(runParams);
    }

    if (startTime < getTimestamp()) {
      return;
    }

    const endTime = performance.now();
    endTestCollection({
      collectionID,
      endTime,
    });

    collectionID += 1;
  }

  endTestRun(startTime);
};

// iterate through tests synchronously
const runTests: StartLtrTestRun = async (params) => {
  // create results store
  const startTime = updateTimestamp();

  await startLtrTestCollectionRun({ ...params, ...{ startTime } });
  if (startTime < getTimestamp()) {
    // get state
    return;
  }
};

export { Assertions, Test, TestParams, TestCollection, runTests, cancelRun };
