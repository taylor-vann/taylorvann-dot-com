// brian taylor vann

import { unitTestCrawl } from "./builders/build_skeleton/crawl/crawl.test";
import { unitTestRouters } from "./builders/build_skeleton/crawl/routers/routers.test";
import { unitTestBuildSkeleton } from "./builders/build_skeleton/build_skeleton.test";
import { unitTestBuildIntegrals } from "./builders/build_integrals/build_integrals.test";
import { unitTestTextVector } from "./text_vector/text_vector.test";
import { unitTestBuildStructure } from "./builders/build_structure/build_structure.test";

const tests = [
  // unitTestRouters,
  // unitTestCrawl,
  unitTestBuildSkeleton,
  // unitTestBuildIntegrals,
  // unitTestTextVector,
];

export { tests };
