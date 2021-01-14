// brian taylor vann

import { CrawlStatus } from "../../type_flyweight/crawl";

type Routes = Record<string, CrawlStatus>;
type Routers = Partial<Record<CrawlStatus, Routes>>;
type CreateAlphabetKeys = (route: CrawlStatus) => Routes;

const SPECIAL_CHARACTERS: string[] = ["_", "-", ".", ":"];

const createAlphaNumericKeys: CreateAlphabetKeys = (route) => {
  const alphabetSet: Routes = {};
  const lowercaseLimit = "z".charCodeAt(0);
  const uppercaseLimit = "Z".charCodeAt(0);

  // add letters to seive
  let lowercaseIndex = "a".charCodeAt(0);
  let uppercaseIndex = "A".charCodeAt(0);
  while (lowercaseIndex <= lowercaseLimit) {
    alphabetSet[String.fromCharCode(lowercaseIndex)] = route;
    lowercaseIndex += 1;
  }

  while (uppercaseIndex <= uppercaseLimit) {
    alphabetSet[String.fromCharCode(uppercaseIndex)] = route;
    uppercaseIndex += 1;
  }

  // add numbers
  let numericKey = 0;
  while (numericKey < 10) {
    alphabetSet[numericKey] = route;

    numericKey += 1;
  }

  // add special characters
  for (const specialChar of SPECIAL_CHARACTERS) {
    alphabetSet[specialChar] = route;
  }

  return alphabetSet;
};

const routers: Routers = {
  CONTENT_NODE: {
    "<": "OPEN_NODE",
    DEFAULT: "CONTENT_NODE",
  },
  OPEN_NODE: {
    ...createAlphaNumericKeys("OPEN_NODE_VALID"),
    "<": "OPEN_NODE",
    "/": "CLOSE_NODE",
    DEFAULT: "CONTENT_NODE",
  },
  OPEN_NODE_VALID: {
    "<": "OPEN_NODE",
    "/": "SELF_CLOSING_NODE_VALID",
    ">": "OPEN_NODE_CONFIRMED",
    DEFAULT: "OPEN_NODE_VALID",
  },
  CLOSE_NODE: {
    ...createAlphaNumericKeys("CLOSE_NODE_VALID"),
    "<": "OPEN_NODE",
    DEFAULT: "CONTENT_NODE",
  },
  CLOSE_NODE_VALID: {
    "<": "OPEN_NODE",
    ">": "CLOSE_NODE_CONFIRMED",
    DEFAULT: "CLOSE_NODE_VALID",
  },
  SELF_CLOSING_NODE_VALID: {
    "<": "OPEN_NODE",
    ">": "SELF_CLOSING_NODE_CONFIRMED",
    DEFAULT: "SELF_CLOSING_NODE_VALID",
  },
};

export { CrawlStatus, routers };
