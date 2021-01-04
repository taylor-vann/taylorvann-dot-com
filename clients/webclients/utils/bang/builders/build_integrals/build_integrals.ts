// brian taylor vann
// build integrals

import { StructureRender } from "../../type_flyweight/structure";
import { SkeletonNodes } from "../../type_flyweight/crawl";
import { IntegralRender } from "../../type_flyweight/integrals";
import { Vector } from "../../type_flyweight/text_vector";

import {
  copy,
  decrementTarget,
  incrementTarget,
  hasOriginEclipsedTaraget,
  incrementOrigin,
} from "../../text_vector/text_vector";
import { getCharAtPosition } from "../../text_position/text_position";

interface BuildIntegralsParams<A> {
  template: StructureRender<A>;
  skeleton: SkeletonNodes;
}
type BuildIntegrals = <A>(params: BuildIntegralsParams<A>) => IntegralRender;

type VectorCrawl = <A>(
  template: StructureRender<A>,
  innerXmlBounds: Vector
) => Vector | undefined;

const getFirstAttributeCharacter: VectorCrawl = (template, innerXmlBounds) => {
  const attributeVector: Vector = copy(innerXmlBounds);

  let positionChar = getCharAtPosition(template, attributeVector.target);
  while (positionChar === " " && !hasOriginEclipsedTaraget(attributeVector)) {
    incrementOrigin(template, attributeVector);
  }

  if (hasOriginEclipsedTaraget(attributeVector)) {
    return;
  }

  return attributeVector;
};

const getNextAttributeCharacter: VectorCrawl = (template, innerXmlBounds) => {
  const attributeVector: Vector = copy(innerXmlBounds);

  let positionChar = getCharAtPosition(template, attributeVector.target);
  while (positionChar === " " && !hasOriginEclipsedTaraget(attributeVector)) {
    incrementOrigin(template, attributeVector);
  }

  if (hasOriginEclipsedTaraget(attributeVector)) {
    return;
  }

  return attributeVector;
};

// content node steps

// open node steps

// independent node steps

// close node steps

// the goal is to return a list of building instructions

const buildIntegrals: BuildIntegrals = ({ template, skeleton }) => {
  // for each skeleton step

  // content node steps
  // build content steps

  // open node steps
  // independent node steps
  // // iterate across content bounds

  // close node steps
  // get tagname

  // the goal is to return a list of building instructions

  return [];
};

export { BuildIntegralsParams, buildIntegrals };
