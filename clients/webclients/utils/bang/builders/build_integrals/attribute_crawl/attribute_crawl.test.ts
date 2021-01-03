// brian taylor vann
// build integrals

import { StructureRender } from "../../../type_flyweight/structure";
import { create, incrementTarget } from "../../../text_vector/text_vector";

import { crawlForAttribute } from "./attribute_crawl";

type TextTextInterpolator = <A>(
  templateArray: TemplateStringsArray,
  ...injections: A[]
) => StructureRender<A>;

const RECURSION_SAFETY = 256;

const testTextInterpolator: TextTextInterpolator = (
  templateArray,
  ...injections
) => {
  return { templateArray, injections };
};

const title = "attribute_crawl";
const runTestsAsynchronously = true;

// // but we are seaching between and not incliding '<' '>'
// " " // invalid
//
// // checkbox // invalid
// // checkbox checked // valid
// // checkbox checked  // valid

// // checkbox hello="" // valid
// // checkbox hello="world" // valid
// // checkbox hello="${"world"}" // valid

// // we are looking mainly for ="(-->)"
// //   or we are looking for

const testEmptyString = () => {
  const assertions = [];

  const template = testTextInterpolator``;
  const vector = create();
  const results = crawlForAttribute(template, vector);

  if (results !== undefined) {
    assertions.push("this should have failed");
  }

  return assertions;
};

