import { StructureRender } from "../type_flyweight/structure";
import { Position } from "../type_flyweight/text_vector";

type Create = (position?: Position) => Position;
type Copy = (position: Position) => Position;

type Increment = <A>(
  template: StructureRender<A>,
  position: Position
) => Position | undefined;

type GetTargetChar = <A>(
  template: StructureRender<A>,
  position: Position
) => string | undefined;

const DEFAULT_POSITION: Position = {
  arrayIndex: 0,
  stringIndex: 0,
};

const create: Create = (position = DEFAULT_POSITION) => ({ ...position });

const copy: Copy = create;

const increment: Increment = (template, position) => {
  // template boundaries
  const templateLength = template.templateArray.length;
  const chunkLength = template.templateArray[position.arrayIndex].length;
  if (chunkLength === undefined) {
    return;
  }

  const arrayIndex = position.arrayIndex;
  const stringIndex = position.stringIndex;
  if (arrayIndex >= templateLength - 1 && chunkLength === 0) {
    return;
  }
  if (arrayIndex >= templateLength - 1 && stringIndex >= chunkLength - 1) {
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

const decrement: Increment = (template, position) => {
  const templateLength = template.templateArray.length - 1;
  if (position.arrayIndex > templateLength) {
    return;
  }

  if (position.arrayIndex <= 0 && position.stringIndex <= 0) {
    return;
  }

  position.stringIndex -= 1;
  if (position.arrayIndex >= 0 && position.stringIndex < 0) {
    position.arrayIndex -= 1;

    const chunk = template.templateArray[position.arrayIndex];
    position.stringIndex = chunk.length - 1;

    // undefined case akin to divide by zero
    if (chunk === "") {
      position.stringIndex = chunk.length;
    }
  }

  return position;
};

const getCharFromTarget: GetTargetChar = (template, position) => {
  const templateArray = template.templateArray;

  return templateArray?.[position.arrayIndex]?.[position.stringIndex];
};

export { copy, create, decrement, increment, getCharFromTarget };
