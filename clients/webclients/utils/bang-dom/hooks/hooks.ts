// brian taylor vann

import {
  CreateNode,
  CreateContentNode,
  SetAttribute,
  SetSiblings,
  RemoveSiblings,
  InterfaceHooks,
} from "../../bang/interface_hooks/interface_hooks";

type DocumentNode = Text | Element | Node | HTMLElement;
type NodeFunctor = <A>(params: A) => DocumentNode[];
type AttributeKinds = boolean | string | undefined | NodeFunctor;

const createNode: CreateNode<DocumentNode> = (tag) => {
  return document.createElement(tag);
};

const createContentNode: CreateContentNode<DocumentNode> = (content) => {
  return document.createTextNode(content);
};

const setSiblings: SetSiblings<DocumentNode> = ({
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

const setAttribute: SetAttribute<DocumentNode, AttributeKinds> = ({
  node,
  attribute,
  value,
}) => {
  return node;
};

const hooks: InterfaceHooks<DocumentNode, AttributeKinds> = {
  setAttribute,
  createNode,
  createContentNode,
  setSiblings,
  removeSiblings,
};

export { hooks };
