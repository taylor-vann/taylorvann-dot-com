// brian taylor vann
// build integrals

// we need an injection test or two

import { AttributeValue } from "../../type_flyweight/hooks";
import { samestuff } from "../../../little_test_runner/samestuff/samestuff";
import { BuildIntegralsParams } from "./build_integrals";
import { buildSkeleton } from "../build_skeleton/build_skeleton";
import { buildIntegrals } from "./build_integrals";
import { Integrals } from "../../type_flyweight/integrals";
import { TestAttributes, TestNode } from "../../test_hooks/test_element";

type TextTextInterpolator<N, A> = (
  templateArray: TemplateStringsArray,
  ...injections: AttributeValue<N, A>[]
) => BuildIntegralsParams<N, A>;

const title = "build_integrals";
const runTestsAsynchronously = true;

const testTextInterpolator: TextTextInterpolator<TestNode, TestAttributes> = (
  templateArray,
  ...injections
) => {
  const template = { templateArray, injections };
  return {
    template: template,
    skeleton: buildSkeleton(template),
  };
};

const findParagraph = () => {
  const assertions = [];

  const expectedResults: Integrals = [
    {
      kind: "NODE",
      tagNameVector: {
        origin: {
          arrayIndex: 0,
          stringIndex: 1,
        },
        target: {
          arrayIndex: 0,
          stringIndex: 1,
        },
      },
    },
  ];

  const params = testTextInterpolator`<p>`;
  const results = buildIntegrals(params);

  if (!samestuff(expectedResults, results)) {
    assertions.push("unexpected results found.");
  }

  return assertions;
};

const findParagraphWithAttributes = () => {
  const assertions = [];

  const expectedResults: Integrals = [
    {
      kind: "NODE",
      tagNameVector: {
        origin: {
          arrayIndex: 0,
          stringIndex: 1,
        },
        target: {
          arrayIndex: 0,
          stringIndex: 1,
        },
      },
    },
    {
      kind: "EXPLICIT_ATTRIBUTE",
      attributeVector: {
        origin: {
          arrayIndex: 0,
          stringIndex: 3,
        },
        target: {
          arrayIndex: 0,
          stringIndex: 9,
        },
      },
      valueVector: {
        origin: {
          arrayIndex: 0,
          stringIndex: 11,
        },
        target: {
          arrayIndex: 0,
          stringIndex: 25,
        },
      },
    },
  ];

  const params = testTextInterpolator`<p message="hello, world!">`;
  const results = buildIntegrals(params);

  if (!samestuff(expectedResults, results)) {
    assertions.push("unexpected results found.");
  }

  return assertions;
};

const findParagraphWithImplicitAttribute = () => {
  const assertions = [];

  const expectedResults: Integrals = [
    {
      kind: "NODE",
      tagNameVector: {
        origin: {
          arrayIndex: 0,
          stringIndex: 1,
        },
        target: {
          arrayIndex: 0,
          stringIndex: 1,
        },
      },
    },

    {
      kind: "EXPLICIT_ATTRIBUTE",
      attributeVector: {
        origin: {
          arrayIndex: 0,
          stringIndex: 3,
        },
        target: {
          arrayIndex: 0,
          stringIndex: 9,
        },
      },
      valueVector: {
        origin: {
          arrayIndex: 0,
          stringIndex: 11,
        },
        target: {
          arrayIndex: 0,
          stringIndex: 25,
        },
      },
    },

    {
      kind: "IMPLICIT_ATTRIBUTE",
      attributeVector: {
        origin: {
          arrayIndex: 0,
          stringIndex: 27,
        },
        target: {
          arrayIndex: 0,
          stringIndex: 33,
        },
      },
    },
  ];

  const params = testTextInterpolator`<p message="hello, world!" checked>`;
  const results = buildIntegrals(params);

  if (!samestuff(expectedResults, results)) {
    assertions.push("unexpected results found.");
  }

  return assertions;
};

const findParagraphWithInjectedAttribute = () => {
  const assertions = [];

  const expectedResults: Integrals = [
    {
      kind: "NODE",
      tagNameVector: {
        origin: {
          arrayIndex: 0,
          stringIndex: 1,
        },
        target: {
          arrayIndex: 0,
          stringIndex: 1,
        },
      },
    },
    {
      kind: "INJECTED_ATTRIBUTE",
      attributeVector: {
        origin: {
          arrayIndex: 0,
          stringIndex: 3,
        },
        target: {
          arrayIndex: 0,
          stringIndex: 9,
        },
      },
      valueVector: {
        origin: {
          arrayIndex: 0,
          stringIndex: 11,
        },
        target: {
          arrayIndex: 1,
          stringIndex: 0,
        },
      },
      injectionID: 0,
    },
  ];

  const params = testTextInterpolator`<p message="${"hello, world!"}">`;
  const results = buildIntegrals(params);

  if (!samestuff(expectedResults, results)) {
    assertions.push("unexpected results found.");
  }

  return assertions;
};

