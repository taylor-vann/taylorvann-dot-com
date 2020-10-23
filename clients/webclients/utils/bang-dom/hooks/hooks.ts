// brian taylor vann

import {
  ParseNode,
  ParseResults,
  CreateNode,
  CreateContentNode,
  AddDescendent,
  RemoveDescendent,
  Hooks,
} from "../../bang/bang_hooks/bang_hooks";

type TagNames = keyof HTMLElementTagNameMap;
type DocumentNode = Text | Element | Node | HTMLElement;
type NodeFunctor = <A>(params: A) => DocumentNode[];
type AttributeKinds = boolean | string | undefined | NodeFunctor;
type NodeParams = ParseResults<TagNames, AttributeKinds, EventListener>;

const createNode: CreateNode<NodeParams, DocumentNode> = (parseResults) => {
  return document.createElement(parseResults.tag);
};

const createContentNode: CreateContentNode<DocumentNode> = (content) => {
  return document.createTextNode(content);
};

const addDescendent: AddDescendent<DocumentNode> = (element, descendent) => {
  return element.appendChild(descendent);
};

const removeDescendent: RemoveDescendent<DocumentNode> = (
  element,
  descendent
) => {
  return element.removeChild(descendent);
};

const parseNode: ParseNode<TagNames, AttributeKinds, EventListener> = (
  params
) => {
  return {
    tag: "div",
    kind: "OPEN_NODE",
    attributes: [],
  };
};

const hooks: Hooks<TagNames, DocumentNode, AttributeKinds, EventListener> = {
  parseNode,
  createNode,
  createContentNode,
  addDescendent,
  removeDescendent,
};

export { hooks };
