// brian taylor vann
// render

// N Node
// A Attributables

// attribute injection

import { Template } from "./template";
import { Hooks } from "./hooks";

// content injection
interface ContentInjectionParams<N> {
  content: N[];
  leftSentinel?: N;
  rightSentinel?: N;
  parent?: N;
}
interface ContentInjection<N> {
  kind: "CONTENT";
  params: ContentInjectionParams<N>;
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

type Injection<N, A> = ContentInjection<N> | AttributeInjection<N, A>;

// move those into injection map
type InjectionMap<N, A> = Record<string, Injection<N, A>>;

interface Render<N, A> {
  injections: InjectionMap<N, A>;
  siblings: N[];
}

interface RenderStructure<N, A> {
  hooks: Hooks<N, A>;
  template: Template<A>;
  render: Render<N, A>;
  siblings: N[];
  stack: N[];
}

export { Render, RenderStructure };
