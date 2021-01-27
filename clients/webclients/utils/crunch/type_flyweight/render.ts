// brian taylor vann
// render

// N Node
// A Attributables

// attribute injection

import { Template } from "./template";
import { Hooks } from "./hooks";

// content injection
interface ContentInjectionParams<N> {
  textNode: N;
  left?: N;
  parent?: N;
}
interface ContentInjection<N> {
  kind: "CONTENT";
  params: ContentInjectionParams<N>;
}

interface ContextInjectionParams<N> {
  siblings: N[];
  left?: N;
  parent?: N;
}
interface ContextInjection<N> {
  kind: "CONTEXT";
  params: ContextInjectionParams<N>;
}

interface AttributeInjectionParams<N, A> {
  node: N;
  attribute: string;
  value: A;
}
interface AttributeInjection<N, A> {
  kind: "ATTRIBUTE";
  params: AttributeInjectionParams<N, A>;
}

type Injection<N, A> =
  | ContentInjection<N>
  | ContextInjection<N>
  | AttributeInjection<N, A>;

interface ElementNode<N> {
  kind: "NODE";
  tagName: string;
  node: N;
  selfClosing: boolean;
}

interface TextNode<N> {
  kind: "TEXT";
  node: N;
}

type NodeBit<N> = ElementNode<N> | TextNode<N>;

// move those into injection map
type InjectionMap<N, A> = Record<number, Injection<N, A>>;

interface RenderStructure<N, A> {
  hooks: Hooks<N, A>;
  template: Template<A>;
  injections: InjectionMap<N, A>;
  siblings: N[];
  stack: NodeBit<N>[];
}

export { RenderStructure, Injection, NodeBit, ElementNode, TextNode };
