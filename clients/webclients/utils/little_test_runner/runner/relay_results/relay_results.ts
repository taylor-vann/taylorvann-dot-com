// brian taylor vann

import {
  StartTestRunActionParams,
  StartTestCollectionActionParams,
  StartTestActionParams,
  EndTestActionParams,
  EndTestCollectionActionParams,
  EndTestRunActionParams,
} from "../../results_store/results_store";

import { dispatch } from "../../results_store/results_store";

type StartTestRun = (params: StartTestRunActionParams) => void;
type StartTestCollection = (params: StartTestCollectionActionParams) => void;
type StartTest = (params: StartTestActionParams) => void;
type CancelRun = (params: EndTestRunActionParams) => void;
type SendTestResults = (params: EndTestActionParams) => void;
type EndTestCollection = (params: EndTestCollectionActionParams) => void;
type EndTestRun = CancelRun;

// run tests
const startTestRun: StartTestRun = (params) => {
  dispatch({
    action: "START_TEST_RUN",
    params,
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

const cancelRun: CancelRun = (params) => {
  dispatch({
    action: "CANCEL_RUN",
    params,
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

const endTestRun: EndTestRun = (params) => {
  dispatch({
    action: "END_TEST_RUN",
    params,
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
