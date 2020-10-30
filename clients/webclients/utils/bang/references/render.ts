// brian taylor vann

// N Node
// A Attributables

interface RenderResults<A> {
  templateArray: TemplateStringsArray;
  injections: A[];
}

interface AttributeInjection<N, A> {
  kind: "ATTRIBUTE";
  node: N;
  name: string;
  value: A;
}
interface StructureInjection<N, A> {
  kind: "STRUCTURE";
  siblings: N[];
  parent: N;
  left?: N;
  right?: N;
}
type Injection<N, A> = AttributeInjection<N, A> | StructureInjection<N, A>;
type Injections<N, A> = Injection<N, A>[];
interface StructureRender<N, A> {
  injections: Injections<N, A>;
  siblings: N[];
}

export { Injection, RenderResults, StructureRender };
