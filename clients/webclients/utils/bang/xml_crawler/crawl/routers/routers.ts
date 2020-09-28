// brian taylor vann

type IsAlphabeticCharacter = (char: string) => boolean;

type CrawlStatus =
  | "NOT_FOUND"
  | "OPEN_NODE"
  | "OPEN_NODE_VALID"
  | "OPEN_NODE_CONFIRMED"
  | "CLOSE_NODE"
  | "CLOSE_NODE_VALID"
  | "CLOSE_NODE_CONFIRMED"
  | "INDEPENDENT_NODE"
  | "INDEPENDENT_NODE_VALID"
  | "INDEPENDENT_NODE_CONFIRMED";

type RoutersReducer = (char: string) => CrawlStatus;

const NOT_FOUND = "NOT_FOUND";
const OPEN_NODE = "OPEN_NODE";
const OPEN_NODE_VALID = "OPEN_NODE_VALID";
const OPEN_NODE_CONFIRMED = "OPEN_NODE_CONFIRMED";
const CLOSE_NODE = "CLOSE_NODE";
const CLOSE_NODE_VALID = "CLOSE_NODE_VALID";
const CLOSE_NODE_CONFIRMED = "CLOSE_NODE_CONFIRMED";
const INDEPENDENT_NODE_VALID = "INDEPENDENT_NODE_VALID";
const INDEPENDENT_NODE_CONFIRMED = "INDEPENDENT_NODE_CONFIRMED";

const OPEN_DELIMITER = "<";
const CLOSE_DELIMITER = ">";
const BACKSLASH_DELIMITER = "/";
const ALPHA_CHAR_CODE = "a".charCodeAt(0);
const ZETA_CHAR_CODE = "z".charCodeAt(0);

const isAlphabeticCharacter: IsAlphabeticCharacter = (char) => {
  const charCode = char.charCodeAt(0);
  return ALPHA_CHAR_CODE <= charCode && charCode <= ZETA_CHAR_CODE;
};

const notFound: RoutersReducer = (char) => {
  if (char === OPEN_DELIMITER) {
    return OPEN_NODE;
  }

  return NOT_FOUND;
};

const openNode: RoutersReducer = (char) => {
  if (char === OPEN_DELIMITER) {
    return OPEN_NODE;
  }
  if (char === BACKSLASH_DELIMITER) {
    return CLOSE_NODE;
  }
  if (isAlphabeticCharacter(char)) {
    return OPEN_NODE_VALID;
  }

  return NOT_FOUND;
};

const openNodeValid: RoutersReducer = (char) => {
  if (char === OPEN_DELIMITER) {
    return OPEN_NODE;
  }
  if (char == BACKSLASH_DELIMITER) {
    return INDEPENDENT_NODE_VALID;
  }
  if (char == CLOSE_DELIMITER) {
    return OPEN_NODE_CONFIRMED;
  }

  return OPEN_NODE_VALID;
};

const independentNodeValid: RoutersReducer = (char) => {
  if (char === OPEN_DELIMITER) {
    return OPEN_NODE;
  }
  if (char === CLOSE_DELIMITER) {
    return INDEPENDENT_NODE_CONFIRMED;
  }

  return INDEPENDENT_NODE_VALID;
};

const closeNode: RoutersReducer = (char) => {
  if (char === OPEN_DELIMITER) {
    return OPEN_NODE;
  }
  if (isAlphabeticCharacter(char)) {
    return CLOSE_NODE_VALID;
  }

  return NOT_FOUND;
};

const closeNodeValid: RoutersReducer = (char) => {
  if (char === OPEN_DELIMITER) {
    return OPEN_NODE;
  }
  if (char === CLOSE_DELIMITER) {
    return CLOSE_NODE_CONFIRMED;
  }

  return CLOSE_NODE_VALID;
};

export {
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
};
