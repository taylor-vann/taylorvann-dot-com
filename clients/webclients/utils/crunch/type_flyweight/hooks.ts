// brian taylor vann
// hooks

import { Context } from "./context";

type AttributeValue<N, A> = Context<N, A> | A | string | boolean | undefined;
interface SetAttributeParams<N, A> {
  attribute: string;
  node: N;
  value: AttributeValue<N, A>;
}

type SetAttribute<N, A> = (params: SetAttributeParams<N, A>) => void;

type CreateNode<N> = (tag: string) => N;
type CreateTextNode<N> = (content: string) => N;

type AppendDescendant<N> = (parent: N, descendant: N) => void;
type RemoveDescendant<N> = (parent: N, descendant: N) => void;

interface Hooks<N, A> {
  appendDescendant: AppendDescendant<N>;
  createNode: CreateNode<N>;
  createTextNode: CreateTextNode<N>;
  removeDescendant: RemoveDescendant<N>;
  setAttribute: SetAttribute<N, A>;
}

export {
  AppendDescendant,
  AttributeValue,
  CreateNode,
  CreateTextNode,
  Hooks,
  RemoveDescendant,
  SetAttribute,
  SetAttributeParams,
};
