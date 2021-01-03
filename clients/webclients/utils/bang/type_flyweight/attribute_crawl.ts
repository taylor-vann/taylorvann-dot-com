import { Vector } from "./text_vector";

interface ImplicitAttributeParams {
  attributeVector: Vector;
}
type ImplicitAttributeAction = {
  action: "IMPLICIT_ATTRIBUTE_CONFIRMED";
  params: ImplicitAttributeParams;
};

interface ExplicitAttributeParams {
  attributeVector: Vector;
  valueVector: Vector;
}
type ExplicitAttributeAction = {
  action: "EXPLICIT_ATTRIBUTE_CONFIRMED";
  params: ExplicitAttributeParams;
};

interface InjectedAttributeParams {
  attributeVector: Vector;
  injectionID: number;
}
type InjectedExplicitAttributeAction = {
  action: "INJECTED_EXPLICIT_ATTRIBUTE_CONFIRMED";
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
