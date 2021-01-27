// brian taylor vann

import {
  CreateNode,
  CreateTextNode,
  // SetDescendant,
  AppendDescendant,
  RemoveDescendant,
  SetAttribute,
  SetAttributeParams,
  // SetSiblings,
  // RemoveSiblings,
  Hooks,
} from "../../bang/type_flyweight/hooks";

type DocumentNode = Text | HTMLElement;
type AttributeKinds =
  | EventListenerOrEventListenerObject
  | boolean
  | string
  | undefined;

const createNode = (tag: string) => {
  return document.createElement(tag);
};

const createTextNode = (content: string) => {
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

const appendDescendant: AppendDescendant<DocumentNode> = ({
  descendant,
  parentNode,
  leftNode,
}) => {
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
