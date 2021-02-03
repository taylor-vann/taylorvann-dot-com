// brian taylor vann
// test hooks

import { Hooks } from "../type_flyweight/hooks";
import { TestNode, TestAttributes } from "./test_element";
import { Context } from "../type_flyweight/context";

type SimilarNodeTrees = <N>(expectedSiblings: N[], siblings: N[]) => boolean;

const hooks: Hooks<TestNode, string> = {
  createNode: (tagname) => {
    return { kind: "ELEMENT", attributes: {}, tagname };
  },
  createTextNode: (text) => {
    return { kind: "TEXT", text };
  },
  setAttribute: (params) => {
    const { node, attribute, value } = params;
    if (value instanceof Context) {
      return;
    }

    if (node.kind === "ELEMENT") {
      node.attributes[attribute] = value;
    }
  },
  appendDescendant: (parent, descendant) => {
    if (parent === undefined || parent.kind !== "ELEMENT") {
      return;
    }

    // Add parent to descendant
    descendant.parent = parent;

    // if children exist, add descendant
    if (parent !== undefined && parent.rightChild !== undefined) {
      descendant.left = parent.rightChild;
      parent.rightChild.right = descendant;
      parent.rightChild = descendant;
    }
    // if no children, add initial child
    if (parent !== undefined && parent.leftChild === undefined) {
      parent.leftChild = descendant;
      parent.rightChild = descendant;
    }
  },
  removeDescendant: (parent, descendant) => {
    const leftNode = descendant.left;
    const rightNode = descendant.right;

    // remove descendant references
    descendant.parent = undefined;
    descendant.right = undefined;
    descendant.left = undefined;

    // if descendant is leftChild
    if (leftNode !== undefined) {
      leftNode.right = rightNode;
    }

    // if descendant is rightChild
    if (rightNode !== undefined) {
      rightNode.left = leftNode;
    }

    // if descendant is rightChild
    if (parent.kind === "ELEMENT") {
      if (descendant === parent.leftChild) {
        parent.leftChild = rightNode;
      }
      if (descendant === parent.rightChild) {
        parent.rightChild = leftNode;
      }
    }
  },
};

const similarNodeTrees: SimilarNodeTrees = (expectedSiblings, siblings) => {
  // avoid recursive relationships
  // travers children but do not traverse parent attribute
  // compare properties and attributes

  return false;
};

export { similarNodeTrees, hooks };
