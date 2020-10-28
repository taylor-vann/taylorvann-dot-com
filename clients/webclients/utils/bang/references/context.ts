// brian taylor vann

// N Node
// A Attributables

import { RenderResults } from "./render";

type BangFunc = () => void;
type OnConnectedFunc<R> = (bang: BangFunc) => R;
type OnDisconnectedFunc<R> = (params: R) => void;
type RenderFunc<A, P> = (params: P) => RenderResults<A>;
interface ContextParams<A, P, R> {
  onConnected: OnConnectedFunc<R>;
  onDisconnected: OnDisconnectedFunc<R>;
  render: RenderFunc<A, P>;
}

type Timestamp = number;
type DescendantRecord = Record<number, Timestamp>;
interface Context<A, P, R> {
  id: number;
  structureID: number;
  timestamp: number;
  renderResults: RenderResults<A>;
  connectedResults: R;
  params: P;
  descendants: DescendantRecord;
}

interface Structure {
  id: number;
  timestamp: number;
}
type CreateContext<P> = (params: P) => Structure;
type ContextFactory<A> = <P, R>(
  params: ContextParams<A, P, R>
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
