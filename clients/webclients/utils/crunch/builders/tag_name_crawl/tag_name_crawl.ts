// brian taylor vann
// tag name crawl

import { Template } from "../../type_flyweight/template";
import { Vector } from "../../type_flyweight/text_vector";

import {
  copy,
  decrementTarget,
  hasOriginEclipsedTaraget,
  incrementOrigin,
} from "../../text_vector/text_vector";

import { getCharAtPosition } from "../../text_position/text_position";

type BreakRunes = Record<string, boolean>;

const BREAK_RUNES: BreakRunes = {
  " ": true,
  "\n": true,
};

const crawlForTagName = <N, A>(
  template: Template<N, A>,
  innerXmlBounds: Vector
) => {
  const tagVector: Vector = copy(innerXmlBounds);
  let positionChar = getCharAtPosition(template, tagVector.origin);
  if (positionChar === undefined || BREAK_RUNES[positionChar]) {
    return;
  }

  while (
    BREAK_RUNES[positionChar] !== undefined &&
    !hasOriginEclipsedTaraget(tagVector)
  ) {
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
  if (BREAK_RUNES[positionChar]) {
    decrementTarget(template, adjustedVector);
  }

  return adjustedVector;
};

export { crawlForTagName };
