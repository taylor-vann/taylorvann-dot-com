// brian taylor vann
// structure manager

// N Node
// A Attributables
// P Params

// structure manager keeps foundations for "create renders"

import { Context, ContextParams } from "../../references/context";

interface StructureManagerBase<N, A> {
  createStructure: <P, R>(structureID: ContextParams<N, A, P, R>) => number;
  getStructure: <P, R>(stubID: number) => Context<N, A, P, R>;
}

export { StructureManagerBase };
