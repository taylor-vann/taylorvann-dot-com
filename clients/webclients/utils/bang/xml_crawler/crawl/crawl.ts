// brian taylor vann

import { KeywordTypeNode } from "typescript";

interface BrokenTextPostition {
  stringArrayIndex: number;
  stringIndex: number;
}
interface CrawlParams {
  brokenText: string[];
  startPosition?: BrokenTextPostition;
}

type NodeType = "OPEN_NODE" | "CLOSED_NODE" | "INDEPENDENT_NODE";

type CrawlStatus =
  | "OPEN_NODE"
  | "OPEN_NODE_VALID"
  | "OPEN_NODE_CONFIRMED"
  | "CLOSED_NODE"
  | "CLOSED_NODE_VALID"
  | "CLOSED_NODE_CONFIRMED"
  | "INDEPENDENT_NODE"
  | "INDEPENDENT_NODE_CONFIRMED";

interface TargetTextVector {
  start: BrokenTextPostition;
  end: BrokenTextPostition;
}

interface CrawlResults {
  type?: NodeType;
  target: TargetTextVector;
}

const UNKNOWN = "UNKNOWN";
const OPEN_NODE = "OPEN_NODE";
const OPEN_NODE_VALID = "OPEN_NODE_VALID";
const OPEN_NODE_CONFIRMED = "OPEN_NODE_CONFIRMED";
const CLOSED_NODE = "CLOSED_NODE";
const CLOSED_NODE_VALID = "CLOSED_NODE_VALID";
const CLOSED_NODE_CONFIRMED = "CLOSED_NODE_CONFIRMED";

const SPACE_DELIMITER = " ";
const OPEN_DELIMITER = "<";
const CLOSE_DELIMITER = ">";
const BACKSLASH_CHAR = "/";
const ALPHA_CHAR_CODE = "a".charCodeAt(0);
const ZETA_CHAR_CODE = "z".charCodeAt(0);

type IsAlphabeticCharacter = (char: string) => boolean;
type Crawl = (params: CrawlParams) => TargetTextVector;

const isAlphabeticCharacter: IsAlphabeticCharacter = (char: string) => {
  const charCode = char.charCodeAt(0);
  return ALPHA_CHAR_CODE <= charCode && charCode <= ZETA_CHAR_CODE;
};

const unknownFunc = (char: string) => {
  if (char == OPEN_DELIMITER) {
    return OPEN_NODE;
  }

  return UNKNOWN;
};

const openNodeFunc = (char: string) => {
  if (char == OPEN_DELIMITER) {
    return OPEN_NODE;
  }
  if (isAlphabeticCharacter(char)) {
    return OPEN_NODE_VALID;
  }
  if (char == BACKSLASH_CHAR) {
    return CLOSED_NODE;
  }

  return UNKNOWN;
};

const openNodeValidFunc = (char: string) => {
  if (char == OPEN_DELIMITER) {
    return OPEN_NODE;
  }
  if (char == CLOSE_DELIMITER) {
    return OPEN_NODE_CONFIRMED;
  }

  return OPEN_NODE_VALID;
};

const closedNodeFunc = (char: string) => {
  if (char == OPEN_DELIMITER) {
    return OPEN_NODE;
  }
  if (isAlphabeticCharacter(char)) {
    return CLOSED_NODE_VALID;
  }

  return UNKNOWN;
};

const closedNodeValidFunc = (char: string) => {
  if (char == OPEN_DELIMITER) {
    return OPEN_NODE;
  }
  if (char === CLOSE_DELIMITER) {
    return CLOSED_NODE_CONFIRMED;
  }

  return CLOSED_NODE_VALID;
};

// use GRAPHS to decide what node

// unknown -> "<" -> open_node

// open_node -> " "   -> unknown
//           -> "a-z" -> open_node_valid
//           -> "/"   -> closed_node

// closed_node -> "a-z" -> closed_node_valid
//             -> " "   -> unknown

// open_node_valid -> "/" -> independent_node
//                 -> "<" -> open_node
//                 -> ">" -> open_node_confired

// closed_node_valid -> "<" -> open_node
//                   -> ">" -> closed_node_confirmed

// independent_node -> "/" -> independent_node
//                  -> ">" -> open_node_confired

const crawl = (params: CrawlParams) => {
  const { brokenText, startPosition } = params;

  let stringArrayIndex = startPosition?.stringArrayIndex ?? 0;
  let stringIndex = startPosition?.stringIndex ?? 0;

  while (stringArrayIndex < brokenText.length) {
    const chunk = brokenText[stringArrayIndex];
    while (stringIndex < chunk.length) {
      // if <, update most recent < index

      // if >

      stringIndex += 1;
    }
    stringArrayIndex += 1;
  }
  // iterate from  starting point
  // over every word
  // over every letter
  //   -> look for the most recent < and the first >
  //   -> must start with a letter or /

  // two whiles
};

export { crawl };
