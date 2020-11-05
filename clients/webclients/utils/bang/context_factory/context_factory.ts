// brian taylor vann
// structure manager

import { Context } from "../context/context";
import { InterfaceHooks } from "../interface_hooks/interface_hooks";
import { Structure } from "../references/context";

class ContextFactory<N, A, P, R> {
  private hooks: InterfaceHooks<N, A>;
  private structure: Structure<N, A, P, R>;

  constructor(
    interfaceHooks: InterfaceHooks<N, A>,
    structure: Structure<N, A, P, R>
  ) {
    this.hooks = interfaceHooks;
    this.structure = structure;
  }

  createContext(params: P) {
    const ctx = new Context(this.hooks, this.structure);
    ctx.update(params);

    return ctx;
  }
}

export { ContextFactory };
