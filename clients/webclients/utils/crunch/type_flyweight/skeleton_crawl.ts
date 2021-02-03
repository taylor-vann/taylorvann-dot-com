// brian taylor vann
// crawl

import { Vector } from "./text_vector";

type CrawlStatus =
  | "CONTENT_NODE"
  | "OPEN_NODE"
  | "OPEN_NODE_VALID"
  | "OPEN_NODE_CONFIRMED"
  | "CLOSE_NODE"
  | "CLOSE_NODE_VALID"
  | "CLOSE_NODE_CONFIRMED"
  | "SELF_CLOSING_NODE"
  | "SELF_CLOSING_NODE_VALID"
  | "SELF_CLOSING_NODE_CONFIRMED";

interface CrawlResults {
  nodeType: CrawlStatus;
  vector: Vector;
}

type SkeletonNodes = CrawlResults[];

export { CrawlResults, CrawlStatus, SkeletonNodes };
