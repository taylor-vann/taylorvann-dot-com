// brian taylor vann

import { Assertions } from "../../results_store/results_store";
import { dispatch } from "../../conductor/conductor";

interface StartTestCollectionParams {
  collectionID: number;
  startTime: number;
}
type StartTestCollection = (params: StartTestCollectionParams) => void;

interface StartTestParams {
  collectionID: number;
  testID: number;
  startTime: number;
}
type StartTest = (params: StartTestParams) => void;

type SendTestResultsParams = {
  startTime: number;
  endTime: number;
  assertions: Assertions;
  collectionID: number;
  testID: number;
};
type SendTestResults = (params: SendTestResultsParams) => void;

interface EndTestCollectionParams {
  collectionID: number;
  endTime: number;
}
type EndTestCollection = (params: EndTestCollectionParams) => void;

type RaceCheck = (time: number) => void;

// run tests
const startTestRun: RaceCheck = (startTime) => {
  dispatch({
    action: "START_TEST_RUN",
    params: { startTime },
  });
};

const startTestCollection: StartTestCollection = (params) => {
  dispatch({
    action: "START_TEST_COLLECTION",
    params,
  });
};

const startTest: StartTest = (params) => {
  dispatch({
    action: "START_TEST",
    params,
  });
};

const cancelRun: RaceCheck = (endTime) => {
  dispatch({
    action: "CANCEL_RUN",
    params: { endTime },
  });
};

const sendTestResult: SendTestResults = (params) => {
  dispatch({
    action: "END_TEST",
    params,
  });
};

const endTestCollection: EndTestCollection = (params) => {
  dispatch({
    action: "END_TEST_COLLECTION",
    params,
  });
};

const endTestRun: RaceCheck = (endTime) => {
  dispatch({
    action: "END_TEST_RUN",
    params: { endTime },
  });
};

export {
  startTestRun,
  startTestCollection,
  startTest,
  cancelRun,
  sendTestResult,
  endTestCollection,
  endTestRun,
};
