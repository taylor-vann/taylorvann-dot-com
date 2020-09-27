// brian taylor vann

// xml-crawl tests

import { XMLCrawler } from "./xml_crawler";
const title = "XML Crawl";

const defaultFailTest = () => {
  return ["fail xml crawl immediately"];
};

const unitTestXMLCrawler = {
  title,
  tests: [defaultFailTest],
};

export { unitTestXMLCrawler };
