import {
  copy,
  create,
  decrement,
  increment,
  getCharAtPosition,
} from "./text_position";
import { Position } from "../type_flyweight/text_vector";
import { StructureRender } from "../type_flyweight/structure";

type TestTextInterpolator = <A>(
  templateArray: TemplateStringsArray,
  ...injections: A[]
) => StructureRender<A>;

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
  const vector = create();

  if (vector.stringIndex !== 0) {
    assertions.push("text position string index does not match");
  }
  if (vector.arrayIndex !== 0) {
    assertions.push("text position array index does not match");
  }

  return assertions;
};

const createTextPositionFromPosition = () => {
  const assertions = [];
  const prevPosition = {
    stringIndex: 3,
    arrayIndex: 4,
  };
  const vector = create(prevPosition);

  if (vector.stringIndex !== 3) {
    assertions.push("text position string index does not match");
  }
  if (vector.arrayIndex !== 4) {
    assertions.push("text position array index does not match");
  }

  return assertions;
};

const copyTextPosition = () => {
  const assertions = [];
  const position: Position = { arrayIndex: 2, stringIndex: 3 };

  const copiedPosition = copy(position);

  if (position.stringIndex !== copiedPosition.stringIndex) {
    assertions.push("text position string index does not match");
  }
  if (position.arrayIndex !== copiedPosition.arrayIndex) {
    assertions.push("text position array index does not match");
  }
  return assertions;
};

const incrementTextPosition = () => {
  const assertions = [];
  const structureRender = testTextInterpolator`hello`;
  const position: Position = create();

  increment(structureRender, position);

  if (position.stringIndex !== 1) {
    assertions.push("text position string index does not match");
  }
  if (position.arrayIndex !== 0) {
    assertions.push("text position array index does not match");
  }

  return assertions;
};

const incrementMultiTextPosition = () => {
  const assertions = [];
  const structureRender = testTextInterpolator`hey${"world"}, how are you?`;
  const position: Position = create();

  increment(structureRender, position);
  increment(structureRender, position);
  increment(structureRender, position);
  increment(structureRender, position);
  increment(structureRender, position);

  if (position.stringIndex !== 2) {
    assertions.push("text position string index does not match");
  }
  if (position.arrayIndex !== 1) {
    assertions.push("text position array index does not match");
  }

  return assertions;
};

const incrementEmptyTextPosition = () => {
  const assertions = [];
  const structureRender = testTextInterpolator`${"hey"}${"world"}${"!!"}`;
  const position: Position = create();

  increment(structureRender, position);
  increment(structureRender, position);
  increment(structureRender, position);

  if (increment(structureRender, position) !== undefined) {
    assertions.push("should not return after traversed");
  }

  if (position.stringIndex !== 0) {
    assertions.push("text position string index does not match");
  }
  if (position.arrayIndex !== 3) {
    assertions.push("text position array index does not match");
  }

  return assertions;
};

const incrementTextPositionTooFar = () => {
  const assertions = [];
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

  if (position.stringIndex !== 13) {
    assertions.push("text position string index does not match");
  }
  if (position.arrayIndex !== 1) {
    assertions.push("text position array index does not match");
  }

  return assertions;
};

const decrementTextPosition = () => {
  const assertions = [];
  const structureRender = testTextInterpolator`hello`;
  const arrayLength = structureRender.templateArray.length - 1;
  const stringLength = structureRender.templateArray[arrayLength].length - 1;
  const position: Position = copy({
    arrayIndex: arrayLength,
    stringIndex: stringLength,
  });

  decrement(structureRender, position);
  if (position.stringIndex !== 3) {
    assertions.push("text position string index does not match");
  }
  if (position.arrayIndex !== 0) {
    assertions.push("text position array index does not match");
  }

  return assertions;
};

const decrementMultiTextPosition = () => {
  const assertions = [];
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

  if (position.stringIndex !== 1) {
    assertions.push("text position string index does not match");
  }
  if (position.arrayIndex !== 0) {
    assertions.push("text position array index does not match");
  }

  return assertions;
};

const decrementEmptyTextPosition = () => {
  const assertions = [];
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

  if (position.stringIndex !== 0) {
    assertions.push("text position string index does not match");
  }
  if (position.arrayIndex !== 0) {
    assertions.push("text position array index does not match");
  }

  return assertions;
};

const decrementTextPositionTooFar = () => {
  const assertions = [];
  const structureRender = testTextInterpolator`hey${"world"}, how are you?`;
  const position: Position = create();

  const MAX_DEPTH = 20;
  let safety = 0;
  while (decrement(structureRender, position) && safety < MAX_DEPTH) {
    // iterate across structure
    safety += 1;
  }

  if (position.stringIndex !== 0) {
    assertions.push("text position string index does not match");
  }
  if (position.arrayIndex !== 0) {
    assertions.push("text position array index does not match");
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
