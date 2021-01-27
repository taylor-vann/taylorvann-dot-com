// brian taylor vann
// attribute crawl

import { samestuff } from "../../../little_test_runner/samestuff/samestuff";
import { Template } from "../../type_flyweight/template";
import { create, incrementTarget } from "../../text_vector/text_vector";
import { crawlForAttribute } from "./attribute_crawl";
import {
  ExplicitAttributeAction,
  InjectedAttributeAction,
} from "../../type_flyweight/attribute_crawl";

type TextTextInterpolator = <A>(
  templateArray: TemplateStringsArray,
  ...injections: A[]
) => Template<A>;

const RECURSION_SAFETY = 256;

const testTextInterpolator: TextTextInterpolator = (
  templateArray,
  ...injections
) => {
  return { templateArray, injections };
};

const title = "attribute_crawl";
const runTestsAsynchronously = true;

const emptyString = () => {
  const assertions = [];

  const template = testTextInterpolator``;
  const vector = create();

  const results = crawlForAttribute(template, vector);

  if (results !== undefined) {
    assertions.push("this should have failed");
  }

  return assertions;
};

const emptySpaceString = () => {
  const assertions = [];

  const template = testTextInterpolator` `;
  const vector = create();

  let safety = 0;
  while (incrementTarget(template, vector) && safety < RECURSION_SAFETY) {
    safety += 1;
  }

  const results = crawlForAttribute(template, vector);

  if (results !== undefined) {
    assertions.push("this should have failed");
  }

  return assertions;
};

const emptyMultiSpaceString = () => {
  const assertions = [];

  const template = testTextInterpolator`   `;
  const vector = create();

  let safety = 0;
  while (incrementTarget(template, vector) && safety < RECURSION_SAFETY) {
    safety += 1;
  }

  const results = crawlForAttribute(template, vector);

  if (results !== undefined) {
    assertions.push("this should have failed");
  }

  return assertions;
};

const implicitString = () => {
  const assertions = [];

  const expectedResults = {
    kind: "IMPLICIT_ATTRIBUTE",

    attributeVector: {
      origin: {
        arrayIndex: 0,
        stringIndex: 0,
      },
      target: {
        arrayIndex: 0,
        stringIndex: 6,
      },
    },
  };

  const template = testTextInterpolator`checked`;
  const vector = create();

  let safety = 0;
  while (incrementTarget(template, vector) && safety < RECURSION_SAFETY) {
    safety += 1;
  }

  const results = crawlForAttribute(template, vector);

  if (!samestuff(expectedResults, results)) {
    assertions.push("unexpected results found.");
  }

  return assertions;
};

const implicitStringWithTrailingSpaces = () => {
  const assertions = [];

  const expectedResults = {
    kind: "IMPLICIT_ATTRIBUTE",

    attributeVector: {
      origin: {
        arrayIndex: 0,
        stringIndex: 0,
      },
      target: {
        arrayIndex: 0,
        stringIndex: 6,
      },
    },
  };

  const template = testTextInterpolator`checked    `;
  const vector = create();

  let safety = 0;
  while (incrementTarget(template, vector) && safety < RECURSION_SAFETY) {
    safety += 1;
  }

  const results = crawlForAttribute(template, vector);

  if (!samestuff(expectedResults, results)) {
    assertions.push("unexpected results found.");
  }

  return assertions;
};

const malformedExplicitString = () => {
  const assertions = [];

  const template = testTextInterpolator`checked=`;
  const vector = create();

  let safety = 0;
  while (incrementTarget(template, vector) && safety < RECURSION_SAFETY) {
    safety += 1;
  }

  const results = crawlForAttribute(template, vector);

  if (results !== undefined) {
    assertions.push("this should not have returned results");
  }

  return assertions;
};

const almostExplicitString = () => {
  const assertions = [];

  const template = testTextInterpolator`checked="`;
  const vector = create();

  let safety = 0;
  while (incrementTarget(template, vector) && safety < RECURSION_SAFETY) {
    safety += 1;
  }

  const results = crawlForAttribute(template, vector);
  if (results !== undefined) {
    assertions.push("this should not have returned results");
  }

  return assertions;
};

const emptyExplicitString = () => {
  const assertions = [];

  const expectedResults: ExplicitAttributeAction = {
    kind: "EXPLICIT_ATTRIBUTE",
    attributeVector: {
      origin: {
        arrayIndex: 0,
        stringIndex: 0,
      },
      target: {
        arrayIndex: 0,
        stringIndex: 6,
      },
    },
    valueVector: {
      origin: {
        arrayIndex: 0,
        stringIndex: 8,
      },
      target: {
        arrayIndex: 0,
        stringIndex: 9,
      },
    },
  };

  const template = testTextInterpolator`checked=""`;
  const vector = create();

  let safety = 0;
  while (incrementTarget(template, vector) && safety < RECURSION_SAFETY) {
    safety += 1;
  }

  const results = crawlForAttribute(template, vector);

  if (!samestuff(expectedResults, results)) {
    assertions.push("unexpected results found.");
  }

  return assertions;
};

