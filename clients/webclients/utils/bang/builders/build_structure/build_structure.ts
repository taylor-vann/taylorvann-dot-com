// brian taylor vann
// build structure

import { InterfaceHooks } from "../../interface_hooks/interface_hooks";
import { Structure } from "../../references/context";
import { StructureRender, RenderResults } from "../../references/render";

// structure render and crawl results

interface BuildStructureParams<N, A, P, R> {
  hooks: InterfaceHooks<N, A>;
  structureRef: Structure<N, A, P, R>;
  params?: P;
}
type BuildStructure = <N, A, P, R>(
  params: BuildStructureParams<N, A, P, R>
) => RenderResults<N, A>;

const buildStructure: BuildStructure = ({ hooks, structureRef, params }) => {
  // lets start building stuff

  return {
    injections: [],
    siblings: [],
  };
};

export { buildStructure };
