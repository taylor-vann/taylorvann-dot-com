// brian taylor vann

import {
  CrawlStatus,
  RoutersReducer,
  NOT_FOUND,
  OPEN_NODE,
  OPEN_NODE_VALID,
  OPEN_NODE_CONFIRMED,
  CLOSE_NODE,
  CLOSE_NODE_VALID,
  CLOSE_NODE_CONFIRMED,
  INDEPENDENT_NODE_VALID,
  INDEPENDENT_NODE_CONFIRMED,
  notFound,
  openNode,
  openNodeValid,
  independentNodeValid,
  closeNode,
  closeNodeValid,
} from "./routers/routers";

type CrawlReducer = Partial<Record<CrawlStatus, RoutersReducer>>;

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
  startPosition?: BrokenTextPostition;
}

type CreateDefaultState = () => CrawlResults;
type Crawl = (params: CrawlParams) => CrawlResults;

const crawlReducer: CrawlReducer = {
  [NOT_FOUND]: notFound,
  [OPEN_NODE]: openNode,
  [OPEN_NODE_VALID]: openNodeValid,
  [CLOSE_NODE]: closeNode,
  [CLOSE_NODE_VALID]: closeNodeValid,
  [INDEPENDENT_NODE_VALID]: independentNodeValid,
};

const validSieve: Sieve = {
  [OPEN_NODE_VALID]: OPEN_NODE_VALID,
  [CLOSE_NODE_VALID]: CLOSE_NODE_VALID,
  [INDEPENDENT_NODE_VALID]: INDEPENDENT_NODE_VALID,
};
const confirmedSieve: Sieve = {
  [OPEN_NODE_CONFIRMED]: OPEN_NODE_CONFIRMED,
  [CLOSE_NODE_CONFIRMED]: CLOSE_NODE_CONFIRMED,
  [INDEPENDENT_NODE_CONFIRMED]: INDEPENDENT_NODE_CONFIRMED,
};

const createDefaultState: CreateDefaultState = () => {
  return {
    nodeType: "NOT_FOUND",
    target: {
      startPosition: {
        stringArrayIndex: 0,
        stringIndex: 0,
      },
      endPosition: {
        stringArrayIndex: 0,
        stringIndex: 0,
      },
    },
  };
};

const setStartStateProperties: Crawl = (params) => {
  const cState = createDefaultState();
  const { startPosition } = params;
  if (startPosition === undefined) {
    return cState;
  }

  const { stringArrayIndex, stringIndex } = startPosition;
  if (cState.target) {
    cState.target.startPosition.stringArrayIndex = stringArrayIndex;
    cState.target.endPosition.stringIndex = stringIndex;
  }

  return cState;
};

const setStartPosition = (
  results: CrawlResults,
  stringArrayIndex: number,
  stringIndex: number
) => {
  results.target.startPosition.stringArrayIndex = stringArrayIndex;
  results.target.startPosition.stringIndex = stringIndex;
};

const setEndPosition = (
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

  let { stringIndex, stringArrayIndex } = cState.target.startPosition;
  while (stringArrayIndex < brokenText.length) {
    const chunk = brokenText[stringArrayIndex];

    if (validSieve[cState.nodeType]) {
      cState.nodeType = NOT_FOUND;
    }

    while (stringIndex < chunk.length) {
      const crawlFunc = crawlReducer[cState.nodeType];
      if (crawlFunc) {
        cState.nodeType = crawlFunc(chunk.charAt(stringIndex));
      }
      if (confirmedSieve[cState.nodeType]) {
        setEndPosition(cState, stringArrayIndex, stringIndex);
        return cState;
      }
      if (cState.nodeType === OPEN_NODE) {
        setStartPosition(cState, stringArrayIndex, stringIndex);
      }

      stringIndex += 1;
    }

    // skip to next chunk
    stringIndex = 0;
    stringArrayIndex += 1;
  }

  // finished walk without results
  return createDefaultState();
};

export { crawl };
