import { Hooks } from "../type_flyweight/hooks";

type Attributes = string | Function;

interface Element {
  kind: "ELEMENT";
  tagname: string;
  attributes: Record<string, Attributes>;
  children?: Node;
  parent?: Element;
  left?: Node;
  right?: Node;
}
interface Text {
  kind: "TEXT";
  text: string;
  parent?: Element;
  left?: Node;
  right?: Node;
}

type Node = Element | Text;

const TestHooks: Hooks<Node, Attributes> = {
  createNode: (tagname) => {
    return { kind: "ELEMENT", tagname, attributes: {} };
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
