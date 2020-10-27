// brian taylor vann

// N Node
// A Attributables

type Timestamp = number;
type DescendantRecord = Record<number, Timestamp>;

interface RenderResults<A> {
  templateArray: TemplateStringsArray;
  injections: A[];
}

type BangFunc = () => void;
type OnConnectedFunc<R> = () => R;
type OnDisconnectedFunc<R> = (params: R) => void;
type RenderFunc<A, P> = (params: P, bang: BangFunc) => RenderResults<A>;

interface ContextParams<A, P, R> {
  onConnected: OnConnectedFunc<R>;
  onDisconnected: OnDisconnectedFunc<R>;
  render: RenderFunc<A, P>;
}

interface Context<A, P, R> {
  onConnected: () => R;
  onDisconnected: (params: R) => void;
  renderStructure: (params: P) => void;
  id: number;
  connectedResults: P;
  descendants: DescendantRecord;
  params: P;
  renderResults: RenderResults<A>;
}

interface Structure {
  id: number;
}
type CreateContext<A, P, R> = (params: P) => Context<A, P, R>;
type ContextFactory<A> = <P, R>(
  params: ContextParams<A, P, R>
) => CreateContext<A, P, R>;

export {
  Context,
  ContextParams,
  CreateContext,
  ContextFactory,
  DescendantRecord,
  RenderResults,
  Structure,
  Timestamp,
};