const findParagraphWithInjectedAndImplicitAttributes = () => {
  const assertions = [];

  const expectedResults: Integrals = [
    {
      kind: "NODE",
      tagNameVector: {
        origin: {
          arrayIndex: 0,
          stringIndex: 1,
        },
        target: {
          arrayIndex: 0,
          stringIndex: 1,
        },
      },
    },
    {
      kind: "INJECTED_ATTRIBUTE",
      attributeVector: {
        origin: {
          arrayIndex: 0,
          stringIndex: 3,
        },
        target: {
          arrayIndex: 0,
          stringIndex: 9,
        },
      },
      valueVector: {
        origin: {
          arrayIndex: 0,
          stringIndex: 11,
        },
        target: {
          arrayIndex: 1,
          stringIndex: 0,
        },
      },
      injectionID: 0,
    },

    {
      kind: "IMPLICIT_ATTRIBUTE",
      attributeVector: {
        origin: {
          arrayIndex: 1,
          stringIndex: 2,
        },
        target: {
          arrayIndex: 1,
          stringIndex: 8,
        },
      },
    },
  ];

  const params = testTextInterpolator`<p message="${"hello, world!"}" checked>`;
  const results = buildIntegrals(params);

  if (!samestuff(expectedResults, results)) {
    assertions.push("unexpected results found.");
  }

  return assertions;
};

const testFindCloseParagraph = () => {
  const assertions = [];

  const expectedResults: Integrals = [
    {
      kind: "CLOSE_NODE",
      tagNameVector: {
        origin: {
          arrayIndex: 0,
          stringIndex: 2,
        },
        target: {
          arrayIndex: 0,
          stringIndex: 2,
        },
      },
    },
  ];

  const params = testTextInterpolator`</p>`;
  const results = buildIntegrals(params);

  if (!samestuff(expectedResults, results)) {
    assertions.push("unexpected results found.");
  }

  return assertions;
};

const testFindCloseH1 = () => {
  const assertions = [];

  const expectedResults: Integrals = [
    {
      kind: "CLOSE_NODE",
      tagNameVector: {
        origin: {
          arrayIndex: 0,
          stringIndex: 2,
        },
        target: {
          arrayIndex: 0,
          stringIndex: 3,
        },
      },
    },
  ];

  const params = testTextInterpolator`</h1>`;
  const results = buildIntegrals(params);

  if (!samestuff(expectedResults, results)) {
    assertions.push("unexpected results found.");
  }

  return assertions;
};

const testFindCloseParagraphWithTrailingSpaces = () => {
  const assertions = [];

  const expectedResults: Integrals = [
    {
      kind: "CLOSE_NODE",
      tagNameVector: {
        origin: {
          arrayIndex: 0,
          stringIndex: 2,
        },
        target: {
          arrayIndex: 0,
          stringIndex: 3,
        },
      },
    },
  ];

  const params = testTextInterpolator`</h1        >`;
  const results = buildIntegrals(params);

  if (!samestuff(expectedResults, results)) {
    assertions.push("unexpected results found.");
  }

  return assertions;
};

const testFindContent = () => {
  const assertions = [];

  const expectedResults: Integrals = [
    {
      kind: "TEXT",
      textVector: {
        origin: {
          arrayIndex: 0,
          stringIndex: 0,
        },
        target: {
          arrayIndex: 0,
          stringIndex: 11,
        },
      },
    },
  ];

  const params = testTextInterpolator`hello world!`;
  const results = buildIntegrals(params);

  if (!samestuff(expectedResults, results)) {
    assertions.push("unexpected results found.");
  }

  return assertions;
};

const testFindContentWithInjection = () => {
  const assertions = [];

  const expectedResults: Integrals = [
    {
      kind: "TEXT",
      textVector: {
        origin: {
          arrayIndex: 0,
          stringIndex: 0,
        },
        target: {
          arrayIndex: 0,
          stringIndex: -1,
        },
      },
    },
    {
      kind: "CONTEXT_INJECTION",
      injectionID: 0,
    },
    {
      kind: "TEXT",
      textVector: {
        origin: {
          arrayIndex: 1,
          stringIndex: 0,
        },
        target: {
          arrayIndex: 1,
          stringIndex: 11,
        },
      },
    },
  ];

  const params = testTextInterpolator`${"hello"}hello world!`;
  const results = buildIntegrals(params);

  if (!samestuff(expectedResults, results)) {
    assertions.push("unexpected results found.");
  }

  return assertions;
};

