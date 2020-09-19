// brian taylor vann

import { Assertions, dispatch } from "../results_store/results_store";

type GetTimestamp = () => number;
type UpdateTimestamp = GetTimestamp;
type TestResultsDispatchParams = {
  startTime: number;
  endTime: number;
  assertions: Assertions;
  collectionID: string;
  testID: string;
};
type DispatchTestResults = (params: TestResultsDispatchParams) => void;
type RaceCheckFunction = (time: number) => void;
type VoidFunction = () => void;

// timestamps
let currentTestTimestamp = performance.now();
const getTimestamp: GetTimestamp = () => {
  return currentTestTimestamp;
};
const updateTimestamp: UpdateTimestamp = () => {
  currentTestTimestamp = performance.now();
  return currentTestTimestamp;
};

// run tests
const startTestCollectionRun: RaceCheckFunction = (startTime) => {
  dispatch({
    action: "START_RUN",
    params: { startTime },
  });
};

const endTestCollectionRun: RaceCheckFunction = (startTime) => {
  if (startTime < getTimestamp()) {
    return;
  }
  const endTime = performance.now();
  dispatch({
    action: "END_RUN",
    params: { endTime },
  });
};

const sendTestResult: DispatchTestResults = (params) => {
  dispatch({
    action: "END_TEST",
    params,
  });
};

const cancelRun: VoidFunction = () => {
  const endTime = updateTimestamp();
  dispatch({ action: "CANCEL_RUN", params: { endTime } });
};

export {
  getTimestamp,
  updateTimestamp,
  startTestCollectionRun,
  endTestCollectionRun,
  sendTestResult,
  cancelRun,
};
