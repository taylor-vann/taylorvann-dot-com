// brian taylor vann
// skeleton routers

import { CrawlStatus } from "../../type_flyweight/skeleton_crawl";

type Routes = Record<string, CrawlStatus>;
type Routers = Partial<Record<CrawlStatus, Routes>>;

const routers: Routers = {
  CONTENT_NODE: {
    "<": "OPEN_NODE",
    DEFAULT: "CONTENT_NODE",
  },
  OPEN_NODE: {
    " ": "CONTENT_NODE",
    "\n": "CONTENT_NODE",
    "<": "OPEN_NODE",
    "/": "CLOSE_NODE",
    DEFAULT: "OPEN_NODE_VALID",
  },
  OPEN_NODE_VALID: {
    "<": "OPEN_NODE",
    "/": "SELF_CLOSING_NODE_VALID",
    ">": "OPEN_NODE_CONFIRMED",
    DEFAULT: "OPEN_NODE_VALID",
  },
  CLOSE_NODE: {
    " ": "CONTENT_NODE",
    "\n": "CONTENT_NODE",
    "<": "OPEN_NODE",
    DEFAULT: "CLOSE_NODE_VALID",
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
