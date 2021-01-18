// brian taylor vann
// interface hooks

// N Node
// A Attributables

interface SetAttributeParams<N, A> {
  node: N;
  attribute: string;
  value: A;
}
type SetAttribute<N, A> = (params: SetAttributeParams<N, A>) => N;

type CreateNode<N> = (tag: string) => N;
type CreateTextNode<N> = (content: string) => N;
type RemoveDescendant<N> = (element: N, descendant: N) => N;

// we need
type AppendDescendent<N> = (element: N, descendant: N) => N;
// type RemoveDescendant

// set attribute is a problem
// Use this to create new Bang Interfaces

//

// we eneed a demo / test object for builds

interface TestNode {
  tagname: string;
  attributes: Record<string, string>;
  children: TestNode[];
}

interface Hooks<N, A> {
  setAttribute: SetAttribute<N, A>;
  createNode: CreateNode<N>;
  createTextNode: CreateTextNode<N>;
}

export {
  SetAttributeParams,
  SetAttribute,
  CreateNode,
  CreateTextNode,
  RemoveDescendant,
  Hooks,
};
