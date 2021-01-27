// brian taylor vann
// test hooks

import { Hooks } from "../type_flyweight/hooks";
import { TestNode, TestAttributes } from "./test_element";

// these names need to change

const TestHooks: Hooks<TestNode, TestAttributes> = {
  createNode: (tagname) => {
    return { kind: "ELEMENT", attributes: {}, tagname };
  },
  createTextNode: (text) => {
    return { kind: "TEXT", text };
  },
  setAttribute: (params) => {
    const { node, attribute, value } = params;
    if (node.kind === "ELEMENT") {
      node.attributes[attribute] = value;
    }
  },
  appendDescendant: (params) => {
    const { descendant, leftNode } = params;

    const rightNode = leftNode?.right;
    if (leftNode !== undefined) {
      leftNode.right = descendant;
    }
    descendant.right = rightNode;
  },
  removeDescendant: (descendant) => {
    const leftNode = descendant.left;
    const rightNode = descendant.right;

    if (leftNode !== undefined) {
      leftNode.right = rightNode;
    }

    if (rightNode !== undefined) {
      rightNode.left = leftNode;
    }
  },
};
