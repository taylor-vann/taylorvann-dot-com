// brian taylor vann

import { crawl } from "../skeleton_crawl/skeleton_crawl";
import { CrawlResults, SkeletonNodes } from "../../type_flyweight/crawl";
import { Template } from "../../type_flyweight/template";
import { copy, decrement, increment } from "../../text_position/text_position";

type NodeType =
  | "OPEN_NODE"
  | "INDEPENDENT_NODE"
  | "CLOSE_NODE"
  | "CONTENT_NODE";

type CreateDefaultVector = () => CrawlResults;
interface BuildMissingStringNodeParams<A> {
  template: Template<A>;
  currentCrawl: CrawlResults;
  previousCrawl?: CrawlResults;
}
type BuildMissingStringNode = <A>(
  params: BuildMissingStringNodeParams<A>
) => CrawlResults | void;

type BuildSkeletonSieve = Record<string, NodeType>;

type BuildSkeleton = <A>(template: Template<A>) => SkeletonNodes;

const MAX_DEPTH = 128;

const DEFAULT_CRAWL_RESULTS: CrawlResults = {
  nodeType: "CONTENT_NODE",
  vector: {
    origin: { arrayIndex: 0, stringIndex: 0 },
    target: { arrayIndex: 0, stringIndex: 0 },
  },
};

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
  // get position values
  const originPos =
    previousCrawl !== undefined
      ? previousCrawl.vector.target
      : DEFAULT_CRAWL_RESULTS.vector.target;

  const targetPos = currentCrawl.vector.origin;

  // justify text vector distance
  const stringDistance = Math.abs(
    targetPos.stringIndex - originPos.stringIndex
  );
  const stringArrayDistance = Math.abs(
    targetPos.arrayIndex - originPos.arrayIndex
  );
  if (stringDistance + stringArrayDistance < 2) {
    return;
  }

  // copy and correlate position values
  const origin =
    previousCrawl === undefined
      ? copy(DEFAULT_CRAWL_RESULTS.vector.target)
      : copy(previousCrawl.vector.target);

  const target = copy(currentCrawl.vector.origin);

  decrement(template, target);
  if (previousCrawl !== undefined) {
    increment(template, origin);
  }

  return {
    nodeType: "CONTENT_NODE",
    vector: {
      origin,
      target,
    },
  };
};

const buildSkeleton: BuildSkeleton = (template) => {
  const skeleton: SkeletonNodes = [];
  let depth = 0;

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
