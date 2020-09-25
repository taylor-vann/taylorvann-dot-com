import { TestRunResults } from "../state_types/state_types";
import { EndTestRunActionParams } from "../../action_types/actions_types";

type CancelRun = (
  results: TestRunResults,
  params: EndTestRunActionParams
) => TestRunResults;

const cancelRunState: CancelRun = (runResults, params) => {
  const { endTime } = params;
  runResults.endTime = endTime;
  runResults.status = "cancelled";

  const collectionResults = runResults.results;
  if (collectionResults) {
    for (const collection of collectionResults) {
      if (collection.status === "submitted") {
        collection.status = "cancelled";
      }

      const testResults = collection.results;
      if (testResults) {
        for (const result of testResults) {
          if (result.status === "submitted") {
            result.status = "cancelled";
          }
        }
      }
    }
  }

  return runResults;
};

export { cancelRunState };
