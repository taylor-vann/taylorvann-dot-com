// brian taylor vann

import {
  CreateNode,
  CreateContentNode,
  SetDescendant,
  RemoveDescendant,
  SetAttribute,
  SetSiblings,
  RemoveSiblings,
  InterfaceHooks,
} from "../../bang/interface_hooks/interface_hooks";

type DocumentNode = Text | Element | Node | HTMLElement;
type AttributeKinds = boolean | string | undefined;

const createNode: CreateNode<DocumentNode> = (tag) => {
  return document.createElement(tag);
};

const createContentNode: CreateContentNode<DocumentNode> = (content) => {
  return document.createTextNode(content);
};

const setDescendant: SetDescendant<DocumentNode> = (element, descendant) => {
  return element.appendChild(descendant);
};

const removeDescendant: RemoveDescendant<DocumentNode> = (
  element,
  descendant
) => {
  return element.removeChild(descendant);
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
  if (value === undefined) {
    return node;
  }

  if (node instanceof Element) {
    node.setAttribute(attribute, value.toString());
  }

  return node;
};

const hooks: InterfaceHooks<DocumentNode, AttributeKinds> = {
  setAttribute,
  createNode,
  createContentNode,
  setDescendant,
  removeDescendant,
  setSiblings,
  removeSiblings,
};

export { DocumentNode, AttributeKinds, hooks };
