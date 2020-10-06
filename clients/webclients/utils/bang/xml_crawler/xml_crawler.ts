// brian taylor vann

import { SkeletonNodes, buildSkeleton } from "./build_skeleton/build_skeleton";

class XMLCrawler {
  private brokenText: TemplateStringsArray;
  private injections: string[];
  private results: SkeletonNodes;

  constructor(brokenText: TemplateStringsArray, ...injections: string[]) {
    this.brokenText = brokenText;
    this.injections = injections;
    this.results = buildSkeleton(brokenText, ...injections);
  }

  getResults(): SkeletonNodes {
    return this.results;
  }
}

export { XMLCrawler };
