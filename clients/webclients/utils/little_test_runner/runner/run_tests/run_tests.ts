// little test runner
// brian taylor vann

import { Assertions, Test } from "../../results_store/results_store";
import { startTest, sendTestResult } from "../relay_results/relay_results";
import { getTimestamp } from "../timestamp_sieve/timestamp_sieve";

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
  tests: Test[];
  collectionID: number;
  startTime: number;
}
type RunTests = (params: RunTestsParams) => Promise<void>;

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
    if (issuedAt < getTimestamp()) {
      return;
    }

    const startTime = performance.now();
    startTest({
      collectionID,
      testID,
      startTime,
    });

    const assertions = await Promise.race([
      params.testFunc(),
      createTestTimeout(timeoutInterval),
    ]);

    if (issuedAt < getTimestamp()) {
      return;
    }
    const endTime = performance.now();
    sendTestResult({
      endTime,
      assertions,
      collectionID,
      testID,
    });
  };
};

const runTestsInOrder: RunTests = async ({
  startTime,
  collectionID,
  tests,
  timeoutInterval,
}) => {
  const builtAsyncTests = [];

  let testID = 0;
  for (const testFunc of tests) {
    builtAsyncTests.push(
      buildTest({
        collectionID,
        issuedAt: startTime,
        testFunc,
        testID,
        timeoutInterval,
      })
    );
    testID += 1;
  }

  if (startTime < getTimestamp()) {
    return;
  }
  await Promise.all(builtAsyncTests);
};

const runTestsAllAtOnce: RunTests = async ({
  startTime,
  collectionID,
  tests,
  timeoutInterval,
}) => {
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
    testID += 1;
  }
};

export { runTestsInOrder, runTestsAllAtOnce };
