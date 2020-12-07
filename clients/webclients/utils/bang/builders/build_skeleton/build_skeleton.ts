// brian taylor vann

import { crawl } from "./crawl/crawl";
import { CrawlResults, SkeletonNodes } from "../../type_flyweight/crawl";
import { StructureRender } from "../../type_flyweight/render";
import { copy, decrement, increment } from "../../text_vector/text_vector";

type NodeType =
  | "OPEN_NODE"
  | "INDEPENDENT_NODE"
  | "CLOSE_NODE"
  | "CONTENT_NODE";

interface BuildMissingStringNodeParams<A> {
  template: StructureRender<A>;
  currentCrawl: CrawlResults;
  previousCrawl?: CrawlResults;
}
type BuildMissingStringNode = <A>(
  params: BuildMissingStringNodeParams<A>
) => CrawlResults | void;

type BuildSkeletonSieve = Record<string, NodeType>;

type BuildSkeleton = <A>(template: StructureRender<A>) => SkeletonNodes;

const MAX_DEPTH = 128;

const SKELETON_SIEVE: BuildSkeletonSieve = {
  ["OPEN_NODE_CONFIRMED"]: "OPEN_NODE",
  ["INDEPENDENT_NODE_CONFIRMED"]: "INDEPENDENT_NODE",
  ["CLOSE_NODE_CONFIRMED"]: "CLOSE_NODE",
  ["CONTENT_NODE"]: "CONTENT_NODE",
};

const buildMissingStringNode: BuildMissingStringNode = ({
  template,
  currentCrawl,
  previousCrawl,
}) => {
  if (previousCrawl === undefined) {
    return;
  }
  const target = previousCrawl.vector.target;
  const origin = currentCrawl.vector.origin;

  const stringDistance = Math.abs(origin.stringIndex - target.stringIndex);
  const stringArrayDistance = origin.arrayIndex - target.arrayIndex;
  if (2 > stringArrayDistance + stringDistance) {
    return;
  }

  // copy
  const previousVectorCopy = copy(previousCrawl.vector);
  const currentVectorCopy = copy(currentCrawl.vector);

  const contentStart = increment(previousVectorCopy.target, template);
  const contentEnd = decrement(currentVectorCopy.origin, template);
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

const buildSkeleton: BuildSkeleton = (template) => {
  let depth = 0;
  const skeleton: SkeletonNodes = [];

  let previousCrawl: CrawlResults | undefined;
  let currentCrawl = crawl(template, previousCrawl);

  while (currentCrawl && depth < MAX_DEPTH) {
    // get string in between crawls
    const stringBone = buildMissingStringNode({
      template,
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
    currentCrawl = crawl(template, previousCrawl);

    depth += 1;
  }

  return skeleton;
};

export { SkeletonNodes, buildSkeleton };
