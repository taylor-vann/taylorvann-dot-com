// brian taylor vann

// bang - stateless document builder (build a node graph)

// N Node
// A Attributables

import { Structure } from "./references/context";
import { InterfaceHooks } from "./interface_hooks/interface_hooks";
import { ContextFactory } from "./context_factory/context_factory";
import { Context } from "../bang/context/context";

type ContextFunc<N, A, P, R> = (params: P) => Context<N, A, P, R>;

interface BangBase<N, A> {
  createContextFactory<P, R>(
    structure: Structure<N, A, P, R>
  ): ContextFunc<N, A, P, R>;
}

class Bang<N, A> implements BangBase<N, A> {
  private hooks: InterfaceHooks<N, A>;

  constructor(interfaceHooks: InterfaceHooks<N, A>) {
    this.hooks = interfaceHooks;
  }

  createContextFactory<P, R>(structure: Structure<N, A, P, R>) {
    const contextFactory = new ContextFactory(this.hooks, structure);
    const createContext: ContextFunc<N, A, P, R> = (params: P) =>
      contextFactory.createContext(params);

    return createContext;
  }
}

export { Bang };
