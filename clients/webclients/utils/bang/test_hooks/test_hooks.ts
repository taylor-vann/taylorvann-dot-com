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
    return node;
  },
  appendDescendant: (params) => {
    const { descendant, parentNode, leftNode } = params;
    // append descendent logic
    return descendant;
  },
  removeDescendant: (descendant) => {
    // remove descendent l
    return descendant;
  },
};
