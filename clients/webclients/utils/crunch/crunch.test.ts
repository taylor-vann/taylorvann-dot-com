// brian taylor vann
// bang

import { unitTestAttributeCrawl } from "./builders/attribute_crawl/attribute_crawl.test";
import { unitTestBuildIntegrals } from "./builders/build_integrals/build_integrals.test";
import { unitTestBuildRender } from "./builders/build_render/build_render.test";
import { unitTestBuildSkeleton } from "./builders/build_skeleton/build_skeleton.test";
import { unitTestCrawlRouters } from "./builders/skeleton_routers/skeleton_routers.test";
import { unitTestSkeletonCrawl } from "./builders/skeleton_crawl/skeleton_crawl.test";
import { unitTestTagNameCrawl } from "./builders/tag_name_crawl/tag_name_crawl.test";
import { unitTestTestHooks } from "./test_hooks/test_hooks.test";
import { unitTestTextPosition } from "./text_position/text_position.test";
import { unitTestTextVector } from "./text_vector/text_vector.test";

const tests = [
  // unitTestAttributeCrawl,
  // unitTestBuildIntegrals,
  unitTestBuildRender,
  // unitTestBuildSkeleton,
  // unitTestCrawlRouters,
  // unitTestSkeletonCrawl,
  // unitTestTagNameCrawl,
  // unitTestTestHooks,
  // unitTestTextPosition,
  // unitTestTextVector,
];

export { tests };
