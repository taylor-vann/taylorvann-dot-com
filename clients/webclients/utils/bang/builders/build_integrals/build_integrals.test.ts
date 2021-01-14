// brian taylor vann
// build integrals

import { samestuff } from "../../../little_test_runner/samestuff/samestuff";
import { BuildIntegralsParams } from "./build_integrals";
import { buildSkeleton } from "../build_skeleton/build_skeleton";
import { buildIntegrals } from "../build_integrals/build_integrals";

type TextTextInterpolator = <A>(
  templateArray: TemplateStringsArray,
  ...injections: A[]
) => BuildIntegralsParams<A>;

const testTextInterpolator: TextTextInterpolator = (
  templateArray,
  ...injections
) => {
  const template = { templateArray, injections };
  return {
    template: template,
    skeleton: buildSkeleton(template),
  };
};

const title = "build_integrals";
const runTestsAsynchronously = true;

const testFindOpenParagraph = () => {
  const assertions = [];
  const params = testTextInterpolator`<p>`;
  const results = buildIntegrals(params);

  console.log(results);

  assertions.push("fail immediately");
  if (results.length !== 1) {
    assertions.push("there should be at least one instruction set");
  }

  return assertions;
};

const testFindOpenParagraphWithAttributes = () => {
  const assertions = [];
  const params = testTextInterpolator`<p message="hello, world!">`;
  const results = buildIntegrals(params);

  console.log(results);

  assertions.push("fail immediately");
  if (results.length !== 1) {
    assertions.push("there should be at least one instruction set");
  }

  return assertions;
};

const testFindOpenParagraphWithTrailingImplicitAttribute = () => {
  const assertions = [];
  const params = testTextInterpolator`<p message="hello, world!" checked>`;
  const results = buildIntegrals(params);

  console.log(results);

  assertions.push("fail immediately");
  if (results.length !== 1) {
    assertions.push("there should be at least one instruction set");
  }

  return assertions;
};

const testFindOpenParagraphWithInjectedAttribute = () => {
  const assertions = [];
  const params = testTextInterpolator`<p message="${"hello, world!"}">`;
  const results = buildIntegrals(params);

  console.log(results);

  assertions.push("fail immediately");
  if (results.length !== 1) {
    assertions.push("there should be at least one instruction set");
  }

  return assertions;
};

const testFindOpenParagraphWithInjectedAndTrailingImplicitAttributes = () => {
  const assertions = [];
  const params = testTextInterpolator`<p message="${"hello, world!"}" checked>`;
  const results = buildIntegrals(params);

  console.log(results);

  assertions.push("fail immediately");
  if (results.length !== 1) {
    assertions.push("there should be at least one instruction set");
  }

  return assertions;
};

const testFindCloseParagraph = () => {
  const assertions = [];
  const params = testTextInterpolator`</p>`;
  const results = buildIntegrals(params);

  assertions.push("fail immediately");
  if (results.length !== 1) {
    assertions.push("there should be at least one instruction set");
  }

  return assertions;
};

const testFindCloseH1 = () => {
  const assertions = [];
  const params = testTextInterpolator`</h1>`;
  const results = buildIntegrals(params);

  console.log(results);
  assertions.push("fail immediately");
  if (results.length !== 1) {
    assertions.push("there should be at least one instruction set");
  }

  return assertions;
};

const testFindCloseParagraphWithTrailingSpaces = () => {
  const assertions = [];
  const params = testTextInterpolator`</h1        >`;
  const results = buildIntegrals(params);

  assertions.push("fail immediately");
  if (results.length !== 1) {
    assertions.push("there should be at least one instruction set");
  }

  return assertions;
};

const testFindContent = () => {
  const assertions = [];
  const params = testTextInterpolator`hello world!`;
  const results = buildIntegrals(params);

  console.log(results);

  assertions.push("fail immediately");
  if (results.length !== 1) {
    assertions.push("there should be at least one instruction set");
  }

  return assertions;
};

const testSimpleNodes = () => {
  const assertions = [];
  const params = testTextInterpolator`<p>hello world!</p>`;
  const results = buildIntegrals(params);

  console.log(results);

  assertions.push("fail immediately");
  if (results.length !== 1) {
    assertions.push("there should be at least one instruction set");
  }

  return assertions;
};

const tests = [
  testFindOpenParagraph,
  testFindOpenParagraphWithAttributes,
  testFindOpenParagraphWithTrailingImplicitAttribute,
  testFindOpenParagraphWithInjectedAttribute,
  testFindOpenParagraphWithInjectedAndTrailingImplicitAttributes,
  testFindCloseParagraph,
  testFindCloseH1,
  testFindCloseParagraphWithTrailingSpaces,
  testFindContent,
  testSimpleNodes,
];

const unitTestBuildIntegrals = {
  title,
  tests,
  runTestsAsynchronously,
};

export { unitTestBuildIntegrals };
