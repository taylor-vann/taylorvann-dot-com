// brian taylor vann
// build integrals

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

type AttributeCrawl = <A>(
  template: Template<A>,
  vectorBounds: Vector
) => AttributeAction | undefined;

type AttributeValueCrawl = <A>(
  template: Template<A>,
  vectorBounds: Vector,
  Attributekind: AttributeAction
) => AttributeAction | undefined;

const ATTRIBUTE_FOUND = "ATTRIBUTE_FOUND";
const ATTRIBUTE_ASSIGNMENT = "ATTRIBUTE_ASSIGNMENT";
const IMPLICIT_ATTRIBUTE = "IMPLICIT_ATTRIBUTE";
const EXPLICIT_ATTRIBUTE = "EXPLICIT_ATTRIBUTE";
const INJECTED_ATTRIBUTE = "INJECTED_ATTRIBUTE";

const getAttributeName: AttributeCrawl = (template, vectorBounds) => {
  const attributeVector: Vector = copy(vectorBounds);

  let positionChar = getCharAtPosition(template, attributeVector.origin);
  if (positionChar === undefined || positionChar === " ") {
    return;
  }

  let tagNameCrawlState = ATTRIBUTE_FOUND;
  if (positionChar === " ") {
    tagNameCrawlState = IMPLICIT_ATTRIBUTE;
  }
  if (positionChar === "=") {
    tagNameCrawlState = ATTRIBUTE_ASSIGNMENT;
  }

  while (
    tagNameCrawlState === ATTRIBUTE_FOUND &&
    !hasOriginEclipsedTaraget(attributeVector)
  ) {
    if (incrementOrigin(template, attributeVector) === undefined) {
      return;
    }

    positionChar = getCharAtPosition(template, attributeVector.origin);
    if (positionChar === undefined) {
      return;
    }
    tagNameCrawlState = ATTRIBUTE_FOUND;
    if (positionChar === " ") {
      tagNameCrawlState = IMPLICIT_ATTRIBUTE;
    }
    if (positionChar === "=") {
      tagNameCrawlState = ATTRIBUTE_ASSIGNMENT;
    }
  }

  // we have found a tag, copy vector
  const adjustedVector: Vector = {
    origin: { ...vectorBounds.origin },
    target: { ...attributeVector.origin },
  };

  if (tagNameCrawlState === ATTRIBUTE_FOUND) {
    return {
      kind: IMPLICIT_ATTRIBUTE,
      attributeVector: adjustedVector,
    };
  }

  if (tagNameCrawlState === IMPLICIT_ATTRIBUTE) {
    if (positionChar === " ") {
      decrementTarget(template, adjustedVector);
    }
    return {
      kind: IMPLICIT_ATTRIBUTE,
      attributeVector: adjustedVector,
    };
  }

  if (tagNameCrawlState === ATTRIBUTE_ASSIGNMENT) {
    decrementTarget(template, adjustedVector);
    return {
      kind: EXPLICIT_ATTRIBUTE,
      attributeVector: adjustedVector,
      valueVector: adjustedVector,
    };
  }
};

const getAttributeQuality: AttributeValueCrawl = (
  template,
  vectorBounds,
  attributeAction
) => {
  // make sure explicity attribute follows (=")
  const attributeVector = copy(vectorBounds);

  let positionChar = getCharAtPosition(template, attributeVector.origin);
  if (positionChar !== "=") {
    return;
  }

  incrementOrigin(template, attributeVector);
  if (hasOriginEclipsedTaraget(attributeVector)) {
    return;
  }

  positionChar = getCharAtPosition(template, attributeVector.origin);
  if (positionChar !== '"') {
    return;
  }

  // we have an attribute!
  const attributeQualityVector = copy(attributeVector);

  // check for injected attribute
  const arrayIndex = attributeVector.origin.arrayIndex;
  if (incrementOrigin(template, attributeQualityVector) === undefined) {
    return;
  }
  positionChar = getCharAtPosition(template, attributeQualityVector.origin);
  if (positionChar === undefined) {
    return;
  }

  // check if there is a valid injection
  const arrayIndexDistance = Math.abs(
    arrayIndex - attributeQualityVector.origin.arrayIndex
  );
  if (arrayIndexDistance > 0 && positionChar !== '"') {
    return;
  }

  if (arrayIndexDistance === 1 && positionChar === '"') {
    // we have an injected attribute
    const injectionVector: Vector = {
      origin: { ...attributeVector.origin },
      target: { ...attributeQualityVector.origin },
    };

    const attributeVectorCopy = copy(attributeAction.attributeVector);

    return {
      kind: INJECTED_ATTRIBUTE,
      attributeVector: attributeVectorCopy,
      valueVector: injectionVector,
      injectionID: arrayIndex,
    };
  }

  // explore potential explicit attribute
  while (
    positionChar !== '"' &&
    !hasOriginEclipsedTaraget(attributeQualityVector)
  ) {
    if (incrementOrigin(template, attributeQualityVector) === undefined) {
      return;
    }
    // check if valid injection
    if (arrayIndex < attributeQualityVector.origin.arrayIndex) {
      return;
    }

    positionChar = getCharAtPosition(template, attributeQualityVector.origin);
    if (positionChar === undefined) {
      return;
    }
  }

  // check if bounds are valid
  if (positionChar === '"') {
    const explicitVector: Vector = {
      origin: { ...attributeVector.origin },
      target: { ...attributeQualityVector.origin },
    };
    const attributeVectorCopy = copy(attributeAction.attributeVector);

    return {
      kind: "EXPLICIT_ATTRIBUTE",
      attributeVector: attributeVectorCopy,
      valueVector: explicitVector,
    };
  }
};

const crawlForAttribute: AttributeCrawl = (template, vectorBounds) => {
  // get first character of attribute or return
  const attributeNameResults = getAttributeName(template, vectorBounds);
  if (attributeNameResults === undefined) {
    return;
  }
  if (attributeNameResults.kind === "IMPLICIT_ATTRIBUTE") {
    return attributeNameResults;
  }

  // get bounding vector
  let qualityVector: Vector = copy(vectorBounds);
  qualityVector.origin = {
    ...attributeNameResults.attributeVector.target,
  };
  incrementOrigin(template, qualityVector);

  return getAttributeQuality(template, qualityVector, attributeNameResults);
};

export { crawlForAttribute };
