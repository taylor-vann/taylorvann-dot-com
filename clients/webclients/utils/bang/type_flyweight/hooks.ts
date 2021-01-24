// brian taylor vann
// hooks

type AttributeValue<A> = A | string | boolean | undefined;
interface SetAttributeParams<N, A> {
  node: N;
  attribute: string;
  value: AttributeValue<A>;
}

type SetAttribute<N, A> = (params: SetAttributeParams<N, A>) => N;

type CreateNode<N> = (tag: string) => N;
type CreateTextNode<N> = (content: string) => N;

interface DescendantParams<N> {
  descendant: N;
  parentNode?: N;
  leftNode?: N;
}
type AppendDescendant<N> = (params: DescendantParams<N>) => N;
type RemoveDescendant<N> = (parent: N, descendant: N) => N;

interface Hooks<N, A> {
  setAttribute: SetAttribute<N, A>;
  createNode: CreateNode<N>;
  createTextNode: CreateTextNode<N>;
  appendDescendant: AppendDescendant<N>;
  removeDescendant: RemoveDescendant<N>;
}

export {
  AttributeValue,
  SetAttributeParams,
  SetAttribute,
  CreateNode,
  CreateTextNode,
  AppendDescendant,
  RemoveDescendant,
  Hooks,
};
