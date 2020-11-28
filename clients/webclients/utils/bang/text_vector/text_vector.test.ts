import { copy, increment, getCharFromTarget } from "./text_vector";
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
  const vector: Vector = {
    origin: { arrayIndex: 0, stringIndex: 0 },
    target: { arrayIndex: 0, stringIndex: 0 },
  };
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
  const vector: Vector = {
    origin: { arrayIndex: 0, stringIndex: 0 },
    target: { arrayIndex: 0, stringIndex: 0 },
  };
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
  const structureRender = testTextInterpolator`hello`;
  const vector: Vector = {
    origin: { arrayIndex: 0, stringIndex: 0 },
    target: { arrayIndex: 0, stringIndex: 0 },
  };

  increment(vector, structureRender);
  increment(vector, structureRender);
  increment(vector, structureRender);
  increment(vector, structureRender);
  increment(vector, structureRender);
  increment(vector, structureRender);

  if (vector.target.stringIndex !== 4) {
    assertions.push("text vector string index does not match");
  }
  if (vector.target.arrayIndex !== 0) {
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
