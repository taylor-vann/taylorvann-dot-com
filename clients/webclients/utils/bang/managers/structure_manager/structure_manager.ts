// brian taylor vann
// structure manager

// N Node
// A Attributables
// P Params

// structure manager keeps foundations for "create renders"

import { Context, ContextParams } from "../../references/context";

interface StructureManagerBase<A> {
  createStructure: <P, R>(structureID: ContextParams<A, P, R>) => number;
  getStructure: <P, R>(stubID: number) => Context<A, P, R>;
}

export { StructureManagerBase };
