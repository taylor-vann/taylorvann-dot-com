// brian taylor vann

// N Node
// A Attributables

import { RenderResults, StructureRender } from "./render";

type BangFunc = () => void;
type OnConnectedFunc<R> = (bang: BangFunc) => R;
type OnDisconnectedFunc<R> = (params: R) => void;
type RenderFunc<N, A, P> = (params: P) => RenderResults<N, A>;

type Timestamp = number;
type DescendantRecord = Record<number, Timestamp>;
interface Context<N, A, P, R> {
  id: number;
  structureID: number;
  timestamp: number;
  connectedResults: R;
  params: P;
  structureResults: StructureRender<A>;
  renderResults: RenderResults<N, A>;
}

interface Structure {
  id: number;
  timestamp: number;
}
type CreateContext<P> = (params: P) => Structure;
interface ContextParams<N, A, P, R> {
  onConnected: OnConnectedFunc<R>;
  onDisconnected: OnDisconnectedFunc<R>;
  render: RenderFunc<N, A, P>;
}
type ContextFactory<N, A> = <P, R>(
  params: ContextParams<N, A, P, R>
) => CreateContext<P>;

export {
  Context,
  ContextParams,
  CreateContext,
  ContextFactory,
  DescendantRecord,
  Structure,
  Timestamp,
};
