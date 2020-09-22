// little test runner
// brian taylor vann

import { ResultsStoreAction } from "../action_types/actions_types";
import {
  buildResults,
  startTestCollection,
  startTest,
  cancelRun,
  endTest,
  endTestCollection,
  endTestRun,
  getResults,
} from "../state_store/state_store";
import { broadcast } from "../publisher/publisher";

type Consolidate = (action: ResultsStoreAction) => void;

const START_TEST_RUN = "START_TEST_RUN";
const START_TEST_COLLECTION = "START_TEST_COLLECTION";
const START_TEST = "START_TEST";
const CANCEL_RUN = "CANCEL_RUN";
const END_TEST = "END_TEST";
const END_TEST_COLLECTION = "END_TEST_COLLECTION";
const END_TEST_RUN = "END_TEST_RUN";

const consolidate: Consolidate = (action) => {
  switch (action.action) {
    case START_TEST_RUN:
      buildResults(action.params);
      break;
    case START_TEST_COLLECTION:
      startTestCollection(action.params);
      break;
    case START_TEST:
      startTest(action.params);
      break;
    case CANCEL_RUN:
      cancelRun(action.params);
      break;
    case END_TEST:
      endTest(action.params);
      break;
    case END_TEST_COLLECTION:
      endTestCollection(action.params);
      break;
    case END_TEST_RUN:
      endTestRun(action.params);
      break;
    default:
      break;
  }

  broadcast(getResults());
};

const dispatch: Consolidate = (action) => {
  consolidate(action);
};

export { dispatch };
