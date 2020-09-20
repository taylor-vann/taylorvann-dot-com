// brian taylor vann
import { copycopy } from "../../copycopy/copycopy";

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
type TestRunResults = {
  status: TestStatus;
  results?: CollectionResults;
};

type GetResults = () => TestRunResults;

const defaultResultsState: TestRunResults = {
  status: "untested",
};

const getResults: GetResults = () => {
  return copycopy<TestRunResults>(defaultResultsState);
};

export { Assertions, TestRunResults, getResults };
