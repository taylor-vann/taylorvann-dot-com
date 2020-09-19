// little test runner
// brian taylor vann

import { Assertions } from "../results_store/results_store";
import { Test, getTestCollections } from "../test_store/test_store";
import {
  getTimestamp,
  updateTimestamp,
  startTestCollectionRun,
  endTestCollectionRun,
  sendTestResult,
  cancelRun,
} from "./runner_publisher";

type CreateTestTimeout = (requestedInterval?: number) => Promise<Assertions>;
type BuildLtrTestParams = {
  collectionID: string;
  issuedAt: number;
  testFunc: Test;
  testID: string;
  timeoutInterval?: number;
};
type LtrTest = () => Promise<void>;
type BuildLtrTest = (params: BuildLtrTestParams) => LtrTest;
interface RunTestParams {
  timeoutInterval?: number;
}
type RunTests = (params: RunTestParams) => void;

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

const runTestsInOrder: RunTests = async ({ timeoutInterval }) => {
  const testCollections = getTestCollections();
  const startTime = updateTimestamp();

  startTestCollectionRun(startTime);
  for (const collectionID in testCollections) {
    const { tests } = testCollections[collectionID];
    for (const testID in tests) {
      if (startTime < getTimestamp()) {
        return;
      }
      const testFunc = tests[testID];
      const builtTest = buildTest({
        collectionID,
        issuedAt: startTime,
        testFunc,
        testID,
        timeoutInterval,
      });

      await builtTest();
    }
  }

  endTestCollectionRun(startTime);
};

const runTests: RunTests = (params) => {
  runTestsInOrder(params);
};

export { runTests, cancelRun };
