// brian taylor vann
// crawl

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

interface Position {
  arrayIndex: number;
  stringIndex: number;
}
interface Vector {
  start: Position;
  end: Position;
}
interface CrawlResults {
  nodeType: CrawlStatus;
  target: Vector;
}

type SkeletonNodes = CrawlResults[];

export { CrawlResults, CrawlStatus, Position, SkeletonNodes, Vector };
