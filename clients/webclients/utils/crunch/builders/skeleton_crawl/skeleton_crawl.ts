// brian taylor vann
// skeleton crawl

import { routers } from "../skeleton_routers/skeleton_routers";
import { Template } from "../../type_flyweight/template";
import { CrawlResults, CrawlStatus } from "../../type_flyweight/skeleton_crawl";
import { Position } from "../../type_flyweight/text_vector";
import {
  create,
  createFollowingVector,
  incrementTarget,
} from "../../text_vector/text_vector";
import {
  copy as copyPosition,
  getCharAtPosition,
} from "../../text_position/text_position";

type Sieve = Partial<Record<CrawlStatus, CrawlStatus>>;
type SetNodeType = <N, A>(
  template: Template<N, A>,
  results: CrawlResults
) => void;
type SetStartStateProperties = <N, A>(
  template: Template<N, A>,
  previousCrawl?: CrawlResults
) => CrawlResults | undefined;
type Crawl = <N, A>(
  template: Template<N, A>,
  previousCrawl?: CrawlResults
) => CrawlResults | undefined;

const DEFAULT = "DEFAULT";
const CONTENT_NODE = "CONTENT_NODE";
const OPEN_NODE = "OPEN_NODE";

const validSieve: Sieve = {
  ["OPEN_NODE_VALID"]: "OPEN_NODE_VALID",
  ["CLOSE_NODE_VALID"]: "CLOSE_NODE_VALID",
  ["SELF_CLOSING_NODE_VALID"]: "SELF_CLOSING_NODE_VALID",
};

const confirmedSieve: Sieve = {
  ["OPEN_NODE_CONFIRMED"]: "OPEN_NODE_CONFIRMED",
  ["CLOSE_NODE_CONFIRMED"]: "CLOSE_NODE_CONFIRMED",
  ["SELF_CLOSING_NODE_CONFIRMED"]: "SELF_CLOSING_NODE_CONFIRMED",
};

const setStartStateProperties: SetStartStateProperties = (
  template,
  previousCrawl
) => {
  if (previousCrawl === undefined) {
    return {
      nodeType: CONTENT_NODE,
      vector: create(),
    };
  }

  const followingVector = createFollowingVector(template, previousCrawl.vector);
  if (followingVector === undefined) {
    return;
  }

  const crawlState: CrawlResults = {
    nodeType: CONTENT_NODE,
    vector: followingVector,
  };

  return crawlState;
};

const setNodeType: SetNodeType = (template, crawlState) => {
  const nodeStates = routers[crawlState.nodeType];
  const char = getCharAtPosition(template, crawlState.vector.target);

  if (nodeStates !== undefined && char !== undefined) {
    const defaultNodeType = nodeStates[DEFAULT] ?? CONTENT_NODE;
    crawlState.nodeType = nodeStates[char] ?? defaultNodeType;
  }

  return crawlState;
};

const crawl: Crawl = (template, previousCrawl) => {
  const crawlState = setStartStateProperties(template, previousCrawl);
  if (crawlState === undefined) {
    return;
  }

  let openPosition: Position | undefined;
  setNodeType(template, crawlState);

  while (incrementTarget(template, crawlState.vector)) {
    // default to content_node on each cycle
    if (
      validSieve[crawlState.nodeType] === undefined &&
      crawlState.vector.target.stringIndex === 0
    ) {
      crawlState.nodeType = CONTENT_NODE;
    }

    setNodeType(template, crawlState);

    if (crawlState.nodeType === OPEN_NODE) {
      openPosition = copyPosition(crawlState.vector.target);
    }

    if (confirmedSieve[crawlState.nodeType]) {
      if (openPosition !== undefined) {
        crawlState.vector.origin = openPosition;
      }
      break;
    }
  }

  return crawlState;
};

export { crawl };