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
  valueVector: Vector;
  injectionID: number;
}
type InjectedAttributeAction = {
  action: "APPEND_INJECTED_ATTRIBUTE";
  params: InjectedAttributeParams;
};

type AttributeAction =
  | ImplicitAttributeAction
  | ExplicitAttributeAction
  | InjectedAttributeAction;

export {
  ImplicitAttributeParams,
  ImplicitAttributeAction,
  ExplicitAttributeParams,
  ExplicitAttributeAction,
  InjectedAttributeParams,
  InjectedAttributeAction,
  AttributeAction,
};
