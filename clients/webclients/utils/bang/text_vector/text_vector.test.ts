import { copy, create, increment, getCharFromTarget } from "./text_vector";
import { Vector } from "../type_flyweight/text_vector";
import { StructureRender } from "../type_flyweight/structure";

type TextTextInterpolator = <A>(
  templateArray: TemplateStringsArray,
  ...injections: A[]
) => StructureRender<A>;

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

  const vector = create();

  if (vector.origin.stringIndex !== 0) {
    assertions.push("text vector string index does not match");
  }
  if (vector.origin.arrayIndex !== 0) {
    assertions.push("text vector array index does not match");
  }

  return assertions;
};

const createTextVectorFromPosition = () => {
  const assertions = [];
  const prevPosition = {
    stringIndex: 3,
    arrayIndex: 4,
  };
  const vector = create(prevPosition);

  if (vector.origin.stringIndex !== 3) {
    assertions.push("text vector string index does not match");
  }
  if (vector.origin.arrayIndex !== 4) {
    assertions.push("text vector array index does not match");
  }

  return assertions;
};

const copyTextVector = () => {
  const assertions = [];
  const vector: Vector = {
    origin: { arrayIndex: 0, stringIndex: 1 },
    target: { arrayIndex: 2, stringIndex: 3 },
  };

  const copiedVector = copy(vector);

  if (vector.origin.stringIndex !== copiedVector.origin.stringIndex) {
    assertions.push("text vector string index does not match");
  }
  if (vector.origin.arrayIndex !== copiedVector.origin.arrayIndex) {
    assertions.push("text vector array index does not match");
  }
  return assertions;
};

const incrementTextVector = () => {
  const assertions = [];
  const structureRender = testTextInterpolator`hello`;
  const vector: Vector = create();
  increment(vector, structureRender);

  if (vector.target.stringIndex !== 1) {
    assertions.push("text vector string index does not match");
  }
  if (vector.target.arrayIndex !== 0) {
    assertions.push("text vector array index does not match");
  }

  return assertions;
};

const incrementMultiTextVector = () => {
  const assertions = [];
  const structureRender = testTextInterpolator`hey${"world"}, how are you?`;
  const vector: Vector = create();
  increment(vector, structureRender);
  increment(vector, structureRender);
  increment(vector, structureRender);
  increment(vector, structureRender);
  increment(vector, structureRender);

  if (vector.target.stringIndex !== 2) {
    assertions.push("text vector string index does not match");
  }
  if (vector.target.arrayIndex !== 1) {
    assertions.push("text vector array index does not match");
  }

  return assertions;
};

const incrementTextVectorTooFar = () => {
  const assertions = [];
  const structureRender = testTextInterpolator`hey${"world"}, how are you?`;
  const vector: Vector = create();

  const MAX_DEPTH = 20;
  let safety = 0;
  while (increment(vector, structureRender) && safety < MAX_DEPTH) {
    // iterate across structure
    safety += 1;
  }
  console.log("increment text vector too far");
  console.log(vector);
  if (vector.target.stringIndex !== 13) {
    assertions.push("text vector string index does not match");
  }
  if (vector.target.arrayIndex !== 1) {
    assertions.push("text vector array index does not match");
  }

  return assertions;
};

const getCharFromTemplate = () => {
  const assertions = [];
  const structureRender = testTextInterpolator`hello`;
  const vector: Vector = {
    origin: { arrayIndex: 0, stringIndex: 0 },
    target: { arrayIndex: 0, stringIndex: 2 },
  };

  const char = getCharFromTarget(vector, structureRender);

  if (char !== "l") {
    assertions.push("textVector target is not 'l'");
  }

  return assertions;
};

const tests = [
  createTextVector,
  createTextVectorFromPosition,
  copyTextVector,
  incrementTextVector,
  incrementMultiTextVector,
  incrementTextVectorTooFar,
  getCharFromTemplate,
];

const unitTestTextVector = {
  title,
  tests,
  runTestsAsynchronously,
};

export { unitTestTextVector };
