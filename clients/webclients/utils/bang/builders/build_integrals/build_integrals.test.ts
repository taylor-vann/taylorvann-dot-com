// brian taylor vann
// build integrals

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

const tests = [
  testFindCloseParagraph,
  testFindCloseH1,
  testFindCloseParagraphWithTrailingSpaces,
];

const unitTestBuildIntegrals = {
  title,
  tests,
  runTestsAsynchronously,
};

export { unitTestBuildIntegrals };
