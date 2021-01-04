import {
  Results,
  TestResults,
  TestRunResults,
} from "../state_types/state_types";
import { StartTestRunActionParams } from "../../action_types/actions_types";

// type buildCollectionState = (tests:  )
type BuildResultsState = (params: StartTestRunActionParams) => TestRunResults;

const buildResultsState: BuildResultsState = ({
  testCollection,
  startTime,
  stub,
}) => {
  const nextState: TestRunResults = {
    status: "submitted",
    results: [],
    startTime,
  };

  for (const collection of testCollection) {
    const { tests, title } = collection;
    const collectionResults: TestResults = {
      title,
      status: "unsubmitted",
    };

    const results: Results = [];
    for (const test of tests) {
      const { name } = test;
      results.push({
        status: "unsubmitted",
        name,
      });
    }

    if (nextState.results) {
      nextState.results.push({ ...collectionResults, ...{ results } });
    }
  }

  return nextState;
};

export { buildResultsState };
