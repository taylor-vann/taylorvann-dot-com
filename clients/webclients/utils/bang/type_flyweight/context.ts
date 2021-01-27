import { RenderStructure } from "./render";

interface ContextParams<N, A> {
  renderStructure: RenderStructure<N, A>;
}

class Context<N, A> {
  private rs: RenderStructure<N, A>;

  constructor(params: ContextParams<N, A>) {
    this.rs = params.renderStructure;
  }

  getSiblings(): N[] {
    return this.rs.siblings;
  }
}

export { Context };
