// brian taylor vann
// tag name crawl

import { samestuff } from "../../../little_test_runner/samestuff/samestuff";
import { StructureRender } from "../../type_flyweight/structure";
import { Vector } from "../../type_flyweight/text_vector";

import {
  copy,
  decrementTarget,
  hasOriginEclipsedTaraget,
  incrementOrigin,
} from "../../text_vector/text_vector";

import { getCharAtPosition } from "../../text_position/text_position";

const crawlForTagName = <A>(
  template: StructureRender<A>,
  innerXmlBounds: Vector
) => {
  const tagVector: Vector = copy(innerXmlBounds);
  let positionChar = getCharAtPosition(template, tagVector.origin);
  if (positionChar === undefined || positionChar === " ") {
    return;
  }

  while (positionChar !== " " && !hasOriginEclipsedTaraget(tagVector)) {
    if (incrementOrigin(template, tagVector) === undefined) {
      return;
    }

    positionChar = getCharAtPosition(template, tagVector.origin);
    if (positionChar === undefined) {
      return;
    }
  }

  const adjustedVector: Vector = {
    origin: { ...innerXmlBounds.origin },
    target: { ...tagVector.origin },
  };

  // walk back a step if successive space found
  if (positionChar === " ") {
    decrementTarget(template, adjustedVector);
  }

  return adjustedVector;
};

export { crawlForTagName };
