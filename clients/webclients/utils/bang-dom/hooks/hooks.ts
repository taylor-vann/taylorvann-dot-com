// brian taylor vann

import {
  ParseNode,
  ParseResults,
  CreateNode,
  CreateContentNode,
  AddDescendent,
  RemoveDescendent,
  InterfaceHooks,
} from "../../bang/interface_hooks/interface_hooks";

// Structure
// {
//   id,
// }

type DocumentNode = Text | Element | Node | HTMLElement;
type NodeFunctor = <A>(params: A) => DocumentNode[];
type AttributeKinds = boolean | string | undefined | NodeFunctor;
type NodeParams = ParseResults<AttributeKinds>;

const createNode: CreateNode<DocumentNode, AttributeKinds> = (parseResults) => {
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

const parseNode: ParseNode<AttributeKinds> = (params) => {
  return {
    tag: "div",
    kind: "OPEN_NODE",
    attributes: {},
  };
};

const hooks: InterfaceHooks<DocumentNode, AttributeKinds> = {
  parseNode,
  createNode,
  createContentNode,
  addDescendent,
  removeDescendent,
};

export { hooks };
