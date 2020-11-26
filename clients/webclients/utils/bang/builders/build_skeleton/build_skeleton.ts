// brian taylor vann

import { crawl } from "./crawl/crawl";
import { CrawlResults, SkeletonNodes } from "../../type_flyweight/crawl";
import { Position } from "../../type_flyweight/text_vector";

type NodeType =
  | "OPEN_NODE"
  | "INDEPENDENT_NODE"
  | "CLOSE_NODE"
  | "CONTENT_NODE";

type GetStringBonePosition = (
  brokenText: TemplateStringsArray,
  crawlResult: CrawlResults
) => Position | void;

interface BuildSkeletonStringBoneParams {
  brokenText: TemplateStringsArray;
  currentCrawl: CrawlResults;
  previousCrawl?: CrawlResults;
}
type BuildSkeletonStringBone = (
  params: BuildSkeletonStringBoneParams
) => CrawlResults | void;

type BuildSkeletonSieve = Record<string, NodeType>;

type BuildSkeleton = <A>(
  brokenText: TemplateStringsArray,
  ...injections: A[]
) => SkeletonNodes;

const MAX_DEPTH = 128;

const SKELETON_SIEVE: BuildSkeletonSieve = {
  ["OPEN_NODE_CONFIRMED"]: "OPEN_NODE",
  ["INDEPENDENT_NODE_CONFIRMED"]: "INDEPENDENT_NODE",
  ["CLOSE_NODE_CONFIRMED"]: "CLOSE_NODE",
  ["CONTENT_NODE"]: "CONTENT_NODE",
};

const getStringBoneStart: GetStringBonePosition = (
  brokenText,
  previousCrawl
) => {
  let { arrayIndex, stringIndex } = previousCrawl.vector.target;
  stringIndex += 1;
  stringIndex %= brokenText[arrayIndex].length;
  if (stringIndex === 0) {
    arrayIndex += 1;
  }

  return {
    arrayIndex,
    stringIndex,
  };
};

const getStringBoneEnd: GetStringBonePosition = (brokenText, currentCrawl) => {
  let { arrayIndex, stringIndex } = currentCrawl.vector.origin;
  stringIndex -= 1;
  if (stringIndex === -1) {
    arrayIndex -= 1;
    stringIndex += brokenText[arrayIndex].length;
  }

  return {
    arrayIndex,
    stringIndex,
  };
};

const buildSkeletonStringBone: BuildSkeletonStringBone = ({
  brokenText,
  currentCrawl,
  previousCrawl,
}) => {
  if (previousCrawl === undefined) {
    return;
  }
  const { target } = previousCrawl.vector;
  const { origin } = currentCrawl.vector;

  const stringDistance = Math.abs(origin.stringIndex - target.stringIndex);
  const stringArrayDistance = origin.arrayIndex - target.arrayIndex;
  if (2 > stringArrayDistance + stringDistance) {
    return;
  }

  const contentStart = getStringBoneStart(brokenText, previousCrawl);
  const contentEnd = getStringBoneEnd(brokenText, currentCrawl);
  if (contentStart && contentEnd) {
    return {
      nodeType: "CONTENT_NODE",
      vector: {
        origin: contentStart,
        target: contentEnd,
      },
    };
  }
};

const buildSkeleton: BuildSkeleton = (brokenText) => {
  let depth = 0;
  const skeleton: SkeletonNodes = [];

  let previousCrawl: CrawlResults | undefined;
  let currentCrawl = crawl(brokenText, previousCrawl);

  while (currentCrawl && depth < MAX_DEPTH) {
    // get string in between crawls
    const stringBone = buildSkeletonStringBone({
      brokenText,
      previousCrawl,
      currentCrawl,
    });
    if (stringBone) {
      skeleton.push(stringBone);
    }

    if (SKELETON_SIEVE[currentCrawl.nodeType]) {
      skeleton.push(currentCrawl);
    }
    previousCrawl = currentCrawl;
    currentCrawl = crawl(brokenText, previousCrawl);

    depth += 1;
  }

  return skeleton;
};

export { SkeletonNodes, buildSkeleton };
