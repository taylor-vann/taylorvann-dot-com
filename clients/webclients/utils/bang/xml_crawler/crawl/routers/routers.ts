// brian taylor vann

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

type Routes = Record<string, CrawlStatus>;
type Routers = Partial<Record<CrawlStatus, Routes>>;

type CreateAlphabetKeys = (route: CrawlStatus) => Routes;

const createAlphabetKeys: CreateAlphabetKeys = (route) => {
  const alphabetSet: Routes = {};

  let lowercaseIndex = "a".charCodeAt(0);
  const lowercaseLimit = "z".charCodeAt(0);

  let uppercaseIndex = "A".charCodeAt(0);
  const uppercaseLimit = "Z".charCodeAt(0);

  while (lowercaseIndex <= lowercaseLimit) {
    alphabetSet[String.fromCharCode(lowercaseIndex)] = route;
    lowercaseIndex += 1;
  }

  while (uppercaseIndex <= uppercaseLimit) {
    alphabetSet[String.fromCharCode(uppercaseIndex)] = route;
    uppercaseIndex += 1;
  }

  return alphabetSet;
};

const routers: Routers = {
  NOT_FOUND: {
    "<": "OPEN_NODE",
    DEFAULT: "NOT_FOUND",
  },
  OPEN_NODE: {
    ...createAlphabetKeys("OPEN_NODE_VALID"),
    "<": "OPEN_NODE",
    "/": "CLOSE_NODE",
    DEFAULT: "NOT_FOUND",
  },
  OPEN_NODE_VALID: {
    "<": "OPEN_NODE",
    "/": "INDEPENDENT_NODE_VALID",
    ">": "OPEN_NODE_CONFIRMED",
    DEFAULT: "OPEN_NODE_VALID",
  },
  CLOSE_NODE: {
    ...createAlphabetKeys("CLOSE_NODE_VALID"),
    "<": "OPEN_NODE",
    DEFAULT: "NOT_FOUND",
  },
  CLOSE_NODE_VALID: {
    "<": "OPEN_NODE",
    ">": "CLOSE_NODE_CONFIRMED",
    DEFAULT: "CLOSE_NODE_VALID",
  },
  INDEPENDENT_NODE_VALID: {
    "<": "OPEN_NODE",
    ">": "INDEPENDENT_NODE_CONFIRMED",
    DEFAULT: "INDEPENDENT_NODE_VALID",
  },
};

export { CrawlStatus, routers };
