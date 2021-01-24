import { AttributeValue } from "../type_flyweight/hooks";

// boolean, string, and undefined are included in Attributes by default
type TestAttributes = Function;

interface Element {
  kind: "ELEMENT";
  tagname: string;
  attributes: Record<string, AttributeValue<TestAttributes>>;
  parent?: Element;
  left?: Node;
  right?: Node;
  leftChild?: Element;
  rightChild?: Element;
}

interface Text {
  kind: "TEXT";
  text: string;
  parent?: Element;
  left?: Node;
  right?: Node;
  leftChild?: Node;
  rightChild?: Node;
}

type TestNode = Element | Text;

export { TestAttributes, TestNode };
