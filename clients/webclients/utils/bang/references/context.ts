// brian taylor vann

// N Node
// A Attributables

import { RenderResults, StructureRender } from "./render";

type Timestamp = number;
type DescendantRecord = Record<number, Timestamp>;
// might need to rethink context
// we need descendants and we need to remove them

type BangFunc = () => void;
type OnConnectedFunc<R> = (bang: BangFunc) => R;
type OnDisconnectedFunc<R> = (params: R) => void;
type RenderFunc<N, A, P> = (params: P) => RenderResults<N, A>;
interface Structure<N, A, P, R> {
  onConnected: OnConnectedFunc<R>;
  onDisconnected: OnDisconnectedFunc<R>;
  render: RenderFunc<N, A, P>;
}
type StructureFactory<N, A> = <P, R>() => Structure<N, A, P, R>;

interface ContextBase<N, A, P, R> {
  readonly id: number;
  readonly timestamp: Timestamp;
  update(params?: P): N[];
  disconnect(): void;
  getBang(): BangFunc;
}

export {
  BangFunc,
  ContextBase,
  Structure,
  StructureFactory,
  DescendantRecord,
  Timestamp,
};
