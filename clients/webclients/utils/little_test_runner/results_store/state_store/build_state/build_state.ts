import {
  Results,
  TestResults,
  TestRunResults,
} from "../state_types/state_types";
import { TestCollection } from "../../test_types/test_types";

interface BuildResultsStateParams {
  startTime: number;
  testCollection: TestCollection;
}
type BuildResultsState = (params: BuildResultsStateParams) => TestRunResults;

const buildResultsState: BuildResultsState = ({
  startTime,
  testCollection,
}) => {
  const nextState: TestRunResults = {
    status: "untested",
    results: [],
    startTime,
  };

  for (const collection of testCollection) {
    const { tests, title } = collection;
    const collectionResults: TestResults = {
      title,
      status: "untested",
    };

    const results: Results = [];
    for (const test of tests) {
      const { name } = test;
      results.push({
        status: "untested",
        name,
      });
    }

    nextState.results.push({ ...collectionResults, ...{ results } });
  }

  return nextState;
};

export { BuildResultsStateParams, BuildResultsState, buildResultsState };
