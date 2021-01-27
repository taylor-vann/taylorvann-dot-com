// brian taylor vann
// content crawl

import { Vector } from "./text_vector";

type ContentCrawlAction = {
  kind: "CREATE_CONTENT";
  contentVector: Vector;
};

export { ContentCrawlAction };
