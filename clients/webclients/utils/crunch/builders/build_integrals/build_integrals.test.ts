// brian taylor vann
// build integrals

import { samestuff } from "../../../little_test_runner/samestuff/samestuff";
import { BuildIntegralsParams } from "./build_integrals";
import { buildSkeleton } from "../build_skeleton/build_skeleton";
import { buildIntegrals } from "./build_integrals";
import { Integrals } from "../../type_flyweight/integrals";

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

const testFindOpenParagraphWithAttributes = () => {
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

const testFindOpenParagraphWithTrailingImplicitAttribute = () => {
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

const testFindOpenParagraphWithInjectedAttribute = () => {
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

const testFindOpenParagraphWithInjectedAndTrailingImplicitAttributes = () => {
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
  console.log(results);
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
