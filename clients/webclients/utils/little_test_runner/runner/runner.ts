// little test runner
// brian taylor vann

import { Assertions } from "../results_store/results_store";
import {
  Test,
  Tests,
  TestCollections,
  getTestCollections,
} from "../test_store/test_store";
import {
  getTimestamp,
  updateTimestamp,
  startTestCollectionsRun,
  endTestCollectionsRun,
  sendTestResult,
  cancelRun,
} from "./relay_results/relay_results";

type CreateTestTimeout = (requestedInterval?: number) => Promise<Assertions>;
type LtrTest = () => Promise<void>;
type BuildLtrTestParams = {
  collectionID: number;
  issuedAt: number;
  testFunc: Test;
  testID: number;
  timeoutInterval?: number;
};
type BuildLtrTest = (params: BuildLtrTestParams) => LtrTest;
interface RunTestsParams {
  timeoutInterval?: number;
  tests: Tests;
  collectionID: number;
  startTime: number;
}
type RunTests = (params: RunTestsParams) => Promise<void>;
interface RunTestCollectionsParams {
  testCollections: TestCollections;
}
type RunTestCollections = (params: RunTestCollectionsParams) => Promise<void>;

// create a test collection

const sleep = async (time: number) => {
  return new Promise((resolve) => {
    setTimeout(() => {
      resolve();
    }, time);
  });
};

const defaultTimeoutInterval = 10000;
const getTimeoutAssertions = (timeoutInterval: number) => [
  `timed out at: ${timeoutInterval}`,
];

const createTestTimeout: CreateTestTimeout = async (
  timeoutInterval?: number
) => {
  const interval = timeoutInterval ?? defaultTimeoutInterval;
  await sleep(interval);
  return getTimeoutAssertions(interval);
};

const buildTest: BuildLtrTest = (params) => {
  const { issuedAt, testID, collectionID, timeoutInterval } = params;
  return async () => {
    const startTime = performance.now();
    const assertions = await Promise.race([
      params.testFunc(),
      createTestTimeout(timeoutInterval),
    ]);
    if (issuedAt < getTimestamp()) {
      return;
    }
    const endTime = performance.now();
    sendTestResult({
      startTime,
      endTime,
      assertions,
      collectionID,
      testID,
    });
  };
};

const asynchonouslyRunTests: RunTests =  async ({startTime, collectionID, tests, timeoutInterval}) => {
  const builtAsyncTests = [];
  let testID = 0;
  for (const testFunc of tests) {
    const testFunc = tests[testID];
    builtAsyncTests.push(buildTest({
      collectionID,
      issuedAt: startTime,
      testFunc,
      testID,
      timeoutInterval,
    }));
    testID += 1;
  }

  if (startTime < getTimestamp()) {
    return;
  }
  await Promise.all(builtAsyncTests)
}

const iterateThroughTests: RunTests = async ({startTime, collectionID, tests, timeoutInterval}) => {
  let testID = 0;
  for (const testFunc of tests) {
    if (startTime < getTimestamp()) {
      return;
    }
    const builtTest = buildTest({
      collectionID,
      issuedAt: startTime,
      testFunc,
      testID,
      timeoutInterval,
    });

    await builtTest();
    testID += 1
  }
}

const runTestCollections: RunTestCollections = async ({testCollections }) => {
  const startTime = updateTimestamp();
  startTestCollectionsRun(startTime);

  let collectionID = 0;
  for (const collection of testCollections) {
    const { tests, runTestsAsynchronously, timeoutInterval } = collection;
    const runParams: RunTestsParams = {collectionID, tests, startTime, timeoutInterval};
    if (runTestsAsynchronously) {
      await iterateThroughTests(runParams);
    } else {
      await asynchonouslyRunTests(runParams)
    }

    collectionID += 1;
  }

  endTestCollectionsRun(startTime);
};

// iterate through tests synchronously
const runTests: RunTestCollections = async (params) => {
  runTestCollections(params);
};

export { runTests, cancelRun };
