import { StructureRender } from "../type_flyweight/structure";
import { Vector } from "../type_flyweight/text_vector";

type Copy = (vector: Vector) => Vector;
type GetTagetChar = <A>(
  vector: Vector,
  tempalte: StructureRender<A>
) => string | void;

type Increment = <A>(vector: Vector, tempalte: StructureRender<A>) => void;

const copy: Copy = (vector) => {
  return {
    origin: { ...vector.origin },
    target: { ...vector.target },
  };
};

const increment: Increment = <A>(
  vector: Vector,
  template: StructureRender<A>
) => {
  const target = vector.target;
  const arrayLength = template.templateArray[target.arrayIndex].length;
  const templateLength = template.templateArray.length;

  if (
    target.arrayIndex >= templateLength - 1 &&
    target.stringIndex >= arrayLength - 1
  ) {
    return;
  }

  target.stringIndex += 1;
  target.stringIndex %= arrayLength;
  if (target.stringIndex === 0 && target.arrayIndex < templateLength - 1) {
    target.arrayIndex += 1;
  }
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

export { copy, increment, getCharFromTarget };
