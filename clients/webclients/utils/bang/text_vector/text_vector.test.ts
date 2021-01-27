// brian taylor vann
// text vector

import { samestuff } from "../../little_test_runner/samestuff/samestuff";
import {
  copy,
  create,
  createFollowingVector,
  incrementTarget,
  hasOriginEclipsedTaraget,
  getText,
} from "./text_vector";
import { Vector } from "../type_flyweight/text_vector";
import { Template } from "../type_flyweight/template";

type TextTextInterpolator = <A>(
  templateArray: TemplateStringsArray,
  ...injections: A[]
) => Template<A>;

const testTextInterpolator: TextTextInterpolator = (
  templateArray,
  ...injections
) => {
  return { templateArray, injections };
};

const title = "text_vector";
const runTestsAsynchronously = true;

const createTextVector = () => {
  const assertions = [];

  const expectedResults = {
    origin: { arrayIndex: 0, stringIndex: 0 },
    target: { arrayIndex: 0, stringIndex: 0 },
  };

  const vector = create();

  if (!samestuff(expectedResults, vector)) {
    assertions.push("unexpected results found.");
  }

  return assertions;
};

const createTextVectorFromPosition = () => {
  const assertions = [];

  const expectedResults = {
    origin: { arrayIndex: 4, stringIndex: 3 },
    target: { arrayIndex: 4, stringIndex: 3 },
  };

  const vector = create({
    stringIndex: 3,
    arrayIndex: 4,
  });

  if (!samestuff(expectedResults, vector)) {
    assertions.push("unexpected results found.");
  }

  return assertions;
};

const copyTextVector = () => {
  const assertions = [];

  const expectedResults = {
    origin: { arrayIndex: 0, stringIndex: 1 },
    target: { arrayIndex: 2, stringIndex: 3 },
  };

  const copiedVector = copy(expectedResults);
  if (!samestuff(expectedResults, copiedVector)) {
    assertions.push("unexpected results found.");
  }

  return assertions;
};

const incrementTextVector = () => {
  const assertions = [];

  const expectedResults = {
    origin: { arrayIndex: 0, stringIndex: 0 },
    target: { arrayIndex: 0, stringIndex: 1 },
  };

  const structureRender = testTextInterpolator`hello`;
  const vector: Vector = create();

  incrementTarget(structureRender, vector);

  if (!samestuff(expectedResults, vector)) {
    assertions.push("unexpected results found.");
  }

  return assertions;
};

const incrementMultiTextVector = () => {
  const assertions = [];

  const expectedResults = {
    origin: { arrayIndex: 0, stringIndex: 0 },
    target: { arrayIndex: 1, stringIndex: 2 },
  };

  const structureRender = testTextInterpolator`hey${"world"}, how are you?`;
  const vector: Vector = create();

  incrementTarget(structureRender, vector);
  incrementTarget(structureRender, vector);
  incrementTarget(structureRender, vector);
  incrementTarget(structureRender, vector);
  incrementTarget(structureRender, vector);

  if (!samestuff(expectedResults, vector)) {
    assertions.push("unexpected results found.");
  }
  if (vector.target.stringIndex !== 2) {
    assertions.push("text vector string index does not match");
  }
  if (vector.target.arrayIndex !== 1) {
    assertions.push("text vector array index does not match");
  }

  return assertions;
};

const incrementEmptyTextVector = () => {
  const assertions = [];

  const expectedResults = {
    origin: { arrayIndex: 0, stringIndex: 0 },
    target: { arrayIndex: 3, stringIndex: 0 },
  };

  const structureRender = testTextInterpolator`${"hey"}${"world"}${"!!"}`;
  const vector: Vector = create();

  incrementTarget(structureRender, vector);
  incrementTarget(structureRender, vector);
  incrementTarget(structureRender, vector);

  if (incrementTarget(structureRender, vector) !== undefined) {
    assertions.push("should not return after traversed");
  }

  if (!samestuff(expectedResults, vector)) {
    assertions.push("unexpected results found.");
  }

  return assertions;
};

