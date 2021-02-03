// brian taylor vann

import {
  CreateNode,
  CreateTextNode,
  AppendDescendant,
  RemoveDescendant,
  SetAttribute,
  SetAttributeParams,
  Hooks,
} from "../../crunch/type_flyweight/hooks";

type DocumentNode = Text | HTMLElement;
type AttributeKinds =
  | EventListenerOrEventListenerObject
  | boolean
  | string
  | undefined;

const createNode: CreateNode<HTMLElement> = (tag: string) => {
  return document.createElement(tag);
};

const createTextNode: CreateTextNode<Text> = (content: string) => {
  return document.createTextNode(content);
};

const setAttribute: SetAttribute<DocumentNode, AttributeKinds> = ({
  node,
  attribute,
  value,
}: SetAttributeParams<DocumentNode, AttributeKinds>) => {
  // undefined values
  if (value === undefined) {
    // ? should be removed
    if (node instanceof HTMLElement) {
      node.removeAttribute(attribute);
    }

    // @events should be removed
    if (typeof value === "function") {
      node.removeEventListener(attribute, value);
    }
    return node;
  }

  // @ add an event listener
  if (value instanceof Function) {
    node.addEventListener(attribute, value);
  }

  // ?
  if (node instanceof HTMLElement) {
    node.setAttribute(attribute, value.toString());
  }

  return node;
};

const appendDescendant: AppendDescendant<DocumentNode> = (
  parentNode,
  descendant
) => {
  if (parentNode !== undefined) {
    parentNode.removeChild(descendant);
  }

  return descendant;
};

const removeDescendant: RemoveDescendant<DocumentNode> = (
  element,
  descendant
) => {
  element.removeChild(descendant);

  return descendant;
};

const hooks: Hooks<DocumentNode, AttributeKinds> = {
  createNode,
  createTextNode,
  setAttribute,
  appendDescendant,
  removeDescendant,
};

export { DocumentNode, AttributeKinds, hooks };