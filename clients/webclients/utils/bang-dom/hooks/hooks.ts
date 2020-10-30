// brian taylor vann

import {
  ParseNode,
  CreateNode,
  CreateContentNode,
  AddSiblings,
  RemoveSiblings,
  InterfaceHooks,
} from "../../bang/interface_hooks/interface_hooks";

type DocumentNode = Text | Element | Node | HTMLElement;
type NodeFunctor = <A>(params: A) => DocumentNode[];
type AttributeKinds = boolean | string | undefined | NodeFunctor;

const createNode: CreateNode<DocumentNode, AttributeKinds> = (parseResults) => {
  return document.createElement(parseResults.tag);
};

const createContentNode: CreateContentNode<DocumentNode> = (content) => {
  return document.createTextNode(content);
};

const addSiblings: AddSiblings<DocumentNode> = ({
  siblings,
  parent,
  leftSibling,
  rightSibling,
}) => {
  return [document.createTextNode("test")];
};

const removeSiblings: RemoveSiblings<DocumentNode> = ({
  siblings,
  parent,
  leftSibling,
  rightSibling,
}) => {
  return;
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
  addSiblings,
  removeSiblings,
};

export { hooks };
