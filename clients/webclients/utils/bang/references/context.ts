// brian taylor vann
// context

// U Unique Tag names
// N Node
// A Attributables
// P Params

import { Context } from "../context/context";
import { InterfaceHooks } from "../interface_hooks/interface_hooks";
import { RenderResults, StructureRender } from "./render";

type Timestamp = number;
type DescendantRecord = Record<number, Timestamp>;
// might need to rethink context
// we need descendants and we need to remove them

type BangFunc = () => void;
interface OnConnectedParams<P> {
  params: P;
  bang: BangFunc;
}
type OnConnectedFunc<P, R> = (params: OnConnectedParams<P>) => R;
type OnDisconnectedFunc<R> = (params: R) => void;
type RenderFunc<N, A, P> = (params: P) => RenderResults<N, A>;
interface Structure<N, A, P, R> {
  onConnected: OnConnectedFunc<P, R>;
  onDisconnected: OnDisconnectedFunc<R>;
  render: RenderFunc<N, A, P>;
}
type StructureFactory<N, A> = <P, R>() => Structure<N, A, P, R>;

interface InterfaceBase<N, A, P, R> {
  hooks: InterfaceHooks<N, A>;
  structure: Structure<N, A, P, R>;
}
interface ContextBase<N, A, P, R> {
  id: number;
  timestamp: Timestamp;
  params?: P;
  gambit?: R;
  structureResults?: StructureRender<A>;
  renderResults?: RenderResults<N, A>;
}
interface ContextInterface<N, A, P, R> {
  ctx: ContextBase<N, A, P, R>;
  base: InterfaceBase<N, A, P, R>;
}

interface ContextFactoryBase<N, A, P, R> {
  createContext(params: P): ContextInterface<N, A, P, R>;
}

export {
  BangFunc,
  ContextBase,
  ContextFactoryBase,
  ContextInterface,
  InterfaceBase,
  Structure,
  StructureFactory,
  DescendantRecord,
  Timestamp,
};
