import { Vector } from "./text_vector";

interface ImplicitAttributeParams {
  attributeVector: Vector;
}
type ImplicitAttributeAction = {
  action: "APPEND_IMPLICIT_ATTRIBUTE";
  params: ImplicitAttributeParams;
};

interface ExplicitAttributeParams {
  attributeVector: Vector;
  valueVector: Vector;
}
type ExplicitAttributeAction = {
  action: "APPEND_EXPLICIT_ATTRIBUTE";
  params: ExplicitAttributeParams;
};

interface InjectedAttributeParams {
  attributeVector: Vector;
  injectionID: number;
}
type InjectedExplicitAttributeAction = {
  action: "APPEND_INJECTED_ATTRIBUTE";
  params: InjectedAttributeParams;
};

type AttributeAction =
  | ImplicitAttributeAction
  | ExplicitAttributeAction
  | InjectedExplicitAttributeAction;

export {
  ImplicitAttributeParams,
  ImplicitAttributeAction,
  ExplicitAttributeParams,
  ExplicitAttributeAction,
  InjectedAttributeParams,
  InjectedExplicitAttributeAction,
  AttributeAction,
};
