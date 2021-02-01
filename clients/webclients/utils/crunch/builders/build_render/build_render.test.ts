// brian taylor vann
// build render

import {
  buildIntegrals,
  BuildIntegralsParams,
} from "../build_integrals/build_integrals";
import { buildSkeleton } from "../build_skeleton/build_skeleton";
import { hooks } from "../../test_hooks/test_hooks";
import { Context } from "../../type_flyweight/context";
import { Integrals } from "../../type_flyweight/integrals";

import { samestuff } from "../../../little_test_runner/samestuff/samestuff";
import { Template } from "../../type_flyweight/template";
import { buildRender } from "../build_render/build_render";
import { TestAttributes, TestNode } from "../../test_hooks/test_element";

interface InterpolatorResults<N, A> {
  template: Template<N, A>;
  integrals: Integrals;
}
type TextTextInterpolator<N, A> = (
  templateArray: TemplateStringsArray,
  ...injections: (Context<N, A> | A)[]
) => InterpolatorResults<N, A>;

const testTextInterpolator: TextTextInterpolator<TestNode, string> = (
  templateArray,
  ...injections
) => {
  const template = { templateArray, injections };
  const params = {
    skeleton: buildSkeleton(template),
    template,
  };

  return {
    template,
    integrals: buildIntegrals(params),
  };
};

const title = "build_render";
const runTestsAsynchronously = true;

// createNode,
const testCreateNode = () => {
  const assertions: string[] = [];

  const { template, integrals } = testTextInterpolator`<p>`;
  const results = buildRender({
    hooks,
    integrals,
    template,
  });

  console.log(results);

  return assertions;
};

const testCloseNode = () => {
  const assertions: string[] = [];

  const { template, integrals } = testTextInterpolator`<p></p>`;
  const results = buildRender({
    hooks,
    integrals,
    template,
  });

  console.log(results);

  return assertions;
};

const testTextNode = () => {
  const assertions: string[] = [];

  const {
    template,
    integrals,
  } = testTextInterpolator`hello world!<p>It's me!</p>`;
  const results = buildRender({
    hooks,
    integrals,
    template,
  });

  console.log(results);

  return assertions;
};

const testAddAttributesToNodes = () => {
  const assertions: string[] = [];

  const { template, integrals } = testTextInterpolator`
    <p
      checked
      disabled="false"
      skies="${"blue"}">
        Hello world, it's me!
    </p>`;
  const results = buildRender({
    hooks,
    integrals,
    template,
  });

  console.log(results);

  return assertions;
};

const testAddAttributesToMultipleNodes = () => {
  const assertions: string[] = [];

  const { template, integrals } = testTextInterpolator`
    <p>No properties in this paragraph!</p>
    <p
      checked
      disabled="false"
      skies="${"blue"}">
        Hello world, it's me!
    </p>`;
  const results = buildRender({
    hooks,
    integrals,
    template,
  });

  console.log(results);

  return assertions;
};

const testAddContext = () => {
  const assertions: string[] = [];

  const { template, integrals } = testTextInterpolator`
    <p>HelloWorld!</p>
  `;
  const renderStructure = buildRender({
    hooks,
    integrals,
    template,
  });

  const context = new Context({ renderStructure });
  const {
    integrals: contextIntegrals,
    template: contextTemplate,
  } = testTextInterpolator`
    <p>${context}</p>
  `;
  const contextRenderStructure = buildRender({
    hooks,
    integrals: contextIntegrals,
    template: contextTemplate,
  });

  console.log(contextRenderStructure);

  return assertions;
};

const tests = [
  testCreateNode,
  testCloseNode,
  testTextNode,
  testAddAttributesToNodes,
  testAddAttributesToMultipleNodes,
  testAddContext,
];

const unitTestBuildRender = {
  title,
  tests,
  runTestsAsynchronously,
};

export { unitTestBuildRender };
