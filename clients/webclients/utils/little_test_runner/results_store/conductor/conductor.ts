// little test runner
// brian taylor vann

import { ResultsStoreAction } from "../state_store/actions_types";

type Consolidate = (action: ResultsStoreAction) => void;

const START_TEST_RUN = "START_TEST_RUN";
const START_TEST_COLLECTION = "START_TEST_COLLECTION";
const START_TEST = "START_TEST";
const CANCEL_RUN = "CANCEL_RUN";
const END_TEST = "END_TEST";
const END_TEST_COLLECTION = "END_TEST_COLLECTION";
const END_TEST_RUN = "END_TEST_RUN";

const consolidate: Consolidate = (action) => {
  // send to results store
  switch (action.action) {
    case START_TEST_RUN:
      break;
    case START_TEST_COLLECTION:
      break;
    case START_TEST:
      break;
    case CANCEL_RUN:
      break;
    case END_TEST:
      break;
    case END_TEST_COLLECTION:
      break;
    case END_TEST_RUN:
      break;
    default:
      break;
  }

  // dispatch to pubsub
};

const dispatch: Consolidate = (action) => {
  consolidate(action);
  //
};

export { dispatch };
