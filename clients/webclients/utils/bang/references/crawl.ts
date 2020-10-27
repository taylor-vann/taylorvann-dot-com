// brian taylor vann

type CrawlStatus =
  | "CONTENT_NODE"
  | "OPEN_NODE"
  | "OPEN_NODE_VALID"
  | "OPEN_NODE_CONFIRMED"
  | "CLOSE_NODE"
  | "CLOSE_NODE_VALID"
  | "CLOSE_NODE_CONFIRMED"
  | "INDEPENDENT_NODE"
  | "INDEPENDENT_NODE_VALID"
  | "INDEPENDENT_NODE_CONFIRMED";

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
