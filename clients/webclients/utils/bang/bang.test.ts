// brian taylor vann

import { unitTestCrawl } from "./builders/build_skeleton/crawl/crawl.test";
import { unitTestCrawlRouters } from "./builders/build_skeleton/crawl/routers/routers.test";
import { unitTestBuildSkeleton } from "./builders/build_skeleton/build_skeleton.test";
import { unitTestBuildIntegrals } from "./builders/build_integrals/build_integrals.test";
import { unitTestTextPosition } from "./text_position/text_position.test";
import { unitTestTextVector } from "./text_vector/text_vector.test";
import { unitTestBuildStructure } from "./builders/build_structure/build_structure.test";

const tests = [
  unitTestTextPosition,
  // unitTestTextVector,
  // unitTestRouters,
  // unitTestCrawl,
  // unitTestBuildSkeleton,
  // unitTestBuildIntegrals,
];

export { tests };
