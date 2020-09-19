// little test runner
// brian taylor vann

import { TestCollections } from "../test_store/test_store";
import { SubPub } from "../../subpub/subpub";

// Actions
type TestActionParams = {
  assertions?: string[];
  collectionID: number;
  testID: number;
  endTime: number;
  startTime: number;
};
type EndTestAction = {
  action: "END_TEST";
  params: TestActionParams;
};
type RunTestsActionParams = {
  startTime: number;
};
type RunTestsAction = {
  action: "START_RUN";
  params: RunTestsActionParams;
};
type StartTestActionParams = RunTestsActionParams;
type StartTestAction = {
  action: "START_TEST";
  params: StartTestActionParams;
};

type EndTestCollectionsRunParams = {
  endTime: number;
};
type CancelTestsAction = {
  action: "CANCEL_RUN";
  params: EndTestCollectionsRunParams;
};
type EndRunAction = {
  action: "END_RUN";
  params: EndTestCollectionsRunParams;
};
type ResultsStoreAction =
  | StartTestAction
  | RunTestsAction
  | CancelTestsAction
  | EndTestAction
  | EndRunAction
  | EndTestCollectionsRunParams;

const START_RUN = "START_RUN";
const END_RUN = "END_RUN";
const START_TEST = "START_TEST";
const END_TEST = "END_TEST";
const CANCEL_RUN = "CANCEL_RUN";

const resultsState: RunResultsState = { ...defaultResultsState };

const consolidate = (action: ResultsStoreAction) => {
  switch (action.action) {
    case START_RUN:
      break;
    case START_TEST:
      break;
    case END_TEST:
      break;
    case END_RUN:
      break;
    case CANCEL_RUN:
      break;
    default:
      break;
  }

  // dispatch to pubsub 
};

const dispatch = (action: ResultsStoreAction) => {
  consolidate(action);
  //
};

const getResults: GetStore = () => {
  // from results_state
  return resultsState;
};

export {
  getResults,
  dispatch,
};
