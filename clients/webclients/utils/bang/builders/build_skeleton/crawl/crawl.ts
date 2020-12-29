// brian taylor vann

import { routers } from "./routers/routers";
import { StructureRender } from "../../../type_flyweight/structure";
import { CrawlResults, CrawlStatus } from "../../../type_flyweight/crawl";
import { Position, Vector } from "../../../type_flyweight/text_vector";
import {
  create,
  createFollowingVector,
  incrementTarget,
} from "../../../text_vector/text_vector";
import {
  copy as copyPosition,
  getCharFromTarget,
} from "../../../text_position/text_position";

type Sieve = Partial<Record<CrawlStatus, CrawlStatus>>;
type SetNodeType = <A>(
  template: StructureRender<A>,
  results: CrawlResults
) => void;
type SetStartStateProperties = <A>(
  template: StructureRender<A>,
  previousCrawl?: CrawlResults
) => CrawlResults | undefined;
type Crawl = <A>(
  template: StructureRender<A>,
  previousCrawl?: CrawlResults
) => CrawlResults | undefined;

const DEFAULT = "DEFAULT";
const CONTENT_NODE = "CONTENT_NODE";
const OPEN_NODE = "OPEN_NODE";

const validSieve: Sieve = {
  ["OPEN_NODE_VALID"]: "OPEN_NODE_VALID",
  ["CLOSE_NODE_VALID"]: "CLOSE_NODE_VALID",
  ["INDEPENDENT_NODE_VALID"]: "INDEPENDENT_NODE_VALID",
};

const confirmedSieve: Sieve = {
  ["OPEN_NODE_CONFIRMED"]: "OPEN_NODE_CONFIRMED",
  ["CLOSE_NODE_CONFIRMED"]: "CLOSE_NODE_CONFIRMED",
  ["INDEPENDENT_NODE_CONFIRMED"]: "INDEPENDENT_NODE_CONFIRMED",
};

const setStartStateProperties: SetStartStateProperties = (
  template,
  previousCrawl
) => {
  let followingVector: Vector | undefined = create();
  if (previousCrawl !== undefined) {
    followingVector = createFollowingVector(template, previousCrawl.vector);
  }
  if (followingVector === undefined) {
    return;
  }

  const crawlState: CrawlResults = {
    nodeType: CONTENT_NODE,
    vector: followingVector,
  };
  setNodeType(template, crawlState);

  return crawlState;
};

const setNodeType: SetNodeType = (template, crawlState) => {
  const nodeStates = routers[crawlState.nodeType];
  const char = getCharFromTarget(template, crawlState.vector.target);

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
