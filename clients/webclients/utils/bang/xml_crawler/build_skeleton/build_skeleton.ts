import { BrokenTextPostition, crawl, CrawlResults } from "../crawl/crawl";

type NodeType =
  | "OPEN_NODE"
  | "INDEPENDENT_NODE"
  | "CLOSE_NODE"
  | "CONTENT_NODE";

type SkeletonNodes = CrawlResults[];

type GetStringBonePosition = (
  brokenText: TemplateStringsArray,
  crawlResult: CrawlResults
) => BrokenTextPostition | void;

interface BuildSkeletonStringBoneParams {
  brokenText: TemplateStringsArray;
  currentCrawl: CrawlResults;
  previousCrawl?: CrawlResults;
}
type BuildSkeletonStringBone = (
  params: BuildSkeletonStringBoneParams
) => CrawlResults | void;

type BuildSkeletonSieve = Record<string, NodeType>;

type BuildSkeleton = (
  brokenText: TemplateStringsArray,
  ...injections: string[]
) => SkeletonNodes;

const MAX_RECURSION = 128;

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
  let { arrayIndex, stringIndex } = previousCrawl.target.end;
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
  let { arrayIndex, stringIndex } = currentCrawl.target.start;
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
  const { end } = previousCrawl.target;
  const { start } = currentCrawl.target;

  const stringDistance = Math.abs(start.stringIndex - end.stringIndex);
  const stringArrayDistance = start.arrayIndex - end.arrayIndex;
  if (2 > stringArrayDistance + stringDistance) {
    return;
  }

  const contentStart = getStringBoneStart(brokenText, previousCrawl);
  const contentEnd = getStringBoneEnd(brokenText, currentCrawl);
  if (contentStart && contentEnd) {
    return {
      nodeType: "CONTENT_NODE",
      target: {
        start: contentStart,
        end: contentEnd,
      },
    };
  }
};

const buildSkeleton: BuildSkeleton = (brokenText, ...injections) => {
  let depth = 0;
  const skeleton: SkeletonNodes = [];

  let previousCrawl: CrawlResults | undefined;
  let currentCrawl = crawl(brokenText, previousCrawl);

  while (currentCrawl && depth < MAX_RECURSION) {
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
