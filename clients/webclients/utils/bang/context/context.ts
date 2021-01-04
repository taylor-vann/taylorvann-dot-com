// brian taylor vann

// U Unique Tag names
// N Node
// A Attributables
// P Params

import { ContextBase } from "../type_flyweight/context";
import { RenderResults, StructureRender } from "../type_flyweight/render";

type Timestamp = number;
type DescendantRecord = Record<number, Timestamp>;
type CreateContext = <N, A, P, R>() => ContextBase<N, A, P, R>;

class Context<N, A, P, R> implements ContextBase<N, A, P, R> {
  readonly id: number;
  timestamp: number;

  params?: P;
  gambit?: R;
  structureResults?: StructureRender<A>;
  renderResults?: RenderResults<N, A>;

  constructor() {
    this.id = -1;
    this.timestamp = performance.now();
  }
}

export { Context, CreateContext, DescendantRecord, RenderResults, Timestamp };
