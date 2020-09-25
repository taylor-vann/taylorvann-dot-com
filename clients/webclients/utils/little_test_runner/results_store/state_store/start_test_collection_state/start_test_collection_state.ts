import { TestRunResults } from "../state_types/state_types";
import { StartTestCollectionActionParams } from "../../action_types/actions_types";

type StartTestCollection = (
  results: TestRunResults,
  params: StartTestCollectionActionParams
) => TestRunResults;

const startTestCollectionState: StartTestCollection = (runResults, params) => {
  if (runResults.results === undefined) {
    return runResults;
  }

  const { startTime, collectionID } = params;

  const collectionResult = runResults?.results?.[collectionID];
  if (collectionResult) {
    collectionResult.status = "submitted";
    collectionResult.startTime = startTime;
  }

  return runResults;
};

export { startTestCollectionState };
