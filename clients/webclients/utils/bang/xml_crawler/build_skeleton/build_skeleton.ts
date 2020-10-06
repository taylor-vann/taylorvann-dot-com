import { CrawlResults, crawl, createNotFoundCrawlState } from "../crawl/crawl";

type NodeType =
  | "OPEN_NODE"
  | "INDEPENDENT_NODE"
  | "CLOSE_NODE"
  | "CONTENT_NODE";

type SkeletonNodes = CrawlResults[];

interface BuildSkeletonStringBoneParams {
  brokenText: TemplateStringsArray;
  previousCrawl: CrawlResults;
  currentCrawl: CrawlResults;
}
type BuildSkeletonStringBone = (
  params: BuildSkeletonStringBoneParams
) => CrawlResults | void;

type HasReachedTheEnd = (
  brokenText: TemplateStringsArray,
  currentCrawl: CrawlResults
) => boolean;

type BuildSkeletonSieve = Record<string, NodeType>;

type BuildSkeleton = (
  brokenText: TemplateStringsArray,
  ...injections: string[]
) => SkeletonNodes;

const MAX_RECURSION = 1000;

const SKELETON_SIEVE: BuildSkeletonSieve = {
  ["OPEN_NODE_CONFIRMED"]: "OPEN_NODE",
  ["INDEPENDENT_NODE_CONFRIMED"]: "INDEPENDENT_NODE",
  ["CLOSE_NODE_CONFIRMED"]: "CLOSE_NODE",
  ["CONTENT_NODE"]: "CONTENT_NODE",
};

const crawlIsNotComplete: HasReachedTheEnd = (brokenText, currentCrawl) => {
  const { stringArrayIndex, stringIndex } = currentCrawl.target.endPosition;
  const brokenTextSafeLength = brokenText.length;
  return (
    stringArrayIndex < brokenTextSafeLength &&
    stringIndex < brokenText[stringArrayIndex].length - 1
  );
};

const buildSkeletonStringBone: BuildSkeletonStringBone = ({
  brokenText,
  previousCrawl,
  currentCrawl,
}) => {
  // this can be a single function
  const { endPosition } = previousCrawl.target;
  const { startPosition } = currentCrawl.target;
  const stringDistance = startPosition.stringIndex - endPosition.stringIndex;
  const stringArrayDistance =
    startPosition.stringArrayIndex - endPosition.stringArrayIndex;

  if (stringArrayDistance + stringDistance < 2) {
    return;
  }

  // at least distance of 2
  // this can be a single function
  let endStringArrayIndex = startPosition.stringArrayIndex;
  let endStringIndex = startPosition.stringIndex - 1;
  if (endStringIndex < 0 && 0 < endStringArrayIndex) {
    endStringIndex = startPosition.stringIndex - 1;

    endStringArrayIndex = endStringArrayIndex - 1;
    endStringIndex = brokenText[endStringArrayIndex].length + endStringIndex;
  }

  return {
    nodeType: "CONTENT_NODE",
    target: {
      startPosition: {
        stringArrayIndex: endPosition.stringArrayIndex,
        stringIndex: endPosition.stringIndex,
      },
      endPosition: {
        stringArrayIndex: 0,
        stringIndex: 0,
      },
    },
  };
};

const buildSkeleton: BuildSkeleton = (brokenText, ...injections) => {
  // iterate through brokenText and injectsion
  // add string and nodes to results array
  const skeleton: SkeletonNodes = [];
  let previousCrawl: CrawlResults | undefined;
  let currentCrawl: CrawlResults | undefined = createNotFoundCrawlState();

  let currRecursionDepth = 0;

  while (
    crawlIsNotComplete(brokenText, currentCrawl) &&
    currRecursionDepth < MAX_RECURSION
  ) {
    let { endPosition } = currentCrawl.target;
    previousCrawl = currentCrawl;
    currentCrawl = crawl({ brokenText, previousCrawl });
    if (currentCrawl === undefined) {
      break;
    }

    const stringNodeBone = buildSkeletonStringBone({
      brokenText,
      previousCrawl,
      currentCrawl,
    });
    // if (stringNodeBone) {
    //   skeleton.push(stringNodeBone);
    // }

    const nodeType = SKELETON_SIEVE[currentCrawl.nodeType];

    if (nodeType) {
      skeleton.push(currentCrawl);
    }

    currRecursionDepth += 1;
  }

  return skeleton;
};

export { SkeletonNodes, buildSkeleton };
