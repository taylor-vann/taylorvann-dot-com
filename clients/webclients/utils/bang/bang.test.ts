// brian taylor vann

import { unitTestCrawl } from "./crawl/crawl.test";
import { unitTestRouters } from "./crawl/routers/routers.test";
import { unitTestBuildSkeleton } from "./build_skeleton/build_skeleton.test";

// import { unitTestXMLCrawler } from "./xml_crawler/xml_crawler.test";

const tests = [unitTestRouters, unitTestCrawl, unitTestBuildSkeleton];

export { tests };
