// brian taylor vann

import {
  CreateNode,
  CreateTextNode,
  // SetDescendant,
  RemoveDescendant,
  SetAttribute,
  SetAttributeParams,
  // SetSiblings,
  // RemoveSiblings,
  Hooks,
} from "../../bang/hooks/hooks";

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

// const setDescendant: SetDescendant<DocumentNode> = (element, descendant) => {
//   return element.appendChild(descendant);
// };

const removeDescendant: RemoveDescendant<DocumentNode> = (
  element,
  descendant
) => {
  return element.removeChild(descendant);
};

// const setSiblings: SetSiblings<DocumentNode> = ({
//   siblings,
//   parent,
//   leftSibling,
//   rightSibling,
// }) => {
//   return [document.createTextNode("test")];
// };

// const removeSiblings: RemoveSiblings<DocumentNode> = ({
//   siblings,
//   parent,
//   leftSibling,
//   rightSibling,
// }) => {
//   return;
// };

const hooks: Hooks<DocumentNode, AttributeKinds> = {
  createNode,
  createTextNode,
  setAttribute,
  // setDescendant,
  // removeDescendant,
  // setSiblings,
  // removeSiblings,
};

export { DocumentNode, AttributeKinds, hooks };
