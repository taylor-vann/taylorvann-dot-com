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

interface CreateOpenNode {
  action: "CREATE_OPEN_NODE";
}
interface CreateIndependentNode {
  action: "CREATE_INDEPENDENT_NODE";
}
interface CreateAttribute {
  action: "CREATE_ATTRIBUTE";
}
interface CreateClosedNode {
  action: "CREATE_CLOSED_NODE";
}
interface CreateContentNode {
  action: "CREATE_CONTENT_NODE";
}

type IntegralAction =
  | CreateOpenNode
  | CreateIndependentNode
  | CreateAttribute
  | CreateClosedNode
  | CreateContentNode;

type IntegralRender = IntegralAction[];

export { IntegralAction, IntegralRender };
