// brian taylor vann
// bang - document builder (build a node graph)

// N Node
// A Attributables

import { Structure } from "./references/context";
import { InterfaceHooks } from "./interface_hooks/interface_hooks";
import { ContextFactory } from "./context_factory/context_factory";

interface BangBase<N, A> {
  createContextFactory<P, R>(
    structure: Structure<N, A, P, R>
  ): ContextFactory<N, A, P, R>;
}

class Bang<N, A> implements BangBase<N, A> {
  private hooks: InterfaceHooks<N, A>;

  constructor(interfaceHooks: InterfaceHooks<N, A>) {
    this.hooks = interfaceHooks;
  }

  createContextFactory<P, R>(structure: Structure<N, A, P, R>) {
    const contextFactory = new ContextFactory({ hooks: this.hooks, structure });
    return contextFactory;
  }
}

export { Bang };