const createFollowingTextVector = () => {
  const assertions = [];

  const expectedResults = {
    origin: { arrayIndex: 0, stringIndex: 5 },
    target: { arrayIndex: 0, stringIndex: 5 },
  };

  const structureRender = testTextInterpolator`supercool`;
  const vector: Vector = create();

  incrementTarget(structureRender, vector);
  incrementTarget(structureRender, vector);
  incrementTarget(structureRender, vector);
  incrementTarget(structureRender, vector);

  const results = createFollowingVector(structureRender, vector);

  if (!samestuff(expectedResults, results)) {
    assertions.push("unexpected results found.");
  }

  return assertions;
};

const incrementTextVectorTooFar = () => {
  const assertions = [];

  const expectedResults = {
    origin: { arrayIndex: 0, stringIndex: 0 },
    target: { arrayIndex: 1, stringIndex: 13 },
  };

  const structureRender = testTextInterpolator`hey${"world"}, how are you?`;
  const results: Vector = create();

  const MAX_DEPTH = 20;
  let safety = 0;
  while (incrementTarget(structureRender, results) && safety < MAX_DEPTH) {
    // iterate across structure
    safety += 1;
  }

  if (!samestuff(expectedResults, results)) {
    assertions.push("unexpected results found.");
  }

  return assertions;
};

const testHasOriginEclipsedTaraget = () => {
  const assertions = [];

  const vector: Vector = create();
  const results = hasOriginEclipsedTaraget(vector);

  if (results !== true) {
    assertions.push("orign eclipsed target");
  }

  return assertions;
};

const testHasOriginNotEclipsedTaraget = () => {
  const assertions = [];

  const structureRender = testTextInterpolator`hey${"world"}, how are you?`;
  const vector: Vector = create();

  incrementTarget(structureRender, vector);
  incrementTarget(structureRender, vector);
  incrementTarget(structureRender, vector);
  incrementTarget(structureRender, vector);

  const results = hasOriginEclipsedTaraget(vector);

  if (results !== false) {
    assertions.push("orign has not eclipsed target");
  }

  return assertions;
};

const testGetTextReturnsActualText = () => {
  const expectedResult = "world";
  const assertions = [];

  const structureRender = testTextInterpolator`hey world, how are you?`;
  const vector: Vector = {
    origin: {
      arrayIndex: 0,
      stringIndex: 4,
    },
    target: {
      arrayIndex: 0,
      stringIndex: 8,
    },
  };

  const results = getText(structureRender, vector);

  if (expectedResult !== results) {
    assertions.push("text should say 'world'");
  }

  return assertions;
};

const testGetTextOverTemplate = () => {
  const expectedResult = "how  you";
  const assertions = [];

  const structureRender = testTextInterpolator`hey ${"world"}, how ${"are"} you?`;
  const vector: Vector = {
    origin: {
      arrayIndex: 1,
      stringIndex: 2,
    },
    target: {
      arrayIndex: 2,
      stringIndex: 3,
    },
  };

  const results = getText(structureRender, vector);

  if (expectedResult !== results) {
    assertions.push("text should say 'world'");
  }

  return assertions;
};

const testGetTextOverChonkyTemplate = () => {
  const expectedResult = "how  you  buster";
  const assertions = [];

  const structureRender = testTextInterpolator`hey ${"world"}, how ${"are"} you ${"doing"} buster?`;
  const vector: Vector = {
    origin: {
      arrayIndex: 1,
      stringIndex: 2,
    },
    target: {
      arrayIndex: 3,
      stringIndex: 6,
    },
  };

  const results = getText(structureRender, vector);

  if (expectedResult !== results) {
    assertions.push("text should say 'world'");
  }

  return assertions;
};

const tests = [
  createTextVector,
  createTextVectorFromPosition,
  createFollowingTextVector,
  copyTextVector,
  incrementTextVector,
  incrementMultiTextVector,
  incrementEmptyTextVector,
  incrementTextVectorTooFar,
  testHasOriginEclipsedTaraget,
  testHasOriginNotEclipsedTaraget,
  testGetTextReturnsActualText,
  testGetTextOverTemplate,
  testGetTextOverChonkyTemplate,
];

const unitTestTextVector = {
  title,
  tests,
  runTestsAsynchronously,
};

export { unitTestTextVector };
