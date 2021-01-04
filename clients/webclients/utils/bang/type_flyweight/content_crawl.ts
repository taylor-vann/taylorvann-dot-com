import { Vector } from "./text_vector";

interface ContentCrawlParams {
  contentVector: Vector;
}
type ContentCrawlAction = {
  action: "CREATE_CONTENT";
  params: ContentCrawlParams;
};

export { ContentCrawlAction };
