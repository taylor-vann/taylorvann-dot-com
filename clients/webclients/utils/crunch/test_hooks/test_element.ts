// brian taylor vann
// test element

import { AttributeValue } from "../type_flyweight/hooks";

// boolean, string, and undefined are included in Attributes by default
type TestAttributes = string;

interface TestElement {
  kind: "ELEMENT";
  tagname: string;
  attributes: Record<string, AttributeValue<TestAttributes>>;
  parent?: TestElement;
  left?: TestNode;
  right?: TestNode;
  leftChild?: TestNode;
  rightChild?: TestNode;
}

interface TestText {
  kind: "TEXT";
  text: string;
  parent?: TestElement;
  left?: TestNode;
  right?: TestNode;
}

type TestNode = TestElement | TestText;

export { TestAttributes, TestNode };
