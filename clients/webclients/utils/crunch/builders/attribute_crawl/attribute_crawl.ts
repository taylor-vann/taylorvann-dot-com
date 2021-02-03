// brian taylor vann
// attribute crawl

import { Template } from "../../type_flyweight/template";
import { Vector } from "../../type_flyweight/text_vector";
import { AttributeAction } from "../../type_flyweight/attribute_crawl";
import {
  copy,
  decrementTarget,
  hasOriginEclipsedTaraget,
  incrementOrigin,
} from "../../text_vector/text_vector";
import { getCharAtPosition } from "../../text_position/text_position";

type AttributeCrawl = <N, A>(
  template: Template<N, A>,
  vectorBounds: Vector
) => AttributeAction | undefined;

type AttributeValueCrawl = <N, A>(
  template: Template<N, A>,
  vectorBounds: Vector,
  Attributekind: AttributeAction
) => AttributeAction | undefined;

type BreakRunes = Record<string, boolean>;

const QUOTE_RUNE = '"';
const ASSIGN_RUNE = "=";
const ATTRIBUTE_FOUND = "ATTRIBUTE_FOUND";
const ATTRIBUTE_ASSIGNMENT = "ATTRIBUTE_ASSIGNMENT";
const IMPLICIT_ATTRIBUTE = "IMPLICIT_ATTRIBUTE";
const EXPLICIT_ATTRIBUTE = "EXPLICIT_ATTRIBUTE";
const INJECTED_ATTRIBUTE = "INJECTED_ATTRIBUTE";

const BREAK_RUNES: BreakRunes = {
  " ": true,
  "\n": true,
};

const getAttributeName: AttributeCrawl = (template, vectorBounds) => {
  let positionChar = getCharAtPosition(template, vectorBounds.origin);
  if (positionChar === undefined || BREAK_RUNES[positionChar]) {
    return;
  }

  let tagNameCrawlState = ATTRIBUTE_FOUND;
  const bounds: Vector = copy(vectorBounds);

  while (
    tagNameCrawlState === ATTRIBUTE_FOUND &&
    !hasOriginEclipsedTaraget(bounds)
  ) {
    if (incrementOrigin(template, bounds) === undefined) {
      return;
    }

    positionChar = getCharAtPosition(template, bounds.origin);
    if (positionChar === undefined) {
      return;
    }
    tagNameCrawlState = ATTRIBUTE_FOUND;
    if (BREAK_RUNES[positionChar]) {
      tagNameCrawlState = IMPLICIT_ATTRIBUTE;
    }
    if (positionChar === ASSIGN_RUNE) {
      tagNameCrawlState = ATTRIBUTE_ASSIGNMENT;
    }
  }

  // we have found a tag, copy vector
  const attributeVector: Vector = {
    origin: { ...vectorBounds.origin },
    target: { ...bounds.origin },
  };

  // edge case, we've found text but no break runes
  if (tagNameCrawlState === ATTRIBUTE_FOUND) {
    return {
      kind: IMPLICIT_ATTRIBUTE,
      attributeVector,
    };
  }

  // if implict attribute
  if (tagNameCrawlState === IMPLICIT_ATTRIBUTE) {
    if (BREAK_RUNES[positionChar]) {
      decrementTarget(template, attributeVector);
    }
    return {
      kind: IMPLICIT_ATTRIBUTE,
      attributeVector,
    };
  }

  if (tagNameCrawlState === ATTRIBUTE_ASSIGNMENT) {
    decrementTarget(template, attributeVector);
    return {
      kind: EXPLICIT_ATTRIBUTE,
      valueVector: attributeVector,
      attributeVector,
    };
  }
};

const getAttributeQuality: AttributeValueCrawl = (
  template,
  vectorBounds,
  attributeAction
) => {
  let positionChar = getCharAtPosition(template, vectorBounds.origin);
  if (positionChar !== ASSIGN_RUNE) {
    return;
  }

  const bound = copy(vectorBounds);

  incrementOrigin(template, bound);
  if (hasOriginEclipsedTaraget(bound)) {
    return;
  }

  positionChar = getCharAtPosition(template, bound.origin);
  if (positionChar !== QUOTE_RUNE) {
    return;
  }

  // we have an attribute!
  const { arrayIndex } = bound.origin;
  const valVector = copy(bound);

  // check for injected attribute
  if (incrementOrigin(template, valVector) === undefined) {
    return;
  }
  positionChar = getCharAtPosition(template, valVector.origin);
  if (positionChar === undefined) {
    return;
  }

  // check if there is a valid injection
  const arrayIndexDistance = Math.abs(arrayIndex - valVector.origin.arrayIndex);
  if (arrayIndexDistance === 1 && positionChar === QUOTE_RUNE) {
    return {
      kind: INJECTED_ATTRIBUTE,
      injectionID: arrayIndex,
      attributeVector: attributeAction.attributeVector,
      valueVector: {
        origin: { ...bound.origin },
        target: { ...valVector.origin },
      },
    };
  }

  // explore potential for explicit attribute
  while (positionChar !== QUOTE_RUNE && !hasOriginEclipsedTaraget(valVector)) {
    if (incrementOrigin(template, valVector) === undefined) {
      return;
    }

    // return if unexpected injection found
    if (arrayIndex < valVector.origin.arrayIndex) {
      return;
    }

    positionChar = getCharAtPosition(template, valVector.origin);
    if (positionChar === undefined) {
      return;
    }
  }

  // exlpicit attribute found
  if (
    attributeAction.kind === "EXPLICIT_ATTRIBUTE" &&
    positionChar === QUOTE_RUNE
  ) {
    attributeAction.valueVector = {
      origin: { ...bound.origin },
      target: { ...valVector.origin },
    };

    // get text vettor between (")<quotes>(")
    incrementOrigin(template, attributeAction.valueVector);
    decrementTarget(template, attributeAction.valueVector);

    return attributeAction;
  }
};

const crawlForAttribute: AttributeCrawl = (template, vectorBounds) => {
  // get first character of attribute or return
  const attrResults = getAttributeName(template, vectorBounds);
  if (attrResults === undefined) {
    return;
  }
  if (attrResults.kind === "IMPLICIT_ATTRIBUTE") {
    return attrResults;
  }

  // get bounds for attribute value
  let valBounds: Vector = copy(vectorBounds);
  valBounds.origin = {
    ...attrResults.attributeVector.target,
  };
  incrementOrigin(template, valBounds);

  return getAttributeQuality(template, valBounds, attrResults);
};

export { crawlForAttribute };
