// brian taylor vann

// N Node
// A Attributables

interface SetAttributeParams<N, A> {
  node: N;
  attribute: string;
  value: A;
}
type SetAttribute<N, A> = (params: SetAttributeParams<N, A>) => void;
type CreateNode<N> = (tag: string) => N;
type CreateContentNode<N> = (content: string) => N;
interface SetSiblingsParams<N> {
  siblings: N[];
  parent: N;
  leftSibling?: N;
  rightSibling?: N;
}
type SetSiblings<N> = (params: SetSiblingsParams<N>) => N[];
type RemoveSiblingsParams<N> = SetSiblingsParams<N>;
type RemoveSiblings<N> = (params: RemoveSiblingsParams<N>) => void;

// set attribute is a problem
// Use this to create new Bang Interfaces
interface InterfaceHooks<N, A> {
  setAttribute: SetAttribute<N, A>;
  createNode: CreateNode<N>;
  createContentNode: CreateContentNode<N>;
  setSiblings: SetSiblings<N>;
  removeSiblings: RemoveSiblings<N>;
}

export {
  SetAttribute,
  SetSiblings,
  CreateContentNode,
  CreateNode,
  InterfaceHooks,
  RemoveSiblings,
};
