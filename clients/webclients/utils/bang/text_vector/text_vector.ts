import { Template } from "../type_flyweight/template";
import { Position, Vector } from "../type_flyweight/text_vector";
import {
  copy as copyPosition,
  increment,
  decrement,
} from "../text_position/text_position";

type Create = (position?: Position) => Vector;

type CreateFollowingVector = <A>(
  template: Template<A>,
  vector: Vector
) => Vector | undefined;

type Copy = (vector: Vector) => Vector;

type GetTagetChar = <A>(
  template: Template<A>,
  vector: Vector
) => string | undefined;

type Increment = <A>(
  template: Template<A>,
  vector: Vector
) => Vector | undefined;

type HasOriginEclipsedTaraget = (vector: Vector) => boolean;

type GetTextFromVector = <A>(
  template: Template<A>,
  vector: Vector
) => string | void;

const DEFAULT_POSITION: Position = {
  arrayIndex: 0,
  stringIndex: 0,
};

const create: Create = (position = DEFAULT_POSITION) => ({
  origin: { ...position },
  target: { ...position },
});

const createFollowingVector: CreateFollowingVector = (template, vector) => {
  const followingVector = copy(vector);
  if (increment(template, followingVector.target)) {
    followingVector.origin = copyPosition(followingVector.target);
    return followingVector;
  }

  return;
};

const copy: Copy = (vector) => {
  return {
    origin: copyPosition(vector.origin),
    target: copyPosition(vector.target),
  };
};

const incrementOrigin: Increment = (template, vector) => {
  if (increment(template, vector.origin)) {
    return vector;
  }
  return;
};

const decrementOrigin: Increment = (template, vector) => {
  if (decrement(template, vector.origin)) {
    return vector;
  }
  return;
};

const incrementTarget: Increment = (template, vector) => {
  if (increment(template, vector.target)) {
    return vector;
  }
  return;
};

const decrementTarget: Increment = (template, vector) => {
  if (decrement(template, vector.target)) {
    return vector;
  }
  return;
};

const getTextFromTarget: GetTagetChar = (template, vector) => {
  const templateArray = template.templateArray;
  const arrayIndex = vector.target.arrayIndex;
  const stringIndex = vector.target.stringIndex;

  if (arrayIndex > templateArray.length - 1) {
    return;
  }

  if (stringIndex > templateArray[arrayIndex].length - 1) {
    return;
  }

  return templateArray[arrayIndex][stringIndex];
};

const hasOriginEclipsedTaraget: HasOriginEclipsedTaraget = (vector) => {
  if (
    vector.origin.arrayIndex >= vector.target.arrayIndex &&
    vector.origin.stringIndex >= vector.target.stringIndex
  ) {
    return true;
  }

  return false;
};

const getText: GetTextFromVector = (template, vector) => {
  if (hasOriginEclipsedTaraget(vector)) {
    return;
  }

  // edge case, only one array length
  if (vector.target.arrayIndex === vector.origin.arrayIndex) {
    const distance = vector.target.stringIndex - vector.origin.stringIndex + 1;
    const templateText = template.templateArray[vector.origin.arrayIndex];
    const copiedText = templateText.substr(vector.origin.stringIndex, distance);

    return copiedText;
  }

  // otherwise, stack and arrayy
  let texts: string[] = [];

  // get head text
  let templateText = template.templateArray[vector.origin.arrayIndex];
  if (templateText === undefined) {
    return;
  }
  let templateTextIndex = vector.origin.stringIndex;
  let distance = templateText.length - templateTextIndex;
  let copiedText = templateText.substr(templateTextIndex, distance);
  texts.push(copiedText);

  // get in between
  let tail = vector.origin.arrayIndex + 1;
  while (tail < vector.target.arrayIndex) {
    texts.push(template.templateArray[tail]);
    tail += 1;
  }

  // get tail text
  templateText = template.templateArray[vector.target.arrayIndex];
  if (templateText === undefined) {
    return;
  }
  distance = vector.target.stringIndex + 1;
  copiedText = templateText.substr(0, distance);
  texts.push(copiedText);

  return texts.join("");
};

export {
  Vector,
  copy,
  create,
  createFollowingVector,
  decrementOrigin,
  decrementTarget,
  incrementOrigin,
  incrementTarget,
  getText,
  getTextFromTarget,
  hasOriginEclipsedTaraget,
};
