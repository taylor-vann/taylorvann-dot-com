// brian taylor vann

import { unitTestCrawl } from "./xml_crawler/crawl/crawl.test";
import { unitTestRouters } from "./xml_crawler/crawl/routers/routers.test";
import { unitTestBuildSkeleton } from "./xml_crawler/build_skeleton/build_skeleton.test";

// import { unitTestXMLCrawler } from "./xml_crawler/xml_crawler.test";

const tests = [unitTestRouters, unitTestCrawl, unitTestBuildSkeleton];

export { tests };
