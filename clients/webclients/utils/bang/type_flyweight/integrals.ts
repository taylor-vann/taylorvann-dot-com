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

interface ElementParams {
  tagNameVector: Vector;
}
interface CreateElement {
  action: "CREATE_ELEMENT";
  params: ElementParams;
}
interface CreateIndependentElement {
  action: "CREATE_INDEPENDENT_ELEMENT";
  params: ElementParams;
}
interface CloseElement {
  action: "CLOSE_ELEMENT";
  params: ElementParams;
}

type IntegralAction =
  | CreateElement
  | CreateIndependentElement
  | ExplicitAttributeAction
  | ImplicitAttributeParams
  | InjectedAttributeParams
  | CloseElement
  | AttributeAction
  | ContentCrawlAction;

type Integrals = IntegralAction[];

export {
  IntegralAction,
  Integrals,
  CreateElement,
  CreateIndependentElement,
  ExplicitAttributeAction,
  ImplicitAttributeParams,
  InjectedAttributeParams,
  CloseElement,
};
