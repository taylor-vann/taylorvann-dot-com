// brian taylor vann
// build render

import { buildIntegrals } from "../build_integrals/build_integrals";
import { buildRender } from "../build_render/build_render";
import { buildSkeleton } from "../build_skeleton/build_skeleton";
import { Context } from "../../type_flyweight/context";
import { hooks } from "../../test_hooks/test_hooks";
import { Integrals } from "../../type_flyweight/integrals";
import { Template } from "../../type_flyweight/template";
import { TestNode } from "../../test_hooks/test_element";

interface InterpolatorResults<N, A> {
  template: Template<N, A>;
  integrals: Integrals;
}
type TextTextInterpolator<N, A> = (
  templateArray: TemplateStringsArray,
  ...injections: (Context<N, A> | A)[]
) => InterpolatorResults<N, A>;

const title = "build_render";
const runTestsAsynchronously = true;

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

// createNode,
const testCreateNode = () => {
  const assertions: string[] = [];

  const { template, integrals } = testTextInterpolator`<p>`;
  const results = buildRender({
    hooks,
    integrals,
    template,
  });

  if (results.siblings.length !== 1) {
    assertions.push("siblings should have length 1");
    return assertions;
  }

  const sibling = results.siblings[0];

  if (sibling.kind !== "ELEMENT") {
    assertions.push("sibling should be an ELEMENT");
    return assertions;
  }
  if (sibling.tagname !== "p") {
    assertions.push("sibling tagname should be p");
  }

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

  if (results.siblings.length !== 1) {
    assertions.push("siblings should have length 1");
    return assertions;
  }

  const sibling = results.siblings[0];

  if (sibling.kind !== "ELEMENT") {
    assertions.push("sibling should be an ELEMENT");
    return assertions;
  }
  if (sibling.tagname !== "p") {
    assertions.push("sibling tagname should be p");
  }

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

  if (results.siblings.length !== 2) {
    assertions.push("siblings should have length 1");
    return assertions;
  }

  const hopefulText = results.siblings[0];
  const hopefulElement = results.siblings[1];

  if (hopefulText.kind !== "TEXT") {
    assertions.push("sibling should be an TEXT");
  }
  if (hopefulText.kind === "TEXT" && hopefulText.text !== "hello world!") {
    assertions.push("sibling tagname should be p");
  }

  if (hopefulElement.kind !== "ELEMENT") {
    assertions.push("sibling should be an ELEMENT");
  }
  if (hopefulElement.kind === "ELEMENT" && hopefulElement.tagname !== "p") {
    assertions.push("sibling tagname should be p");
  }

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

  if (results.siblings.length !== 2) {
    assertions.push("siblings should have length 2");
    return assertions;
  }

  const sibling = results.siblings[1];

  if (sibling.kind !== "ELEMENT") {
    assertions.push("sibling should be an ELEMENT");
    return assertions;
  }
  if (sibling.tagname !== "p") {
    assertions.push("sibling tagname should be p");
  }
  if (sibling.attributes["checked"] !== true) {
    assertions.push("sibling should be checked");
  }
  if (sibling.attributes["disabled"] !== "false") {
    assertions.push("sibling should be disabled");
  }
  if (sibling.attributes["skies"] !== "blue") {
    assertions.push("sibling skies should be blue");
  }

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
  if (results.siblings.length !== 4) {
    assertions.push("siblings should have length 4");
    return assertions;
  }

  const firstParagraph = results.siblings[1];
  if (firstParagraph.kind !== "ELEMENT") {
    assertions.push("sibling should be an ELEMENT");
  }
  if (firstParagraph.kind === "ELEMENT" && firstParagraph.tagname !== "p") {
    assertions.push("sibling tagname should be p");
  }

  const secondParagraph = results.siblings[2];
  if (secondParagraph.kind !== "ELEMENT") {
    assertions.push("sibling should be an ELEMENT");
    return assertions;
  }
  if (secondParagraph.tagname !== "p") {
    assertions.push("sibling tagname should be p");
  }
  if (secondParagraph.attributes["checked"] !== true) {
    assertions.push("sibling should be checked");
  }
  if (secondParagraph.attributes["disabled"] !== "false") {
    assertions.push("sibling should be checked");
  }
  if (secondParagraph.attributes["disabled"] !== "false") {
    assertions.push("sibling should be checked");
  }

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

  const results = buildRender({
    hooks,
    integrals: contextIntegrals,
    template: contextTemplate,
  });

  if (results.siblings.length !== 3) {
    assertions.push("siblings should have length 3");
    return assertions;
  }

  const paragraph = results.siblings[1];
  if (paragraph.kind !== "ELEMENT") {
    assertions.push("second sibling should be an ELEMENT");
    return assertions;
  }

  const contextParagraph = paragraph?.leftChild?.right;
  if (contextParagraph === undefined) {
    assertions.push("there should be a second sibling");
    return assertions;
  }
  if (contextParagraph.kind !== "ELEMENT") {
    assertions.push("second sibling should be an ELEMENT");
    return assertions;
  }
  if (contextParagraph.tagname !== "p") {
    assertions.push("second sibling tagname should be p");
  }

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
