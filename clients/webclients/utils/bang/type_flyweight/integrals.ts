// brian taylor vann
// integrals

// this is definietly gonna change

import {
  ImplicitAttributeParams,
  ImplicitAttributeAction,
  ExplicitAttributeAction,
  InjectedAttributeAction,
  AttributeAction,
} from "./attribute_crawl";

import { ContentCrawlAction } from "./content_crawl";
import { Vector } from "./text_vector";

interface NodeParams {
  tagNameVector: Vector;
}
interface CreateNodeAction {
  action: "CREATE_NODE";
  params: NodeParams;
}
interface CreateSelfClosingNode {
  action: "CREATE_SELF_CLOSING_NODE";
  params: NodeParams;
}
interface CloseNodeAction {
  action: "CLOSE_NODE";
  params: NodeParams;
}
type CreateContentAction = ContentCrawlAction;

type IntegralAction =
  | CreateNodeAction
  | CreateSelfClosingNode
  | ExplicitAttributeAction
  | CloseNodeAction
  | AttributeAction
  | CreateContentAction
  | ImplicitAttributeAction
  | InjectedAttributeAction;

type Integrals = IntegralAction[];

export {
  AttributeAction,
  CloseNodeAction,
  CreateContentAction,
  CreateNodeAction,
  CreateSelfClosingNode,
  ExplicitAttributeAction,
  ImplicitAttributeAction,
  ImplicitAttributeParams,
  InjectedAttributeAction,
  IntegralAction,
  Integrals,
};
