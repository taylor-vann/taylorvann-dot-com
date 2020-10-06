// brian taylor vann

import { CrawlStatus, routers } from "./routers/routers";

type Sieve = Partial<Record<CrawlStatus, CrawlStatus>>;
interface BrokenTextPostition {
  stringArrayIndex: number;
  stringIndex: number;
}
interface BrokenTextVector {
  startPosition: BrokenTextPostition;
  endPosition: BrokenTextPostition;
}
interface CrawlResults {
  nodeType: CrawlStatus;
  target: BrokenTextVector;
}
interface CrawlParams {
  brokenText: TemplateStringsArray;
  previousCrawl?: CrawlResults;
  startPosition?: BrokenTextPostition;
}

type SetPosition = (
  results: CrawlResults,
  stringArrayIndex: number,
  stringIndex: number
) => void;

type CreateDefaultTarget = () => BrokenTextVector;
type CreateCrawlState = () => CrawlResults;
type SetNodeType = (results: CrawlResults, character: string) => void;
type SetStartStateProperties = (
  params: CrawlParams
) => CrawlResults | undefined;

type Crawl = (params: CrawlParams) => CrawlResults | undefined;

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

const createDefaultTarget: CreateDefaultTarget = () => {
  return {
    startPosition: {
      stringArrayIndex: 0,
      stringIndex: 0,
    },
    endPosition: {
      stringArrayIndex: 0,
      stringIndex: 0,
    },
  };
};

const createNotFoundCrawlState: CreateCrawlState = () => {
  return {
    nodeType: "CONTENT_NODE",
    target: createDefaultTarget(),
  };
};

const setStartStateProperties: SetStartStateProperties = (params) => {
  const cState = createNotFoundCrawlState();
  if (params.previousCrawl) {
    console.log("previous crawl");
    console.log(params.previousCrawl);
    let {
      stringArrayIndex,
      stringIndex,
    } = params.previousCrawl.target.endPosition;

    stringIndex += 1;
    stringIndex %= params.brokenText[stringArrayIndex].length;

    if (stringIndex === 0) {
      stringArrayIndex += 1;
    }
    if (stringArrayIndex >= params.brokenText.length) {
      return;
    }

    cState.target.startPosition.stringArrayIndex = stringArrayIndex;
    cState.target.startPosition.stringIndex = stringIndex;
    cState.target.endPosition.stringArrayIndex = stringArrayIndex;
    cState.target.endPosition.stringIndex = stringIndex;
  }

  return cState;
};

const setNodeType: SetNodeType = (cState, char) => {
  const defaultNodeType = routers[cState.nodeType]?.["DEFAULT"] ?? CONTENT_NODE;
  cState.nodeType = routers[cState.nodeType]?.[char] ?? defaultNodeType;

  return cState;
};

const setStartPosition: SetPosition = (
  results: CrawlResults,
  stringArrayIndex: number,
  stringIndex: number
) => {
  results.target.startPosition.stringArrayIndex = stringArrayIndex;
  results.target.startPosition.stringIndex = stringIndex;
  results.target.endPosition.stringArrayIndex = stringArrayIndex;
  results.target.endPosition.stringIndex = stringIndex;
};

const setEndPosition: SetPosition = (
  results: CrawlResults,
  stringArrayIndex: number,
  stringIndex: number
) => {
  results.target.endPosition.stringArrayIndex = stringArrayIndex;
  results.target.endPosition.stringIndex = stringIndex;
};

const crawl: Crawl = (params) => {
  const { brokenText } = params;
  const cState = setStartStateProperties(params);
  if (cState === undefined) {
    return;
  }

  let { stringIndex, stringArrayIndex } = cState.target.startPosition;
  const potentialPosition: BrokenTextPostition = {
    stringArrayIndex,
    stringIndex,
  };

  while (stringArrayIndex < brokenText.length) {
    if (validSieve[cState.nodeType] === undefined) {
      cState.nodeType = CONTENT_NODE;
    }

    const chunk = brokenText[stringArrayIndex];
    while (stringIndex < chunk.length) {
      setNodeType(cState, chunk.charAt(stringIndex));

      if (confirmedSieve[cState.nodeType]) {
        setStartPosition(
          cState,
          potentialPosition.stringArrayIndex,
          potentialPosition.stringIndex
        );
        setEndPosition(cState, stringArrayIndex, stringIndex);
        return cState;
      }

      if (cState.nodeType === OPEN_NODE) {
        potentialPosition.stringArrayIndex = stringArrayIndex;
        potentialPosition.stringIndex = stringIndex;
      }

      stringIndex += 1;
    }

    // skip to next chunk
    stringIndex = 0;
    stringArrayIndex += 1;
  }

  // finished walk without results
  stringArrayIndex = brokenText.length - 1;
  stringIndex = brokenText[stringArrayIndex].length - 1;
  setEndPosition(cState, stringArrayIndex, stringIndex);

  return cState;
};

export { BrokenTextVector, CrawlResults, crawl, createNotFoundCrawlState };
