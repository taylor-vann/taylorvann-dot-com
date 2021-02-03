// brian taylor vann
// crunch - document builder (create unidirectional chunks)

// N Node
// A Attributables

import { Hooks } from "./type_flyweight/hooks";

class Crunch<N, A> {
  private hooks: Hooks<N, A>;

  constructor(interfaceHooks: Hooks<N, A>) {
    this.hooks = interfaceHooks;
  }
}

export { Crunch };
