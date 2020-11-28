// brian taylor vann

import { routers } from "./routers/routers";
import { StructureRender } from "../../../type_flyweight/structure";
import { CrawlResults, CrawlStatus } from "../../../type_flyweight/crawl";
import { Position } from "../../../type_flyweight/text_vector";
import {
  copy,
  create,
  increment,
  getCharFromTarget,
} from "../../../text_vector/text_vector";

type Sieve = Partial<Record<CrawlStatus, CrawlStatus>>;

type SetPosition = (
  results: CrawlResults,
  arrayIndex: number,
  stringIndex: number
) => void;

type CreateCrawlState = () => CrawlResults;
type SetNodeType = (results: CrawlResults, character: string) => void;

type SetStartStateProperties = <A>(
  template: StructureRender<A>,
  previousCrawl?: CrawlResults
) => CrawlResults | undefined;

type Crawl = <A>(
  template: StructureRender<A>,
  previousCrawl?: CrawlResults
) => CrawlResults | undefined;

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

const createDefaultCrawlState: CreateCrawlState = () => {
  return {
    nodeType: "CONTENT_NODE",
    vector: create(),
  };
};

const setStartStateProperties: SetStartStateProperties = (
  template,
  previousCrawl
) => {
  const cState = createDefaultCrawlState();
  if (previousCrawl === undefined) {
    return cState;
  }

  console.log("set start state properties");
  console.log(previousCrawl.vector);
  increment(previousCrawl.vector, template);
  console.log(previousCrawl.vector);
  cState.vector = previousCrawl.vector;

  return cState;
};

const setNodeType: SetNodeType = (cState, char) => {
  const nodeStates = routers[cState.nodeType];
  if (nodeStates) {
    const defaultNodeType = nodeStates["DEFAULT"] ?? CONTENT_NODE;
    cState.nodeType = nodeStates[char] ?? defaultNodeType;
  }

  return cState;
};

const setStart: SetPosition = (
  results: CrawlResults,
  arrayIndex: number,
  stringIndex: number
) => {
  results.vector.origin.arrayIndex = arrayIndex;
  results.vector.origin.stringIndex = stringIndex;
  results.vector.target.arrayIndex = arrayIndex;
  results.vector.target.stringIndex = stringIndex;
};

const setEnd: SetPosition = (
  results: CrawlResults,
  arrayIndex: number,
  stringIndex: number
) => {
  results.vector.target.arrayIndex = arrayIndex;
  results.vector.target.stringIndex = stringIndex;
};

const crawl: Crawl = (template, previousCrawl) => {
  const cState = setStartStateProperties(template, previousCrawl);
  if (cState === undefined) {
    return;
  }

  let { stringIndex, arrayIndex } = cState.vector.origin;
  // retain most recent postition
  const suspect: Position = {
    arrayIndex,
    stringIndex,
  };

  while (arrayIndex < template.templateArray.length) {
    if (validSieve[cState.nodeType] === undefined) {
      cState.nodeType = CONTENT_NODE;
    }

    const chunk = template.templateArray[arrayIndex];
    while (stringIndex < chunk.length) {
      setNodeType(cState, chunk.charAt(stringIndex));

      if (confirmedSieve[cState.nodeType]) {
        // if confirmed, suspected target is verified
        setStart(cState, suspect.arrayIndex, suspect.stringIndex);
        setEnd(cState, arrayIndex, stringIndex);
        return cState;
      }

      if (cState.nodeType === OPEN_NODE) {
        suspect.arrayIndex = arrayIndex;
        suspect.stringIndex = stringIndex;
      }

      stringIndex += 1;
    }

    // skip to next chunk
    stringIndex = 0;
    arrayIndex += 1;
  }

  // finished walk without results
  arrayIndex = template.templateArray.length - 1;
  stringIndex = template.templateArray[arrayIndex].length - 1;
  setEnd(cState, arrayIndex, stringIndex);

  return cState;
};

export { crawl };
