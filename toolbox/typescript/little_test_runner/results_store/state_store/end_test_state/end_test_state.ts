import { TestRunResults } from "../state_types/state_types";
import { EndTestActionParams } from "../../action_types/actions_types";

type EndTest = (
  results: TestRunResults,
  params: EndTestActionParams
) => TestRunResults;

const endTestState: EndTest = (runResults, params) => {
  if (runResults.results === undefined) {
    return runResults;
  }
  const { assertions, endTime, collectionID, testID } = params;

  const testResult = runResults?.results?.[collectionID]?.results?.[testID];
  if (testResult === undefined) {
    return runResults;
  }

  testResult.status = "failed";
  if (assertions === undefined) {
    testResult.status = "passed";
  }
  if (assertions && assertions.length === 0) {
    testResult.status = "passed";
  }

  testResult.assertions = assertions;
  testResult.endTime = endTime;

  return runResults;
};

export { endTestState };
