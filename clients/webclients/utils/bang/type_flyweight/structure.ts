// brian taylor vann
// structure

import { RenderResults } from "./render";

interface StructureRender<A> {
  templateArray: TemplateStringsArray;
  injections: A[];
}

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

export { Structure, StructureFactory, StructureRender };
