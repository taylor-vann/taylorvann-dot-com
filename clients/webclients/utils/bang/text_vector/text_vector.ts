import { StructureRender } from "../type_flyweight/structure";
import { Position, Vector } from "../type_flyweight/text_vector";
import { increment, decrement } from "../text_position/text_position";

type Create = (position?: Position) => Vector;
type CreateFollowingVector = <A>(
  template: StructureRender<A>,
  vector: Vector
) => Vector | undefined;

type Copy = (vector: Vector) => Vector;
type GetTagetChar = <A>(
  template: StructureRender<A>,
  vector: Vector
) => string | undefined;

type Increment = <A>(
  template: StructureRender<A>,
  vector: Vector
) => Vector | undefined;

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
    followingVector.origin = { ...followingVector.target };
    return followingVector;
  }

  return;
};

const copy: Copy = (vector) => {
  return {
    origin: { ...vector.origin },
    target: { ...vector.target },
  };
};

const incrementOrigin: Increment = (template, vector) => {
  if (increment(template, vector.origin)) {
    return vector;
  }
  return;
};

const incrementTarget: Increment = (template, vector) => {
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

const decrementTarget: Increment = (template, vector) => {
  if (decrement(template, vector.origin)) {
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

export {
  copy,
  create,
  createFollowingVector,
  decrementOrigin,
  decrementTarget,
  incrementOrigin,
  incrementTarget,
  getTextFromTarget,
};