const testEmptySpaceString = () => {
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

const testEmptyMultiSpaceString = () => {
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

const testImplicitString = () => {
  const assertions = [];

  const template = testTextInterpolator`checked`;
  const vector = create();

  let safety = 0;
  while (incrementTarget(template, vector) && safety < RECURSION_SAFETY) {
    safety += 1;
  }

  const results = crawlForAttribute(template, vector);

  if (results === undefined) {
    assertions.push("this should not have returned results");
  }
  if (
    results !== undefined &&
    results.action !== "IMPLICIT_ATTRIBUTE_CONFIRMED"
  ) {
    assertions.push("should return IMPLICIT_ATTRIBUTE_CONFIRMED");
  }

  if (
    results !== undefined &&
    results.action === "IMPLICIT_ATTRIBUTE_CONFIRMED"
  ) {
    if (results.params.attributeVector.origin.arrayIndex !== 0) {
      assertions.push("origin.arrayIndex should be 0.");
    }
    if (results.params.attributeVector.origin.stringIndex !== 0) {
      assertions.push("origin.stringIndex should be 0.");
    }

    if (results.params.attributeVector.target.arrayIndex !== 0) {
      assertions.push("target.arrayIndex should be 0.");
    }
    if (results.params.attributeVector.target.stringIndex !== 6) {
      assertions.push("target.stringIndex should be 6.");
    }
  }

  return assertions;
};

const testImplicitStringWithTrailingSpaces = () => {
  const assertions = [];

  const template = testTextInterpolator`checked    `;
  const vector = create();

  let safety = 0;
  while (incrementTarget(template, vector) && safety < RECURSION_SAFETY) {
    safety += 1;
  }

  const results = crawlForAttribute(template, vector);

  if (results === undefined) {
    assertions.push("this should not have returned results");
  }
  if (
    results !== undefined &&
    results.action !== "IMPLICIT_ATTRIBUTE_CONFIRMED"
  ) {
    assertions.push("should return IMPLICIT_ATTRIBUTE_CONFIRMED");
  }

  if (
    results !== undefined &&
    results.action === "IMPLICIT_ATTRIBUTE_CONFIRMED"
  ) {
    if (results.params.attributeVector.origin.arrayIndex !== 0) {
      assertions.push("origin.arrayIndex should be 0.");
    }
    if (results.params.attributeVector.origin.stringIndex !== 0) {
      assertions.push("origin.stringIndex should be 0.");
    }

    if (results.params.attributeVector.target.arrayIndex !== 0) {
      assertions.push("target.arrayIndex should be 0.");
    }
    if (results.params.attributeVector.target.stringIndex !== 6) {
      assertions.push("target.stringIndex should be 6.");
    }
  }

  return assertions;
};

const testMalformedExplicitString = () => {
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

const testAlmostExplicitString = () => {
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

const testEmptyExplicitString = () => {
  const assertions = [];

  const template = testTextInterpolator`checked=""`;
  const vector = create();

  let safety = 0;
  while (incrementTarget(template, vector) && safety < RECURSION_SAFETY) {
    safety += 1;
  }

  const results = crawlForAttribute(template, vector);
  if (results === undefined) {
    assertions.push("this should have returned results");
  }

  if (
    results !== undefined &&
    results.action === "EXPLICIT_ATTRIBUTE_CONFIRMED"
  ) {
    if (results.params.attributeVector.origin.arrayIndex !== 0) {
      assertions.push("attributeVector origin.arrayIndex should be 0.");
    }
    if (results.params.attributeVector.origin.stringIndex !== 0) {
      assertions.push("attributeVector origin.stringIndex should be 0.");
    }

    if (results.params.attributeVector.target.arrayIndex !== 0) {
      assertions.push("attributeVector target.arrayIndex should be 0.");
    }
    if (results.params.attributeVector.target.stringIndex !== 6) {
      assertions.push("attributeVector target.stringIndex should be 6.");
    }

    if (results.params.valueVector.origin.arrayIndex !== 0) {
      assertions.push("valueVector origin.arrayIndex should be 0.");
    }
    if (results.params.valueVector.origin.stringIndex !== 8) {
      assertions.push("valueVector origin.stringIndex should be 0.");
    }

    if (results.params.valueVector.target.arrayIndex !== 0) {
      assertions.push("valueVector target.arrayIndex should be 0.");
    }
    if (results.params.valueVector.target.stringIndex !== 9) {
      assertions.push("valueVector target.stringIndex should be 6.");
    }
  }

  return assertions;
};

const testValidExplicitString = () => {
  const assertions = [];

  const template = testTextInterpolator`checked="checked"`;
  const vector = create();

  let safety = 0;
  while (incrementTarget(template, vector) && safety < RECURSION_SAFETY) {
    safety += 1;
  }

  const results = crawlForAttribute(template, vector);
  if (results === undefined) {
    assertions.push("this should have returned results");
  }

  if (
    results !== undefined &&
    results.action === "EXPLICIT_ATTRIBUTE_CONFIRMED"
  ) {
    if (results.params.attributeVector.origin.arrayIndex !== 0) {
      assertions.push("attributeVector origin.arrayIndex should be 0.");
    }
    if (results.params.attributeVector.origin.stringIndex !== 0) {
      assertions.push("attributeVector origin.stringIndex should be 0.");
    }

    if (results.params.attributeVector.target.arrayIndex !== 0) {
      assertions.push("attributeVector target.arrayIndex should be 0.");
    }
    if (results.params.attributeVector.target.stringIndex !== 6) {
      assertions.push("attributeVector target.stringIndex should be 6.");
    }

    if (results.params.valueVector.origin.arrayIndex !== 0) {
      assertions.push("valueVector origin.arrayIndex should be 0.");
    }
    if (results.params.valueVector.origin.stringIndex !== 8) {
      assertions.push("valueVector origin.stringIndex should be 0.");
    }

    if (results.params.valueVector.target.arrayIndex !== 0) {
      assertions.push("valueVector target.arrayIndex should be 0.");
    }
    if (results.params.valueVector.target.stringIndex !== 16) {
      assertions.push("valueVector target.stringIndex should be 16.");
    }
  }

  return assertions;
};

const testValidExplicitStringWithTrailingSpaces = () => {
  const assertions = [];

  const template = testTextInterpolator`checked="checked   "`;
  const vector = create();

  let safety = 0;
  while (incrementTarget(template, vector) && safety < RECURSION_SAFETY) {
    safety += 1;
  }

  const results = crawlForAttribute(template, vector);
  if (results === undefined) {
    assertions.push("this should have returned results");
  }

  if (
    results !== undefined &&
    results.action === "EXPLICIT_ATTRIBUTE_CONFIRMED"
  ) {
    if (results.params.attributeVector.origin.arrayIndex !== 0) {
      assertions.push("attributeVector origin.arrayIndex should be 0.");
    }
    if (results.params.attributeVector.origin.stringIndex !== 0) {
      assertions.push("attributeVector origin.stringIndex should be 0.");
    }

    if (results.params.attributeVector.target.arrayIndex !== 0) {
      assertions.push("attributeVector target.arrayIndex should be 0.");
    }
    if (results.params.attributeVector.target.stringIndex !== 6) {
      assertions.push("attributeVector target.stringIndex should be 6.");
    }

    if (results.params.valueVector.origin.arrayIndex !== 0) {
      assertions.push("valueVector origin.arrayIndex should be 0.");
    }
    if (results.params.valueVector.origin.stringIndex !== 8) {
      assertions.push("valueVector origin.stringIndex should be 0.");
    }

    if (results.params.valueVector.target.arrayIndex !== 0) {
      assertions.push("valueVector target.arrayIndex should be 0.");
    }
    if (results.params.valueVector.target.stringIndex !== 19) {
      assertions.push("valueVector target.stringIndex should be 19.");
    }
  }

  return assertions;
};

const testInjectedString = () => {
  const assertions = [];

  const template = testTextInterpolator`checked="${"hello"}"`;
  const vector = create();

  let safety = 0;
  while (incrementTarget(template, vector) && safety < RECURSION_SAFETY) {
    safety += 1;
  }

  const results = crawlForAttribute(template, vector);
  if (results === undefined) {
    assertions.push("this should have returned results");
  }

  if (
    results !== undefined &&
    results.action === "EXPLICIT_ATTRIBUTE_CONFIRMED"
  ) {
    if (results.params.attributeVector.origin.arrayIndex !== 0) {
      assertions.push("attributeVector origin.arrayIndex should be 0.");
    }
    if (results.params.attributeVector.origin.stringIndex !== 0) {
      assertions.push("attributeVector origin.stringIndex should be 0.");
    }

    if (results.params.attributeVector.target.arrayIndex !== 0) {
      assertions.push("attributeVector target.arrayIndex should be 0.");
    }
    if (results.params.attributeVector.target.stringIndex !== 6) {
      assertions.push("attributeVector target.stringIndex should be 6.");
    }

    if (results.params.valueVector.origin.arrayIndex !== 0) {
      assertions.push("valueVector origin.arrayIndex should be 0.");
    }
    if (results.params.valueVector.origin.stringIndex !== 8) {
      assertions.push("valueVector origin.stringIndex should be 0.");
    }

    if (results.params.valueVector.target.arrayIndex !== 0) {
      assertions.push("valueVector target.arrayIndex should be 0.");
    }
    if (results.params.valueVector.target.stringIndex !== 19) {
      assertions.push("valueVector target.stringIndex should be 19.");
    }
  }

  return assertions;
};

const testMalformedInjectedString = () => {
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

const testMalformedInjectedStringWithTrailingSpaces = () => {
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

const testMalformedInjectedStringWithStartingSpaces = () => {
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
  testEmptyString,
  testEmptySpaceString,
  testEmptyMultiSpaceString,
  testImplicitString,
  testImplicitStringWithTrailingSpaces,
  testMalformedExplicitString,
  testAlmostExplicitString,
  testEmptyExplicitString,
  testValidExplicitString,
  testValidExplicitStringWithTrailingSpaces,
  testInjectedString,
  testMalformedInjectedString,
  testMalformedInjectedStringWithTrailingSpaces,
  testMalformedInjectedStringWithStartingSpaces,
  // testSingleCharacterString,
  // testCharaceterString,
];

const unitTestAttributeCrawl = {
  title,
  tests,
  runTestsAsynchronously,
};

export { unitTestAttributeCrawl };
