// brian taylor vann

import { unitTestAttributeCrawl } from "./builders/attribute_crawl/attribute_crawl.test";
import { unitTestBuildIntegrals } from "./builders/build_integrals/build_integrals.test";
import { unitTestBuildRender } from "./builders/build_render/build_render.test";
import { unitTestBuildSkeleton } from "./builders/build_skeleton/build_skeleton.test";
import { unitTestCrawl } from "./builders/skeleton_crawl/skeleton_crawl.test";
import { unitTestCrawlRouters } from "./builders/skeleton_routers/skeleton_routers.test";
import { unitTestTagNameCrawl } from "./builders/tag_name_crawl/tag_name_crawl.test";
import { unitTestTextPosition } from "./text_position/text_position.test";
import { unitTestTextVector } from "./text_vector/text_vector.test";

const tests = [
  unitTestAttributeCrawl,
  unitTestBuildIntegrals,
  unitTestBuildRender,
  unitTestBuildSkeleton,
  unitTestCrawl,
  unitTestCrawlRouters,
  unitTestTagNameCrawl,
  unitTestTextPosition,
  unitTestTextVector,
];

export { tests };
