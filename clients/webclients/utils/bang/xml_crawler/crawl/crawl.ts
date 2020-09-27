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

const SPACE_DELIMITER = " ";
const OPEN_DELIMITER = "<";
const CLOSE_DELIMITER = ">";
const ALPHA_CHAR_CODE = "a".charCodeAt(0);
const ZETA_CHAR_CODE = "z".charCodeAt(0);
const BACKSLASH_CHAR_CODE = "/".charCodeAt(0);

type IsAlphabeticCharacter = (targetStr: string, index?: number) => boolean;
type Crawl = (params: CrawlParams) => TargetTextVector;

const isAlphabeticCharacter: IsAlphabeticCharacter = (targetStr, index = 0) => {
  const charCode = targetStr.charCodeAt(index);
  return ALPHA_CHAR_CODE <= charCode && charCode <= ZETA_CHAR_CODE;
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
