// little test runner
// brian taylor vann

// Actions
type StartTestRunActionParams = {
  startTime: number;
};
type StartTestRunAction = {
  action: "START_TEST_RUN";
  params: StartTestRunActionParams;
};

type StartTestCollectionActionParams = {
  collectionID: number;
  startTime: number;
};
type StartTestCollectionAction = {
  action: "START_TEST_COLLECTION";
  params: StartTestCollectionActionParams;
};

type StartTestActionParams = {
  collectionID: number;
  testID: number;
  startTime: number;
};
type StartTestAction = {
  action: "START_TEST";
  params: StartTestActionParams;
};

type EndTestActionParams = {
  assertions?: string[];
  collectionID: number;
  testID: number;
  endTime: number;
  startTime: number;
};
type EndTestAction = {
  action: "END_TEST";
  params: EndTestActionParams;
};

type EndTestCollectionActionParams = {
  collectionID: number;
  endTime: number;
};
type EndTestCollectionAction = {
  action: "END_TEST_COLLECTION";
  params: EndTestCollectionActionParams;
};

type EndTestRunParams = {
  endTime: number;
};
type CancelTestsAction = {
  action: "CANCEL_RUN";
  params: EndTestRunParams;
};
type EndTestRunAction = {
  action: "END_TEST_RUN";
  params: EndTestRunParams;
};
type ResultsStoreAction =
  | StartTestRunAction
  | StartTestCollectionAction
  | StartTestAction
  | CancelTestsAction
  | EndTestAction
  | EndTestCollectionAction
  | EndTestRunAction;

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