const validExplicitString = () => {
  const assertions = [];

  const expectedResults: ExplicitAttributeAction = {
    kind: "EXPLICIT_ATTRIBUTE",

    attributeVector: {
      origin: {
        arrayIndex: 0,
        stringIndex: 0,
      },
      target: {
        arrayIndex: 0,
        stringIndex: 6,
      },
    },
    valueVector: {
      origin: {
        arrayIndex: 0,
        stringIndex: 8,
      },
      target: {
        arrayIndex: 0,
        stringIndex: 16,
      },
    },
  };

  const template = testTextInterpolator`checked="checked"`;
  const vector = create();

  let safety = 0;
  while (incrementTarget(template, vector) && safety < RECURSION_SAFETY) {
    safety += 1;
  }

  const results = crawlForAttribute(template, vector);

  if (!samestuff(expectedResults, results)) {
    assertions.push("unexpected results found.");
  }

  return assertions;
};

const validExplicitStringWithTrailingSpaces = () => {
  const assertions = [];

  const expectedResults: ExplicitAttributeAction = {
    kind: "EXPLICIT_ATTRIBUTE",

    attributeVector: {
      origin: {
        arrayIndex: 0,
        stringIndex: 0,
      },
      target: {
        arrayIndex: 0,
        stringIndex: 6,
      },
    },
    valueVector: {
      origin: {
        arrayIndex: 0,
        stringIndex: 8,
      },
      target: {
        arrayIndex: 0,
        stringIndex: 19,
      },
    },
  };

  const template = testTextInterpolator`checked="checked   "`;
  const vector = create();

  let safety = 0;
  while (incrementTarget(template, vector) && safety < RECURSION_SAFETY) {
    safety += 1;
  }

  const results = crawlForAttribute(template, vector);

  if (!samestuff(expectedResults, results)) {
    assertions.push("unexpected results found.");
  }

  return assertions;
};

const injectedString = () => {
  const assertions = [];

  const expectedResults: InjectedAttributeAction = {
    kind: "INJECTED_ATTRIBUTE",

    attributeVector: {
      origin: {
        arrayIndex: 0,
        stringIndex: 0,
      },
      target: {
        arrayIndex: 0,
        stringIndex: 6,
      },
    },
    valueVector: {
      origin: {
        arrayIndex: 0,
        stringIndex: 8,
      },
      target: {
        arrayIndex: 1,
        stringIndex: 0,
      },
    },
    injectionID: 0,
  };

  const template = testTextInterpolator`checked="${"hello"}"`;
  const vector = create();

  let safety = 0;
  while (incrementTarget(template, vector) && safety < RECURSION_SAFETY) {
    safety += 1;
  }

  const results = crawlForAttribute(template, vector);

  if (!samestuff(expectedResults, results)) {
    assertions.push("unexpected results found.");
  }

  if (results === undefined) {
    assertions.push("this should have returned results");
  }

  return assertions;
};

const malformedInjectedString = () => {
  const assertions = [];

  const template = testTextInterpolator`checked="${"hello"}`;
  const vector = create();

  let safety = 0;
  while (incrementTarget(template, vector) && safety < RECURSION_SAFETY) {
    safety += 1;
  }

  const results = crawlForAttribute(template, vector);
  if (results !== undefined) {
    assertions.push("this should have returned results");
  }

  return assertions;
};

const malformedInjectedStringWithTrailingSpaces = () => {
  const assertions = [];

  const template = testTextInterpolator`checked="${"hello"} "`;
  const vector = create();

  let safety = 0;
  while (incrementTarget(template, vector) && safety < RECURSION_SAFETY) {
    safety += 1;
  }

  const results = crawlForAttribute(template, vector);
  if (results !== undefined) {
    assertions.push("this should not have returned results");
  }

  return assertions;
};

const malformedInjectedStringWithStartingSpaces = () => {
  const assertions = [];

  const template = testTextInterpolator`checked=" ${"hello"}"`;
  const vector = create();

  let safety = 0;
  while (incrementTarget(template, vector) && safety < RECURSION_SAFETY) {
    safety += 1;
  }

  const results = crawlForAttribute(template, vector);
  if (results !== undefined) {
    assertions.push("this should not have returned results");
  }

  return assertions;
};

const tests = [
  emptyString,
  emptySpaceString,
  emptyMultiSpaceString,
  implicitString,
  implicitStringWithTrailingSpaces,
  malformedExplicitString,
  almostExplicitString,
  emptyExplicitString,
  validExplicitString,
  validExplicitStringWithTrailingSpaces,
  injectedString,
  malformedInjectedString,
  malformedInjectedStringWithTrailingSpaces,
  malformedInjectedStringWithStartingSpaces,
];

const unitTestAttributeCrawl = {
  title,
  tests,
  runTestsAsynchronously,
};

export { unitTestAttributeCrawl };
