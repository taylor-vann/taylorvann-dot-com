// brian taylor vann
// integrals

// this is definietly gonna change

import {
  ImplicitAttributeAction,
  ExplicitAttributeAction,
  InjectedAttributeAction,
  AttributeAction,
} from "./attribute_crawl";

import { Vector } from "./text_vector";

interface NodeAction {
  kind: "NODE";
  tagNameVector: Vector;
}
interface SelfClosingNodeAction {
  kind: "SELF_CLOSING_NODE";
  tagNameVector: Vector;
}
interface CloseNodeAction {
  kind: "CLOSE_NODE";
  tagNameVector: Vector;
}

interface TextAction {
  kind: "TEXT";
  textVector: Vector;
}

interface ContextInjectionAction {
  kind: "CONTEXT_INJECTION";
  injectionID: number;
}

type IntegralAction =
  | AttributeAction
  | CloseNodeAction
  | ContextInjectionAction
  | ExplicitAttributeAction
  | ImplicitAttributeAction
  | InjectedAttributeAction
  | NodeAction
  | SelfClosingNodeAction
  | TextAction;

type Integrals = IntegralAction[];

export {
  AttributeAction,
  CloseNodeAction,
  ContextInjectionAction,
  ExplicitAttributeAction,
  ImplicitAttributeAction,
  InjectedAttributeAction,
  IntegralAction,
  Integrals,
  NodeAction,
  SelfClosingNodeAction,
  TextAction,
};