const testFindContentWithInitialMultipleInjections = () => {
  const assertions = [];

  const expectedResults: Integrals = [
    {
      kind: "TEXT",
      textVector: {
        origin: {
          arrayIndex: 0,
          stringIndex: 0,
        },
        target: {
          arrayIndex: 0,
          stringIndex: -1,
        },
      },
    },
    {
      kind: "CONTEXT_INJECTION",
      injectionID: 0,
    },
    {
      kind: "TEXT",
      textVector: {
        origin: {
          arrayIndex: 1,
          stringIndex: 0,
        },
        target: {
          arrayIndex: 1,
          stringIndex: 11,
        },
      },
    },
    {
      kind: "CONTEXT_INJECTION",
      injectionID: 1,
    },
    {
      kind: "TEXT",
      textVector: {
        origin: {
          arrayIndex: 2,
          stringIndex: 0,
        },
        target: {
          arrayIndex: 2,
          stringIndex: 0,
        },
      },
    },
  ];

  const params = testTextInterpolator`${"heyyo"}hello world,${"you're awesome"}!`;
  const results = buildIntegrals(params);

  if (!samestuff(expectedResults, results)) {
    assertions.push("unexpected results found.");
  }

  return assertions;
};

const testFindContentWithEdgeCaseInjections = () => {
  const assertions = [];

  const expectedResults: Integrals = [
    {
      kind: "CONTEXT_INJECTION",
      injectionID: 0,
    },
    {
      kind: "NODE",
      tagNameVector: {
        origin: {
          arrayIndex: 1,
          stringIndex: 1,
        },
        target: {
          arrayIndex: 1,
          stringIndex: 1,
        },
      },
    },
    {
      kind: "TEXT",
      textVector: {
        origin: {
          arrayIndex: 1,
          stringIndex: 3,
        },
        target: {
          arrayIndex: 1,
          stringIndex: 14,
        },
      },
    },
    {
      kind: "CONTEXT_INJECTION",
      injectionID: 1,
    },
    {
      kind: "CLOSE_NODE",
      tagNameVector: {
        origin: {
          arrayIndex: 2,
          stringIndex: 2,
        },
        target: {
          arrayIndex: 2,
          stringIndex: 2,
        },
      },
    },
    {
      kind: "NODE",
      tagNameVector: {
        origin: {
          arrayIndex: 2,
          stringIndex: 5,
        },
        target: {
          arrayIndex: 2,
          stringIndex: 9,
        },
      },
    },
    {
      kind: "INJECTED_ATTRIBUTE",
      attributeVector: {
        origin: {
          arrayIndex: 2,
          stringIndex: 11,
        },
        target: {
          arrayIndex: 2,
          stringIndex: 13,
        },
      },
      valueVector: {
        origin: {
          arrayIndex: 2,
          stringIndex: 15,
        },
        target: {
          arrayIndex: 3,
          stringIndex: 0,
        },
      },
      injectionID: 2,
    },
  ];

  const params = testTextInterpolator`${"heyyo"}<p>hello world,${"you're awesome"}</p><image src="${"hello_world"}">`;
  const results = buildIntegrals(params);

  if (!samestuff(expectedResults, results)) {
    assertions.push("unexpected results found.");
  }

  return assertions;
};

const testSimpleNodes = () => {
  const assertions = [];

  const expectedResults: Integrals = [
    {
      kind: "NODE",
      tagNameVector: {
        origin: {
          arrayIndex: 0,
          stringIndex: 1,
        },
        target: {
          arrayIndex: 0,
          stringIndex: 1,
        },
      },
    },
    {
      kind: "TEXT",
      textVector: {
        origin: {
          arrayIndex: 0,
          stringIndex: 3,
        },
        target: {
          arrayIndex: 0,
          stringIndex: 14,
        },
      },
    },
    {
      kind: "CLOSE_NODE",
      tagNameVector: {
        origin: {
          arrayIndex: 0,
          stringIndex: 17,
        },
        target: {
          arrayIndex: 0,
          stringIndex: 17,
        },
      },
    },
  ];

  const params = testTextInterpolator`<p>hello world!</p>`;
  const results = buildIntegrals(params);

  if (!samestuff(expectedResults, results)) {
    assertions.push("unexpected results found.");
  }

  return assertions;
};

const tests = [
  findParagraph,
  testFindContentWithInjection,
  findParagraphWithAttributes,
  findParagraphWithImplicitAttribute,
  findParagraphWithInjectedAttribute,
  findParagraphWithInjectedAndImplicitAttributes,
  testFindContentWithInitialMultipleInjections,
  testFindContentWithEdgeCaseInjections,
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
