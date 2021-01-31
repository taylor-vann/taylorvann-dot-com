// brian taylor vann
// test hooks

import { hooks } from "./test_hooks";

const title = "test_hooks";
const runTestsAsynchronously = true;

const testCreateNode = () => {
  const assertions = [];

  const node = hooks.createNode("hello");
  if (node === undefined) {
    assertions.push("node should not be undefined.");
  }

  if (node.kind !== "ELEMENT") {
    assertions.push("should create an ELEMENT");
  }

  if (node.kind === "ELEMENT" && node.tagname !== "hello") {
    assertions.push("tagname should be 'hello'");
  }
  return assertions;
};

const testCreateTextNode = () => {
  const assertions = [];

  const node = hooks.createTextNode("hello!");
  if (node === undefined) {
    assertions.push("text node should not be undefined.");
  }

  if (node.kind === "TEXT" && node.text !== "hello!") {
    assertions.push("text node should have 'hello!'");
  }

  return assertions;
};

const testSetAttribute = () => {
  const assertions = [];

  const node = hooks.createNode("basic");
  hooks.setAttribute({ node, attribute: "checked", value: true });

  if (node.kind !== "ELEMENT") {
    assertions.push("node should be an ELEMENT");
  }
  if (node.kind === "ELEMENT" && node.attributes["checked"] !== true) {
    assertions.push("text node should not be undefined.");
  }
  return assertions;
};

// append descendarnt
const testAppendDescendant = () => {
  const assertions = [];

  const sunshine = hooks.createNode("sunshine");
  const moonbeam = hooks.createNode("moonbeam");
  const starlight = hooks.createNode("starlight");

  hooks.appendDescendant(sunshine, starlight);
  hooks.appendDescendant(sunshine, moonbeam);

  if (starlight.kind === "ELEMENT" && starlight.left !== undefined) {
    assertions.push("starlight should have no left sibling");
  }
  if (starlight.kind === "ELEMENT" && starlight.right !== moonbeam) {
    assertions.push("starlight should have moonbeam as a sibling");
  }
  if (moonbeam.kind === "ELEMENT" && moonbeam.right !== undefined) {
    assertions.push("moonbeam should have no left sibling");
  }
  if (moonbeam.kind === "ELEMENT" && starlight.parent !== sunshine) {
    assertions.push("starlight should have sunshin as a parent");
  }
  if (moonbeam.kind === "ELEMENT" && moonbeam.parent !== sunshine) {
    assertions.push("moonbean should have sunshin as a parent");
  }
  return assertions;
};

// remove descendarnt
const testRemoveDescendant = () => {
  const assertions = [];

  const sunshine = hooks.createNode("sunshine");
  const moonbeam = hooks.createNode("moonbeam");
  const starlight = hooks.createNode("starlight");

  hooks.appendDescendant(sunshine, starlight);
  hooks.appendDescendant(sunshine, moonbeam);
  hooks.removeDescendant(sunshine, starlight);

  // starlight should not have left or right
  if (starlight.left !== undefined) {
    assertions.push("starlight should not have a left sibling.");
  }
  if (starlight.right !== undefined) {
    assertions.push("starlight should not have a right sibling.");
  }
  if (starlight.parent !== undefined) {
    assertions.push("starlight should not have a parent.");
  }

  if (moonbeam.left !== undefined) {
    assertions.push("moonbeam should not have a left sibling.");
  }
  if (moonbeam.right !== undefined) {
    assertions.push("moonbeam should not have a right sibling.");
  }
  if (moonbeam.parent !== sunshine) {
    assertions.push("moonbeam should have sunshine as a parent.");
  }

  if (sunshine.kind === "ELEMENT" && sunshine.leftChild !== moonbeam) {
    assertions.push("sunshine should have moonbeam as a left child.");
  }
  if (sunshine.kind === "ELEMENT" && sunshine.rightChild !== moonbeam) {
    assertions.push("sunshine should have moonbeam as a right child.");
  }

  return assertions;
};

const testRemoveAllDescendants = () => {
  const assertions = [];

  const sunshine = hooks.createNode("sunshine");
  const moonbeam = hooks.createNode("moonbeam");
  const starlight = hooks.createNode("starlight");

  hooks.appendDescendant(sunshine, starlight);
  hooks.appendDescendant(sunshine, moonbeam);
  hooks.removeDescendant(sunshine, starlight);
  hooks.removeDescendant(sunshine, moonbeam);

  // starlight should be de-referenced
  if (starlight.left !== undefined) {
    assertions.push("starlight should not have a left sibling.");
  }
  if (starlight.right !== undefined) {
    assertions.push("starlight should not have a right sibling.");
  }
  if (starlight.parent !== undefined) {
    assertions.push("starlight should not have a parent.");
  }

  // moonbean should be de-referenced
  if (moonbeam.left !== undefined) {
    assertions.push("moonbeam should not have a left sibling.");
  }
  if (moonbeam.right !== undefined) {
    assertions.push("moonbeam should not have a right sibling.");
  }
  if (moonbeam.parent !== undefined) {
    assertions.push("moonbeam should not have a parent.");
  }

  // parent and sunshine
  if (sunshine.kind === "ELEMENT" && sunshine.leftChild !== undefined) {
    assertions.push("sunshine should not have a left child.");
  }
  if (sunshine.kind === "ELEMENT" && sunshine.rightChild !== undefined) {
    assertions.push("sunshine should not have a right child.");
  }

  return assertions;
};

const tests = [
  testCreateNode,
  testCreateTextNode,
  testSetAttribute,
  testAppendDescendant,
  testRemoveDescendant,
  testRemoveAllDescendants,
];

const unitTestTestHooks = {
  title,
  tests,
  runTestsAsynchronously,
};

export { unitTestTestHooks };
