// brian taylor vann
// render

// N Node
// A Attributables

interface StructureRender<A> {
  templateArray: TemplateStringsArray;
  injections: A[];
}

interface AttributeInjection<N, A> {
  kind: "ATTRIBUTE";
  node: N;
  name: string;
  value: A;
}
interface StructureInjection<N> {
  kind: "STRUCTURE";
  siblings: N[];
  parent: N;
  left?: N;
  right?: N;
}

// maybe ?? // not needed?
interface ContentInjection<N> {
  kind: "CONTENT";
  siblings: N[];
  parent: N;
  left?: N;
  right?: N; // possibly not needed
}
type Injection<N, A> = AttributeInjection<N, A> | StructureInjection<N>;
type Injections<N, A> = Injection<N, A>[];
interface RenderResults<N, A> {
  injections: Injections<N, A>;
  siblings: N[];
}

export { Injection, RenderResults, StructureRender };
