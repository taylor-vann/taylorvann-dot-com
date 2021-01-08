// brian taylor vann
// build integrals

import { StructureRender } from "../../type_flyweight/structure";
import { SkeletonNodes, CrawlResults } from "../../type_flyweight/crawl";
import { AttributeAction } from "../../type_flyweight/attribute_crawl";

import {
  IntegralAction,
  Integrals,
  CreateElement,
  CreateIndependentElement,
  CloseElement,
} from "../../type_flyweight/integrals";
import { Vector } from "../../type_flyweight/text_vector";

import {
  copy,
  createFollowingVector,
  decrementTarget,
  incrementTarget,
  hasOriginEclipsedTaraget,
  incrementOrigin,
} from "../../text_vector/text_vector";
import { getCharAtPosition } from "../../text_position/text_position";
import { crawlForTagName } from "./tag_name_crawl/tag_name_crawl";
import { crawlForAttribute } from "./attribute_crawl/attribute_crawl";

import { ContentCrawlAction } from "../../type_flyweight/content_crawl";

interface BuildIntegralsParams<A> {
  template: StructureRender<A>;
  skeleton: SkeletonNodes;
}
type BuildIntegrals = <A>(params: BuildIntegralsParams<A>) => Integrals;

type VectorCrawl = <A>(
  template: StructureRender<A>,
  innerXmlBounds: Vector
) => Vector | undefined;

interface AppendElementParams<A> {
  integrals: Integrals;
  template: StructureRender<A>;
  chunk: CrawlResults;
}
type AppendElementIntegrals = <A>(
  params: AppendElementParams<A>
) => Integrals | undefined;

const RECURSION_SAFETY = 256;

const incrementToNextSpaceRune: VectorCrawl = (template, innerXmlBounds) => {
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

const incrementToNextCharRune: VectorCrawl = (template, innerXmlBounds) => {
  const attributeVector: Vector = copy(innerXmlBounds);

  let positionChar = getCharAtPosition(template, attributeVector.target);
  while (positionChar !== " " && !hasOriginEclipsedTaraget(attributeVector)) {
    incrementOrigin(template, attributeVector);
  }

  if (hasOriginEclipsedTaraget(attributeVector)) {
    return;
  }

  return attributeVector;
};

const appendElementIntegrals: AppendElementIntegrals = ({
  integrals,
  template,
  chunk,
}) => {
  const innerXmlBounds = copy(chunk.vector);

  // adjust vector
  incrementOrigin(template, innerXmlBounds);
  decrementTarget(template, innerXmlBounds);
  if (chunk.nodeType === "INDEPENDENT_NODE_CONFIRMED") {
    decrementTarget(template, innerXmlBounds);
  }

  // get tag name
  const tagName = crawlForTagName(template, innerXmlBounds);
  if (tagName === undefined) {
    return;
  }

  if (hasOriginEclipsedTaraget(tagName)) {
    return;
  }
  // get vector bounds
  let attributeVectorBounds = createFollowingVector(template, tagName);
  if (attributeVectorBounds === undefined) {
    return;
  }
  attributeVectorBounds.target = { ...innerXmlBounds.target };

  let safety = 0;
  while (attributeVectorBounds !== undefined && safety < RECURSION_SAFETY) {
    // increase safety
    safety += 1;

    // move to the next space rune
    if (
      incrementToNextSpaceRune(template, attributeVectorBounds) === undefined
    ) {
      return;
    }

    // move to the next non-empty rune
    if (
      incrementToNextCharRune(template, attributeVectorBounds) === undefined
    ) {
      return;
    }

    // get attribute
    const attributeCrawlResults:
      | AttributeAction
      | undefined = crawlForAttribute(template, attributeVectorBounds);
    if (attributeCrawlResults === undefined) {
      return;
    }
    // add attribute to integrals
    integrals.push(attributeCrawlResults);

    // gather for next round
    // get vector bounds
    attributeVectorBounds = createFollowingVector(
      template,
      attributeCrawlResults.params.attributeVector
    );
    if (attributeVectorBounds === undefined) {
      return;
    }
    attributeVectorBounds.target = { ...innerXmlBounds.target };
  }

  if (chunk.nodeType === "INDEPENDENT_NODE_CONFIRMED") {
    decrementTarget(template, innerXmlBounds);
  }

  return integrals;
};

const appendCloseElementIntegrals: AppendElementIntegrals = ({
  integrals,
  template,
  chunk,
}) => {
  const innerXmlBounds = copy(chunk.vector);

  // adjust vector

  incrementOrigin(template, innerXmlBounds);
  incrementOrigin(template, innerXmlBounds);
  decrementTarget(template, innerXmlBounds);

  // get tag name
  let tagNameVector: Vector | undefined = copy(innerXmlBounds);
  tagNameVector = crawlForTagName(template, tagNameVector);
  if (tagNameVector === undefined) {
    console.log("no tag name found");
    return;
  }
  // add tag name to
  tagNameVector.origin = { ...innerXmlBounds.origin };

  const integralAction: CloseElement = {
    action: "CLOSE_ELEMENT",
    params: { tagNameVector },
  };

  // append integralAction to integrals
  integrals.push(integralAction);

  return integrals;
};

const appendContentIntegrals: AppendElementIntegrals = ({
  integrals,
  template,
  chunk,
}) => {
  const contentVector = copy(chunk.vector);
  const integralAction: ContentCrawlAction = {
    action: "CREATE_CONTENT",
    params: { contentVector },
  };

  // append integralAction to integrals
  integrals.push(integralAction);

  return integrals;
};

const buildIntegrals: BuildIntegrals = ({ template, skeleton }) => {
  const integrals: Integrals = [];
  for (const chunk of skeleton) {
    const nodeType = chunk.nodeType;
    if (nodeType === "CLOSE_NODE_CONFIRMED") {
      appendCloseElementIntegrals({ integrals, template, chunk });
    }
    if (nodeType === "CONTENT_NODE") {
      appendContentIntegrals({ integrals, template, chunk });
    }
    if (nodeType === "INDEPENDENT_NODE_CONFIRMED") {
      appendElementIntegrals({ integrals, template, chunk });
    }
    if (nodeType === "OPEN_NODE_CONFIRMED") {
      appendElementIntegrals({ integrals, template, chunk });
    }
  }

  return integrals;
};

export { BuildIntegralsParams, buildIntegrals };
