// brian taylor vann
// integrals

interface CreateTextNodeAction {
  action: "CREATE_TEXT_NODE";
  text: string[];
  injections: Record<number, number>;
}

type Attributes<A> = Record<string, A>;
interface CreateNodeAction<A> {
  action: "CREATE_NODE";
  tag: string;
  attributes: Attributes<A>;
  injectionToAttributeAddress: Record<number, string>;
}
interface CloseNodeAction {
  action: "CLOSE_NODE";
  tag: string;
}
interface CreateSiblingsAction {
  action: "CREATE_SIBLINGS";
  injectionID: number;
}

type IntegralAction<A> =
  | CreateTextNodeAction
  | CreateNodeAction<A>
  | CloseNodeAction
  | CreateSiblingsAction;

type IntegralRender<A> = IntegralAction<A>[];

export {
  Attributes,
  CreateTextNodeAction,
  CreateNodeAction,
  CloseNodeAction,
  CreateSiblingsAction,
  IntegralAction,
  IntegralRender,
};
