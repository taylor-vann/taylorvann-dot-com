import { samestuff } from "../../little_test_runner/samestuff/samestuff";
import {
  copy,
  create,
  decrement,
  increment,
  getCharAtPosition,
} from "./text_position";
import { Position } from "../type_flyweight/text_vector";
import { Template } from "../type_flyweight/template";

type TestTextInterpolator = <A>(
  templateArray: TemplateStringsArray,
  ...injections: A[]
) => Template<A>;

const testTextInterpolator: TestTextInterpolator = (
  templateArray,
  ...injections
) => {
  return { templateArray, injections };
};

const title = "text_position";
const runTestsAsynchronously = true;

const createTextPosition = () => {
  const assertions = [];

  const expectedResults = {
    arrayIndex: 0,
    stringIndex: 0,
  };

  const position = create();

  if (!samestuff(expectedResults, position)) {
    assertions.push("unexpected results found.");
  }

  return assertions;
};

const createTextPositionFromPosition = () => {
  const assertions = [];

  const expectedResults = {
    arrayIndex: 3,
    stringIndex: 4,
  };

  const position = create(expectedResults);

  if (!samestuff(expectedResults, position)) {
    assertions.push("unexpected results found.");
  }

  return assertions;
};

const copyTextPosition = () => {
  const assertions = [];

  const expectedResults = {
    arrayIndex: 2,
    stringIndex: 3,
  };

  const position = copy(expectedResults);

  if (!samestuff(expectedResults, position)) {
    assertions.push("unexpected results found.");
  }

  return assertions;
};

const incrementTextPosition = () => {
  const assertions = [];

  const expectedResults = {
    arrayIndex: 0,
    stringIndex: 1,
  };

  const structureRender = testTextInterpolator`hello`;
  const position: Position = create();

  increment(structureRender, position);

  if (!samestuff(expectedResults, position)) {
    assertions.push("unexpected results found.");
  }

  return assertions;
};

const incrementMultiTextPosition = () => {
  const assertions = [];

  const expectedResults = {
    arrayIndex: 1,
    stringIndex: 2,
  };

  const structureRender = testTextInterpolator`hey${"world"}, how are you?`;
  const position: Position = create();

  increment(structureRender, position);
  increment(structureRender, position);
  increment(structureRender, position);
  increment(structureRender, position);
  increment(structureRender, position);

  if (!samestuff(expectedResults, position)) {
    assertions.push("unexpected results found.");
  }

  return assertions;
};

const incrementEmptyTextPosition = () => {
  const assertions = [];

  const expectedResults = {
    arrayIndex: 3,
    stringIndex: 0,
  };

  const structureRender = testTextInterpolator`${"hey"}${"world"}${"!!"}`;
  const position: Position = create();

  increment(structureRender, position);
  increment(structureRender, position);
  increment(structureRender, position);

  if (increment(structureRender, position) !== undefined) {
    assertions.push("should not return after traversed");
  }

  if (!samestuff(expectedResults, position)) {
    assertions.push("unexpected results found.");
  }

  return assertions;
};

const incrementTextPositionTooFar = () => {
  const assertions = [];

  const expectedResults = {
    arrayIndex: 1,
    stringIndex: 13,
  };

  const structureRender = testTextInterpolator`hey${"world"}, how are you?`;
  const arrayLength = structureRender.templateArray.length - 1;
  const stringLength = structureRender.templateArray[arrayLength].length - 1;
  const position: Position = copy({
    arrayIndex: arrayLength,
    stringIndex: stringLength,
  });

  const MAX_DEPTH = 20;
  let safety = 0;
  while (increment(structureRender, position) && safety < MAX_DEPTH) {
    // iterate across structure
    safety += 1;
  }

  if (!samestuff(expectedResults, position)) {
    assertions.push("unexpected results found.");
  }

  return assertions;
};

const decrementTextPosition = () => {
  const assertions = [];

  const expectedResults = {
    arrayIndex: 0,
    stringIndex: 3,
  };

  const structureRender = testTextInterpolator`hello`;
  const arrayLength = structureRender.templateArray.length - 1;
  const stringLength = structureRender.templateArray[arrayLength].length - 1;
  const position: Position = copy({
    arrayIndex: arrayLength,
    stringIndex: stringLength,
  });

  decrement(structureRender, position);

  if (!samestuff(expectedResults, position)) {
    assertions.push("unexpected results found.");
  }

  return assertions;
};

const decrementMultiTextPosition = () => {
  const assertions = [];

  const expectedResults = {
    arrayIndex: 0,
    stringIndex: 1,
  };

  const structureRender = testTextInterpolator`hey${"hello"}bro!`;
  const arrayLength = structureRender.templateArray.length - 1;
  const stringLength = structureRender.templateArray[arrayLength].length - 1;
  const position: Position = copy({
    arrayIndex: arrayLength,
    stringIndex: stringLength,
  });
  decrement(structureRender, position);
  decrement(structureRender, position);
  decrement(structureRender, position);
  decrement(structureRender, position);
  decrement(structureRender, position);

  if (!samestuff(expectedResults, position)) {
    assertions.push("unexpected results found.");
  }

  return assertions;
};

const decrementEmptyTextPosition = () => {
  const assertions = [];

  const expectedResults = {
    arrayIndex: 0,
    stringIndex: 0,
  };

  const structureRender = testTextInterpolator`${"hey"}${"world"}${"!!"}`;
  const arrayLength = structureRender.templateArray.length - 1;
  const stringLength = structureRender.templateArray[arrayLength].length - 1;
  const position: Position = copy({
    arrayIndex: arrayLength,
    stringIndex: stringLength,
  });

  decrement(structureRender, position);
  decrement(structureRender, position);
  decrement(structureRender, position);

  if (decrement(structureRender, position) !== undefined) {
    assertions.push("should not return after traversed");
  }

  if (!samestuff(expectedResults, position)) {
    assertions.push("unexpected results found.");
  }

  return assertions;
};

const decrementTextPositionTooFar = () => {
  const assertions = [];

  const expectedResults = {
    arrayIndex: 0,
    stringIndex: 0,
  };

  const structureRender = testTextInterpolator`hey${"world"}, how are you?`;
  const position: Position = create();

  const MAX_DEPTH = 20;
  let safety = 0;
  while (decrement(structureRender, position) && safety < MAX_DEPTH) {
    // iterate across structure
    safety += 1;
  }

  if (!samestuff(expectedResults, position)) {
    assertions.push("unexpected results found.");
  }

  return assertions;
};

const getCharFromTemplate = () => {
  const assertions = [];
  const structureRender = testTextInterpolator`hello`;
  const position: Position = { arrayIndex: 0, stringIndex: 2 };
  const char = getCharAtPosition(structureRender, position);

  if (char !== "l") {
    assertions.push("textPosition target is not 'l'");
  }

  return assertions;
};

const tests = [
  createTextPosition,
  createTextPositionFromPosition,
  copyTextPosition,
  incrementTextPosition,
  incrementMultiTextPosition,
  incrementEmptyTextPosition,
  incrementTextPositionTooFar,
  decrementTextPosition,
  decrementMultiTextPosition,
  decrementEmptyTextPosition,
  decrementTextPositionTooFar,
  getCharFromTemplate,
];

const unitTestTextPosition = {
  title,
  tests,
  runTestsAsynchronously,
};

export { unitTestTextPosition };
