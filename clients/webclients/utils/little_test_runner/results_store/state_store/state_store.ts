// brian taylor vann

import { copycopy } from "../../copycopy/copycopy";
import { TestRunResults } from "./state_types/state_types";
import {
  StartTestRunActionParams,
  StartTestCollectionActionParams,
  StartTestActionParams,
  EndTestActionParams,
  EndTestCollectionActionParams,
  EndTestRunActionParams,
} from "../action_types/actions_types";
import { buildResultsState } from "./build_state/build_state";
import { startTestCollectionState } from "./start_test_collection_state/start_test_collection_state";
import { startTestState } from "./start_test_state/start_test_state";
import { cancelRunState } from "./cancel_run_state/cancel_run_state";
import { endTestCollectionState } from "./end_test_collection_state/end_test_collection_state";
import { endTestState } from "./end_test_state/end_test_state";
import { endTestRunState } from "./end_test_run_state/end_test_run_state";

type BuildResults = (params: StartTestRunActionParams) => void;
type StartTestCollection = (params: StartTestCollectionActionParams) => void;
type StartTest = (params: StartTestActionParams) => void;
type CancelRun = (params: EndTestRunActionParams) => void;
type EndTest = (params: EndTestActionParams) => void;
type EndTestCollection = (params: EndTestCollectionActionParams) => void;
type EndTestRun = (params: EndTestRunActionParams) => void;

type GetResults = () => TestRunResults;

const defaultResultsState: TestRunResults = {
  status: "unsubmitted",
};

let resultsState: TestRunResults = { ...defaultResultsState };

const buildResults: BuildResults = (params) => {
  resultsState = buildResultsState(params);
};

const startTestCollection: StartTestCollection = (params) => {
  const copyOfResults = copycopy<TestRunResults>(resultsState);
  resultsState = startTestCollectionState(copyOfResults, params);
};

const startTest: StartTest = (params) => {
  const copyOfResults = copycopy<TestRunResults>(resultsState);
  resultsState = startTestState(copyOfResults, params);
};

const cancelRun: CancelRun = (params) => {
  const copyOfResults = copycopy<TestRunResults>(resultsState);
  resultsState = cancelRunState(copyOfResults, params);
};

const endTest: EndTest = (params) => {
  const copyOfResults = copycopy<TestRunResults>(resultsState);
  resultsState = endTestState(copyOfResults, params);
};

const endTestCollection: EndTestCollection = (params) => {
  const copyOfResults = copycopy<TestRunResults>(resultsState);
  resultsState = endTestCollectionState(copyOfResults, params);
};

const endTestRun: EndTestRun = (params) => {
  const copyOfResults = copycopy<TestRunResults>(resultsState);
  resultsState = endTestRunState(copyOfResults, params);
};

const getResults: GetResults = () => {
  return copycopy<TestRunResults>(resultsState);
};

export {
  buildResults,
  startTestCollection,
  startTest,
  cancelRun,
  endTest,
  endTestCollection,
  endTestRun,
  getResults,
};
