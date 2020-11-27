import { StructureRender } from "../type_flyweight/structure";
import { Vector } from "../type_flyweight/text_vector";

type Copy = (vector: Vector) => Vector;
type Increment = <A>(vector: Vector, tempalte: StructureRender<A>) => void;

const copy: Copy = (original) => {
  return {
    origin: { ...original.origin },
    target: { ...original.target },
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

export { copy, increment };
