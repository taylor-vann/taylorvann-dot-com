// little test runner
// brian taylor vann

import { TestCollections } from "../test_store/test_store";
import { RecallSubscription, SubPub } from "../../subpub/subpub";

// Test Results & Test Collection Results
type TestStatus =
  | "untested"
  | "submitted"
  | "passed"
  | "cancelled"
  | "failed"
  | "completed";
type Assertions = string[];
type Result = {
  status: TestStatus;
  assertions?: Assertions;
  startTime?: number;
  endTime?: number;
};
type Results = { [testName: string]: Result };
type TestResults = {
  title: string;
  status: TestStatus;
  results: Results;
};
type CollectionResults = TestResults[];
type ResultsState = {
  status: TestStatus;
  results?: CollectionResults;
};
type GetState = () => ResultsState;
type RunTestsActionParams = {
  startTime: number;
};
type EndTestsActionParams = {
  endTime: number;
};

// Actions
type TestActionParams = {
  assertions?: string[];
  collectionID: string;
  testID: string;
  endTime: number;
  startTime: number;
};
type RunTestsAction = {
  action: "START_RUN";
  params: RunTestsActionParams;
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
type Subscribe<T> = (callback: RecallSubscription<T>) => () => void;

// create pubsub with state
const subpub = new SubPub<ResultsState>();

const subscribe: Subscribe<ResultsState> = (callback) => {
  const stub = subpub.subscribe(callback);
  return () => {
    subpub.unsubscribe(stub);
  };
};

const defaultTestResult: Result = {
  status: "untested",
};

const defaultResultsState: ResultsState = {
  status: "untested",
};

const resultsState: ResultsState = { ...defaultResultsState };

const createResultState = (testCollection: TestCollections) => {
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
};

const getResults: GetState = () => {
  return resultsState;
};

export { TestResults, createResultState, dispatch, getResults, subscribe };
