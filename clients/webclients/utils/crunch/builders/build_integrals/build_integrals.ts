// brian taylor vann
// build integrals

import { Template } from "../../type_flyweight/template";
import {
  SkeletonNodes,
  CrawlResults,
} from "../../type_flyweight/skeleton_crawl";
import { Integrals } from "../../type_flyweight/integrals";
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

interface BuildIntegralsParams<A> {
  template: Template<A>;
  skeleton: SkeletonNodes;
}
type BuildIntegrals = <A>(params: BuildIntegralsParams<A>) => Integrals;

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
    // something has gone wrong and we should stop
    if (attributeCrawlResults === undefined) {
      return;
    }

    // set origin to following position
    if (attributeCrawlResults.kind === "IMPLICIT_ATTRIBUTE") {
      chunk.origin = { ...attributeCrawlResults.attributeVector.target };
    }
    if (attributeCrawlResults.kind === "EXPLICIT_ATTRIBUTE") {
      chunk.origin = { ...attributeCrawlResults.valueVector.target };
    }
    if (attributeCrawlResults.kind === "INJECTED_ATTRIBUTE") {
      chunk.origin = { ...attributeCrawlResults.valueVector.target };
    }

    integrals.push(attributeCrawlResults);
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

  integrals.push({
    kind: "NODE",
    tagNameVector,
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

  integrals.push({
    kind: "SELF_CLOSING_NODE",
    tagNameVector,
  });

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

  // append integralAction to integrals
  integrals.push({
    kind: "CLOSE_NODE",
    tagNameVector,
  });

  return integrals;
};

const appendContentIntegrals: AppendNodeIntegrals = ({
  integrals,
  template,
  chunk,
}) => {
  const { origin, target } = chunk.vector;

  // what if content has no injections?
  if (origin.arrayIndex === target.arrayIndex) {
    integrals.push({ kind: "TEXT", textVector: chunk.vector });
    return;
  }

  // get that beginning text stuff
  let stringIndex = template.templateArray[origin.arrayIndex].length - 1;
  let textVector = {
    origin,
    target: {
      arrayIndex: origin.arrayIndex,
      stringIndex,
    },
  };

  integrals.push({ kind: "TEXT", textVector });
  integrals.push({
    kind: "CONTEXT_INJECTION",
    injectionID: origin.arrayIndex,
  });

  // get that middle text stuff
  let arrayIndex = origin.arrayIndex + 1;
  while (arrayIndex < target.arrayIndex) {
    stringIndex = template.templateArray[arrayIndex].length - 1;
    textVector = {
      origin: {
        arrayIndex,
        stringIndex: 0,
      },
      target: {
        arrayIndex,
        stringIndex,
      },
    };

    integrals.push({ kind: "TEXT", textVector });
    integrals.push({
      kind: "CONTEXT_INJECTION",
      injectionID: arrayIndex,
    });

    arrayIndex += 1;
  }

  // get that end text stuff
  textVector = {
    origin: {
      arrayIndex: target.arrayIndex,
      stringIndex: 0,
    },
    target,
  };

  integrals.push({ kind: "TEXT", textVector });

  return integrals;
};

const buildIntegrals: BuildIntegrals = ({ template, skeleton }) => {
  const integrals: Integrals = [];

  for (const chunk of skeleton) {
    const nodeType = chunk.nodeType;
    const origin = chunk.vector.origin;

    // does injection come before?
    if (origin.stringIndex === 0 && origin.arrayIndex !== 0) {
      integrals.push({
        kind: "CONTEXT_INJECTION",
        injectionID: origin.arrayIndex - 1,
      });
    }

    if (nodeType === "OPEN_NODE_CONFIRMED") {
      appendNodeIntegrals({ integrals, template, chunk });
    }
    if (nodeType === "CLOSE_NODE_CONFIRMED") {
      appendCloseNodeIntegrals({ integrals, template, chunk });
    }
    if (nodeType === "CONTENT_NODE") {
      appendContentIntegrals({ integrals, template, chunk });
    }
    if (nodeType === "SELF_CLOSING_NODE_CONFIRMED") {
      appendSelfClosingNodeIntegrals({ integrals, template, chunk });
    }
  }

  return integrals;
};

export { BuildIntegralsParams, buildIntegrals };
