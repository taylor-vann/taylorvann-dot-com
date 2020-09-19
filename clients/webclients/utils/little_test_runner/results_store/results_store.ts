type TestStatus =
  | "untested"
  | "submitted"
  | "passed"
  | "cancelled"
  | "failed"
  | "completed";
type Assertions = string[] | undefined;
type RunResultsState = {
  status: TestStatus;
  results?: CollectionResults;
};
type GetStore = () => RunResultsState;
type Subscription<T> = (params: T) => void;
type Subscribe<T> = (callback: Subscription<T>) => () => void;

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

const defaultResultsState: RunResultsState = {
  status: "untested",
};


// add test result

// end test result

// create new state based on test collection

const getResults = () => {};
const dispatch = () => {};

export {
  Assertions,
  Result,
  Results,
  TestResults,
  CollectionResults,
  getResults,
  dispatch,
}