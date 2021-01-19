// brian taylor vann
// bang - document builder (build a node graph)

// N Node
// A Attributables

import { Hooks } from "./type_flyweight/hooks";

class Bang<N, A> {
  private hooks: Hooks<N, A>;

  constructor(interfaceHooks: Hooks<N, A>) {
    this.hooks = interfaceHooks;
  }
}

export { Bang };
