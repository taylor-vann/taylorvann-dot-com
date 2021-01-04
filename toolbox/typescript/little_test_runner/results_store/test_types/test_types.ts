// little test runner
// brian taylor vann

import { Assertions } from "../../results_store/results_store";

type SyncTest = () => Assertions;
type AsyncTest = () => Promise<Assertions>;
type Test = SyncTest | AsyncTest;
type TestParams = {
  title: string;
  tests: Test[];
  runTestsAsynchronously?: boolean;
  timeoutInterval?: number;
};
type TestCollection = TestParams[];

export { Test, TestParams, TestCollection };
