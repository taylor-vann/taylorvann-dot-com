// little test runner
// brian taylor vann

import { Test, getTestCollections } from "../test_store/test_store";
import { dispatch } from "../results_store/results_store";

type LtrTest = () => Promise<void>;
type LtrTestCollection = LtrTest[];
type LtrParams = {
  ordered?: boolean;
  timeoutInterval?: number;
};
type BuildLtrTestParams = {
  collectionID: string;
  issuedAt: number;
  testFunc: Test;
  testID: string;
};
type BuildLtrTest = (params: BuildLtrTestParams) => LtrTest;
type UpdateTimestamp = () => number;
type BuildLtrTestCollection = (issuedAt: number) => LtrTestCollection;
type RunTests = () => void;
type RunLtrTests = (params?: LtrParams) => void;

type TestResultsDispatchParams = {
  startTime: number;
  endTime: number;
  assertions: string[];
  collectionID: string;
  testID: string;
};
type DispatchTestResults = (params: TestResultsDispatchParams) => void;
type RaceCheckFunction = (time: number) => void;
type VoidFunction = () => void;

// timestamps
let currentTestTimestamp = performance.now();
const getTimestamp: UpdateTimestamp = () => {
  return currentTestTimestamp;
};
const updateTimestamp: UpdateTimestamp = () => {
  currentTestTimestamp = performance.now();
  return currentTestTimestamp;
};

// timeout intervals
let timeoutInterval = 10000;
const timeoutAssertions = ["timeout"];
const sleep = async (time: number) => {
  return new Promise((resolve) => {
    setTimeout(() => {
      resolve();
    }, time);
  });
};
const testTimeout: Test = async () => {
  await sleep(timeoutInterval);
  return timeoutAssertions;
};

// run tests
const dispatchStartTestCollectionRun: RaceCheckFunction = (startTime) => {
  dispatch({
    action: "START_RUN",
    params: { startTime },
  });
};

const dispatchEndTestCollectionRun: RaceCheckFunction = (startTime) => {
  if (startTime < getTimestamp()) {
    return;
  }
  const endTime = performance.now();
  dispatch({
    action: "END_RUN",
    params: { endTime },
  });
};

const dispatchTestResults: DispatchTestResults = (params) => {
  dispatch({
    action: "END_TEST",
    params,
  });
};

const dispatchCancelRun: RaceCheckFunction = (endTime) => {
  dispatch({ action: "CANCEL_RUN", params: { endTime } });
};

const buildTest: BuildLtrTest = (params) => {
  const { issuedAt, testID, collectionID } = params;
  return async () => {
    const startTime = performance.now();
    const assertions = await Promise.race([params.testFunc(), testTimeout()]);
    if (issuedAt < currentTestTimestamp) {
      return;
    }
    const endTime = performance.now();

    dispatchTestResults({
      startTime,
      endTime,
      assertions,
      collectionID,
      testID,
    });
  };
};

const runTestsInOrder: RunTests = async () => {
  const startTime = updateTimestamp();
  dispatchStartTestCollectionRun(startTime);

  const testCollections = getTestCollections();
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
      });

      await builtTest();
    }
  }

  dispatchEndTestCollectionRun(startTime);
};

const buildAsyncRun: BuildLtrTestCollection = (issuedAt) => {
  const builtTests: LtrTestCollection = [];
  const testCollection = getTestCollections();
  for (const collectionID in testCollection) {
    const { tests } = testCollection[collectionID];
    for (const testID in tests) {
      const testFunc = tests[testID];
      const builtTest = buildTest({
        collectionID,
        testID,
        issuedAt,
        testFunc,
      });
      builtTests.push(builtTest);
    }
  }

  return builtTests;
};

const runTestsAsync: RunTests = async () => {
  const startTime = updateTimestamp();
  dispatchStartTestCollectionRun(startTime);

  const tests = buildAsyncRun(startTime);
  await Promise.all(tests);

  dispatchEndTestCollectionRun(startTime);
};

const cancelTestRun: VoidFunction = () => {
  const endTime = updateTimestamp();
  dispatchCancelRun(endTime);
};

const runTests: RunLtrTests = (params) => {
  timeoutInterval = params?.timeoutInterval ?? timeoutInterval;
  if (params.ordered == true) {
    runTestsInOrder();
    return;
  }

  runTestsAsync();
};

export { runTests, cancelTestRun };
