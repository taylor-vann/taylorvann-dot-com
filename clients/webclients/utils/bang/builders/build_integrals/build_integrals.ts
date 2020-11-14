// brian taylor vann
// build structure

import { StructureRender } from "../../references/structure";
import { CrawlResults } from "../../references/crawl";
import { IntegralRender } from "../../references/integrals";

interface ParseSkeletonParams<A> {
  template: StructureRender<A>;
  crawl: CrawlResults;
}
type ParseSkeleton = <A>(params: ParseSkeletonParams<A>) => IntegralRender<A>;

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

const parseSkeleton: ParseSkeleton = ({}) => {
  return [];
};

export { parseSkeleton };
