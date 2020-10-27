// brian taylor vann

import { unitTestCrawl } from "./builders/build_skeleton/crawl/crawl.test";
import { unitTestRouters } from "./builders/build_skeleton/crawl/routers/routers.test";
import { unitTestBuildSkeleton } from "./builders/build_skeleton/build_skeleton.test";

const tests = [unitTestRouters, unitTestCrawl, unitTestBuildSkeleton];

export { tests };
