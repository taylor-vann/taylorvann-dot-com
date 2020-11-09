// brian taylor vann
// structure manager

// U Unique Tag names
// N Node
// A Attributables
// P Params

import { Context } from "../context/context";
import { contextManager } from "../context/context_manager";
import {
  InterfaceBase,
  ContextInterface,
  ContextFactoryBase,
} from "../references/context";

class ContextFactory<N, A, P, R> implements ContextFactoryBase<N, A, P, R> {
  private base: InterfaceBase<N, A, P, R>;

  constructor(params: InterfaceBase<N, A, P, R>) {
    this.base = params;
  }

  createContext(params: P) {
    const ctx = new Context<N, A, P, R>();
    const xcontext: ContextInterface<N, A, P, R> = { ctx, base: this.base };
    contextManager.connect({ xcontext, params });

    return xcontext;
  }
}

export { ContextFactory };
