import { StructureRender } from "../type_flyweight/structure";
import { Position, Vector } from "../type_flyweight/text_vector";

type Create = (position?: Position) => Vector;
type CreateFollowingVector = <A>(
  vector: Vector,
  tempalte: StructureRender<A>
) => Vector;

type Copy = (vector: Vector) => Vector;
type GetTagetChar = <A>(
  vector: Vector,
  tempalte: StructureRender<A>
) => string | undefined;

type Increment = <A>(
  position: Position,
  tempalte: StructureRender<A>
) => Position | undefined;

const DEFAULT_POSITION: Position = {
  arrayIndex: 0,
  stringIndex: 0,
};

const create: Create = (position = DEFAULT_POSITION) => ({
  origin: { ...position },
  target: { ...position },
});

const createFollowingVector: CreateFollowingVector = (vector, template) => {
  const followingVector = copy(vector);
  increment(followingVector.target, template);
  followingVector.origin = { ...followingVector.target };

  return followingVector;
};

const copy: Copy = (vector) => {
  return {
    origin: { ...vector.origin },
    target: { ...vector.target },
  };
};

const increment: Increment = <A>(
  position: Position,
  template: StructureRender<A>
) => {
  // template boundaries
  const templateLength = template.templateArray.length;
  const chunkLength = template.templateArray[position.arrayIndex].length;
  if (chunkLength === undefined) {
    return;
  }

  // determine if finished
  if (position.arrayIndex >= templateLength - 1 && chunkLength === 0) {
    return;
  }
  if (
    position.arrayIndex >= templateLength - 1 &&
    position.stringIndex >= chunkLength - 1
  ) {
    return;
  }

  // cannot % modulo by 0
  if (chunkLength > 0) {
    position.stringIndex += 1;
    position.stringIndex %= chunkLength;
  }

  if (position.stringIndex === 0) {
    position.arrayIndex += 1;
  }

  return position;
};

// needs to be tested
const decrement: Increment = <A>(
  position: Position,
  template: StructureRender<A>
) => {
  const templateLength = template.templateArray.length;
  if (position.arrayIndex < 0 || position.arrayIndex >= templateLength - 1) {
    return;
  }
  const chunkLength = template.templateArray[position.arrayIndex].length;

  if (position.arrayIndex < 0) {
    return;
  }

  position.stringIndex -= 1;
  if (position.stringIndex < 0 && position.arrayIndex > 0) {
    position.arrayIndex -= 1;
    position.stringIndex =
      template.templateArray[position.arrayIndex].length - 1;
  }

  return position;
};

const getCharFromTarget: GetTagetChar = (vector, template) => {
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
  decrement,
  increment,
  getCharFromTarget,
};
