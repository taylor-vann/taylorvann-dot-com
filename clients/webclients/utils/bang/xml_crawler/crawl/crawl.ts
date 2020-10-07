// brian taylor vann

import { CrawlStatus, routers } from "./routers/routers";

type Sieve = Partial<Record<CrawlStatus, CrawlStatus>>;
interface BrokenTextPostition {
  arrayIndex: number;
  stringIndex: number;
}
interface BrokenTextVector {
  start: BrokenTextPostition;
  end: BrokenTextPostition;
}
interface CrawlResults {
  nodeType: CrawlStatus;
  target: BrokenTextVector;
}

type SetPosition = (
  results: CrawlResults,
  arrayIndex: number,
  stringIndex: number
) => void;

type CreateCrawlState = () => CrawlResults;
type SetNodeType = (results: CrawlResults, character: string) => void;

type SetStartStateProperties = (
  brokenText: TemplateStringsArray,
  previousCrawl?: CrawlResults
) => CrawlResults | undefined;

type Crawl = (
  brokenText: TemplateStringsArray,
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

const createNotFoundCrawlState: CreateCrawlState = () => {
  return {
    nodeType: "CONTENT_NODE",
    target: {
      start: {
        arrayIndex: 0,
        stringIndex: 0,
      },
      end: {
        arrayIndex: 0,
        stringIndex: 0,
      },
    },
  };
};

const setStartStateProperties: SetStartStateProperties = (
  brokenText,
  previousCrawl
) => {
  const cState = createNotFoundCrawlState();
  if (previousCrawl === undefined) {
    return cState;
  }
  
  let { arrayIndex, stringIndex } = previousCrawl.target.end;

  stringIndex += 1;
  stringIndex %= brokenText[arrayIndex].length;
  if (stringIndex === 0) {
    arrayIndex += 1;
  }
  if (arrayIndex >= brokenText.length) {
    return;
  }

  cState.target.start.arrayIndex = arrayIndex;
  cState.target.start.stringIndex = stringIndex;
  cState.target.end.arrayIndex = arrayIndex;
  cState.target.end.stringIndex = stringIndex;

  return cState;
};

const setNodeType: SetNodeType = (cState, char) => {
  const defaultNodeType = routers[cState.nodeType]?.["DEFAULT"] ?? CONTENT_NODE;
  cState.nodeType = routers[cState.nodeType]?.[char] ?? defaultNodeType;

  return cState;
};

const setStart: SetPosition = (
  results: CrawlResults,
  arrayIndex: number,
  stringIndex: number
) => {
  results.target.start.arrayIndex = arrayIndex;
  results.target.start.stringIndex = stringIndex;
  results.target.end.arrayIndex = arrayIndex;
  results.target.end.stringIndex = stringIndex;
};

const setEnd: SetPosition = (
  results: CrawlResults,
  arrayIndex: number,
  stringIndex: number
) => {
  results.target.end.arrayIndex = arrayIndex;
  results.target.end.stringIndex = stringIndex;
};

const crawl: Crawl = (brokenText, previousCrawl) => {
  const cState = setStartStateProperties(brokenText, previousCrawl);
  if (cState === undefined) {
    return;
  }

  let { stringIndex, arrayIndex } = cState.target.start;
  // retain most recent postition
  const suspect: BrokenTextPostition = {
    arrayIndex,
    stringIndex,
  };

  while (arrayIndex < brokenText.length) {
    if (validSieve[cState.nodeType] === undefined) {
      cState.nodeType = CONTENT_NODE;
    }

    const chunk = brokenText[arrayIndex];
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
  arrayIndex = brokenText.length - 1;
  stringIndex = brokenText[arrayIndex].length - 1;
  setEnd(cState, arrayIndex, stringIndex);

  return cState;
};

export { BrokenTextVector, BrokenTextPostition, CrawlResults, crawl };
