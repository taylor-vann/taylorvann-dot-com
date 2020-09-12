// little test runner
// brian taylor vann

// atomic status
type TestStatus = "processing" | "completed"
type Assertions = string[] | undefined

// tests
type SyncTest = () => Assertions
type AsyncTest = () => Promise<Assertions>
type Test = SyncTest | AsyncTest
type Tests = {[testName: string]: Test}
type TestsGroup = {
  title: string,
  tests: Tests,
}
type Collection = {
  [id: string]: TestsGroup,
}

// test results
type Result = {
  status: TestStatus,
  assertions: Assertions,
}
type Results = {[testName: string]: Result}
type TestResults = {
  title: string,
  status: TestStatus,
  results: Results,
}
type CollectionResults = {
  [id: string]: TestResults,
}
type CollectionState = {
  status: TestStatus,
  results: CollectionResults,
}

// Test State
type State = CollectionState
type GetState = () => State

// declare test
type TestParams = TestsGroup
type CreateTest = (params: TestParams) => {}

// cancel and begin tests
type AncillaryCallback = () => {}
type Subscribe = AncillaryCallback
type Cancel = AncillaryCallback
type Start = AncillaryCallback
