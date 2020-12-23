// brian taylor vann

import { routers } from "./routers/routers";
import { StructureRender } from "../../../type_flyweight/structure";
import { CrawlResults, CrawlStatus } from "../../../type_flyweight/crawl";
import { Position } from "../../../type_flyweight/text_vector";
import {
  create,
  createFollowingVector,
  increment,
  getCharFromTarget,
} from "../../../text_vector/text_vector";

type Sieve = Partial<Record<CrawlStatus, CrawlStatus>>;
type SetNodeType = <A>(
  template: StructureRender<A>,
  results: CrawlResults
) => void;
type CrawlHasEnded = <A>(
  template: StructureRender<A>,
  previousCrawl?: CrawlResults
) => boolean;
type SetStartStateProperties = <A>(
  template: StructureRender<A>,
  previousCrawl?: CrawlResults
) => CrawlResults;
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

const crawlHasEnded: CrawlHasEnded = (template, previousCrawl) => {
  if (previousCrawl === undefined) {
    return false;
  }

  const templateLength = template.templateArray.length;
  const chunkLength =
    template.templateArray[previousCrawl.vector.target.arrayIndex].length;

  if (
    previousCrawl.vector.target.arrayIndex >= templateLength - 1 &&
    previousCrawl.vector.target.stringIndex >= chunkLength - 1
  ) {
    return true;
  }

  return false;
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
    cState.vector = createFollowingVector(previousCrawl.vector, template);
  }

  setNodeType(template, cState);
  return cState;
};

const setNodeType: SetNodeType = (template, cState) => {
  const nodeStates = routers[cState.nodeType];
  const char = getCharFromTarget(cState.vector, template);
  if (nodeStates !== undefined && char !== undefined) {
    const defaultNodeType = nodeStates[DEFAULT] ?? CONTENT_NODE;
    cState.nodeType = nodeStates[char] ?? defaultNodeType;
  }

  return cState;
};

const crawl: Crawl = (template, previousCrawl) => {
  if (crawlHasEnded(template, previousCrawl)) {
    return;
  }

  let openPosition: Position | undefined;
  const cState = setStartStateProperties(template, previousCrawl);

  while (increment(cState.vector.target, template)) {
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
