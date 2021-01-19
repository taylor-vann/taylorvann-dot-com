// brian taylor vann
// build integrals

import { Template } from "../../type_flyweight/template";
import { SkeletonNodes, CrawlResults } from "../../type_flyweight/crawl";

import { Integrals, CloseNodeAction } from "../../type_flyweight/integrals";
import { Vector } from "../../type_flyweight/text_vector";

import {
  copy,
  createFollowingVector,
  decrementTarget,
  hasOriginEclipsedTaraget,
  incrementOrigin,
} from "../../text_vector/text_vector";
import { getCharAtPosition } from "../../text_position/text_position";
import { crawlForTagName } from "../tag_name_crawl/tag_name_crawl";
import { crawlForAttribute } from "../attribute_crawl/attribute_crawl";

import { ContentCrawlAction } from "../../type_flyweight/content_crawl";

interface BuildIntegralsParams<A> {
  template: Template<A>;
  skeleton: SkeletonNodes;
}
type BuildIntegrals = <A>(params: BuildIntegralsParams<A>) => Integrals;

type VectorCrawl = <A>(
  template: Template<A>,
  innerXmlBounds: Vector
) => Vector | undefined;

interface AppendNodeParams<A> {
  integrals: Integrals;
  template: Template<A>;
  chunk: CrawlResults;
}
type AppendNodeIntegrals = <A>(
  params: AppendNodeParams<A>
) => Integrals | undefined;

interface AppendNodeAttributeParams<A> {
  integrals: Integrals;
  template: Template<A>;
  chunk: Vector;
}
type AppendNodeAttributeIntegrals = <A>(
  params: AppendNodeAttributeParams<A>
) => Integrals | undefined;

const RECURSION_SAFETY = 256;

// creates a side effect in innerXmlBounds
const incrementOriginToNextSpaceRune: VectorCrawl = (
  template,
  innerXmlBounds
) => {
  let positionChar = getCharAtPosition(template, innerXmlBounds.origin);
  if (positionChar === undefined) {
    return;
  }

  while (positionChar !== " ") {
    if (hasOriginEclipsedTaraget(innerXmlBounds)) {
      return;
    }
    if (incrementOrigin(template, innerXmlBounds) === undefined) {
      return;
    }
    positionChar = getCharAtPosition(template, innerXmlBounds.origin);
    if (positionChar === undefined) {
      return;
    }
  }

  return innerXmlBounds;
};

// creates a side effect in innerXmlBounds
const incrementOriginToNextCharRune: VectorCrawl = (
  template,
  innerXmlBounds
) => {
  let positionChar = getCharAtPosition(template, innerXmlBounds.origin);
  if (positionChar === undefined) {
    return;
  }

  while (positionChar === " ") {
    if (hasOriginEclipsedTaraget(innerXmlBounds)) {
      return;
    }
    if (incrementOrigin(template, innerXmlBounds) === undefined) {
      return;
    }
    positionChar = getCharAtPosition(template, innerXmlBounds.origin);
    if (positionChar === undefined) {
      return;
    }
  }

  return innerXmlBounds;
};

const appendNodeAttributeIntegrals: AppendNodeAttributeIntegrals = ({
  integrals,
  template,
  chunk,
}) => {
  let safety = 0;
  while (!hasOriginEclipsedTaraget(chunk) && safety < RECURSION_SAFETY) {
    safety += 1;

    if (incrementOriginToNextSpaceRune(template, chunk) === undefined) {
      return;
    }
    if (incrementOriginToNextCharRune(template, chunk) === undefined) {
      return;
    }

    const attributeCrawlResults = crawlForAttribute(template, chunk);
    if (attributeCrawlResults === undefined) {
      return;
    }

    // set origin to following position
    if (attributeCrawlResults.action === "APPEND_IMPLICIT_ATTRIBUTE") {
      chunk.origin = { ...attributeCrawlResults.params.attributeVector.target };
    }
    if (attributeCrawlResults.action === "APPEND_EXPLICIT_ATTRIBUTE") {
      chunk.origin = { ...attributeCrawlResults.params.valueVector.target };
    }
    if (attributeCrawlResults.action === "APPEND_INJECTED_ATTRIBUTE") {
      chunk.origin = { ...attributeCrawlResults.params.valueVector.target };
    }

    integrals.push(attributeCrawlResults);

    // add attribute to integrals
  }

  return integrals;
};

const appendNodeIntegrals: AppendNodeIntegrals = ({
  integrals,
  template,
  chunk,
}) => {
  const innerXmlBounds = copy(chunk.vector);

  // adjust vector
  incrementOrigin(template, innerXmlBounds);
  decrementTarget(template, innerXmlBounds);

  // get tag name
  const tagNameVector = crawlForTagName(template, innerXmlBounds);
  if (tagNameVector === undefined) {
    return;
  }

  // add create
  integrals.push({
    action: "CREATE_NODE",
    params: { tagNameVector },
  });

  const followingVector = createFollowingVector(template, tagNameVector);
  if (followingVector === undefined) {
    return;
  }
  followingVector.target = { ...innerXmlBounds.target };

  // more debugs here lol
  appendNodeAttributeIntegrals({ integrals, template, chunk: followingVector });

  return integrals;
};

const appendSelfClosingNodeIntegrals: AppendNodeIntegrals = ({
  integrals,
  template,
  chunk,
}) => {
  const innerXmlBounds = copy(chunk.vector);

  // adjust vector
  incrementOrigin(template, innerXmlBounds);
  decrementTarget(template, innerXmlBounds);
  decrementTarget(template, innerXmlBounds);

  // get tag name
  const tagNameVector = crawlForTagName(template, innerXmlBounds);
  if (tagNameVector === undefined) {
    return;
  }

  // add create
  integrals.push({
    action: "CREATE_SELF_CLOSING_NODE",
    params: { tagNameVector },
  });

  if (hasOriginEclipsedTaraget(tagNameVector)) {
    return;
  }

  const followingChunk = createFollowingVector(template, tagNameVector);
  if (followingChunk === undefined) {
    return;
  }
  followingChunk.target = { ...innerXmlBounds.target };

  // call attribute search

  return integrals;
};

const appendCloseNodeIntegrals: AppendNodeIntegrals = ({
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
    return;
  }
  // add tag name to
  tagNameVector.origin = { ...innerXmlBounds.origin };

  const integralAction: CloseNodeAction = {
    action: "CLOSE_NODE",
    params: { tagNameVector },
  };

  // append integralAction to integrals
  integrals.push(integralAction);

  return integrals;
};

const appendContentIntegrals: AppendNodeIntegrals = ({
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
      appendCloseNodeIntegrals({ integrals, template, chunk });
    }
    if (nodeType === "CONTENT_NODE") {
      appendContentIntegrals({ integrals, template, chunk });
    }
    if (nodeType === "SELF_CLOSING_NODE_CONFIRMED") {
      appendSelfClosingNodeIntegrals({ integrals, template, chunk });
    }
    if (nodeType === "OPEN_NODE_CONFIRMED") {
      appendNodeIntegrals({ integrals, template, chunk });
    }
  }

  return integrals;
};

export { BuildIntegralsParams, buildIntegrals };
