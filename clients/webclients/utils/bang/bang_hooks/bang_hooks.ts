// brian taylor vann

// U Unique Tag names
// N Node
// A Attributables
// E Event Listeners

// P Params = U A E

import { CrawlResults } from "../crawl/crawl";

interface AttributeBase<K, V> {
  kind: K;
  name: string;
  value: V;
}
type StaticAttribute<R> = AttributeBase<"static", R>;
type EventAttribute<R> = AttributeBase<"event", R>;
type Attribute<A, E> = StaticAttribute<A> | EventAttribute<E>;
type Attributes<A, E> = Attribute<A, E>[];

interface CloseParseResults<U> {
  tag: U;
  kind: "CLOSE_NODE";
}
interface NodeParseResults<U, A, E> {
  tag: U;
  kind: "OPEN_NODE" | "INDEPENDENT_NODE" | "CONTENT_NODE";
  attributes: Attributes<A, E>;
}
type ParseResults<U, A, E> = NodeParseResults<U, A, E> | CloseParseResults<U>;
type Injections<A, E> = (A | E)[];
interface ParseNodeParams<A, E> {
  templateStrings: TemplateStringsArray;
  injections: Injections<A, E>;
  crawlResults: CrawlResults;
}
type ParseNode<U, A, E> = (
  params: ParseNodeParams<A, E>
) => ParseResults<U, A, E>;
type CreateNode<P, N> = (params: P) => N;
type CreateContentNode<N> = (content: string) => N;
type AddDescendent<N> = (element: N, descendent: N) => N;
type RemoveDescendent<N> = (element: N, descendent: N) => N;
interface Hooks<U, N, A, E> {
  parseNode: ParseNode<U, A, E>;
  createNode: CreateNode<ParseResults<U, A, E>, N>;
  createContentNode: CreateContentNode<N>;
  addDescendent: AddDescendent<N>;
  removeDescendent: RemoveDescendent<N>;
}

export {
  StaticAttribute,
  EventAttribute,
  Attribute,
  Attributes,
  ParseResults,
  ParseNode,
  CreateNode,
  CreateContentNode,
  AddDescendent,
  RemoveDescendent,
  Hooks,
};
