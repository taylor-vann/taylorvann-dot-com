// brian talor vann

// Bang context manager
// logic_structures return contexts

// the context manager is a key value storage of a Context
// coupled with a publish subscribe model

// C context
// P params

import { LogicStructure } from "../logic_structures/logic_structures";

interface BangContextBase<P> {
  logicStructure: number;
  lifecycle: number;
  params: number;
  descendants: Record<number, number>;
}

interface BangContextManagerBase<C> {
  addContext: (functor: LogicStructure<C>) => number;
  removeContext: (stubID: number) => void;
}

export { BangContextBase, BangContextManagerBase };
