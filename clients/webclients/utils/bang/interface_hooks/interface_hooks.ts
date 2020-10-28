// brian taylor vann

// N Node
// A Attributables
// P Params

// more render results down here

import { RenderResults } from "../references/render";
import { CrawlResults } from "../builders/build_skeleton/crawl/crawl";
import {
  Attributes,
  OpenParseResults,
  ParseResults,
} from "../references/parse";

interface ParseNodeParams<A> {
  renderResults: RenderResults<A>;
  crawlResults: CrawlResults;
}
type ParseNode<A> = (params: ParseNodeParams<A>) => ParseResults<A>;
type CreateNode<N, A> = (params: OpenParseResults<A>) => N;
type CreateContentNode<N> = (content: string) => N;
type AddDescendent<N> = (element: N, descendent: N) => N;
type RemoveDescendent<N> = (element: N, descendent: N) => N;

// Use this to create new Bang Interfaces
interface InterfaceHooks<N, A> {
  parseNode: ParseNode<A>;
  createNode: CreateNode<N, A>;
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
