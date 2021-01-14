// brian taylor vann
// content_crawl

import { StructureRender } from "../../type_flyweight/structure";
import { Vector } from "../../type_flyweight/text_vector";
import { ContentCrawlAction } from "../../type_flyweight/content_crawl";
import { copy } from "../../text_vector/text_vector";

type ContentCrawl = <A>(
  template: StructureRender<A>,
  vectorBounds: Vector
) => ContentCrawlAction | undefined;

const crawlForContent: ContentCrawl = (template, vectorBounds) => {
  return {
    action: "CREATE_CONTENT",
    params: { contentVector: copy(vectorBounds) },
  };
};

export { crawlForContent };
