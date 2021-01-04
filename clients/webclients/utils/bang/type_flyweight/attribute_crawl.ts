import { Vector } from "./text_vector";

interface ImplicitAttributeParams {
  attributeVector: Vector;
}
type ImplicitAttributeAction = {
  action: "CREATE_IMPLICIT_ATTRIBUTE";
  params: ImplicitAttributeParams;
};

interface ExplicitAttributeParams {
  attributeVector: Vector;
  valueVector: Vector;
}
type ExplicitAttributeAction = {
  action: "CREATE_EXPLICIT_ATTRIBUTE";
  params: ExplicitAttributeParams;
};

interface InjectedAttributeParams {
  attributeVector: Vector;
  injectionID: number;
}
type InjectedExplicitAttributeAction = {
  action: "CREATE_INJECTED_EXPLICIT_ATTRIBUTE";
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
