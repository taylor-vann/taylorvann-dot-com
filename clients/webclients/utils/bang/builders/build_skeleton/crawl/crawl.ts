// brian taylor vann

import { routers } from "./routers/routers";
import { StructureRender } from "../../../type_flyweight/structure";
import { CrawlResults, CrawlStatus } from "../../../type_flyweight/crawl";
import { Position } from "../../../type_flyweight/text_vector";
import {
  create,
  createFollowingVector,
} from "../../../text_vector/text_vector";
import {
  increment,
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
  const cState: CrawlResults = {
    nodeType: CONTENT_NODE,
    vector: create(),
  };

  if (previousCrawl !== undefined) {
    const followingVector = createFollowingVector(
      template,
      previousCrawl.vector
    );
    if (followingVector === undefined) {
      return;
    }
  }

  setNodeType(template, cState);
  return cState;
};

const setNodeType: SetNodeType = (template, cState) => {
  const nodeStates = routers[cState.nodeType];
  const char = getCharFromTarget(template, cState.vector.target);
  if (nodeStates !== undefined && char !== undefined) {
    const defaultNodeType = nodeStates[DEFAULT] ?? CONTENT_NODE;
    cState.nodeType = nodeStates[char] ?? defaultNodeType;
  }

  return cState;
};

const crawl: Crawl = (template, previousCrawl) => {
  let openPosition: Position | undefined;
  const cState = setStartStateProperties(template, previousCrawl);
  if (cState === undefined) {
    return;
  }

  while (increment(template, cState.vector.target)) {
    if (
      validSieve[cState.nodeType] === undefined &&
      cState.vector.target.stringIndex === 0
    ) {
      cState.nodeType = CONTENT_NODE;
    }

    setNodeType(template, cState);
    if (confirmedSieve[cState.nodeType]) {
      if (openPosition !== undefined) {
        cState.vector.origin = { ...openPosition };
      }
      break;
    }

    if (cState.nodeType === OPEN_NODE) {
      openPosition = { ...cState.vector.target };
    }
  }

  return cState;
};

export { crawl };
