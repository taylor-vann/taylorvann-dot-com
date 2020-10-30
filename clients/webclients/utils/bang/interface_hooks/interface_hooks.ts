// brian taylor vann

// N Node
// A Attributables

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
interface AddSiblingsParams<N> {
  siblings: N[];
  parent: N;
  leftSibling?: N;
  rightSibling?: N;
}
type AddSiblings<N> = (params: AddSiblingsParams<N>) => N[];
type RemoveSiblingsParams<N> = AddSiblingsParams<N>;
type RemoveSiblings<N> = (params: RemoveSiblingsParams<N>) => void;

// Use this to create new Bang Interfaces
interface InterfaceHooks<N, A> {
  parseNode: ParseNode<A>;
  createNode: CreateNode<N, A>;
  createContentNode: CreateContentNode<N>;
  addSiblings: AddSiblings<N>;
  removeSiblings: RemoveSiblings<N>;
}

export {
  AddSiblings,
  Attributes,
  CreateContentNode,
  CreateNode,
  InterfaceHooks,
  ParseNode,
  ParseResults,
  RemoveSiblings,
};
