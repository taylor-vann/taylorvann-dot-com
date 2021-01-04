// brian taylor vann
// context

// U Unique Tag names
// N Node
// A Attributables
// P Params

import { InterfaceHooks } from "../interface_hooks/interface_hooks";
import { RenderResults, StructureRender } from "./render";
import { Structure } from "./structure";

type Timestamp = number;
type DescendantRecord = Record<number, Timestamp>;
// might need to rethink context
// we need descendants and we need to remove them

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
  ContextBase,
  ContextFactoryBase,
  ContextInterface,
  InterfaceBase,
  DescendantRecord,
  Timestamp,
};
