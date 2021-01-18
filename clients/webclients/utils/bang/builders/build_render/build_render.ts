// brian taylor vann
// build structure

import { Hooks } from "../../hooks/hooks";
import { Template } from "../../type_flyweight/template";
import { Render } from "../../type_flyweight/render";
import { Integrals } from "../../type_flyweight/integrals";

// structure render and crawl results

interface BuildRenderParams<N, A> {
  hooks: Hooks<N, A>;
  template: Template<A>;
  integrals: Integrals;
}

type BuildRender = <N, A>(params: BuildRenderParams<N, A>) => Render<N, A>;

const buildRender: BuildRender = ({ hooks, template, integrals }) => {
  // lets start building stuff

  const siblings = [];

  // see if there's a starting content node
  return {
    injections: {},
    siblings: [],
  };
};

export { buildRender };
