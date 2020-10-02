type EventMap = Record<string, EventListenerOrEventListenerObject>;
type AttributeMap = Record<string, string>;

interface FoundNode {
  nodeType: "OPEN_NODE" | "INDEPENDENT_NODE";
  nodeTag: string;
  attributes: AttributeMap;
  events: EventMap;
}
interface CloseNode {
  nodeType: "CLOSE_NODE";
  nodeTag: string;
}
interface StringNode {
  nodeType: "STRING_NODE";
  content: string;
}
type CrawlNode = FoundNode | CloseNode | StringNode;
type CrawlNodes = CrawlNode[];

const buildCrawlStringNode = () => {};
const buildFoundNode = () => {};
const buildCloseNode = () => {};
const buildSkeleton = () => {
  // iterate through brokenText and injectsion
  // add string and nodes to results array
  return [];
};

export {
  CrawlNodes,
  buildSkeleton,
  buildCrawlStringNode,
  buildFoundNode,
  buildCloseNode,
};
