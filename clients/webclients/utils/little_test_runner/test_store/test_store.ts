// little test runner
// brian taylor vann

// atomic status
type TestStatus = "processing" | "cancelled" | "completed"
type Assertions = string[] | undefined

// tests
type SyncTest = () => Assertions
type AsyncTest = () => Promise<Assertions>
type Test = SyncTest | AsyncTest
type Tests = {[testName: string]: Test}
type TestParams = {
  title: string,
  tests: Tests,
}
type Collection = {
  [id: string]: TestParams,
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
type State = {
  status: TestStatus,
  results: CollectionResults,
}

// declare test
type AncillaryCallback = () => {}

type GetState = () => State
type CreateTest = (params: TestParams) => {}
type Subscribe = AncillaryCallback
type Cancel = AncillaryCallback
type Start = AncillaryCallback
