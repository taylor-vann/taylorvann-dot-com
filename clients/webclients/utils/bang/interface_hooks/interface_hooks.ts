// brian taylor vann

// N Node
// A Attributables
// P Params

import { RenderResults } from "../references/context";
import { CrawlResults } from "../builders/build_skeleton/crawl/crawl";

interface CloseParseResults {
  tag: string;
  kind: "CLOSE_NODE";
}
type Attributes<A> = A[];
interface NodeParseResults<A> {
  tag: string;
  kind: "OPEN_NODE" | "INDEPENDENT_NODE" | "CONTENT_NODE";
  attributes: Attributes<A>;
}
type ParseResults<A> = NodeParseResults<A> | CloseParseResults;
interface ParseNodeParams<A> {
  renderResults: RenderResults<A>;
  crawlResults: CrawlResults;
}
type ParseNode<A> = (params: ParseNodeParams<A>) => ParseResults<A>;
type CreateNode<P, N> = (params: P) => N;
type CreateContentNode<N> = (content: string) => N;
type AddDescendent<N> = (element: N, descendent: N) => N;
type RemoveDescendent<N> = (element: N, descendent: N) => N;

// Use this to create new Bang Interfaces
interface InterfaceHooks<N, A> {
  parseNode: ParseNode<A>;
  createNode: CreateNode<ParseResults<A>, N>;
  createContentNode: CreateContentNode<N>;
  addDescendent: AddDescendent<N>;
  removeDescendent: RemoveDescendent<N>;
}

export {
  AddDescendent,
  Attributes,
  CreateContentNode,
  CreateNode,
  InterfaceHooks,
  ParseNode,
  ParseResults,
  RemoveDescendent,
};
