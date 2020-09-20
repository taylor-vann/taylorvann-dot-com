// brian taylor vann

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

export {
  StartTestRunActionParams,
  StartTestCollectionActionParams,
  StartTestActionParams,
  EndTestActionParams,
  EndTestCollectionActionParams,
  EndTestRunParams,
  ResultsStoreAction,
};
