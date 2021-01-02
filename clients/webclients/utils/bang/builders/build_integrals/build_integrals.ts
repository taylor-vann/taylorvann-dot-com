// brian taylor vann
// build integrals

import { StructureRender } from "../../type_flyweight/structure";
import { SkeletonNodes } from "../../type_flyweight/crawl";
import { IntegralRender } from "../../type_flyweight/integrals";

interface BuildIntegralsParams<A> {
  template: StructureRender<A>;
  skeleton: SkeletonNodes;
}
type BuildIntegrals = <A>(params: BuildIntegralsParams<A>) => IntegralRender<A>;

const buildIntegrals: BuildIntegrals = ({ template, skeleton }) => {
  // for each sekelton step

  // if content node
  // create array []
  // add injections as strings
  // get injectionsID

  // build_content_node {
  //   content: ["hello", injected, "!"]
  //   injectionIDs: {1: 1}
  // }

  // if open node / independent node
  // get tag name
  // get attributes
  //

  // if close node
  // get tag name

  return [];
};

export { BuildIntegralsParams, buildIntegrals };
