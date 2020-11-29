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

// we want to build something that outputs exact instructions

// -> iterate over building structure
//      need to retain parent, prev, curr
//
// -> we need to breakdown initial renders
//      if open node, find attributes
//        if attribute is injection, create injections
//
// -> if injection is a context

// so we are iterating up till the context

const buildIntegrals: BuildIntegrals = ({ template, skeleton }) => {
  return [];
};

export { BuildIntegralsParams, buildIntegrals };