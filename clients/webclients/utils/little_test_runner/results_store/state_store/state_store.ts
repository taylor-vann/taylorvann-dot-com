// brian taylor vann

import {
  BuildResultsStateParams,
  buildResultsState,
} from "./build_state/build_state";
import { copycopy } from "../../copycopy/copycopy";
import { TestRunResults } from "./state_types/state_types";

type BuildResults = (params: BuildResultsStateParams) => void;
type GetResults = () => TestRunResults;

const defaultResultsState: TestRunResults = {
  status: "untested",
  results: [],
};

let resultsState: TestRunResults = { ...defaultResultsState };

// const START_TEST_RUN = "START_TEST_RUN";
const buildResults: BuildResults = (params) => {
  resultsState = buildResultsState(params);
};

// const START_TEST_COLLECTION = "START_TEST_COLLECTION";
// const START_TEST = "START_TEST";
// const CANCEL_RUN = "CANCEL_RUN";
// const END_TEST = "END_TEST";
// const END_TEST_COLLECTION = "END_TEST_COLLECTION";
// const END_TEST_RUN = "END_TEST_RUN";

const getResults: GetResults = () => {
  return copycopy<TestRunResults>(resultsState);
};

export { getResults, buildResults };
