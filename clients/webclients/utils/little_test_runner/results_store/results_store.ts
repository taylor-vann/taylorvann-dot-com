// little test runner
// brian taylor vann

import { TestCollections } from "../test_store/test_store";
import { SubPub } from "../../subpub/subpub";

// Test Results & Test Collection Results
type TestStatus =
  | "untested"
  | "submitted"
  | "passed"
  | "cancelled"
  | "failed"
  | "completed";
type Assertions = string[] | undefined;
type Result = {
  status: TestStatus;
  assertions?: Assertions;
  startTime?: number;
  endTime?: number;
  testName: string;
};
type Results = Result[];
type TestResults = {
  title: string;
  status: TestStatus;
  results: Results;
};
type CollectionResults = TestResults[];
type RunResultsState = {
  status: TestStatus;
  results?: CollectionResults;
};
type GetStore = () => RunResultsState;
type Subscription<T> = (params: T) => void;
type Subscribe<T> = (callback: Subscription<T>) => () => void;

// Actions
type TestActionParams = {
  assertions?: string[];
  collectionID: string;
  testID: string;
  endTime: number;
  startTime: number;
};
type RunTestsActionParams = {
  startTime: number;
};
type RunTestsAction = {
  action: "START_RUN";
  params: RunTestsActionParams;
};
type EndTestsActionParams = {
  endTime: number;
};
type CancelTestsAction = {
  action: "CANCEL_RUN";
  params: EndTestsActionParams;
};
type CompletedTestAction = {
  action: "END_TEST";
  params: TestActionParams;
};
type CompletedsAction = {
  action: "END_RUN";
  params: EndTestsActionParams;
};
type ResultsStoreAction =
  | RunTestsAction
  | CancelTestsAction
  | CompletedTestAction
  | CompletedsAction;

// create pubsub with state
const subpub = new SubPub<RunResultsState>();

const subscribe: Subscribe<RunResultsState> = (callback) => {
  const stub = subpub.subscribe(callback);
  return () => {
    subpub.unsubscribe(stub);
  };
};

const defaultResultsState: RunResultsState = {
  status: "untested",
};

const resultsState: RunResultsState = { ...defaultResultsState };

const createResultsState = (testCollection: TestCollections) => {
  // iterate through testCollection
  // stub id
  //  -> iterate through tests
  //  -> get test name
};

const dispatch = (action: ResultsStoreAction) => {
  console.log(action);
  // send to store reducer
  // will trigger change

  // start
  // stop
  // finished test
  // finished test collection

  // get copy of test state
  // broadcast test state
};

const getResults: GetStore = () => {
  return resultsState;
};

export {
  Assertions,
  TestResults,
  createResultsState,
  getResults,
  dispatch,
  subscribe,
};
