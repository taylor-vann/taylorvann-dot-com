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

type GetResults = () => RunResultsState;
const defaultResultsState: RunResultsState = {
  status: "untested",
};


// need to define what happens on:
// start test run

// start test collection

// end / add test result

// end test result

// create new state based on test collection

const getResults: GetResults = () => {
  // need to return copy of
  return defaultResultsState;
};

export {
  Assertions,
  Result,
  Results,
  TestResults,
  CollectionResults,
  getResults,
};
