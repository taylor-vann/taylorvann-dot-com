// brian taylor vann

import { unitTestCrawl } from "./builders/skeleton_crawl/skeleton_crawl.test";
import { unitTestCrawlRouters } from "./builders/skeleton_routers/skeleton_routers.test";
import { unitTestBuildSkeleton } from "./builders/build_skeleton/build_skeleton.test";
import { unitTestBuildIntegrals } from "./builders/build_integrals/build_integrals.test";
import { unitTestTextPosition } from "./text_position/text_position.test";
import { unitTestTextVector } from "./text_vector/text_vector.test";
import { unitTestBuildRender } from "./builders/build_render/build_render.test";
import { unitTestTagNameCrawl } from "./builders/tag_name_crawl/tag_name_crawl.test";
import { unitTestAttributeCrawl } from "./builders/attribute_crawl/attribute_crawl.test";

import { unitTestContentCrawl } from "./builders/content_crawl/content_crawl.test";

const tests = [
  unitTestTextPosition,
  unitTestTextVector,
  unitTestCrawlRouters,
  unitTestCrawl,
  unitTestBuildSkeleton,
  unitTestTagNameCrawl,
  unitTestAttributeCrawl,
  unitTestContentCrawl,
  unitTestBuildIntegrals,
];

export { tests };
