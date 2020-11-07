// brian taylor vann

// U Unique Tag names
// N Node
// A Attributables
// P Params

import { InterfaceHooks } from "../interface_hooks/interface_hooks";
import { ContextBase, Structure } from "../references/context";

type Timestamp = number;
type DescendantRecord = Record<number, Timestamp>;

interface RenderResults<A> {
  templateArray: TemplateStringsArray;
  injections: A[];
}

// interface Context<N, A, P, R> {
//   id: number;
//   connectedResults: P;
//   descendants: DescendantRecord;
//   params: P;
//   renderResults: RenderResults<A>;
//   onConnected: () => R;
//   onDisconnected: (params: R) => void;
//   renderStructure: (params: P) => void;
// }

type CreateContext<N, A, P, R> = (params: P) => Context<N, A, P, R>;
type ContextFactory<N, A> = <P, R>(
  params: Structure<N, A, P, R>
) => CreateContext<N, A, P, R>;

class Context<N, A, P, R> implements ContextBase<N, A, P, R> {
  private hooksRef: InterfaceHooks<N, A>;
  private structureRef: Structure<N, A, P, R>;

  readonly id: number = -1;
  readonly timestamp: Timestamp = -1;

  constructor(hooks: InterfaceHooks<N, A>, structure: Structure<N, A, P, R>) {
    this.hooksRef = hooks;
    this.structureRef = structure;
    this.timestamp = performance.now();
  }

  update(params?: P) {
    return [];
  }

  getBang() {
    return () => {};
  }

  disconnect() {}
}

export {
  Context,
  CreateContext,
  ContextFactory,
  DescendantRecord,
  RenderResults,
  Timestamp,
};
