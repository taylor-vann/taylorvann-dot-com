// little test runner
// brian taylor vann

type Assertions = string[] | undefined;
type SyncTest = () => Assertions;
type AsyncTest = () => Promise<Assertions>;
type Test = SyncTest | AsyncTest;
type Tests = Test[];
type TestParams = {
  title: string;
  tests: Tests;
  runTestsAsynchronously?: boolean;
  timeoutInterval?: number;
};
type TestCollections = TestParams[];
type SetTestCollections = (testCollection: TestCollections) => TestCollections;
type GetTestCollections = () => TestCollections;

let testCollection: TestCollections = [];

const setTestCollections: SetTestCollections = (newTestCollection) => {
  testCollection = newTestCollection;
  return testCollection;
};

const getTestCollections: GetTestCollections = () => {
  return testCollection;
};

export { TestCollections, Test, Tests, TestParams, setTestCollections, getTestCollections };
