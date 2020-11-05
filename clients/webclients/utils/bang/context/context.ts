// brian talor vann

// Bang context manager
// logic_structures return contexts

// we don't want direct access to contexts
// key value storage meets graph storage

// we can simplify things, it's a tree. no cycles.

// N Node
// A Attributables

// P Params

import { InterfaceHooks } from "../interface_hooks/interface_hooks";
import { BangFunc, ContextBase, Structure } from "../references/context";
import { RenderResults, StructureRender } from "../references/render";

class Context<N, A, P, R> implements ContextBase<N, A, P, R> {
  // across class
  readonly id = -1;
  readonly timestamp = -1;

  private hooksRef: InterfaceHooks<N, A>;
  private structureRef: Structure<N, A, P, R>;
  // deltas
  // private params: P;
  // private structureRef: Structure<N, A, P, R>;
  // private structureResults: StructureRender<A>;
  // private renderResults: RenderResults<N, A>;

  constructor(hooks: InterfaceHooks<N, A>, structure: Structure<N, A, P, R>) {
    this.hooksRef = hooks;
    this.structureRef = structure;
  }

  update(params?: P) {
    return [];
  }

  getBang() {
    return () => {};
  }

  disconnect() {}
}

export { Context };
