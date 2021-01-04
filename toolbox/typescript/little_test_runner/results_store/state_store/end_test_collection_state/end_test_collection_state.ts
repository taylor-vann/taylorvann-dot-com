import { Results, TestRunResults } from "../state_types/state_types";
import { EndTestCollectionActionParams } from "../../action_types/actions_types";

type AllTestsHavePassed = (testResults: Results) => boolean;
type EndTestCollection = (
  results: TestRunResults,
  params: EndTestCollectionActionParams
) => TestRunResults;

const allTestsHavePassed: AllTestsHavePassed = (testResults: Results) => {
  for (const result of testResults) {
    if (result.status !== "passed") {
      return false;
    }
  }
  return true;
};

const endTestCollectionState: EndTestCollection = (runResults, params) => {
  if (runResults.results === undefined) {
    return runResults;
  }
  const { endTime, collectionID } = params;

  const collection = runResults.results[collectionID];
  if (collection === undefined) {
    return runResults;
  }

  collection.endTime = endTime;
  collection.status = "failed";

  const collectionResults = collection.results;
  if (collectionResults && allTestsHavePassed(collectionResults)) {
    collection.status = "passed";
  }

  return runResults;
};

export { endTestCollectionState };
