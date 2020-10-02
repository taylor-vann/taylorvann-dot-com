// brian taylor vann

import { BrokenTextVector, CrawlResults, crawl } from "./crawl/crawl";
import { CrawlNodes, buildSkeleton } from "./build_skeleton/build_skeleton";

class XMLCrawler {
  private brokenText: TemplateStringsArray;
  private injections: string[];
  private results: CrawlNodes;

  constructor(brokenText: TemplateStringsArray, ...injections: string[]) {
    this.brokenText = brokenText;
    this.injections = injections;
    this.results = buildSkeleton();
  }

  getResults(): CrawlNodes {
    return this.results;
  }
}

export { XMLCrawler };
