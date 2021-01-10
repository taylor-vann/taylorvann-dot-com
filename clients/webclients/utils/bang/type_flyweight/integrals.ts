// brian taylor vann
// integrals

// this is definietly gonna change

import {
  ImplicitAttributeParams,
  ImplicitAttributeAction,
  ExplicitAttributeParams,
  ExplicitAttributeAction,
  InjectedAttributeParams,
  InjectedExplicitAttributeAction,
  AttributeAction,
} from "./attribute_crawl";

import { ContentCrawlAction } from "./content_crawl";
import { Vector } from "./text_vector";

interface NodeParams {
  tagNameVector: Vector;
}
interface CreateNode {
  action: "CREATE_NODE";
  params: NodeParams;
}
interface CreateSelfClosingNode {
  action: "CREATE_SELF_CLOSING_NODE";
  params: NodeParams;
}
interface CloseNode {
  action: "CLOSE_NODE";
  params: NodeParams;
}

type IntegralAction =
  | CreateNode
  | CreateSelfClosingNode
  | ExplicitAttributeAction
  | ImplicitAttributeParams
  | InjectedAttributeParams
  | CloseNode
  | AttributeAction
  | ContentCrawlAction;

type Integrals = IntegralAction[];

export {
  IntegralAction,
  Integrals,
  CreateNode,
  CreateSelfClosingNode,
  ExplicitAttributeAction,
  ImplicitAttributeParams,
  InjectedAttributeParams,
  CloseNode,
};
