import {
  copy,
  create,
  createFollowingVector,
  incrementTarget,
} from "./text_vector";
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

  incrementTarget(structureRender, vector);

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

  incrementTarget(structureRender, vector);
  incrementTarget(structureRender, vector);
  incrementTarget(structureRender, vector);
  incrementTarget(structureRender, vector);
  incrementTarget(structureRender, vector);

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
  const structureRender = testTextInterpolator`${"hey"}${"world"}${"!!"}`;
  const vector: Vector = create();

  incrementTarget(structureRender, vector);
  incrementTarget(structureRender, vector);
  incrementTarget(structureRender, vector);

  if (incrementTarget(structureRender, vector) !== undefined) {
    assertions.push("should not return after traversed");
  }

  if (vector.target.stringIndex !== 0) {
    assertions.push("text vector string index does not match");
  }
  if (vector.target.arrayIndex !== 3) {
    assertions.push("text vector array index does not match");
  }

  return assertions;
};

const createFollowingTextVector = () => {
  const assertions = [];
  const structureRender = testTextInterpolator`supercool`;
  const vector: Vector = create();

  incrementTarget(structureRender, vector);
  incrementTarget(structureRender, vector);
  incrementTarget(structureRender, vector);
  incrementTarget(structureRender, vector);

  const followingVector = createFollowingVector(structureRender, vector);
  if (followingVector === undefined) {
    assertions.push("next vector should not be undefined");
  }
  if (followingVector && followingVector.target.stringIndex !== 5) {
    assertions.push("text vector string index does not match");
  }
  if (followingVector && followingVector.target.arrayIndex !== 0) {
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
  while (incrementTarget(structureRender, vector) && safety < MAX_DEPTH) {
    // iterate across structure
    safety += 1;
  }

  if (vector.target.stringIndex !== 13) {
    assertions.push("text vector string index does not match");
  }
  if (vector.target.arrayIndex !== 1) {
    assertions.push("text vector array index does not match");
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
];

const unitTestTextVector = {
  title,
  tests,
  runTestsAsynchronously,
};

export { unitTestTextVector };
