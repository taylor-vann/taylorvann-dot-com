// brian taylor vann

import { unitTestCrawl } from "./builders/build_skeleton/crawl/crawl.test";
import { unitTestCrawlRouters } from "./builders/build_skeleton/crawl/routers/routers.test";
import { unitTestBuildSkeleton } from "./builders/build_skeleton/build_skeleton.test";
import { unitTestBuildIntegrals } from "./builders/build_integrals/build_integrals.test";
import { unitTestTextPosition } from "./text_position/text_position.test";
import { unitTestTextVector } from "./text_vector/text_vector.test";
import { unitTestBuildStructure } from "./builders/build_structure/build_structure.test";
import { unitTestTagNameCrawl } from "./builders/build_integrals/tag_name_crawl/tag_name_crawl.test";
import { unitTestAttributeCrawl } from "./builders/build_integrals/attribute_crawl/attribute_crawl.test";

import { unitTestContentCrawl } from "./builders/build_integrals/content_crawl/content_crawl.test";

const tests = [
  unitTestTextPosition,
  unitTestTextVector,
  unitTestCrawlRouters,
  unitTestCrawl,
  unitTestBuildSkeleton,
  unitTestTagNameCrawl,
  unitTestAttributeCrawl,
  unitTestContentCrawl,
  // unitTestBuildIntegrals,
];

export { tests };
