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

const defaultFunc = () => {
  const assertions = [];
  const params = testTextInterpolator`<p>`;
  const results = buildIntegrals(params);

  assertions.push("fail immediately");

  if (results.length !== 1) {
    assertions.push("there should be at least one instruction set");
  }

  return [];
};

const tests = [defaultFunc];

const unitTestBuildIntegrals = {
  title,
  tests,
  runTestsAsynchronously,
};

export { unitTestBuildIntegrals };
