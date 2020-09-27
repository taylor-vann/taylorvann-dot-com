import {
  notFound,
  openNode,
  openNodeValid,
  independentNodeValid,
  closeNode,
  closeNodeValid,
} from "./routers";

const title = "Routers | Detect node state";
const runTestsAsynchronously = true;

const notFoundReducesCorrectState = () => {
  const assertions: string[] = [];
  if (notFound("<") !== "OPEN_NODE") {
    assertions.push("< should return OPEN_NODE");
  }

  if (notFound(" ") !== "NOT_FOUND") {
    assertions.push("space should return NOT_FOUND");
  }

  return assertions;
};

const openNodeReducesCorrectState = () => {
  const assertions: string[] = [];
  if (openNode("<") !== "OPEN_NODE") {
    assertions.push("< should return OPEN_NODE");
  }

  if (openNode("/") !== "CLOSE_NODE") {
    assertions.push("/ should return CLOSE_NODE");
  }

  if (openNode("b") !== "OPEN_NODE_VALID") {
    assertions.push("b should return OPEN_NODE_VALID");
  }

  if (openNode(" ") !== "NOT_FOUND") {
    assertions.push("space should return NOT_FOUND");
  }

  return assertions;
};

const openNodeValidReducesCorrectState = () => {
  const assertions: string[] = [];
  if (openNodeValid("<") !== "OPEN_NODE") {
    assertions.push("< should return OPEN_NODE");
  }

  if (openNodeValid("/") !== "INDEPENDENT_NODE_VALID") {
    assertions.push("/ should return INDEPENDENT_NODE_VALID");
  }

  if (openNodeValid(">") !== "OPEN_NODE_CONFIRMED") {
    assertions.push("> should return OPEN_NODE_CONFIRMED");
  }

  if (openNodeValid(" ") !== "OPEN_NODE_VALID") {
    assertions.push("space should return OPEN_NODE_VALID");
  }

  return assertions;
};

const independentNodeValidReducesCorrectState = () => {
  const assertions: string[] = [];
  if (independentNodeValid("<") !== "OPEN_NODE") {
    assertions.push("< should return OPEN_NODE");
  }

  if (independentNodeValid(">") !== "INDEPENDENT_NODE_CONFIRMED") {
    assertions.push("/ should return INDEPENDENT_NODE_CONFIRMED");
  }

  if (independentNodeValid(" ") !== "INDEPENDENT_NODE_VALID") {
    assertions.push("space should return INDEPENDENT_NODE_VALID");
  }

  return assertions;
};

const closeNodeReducesCorrectState = () => {
  const assertions: string[] = [];
  if (closeNode("<") !== "OPEN_NODE") {
    assertions.push("< should return OPEN_NODE");
  }

  if (closeNode("a") !== "CLOSE_NODE_VALID") {
    assertions.push("'a' should return CLOSE_NODE_VALID");
  }

  if (closeNode(" ") !== "NOT_FOUND") {
    assertions.push("space should return CLOSE_NODE_VALID");
  }

  return assertions;
};

const closeNodeValidReducesCorrectState = () => {
  const assertions: string[] = [];
  if (closeNodeValid("<") !== "OPEN_NODE") {
    assertions.push("< should return OPEN_NODE");
  }

  if (closeNodeValid(">") !== "CLOSE_NODE_CONFIRMED") {
    assertions.push("> should return CLOSE_NODE_CONFIRMED");
  }

  if (closeNodeValid(" ") !== "CLOSE_NODE_VALID") {
    assertions.push("space should return CLOSE_NODE_VALID");
  }

  return assertions;
};

const tests = [
  notFoundReducesCorrectState,
  openNodeReducesCorrectState,
  openNodeValidReducesCorrectState,
  independentNodeValidReducesCorrectState,
  closeNodeReducesCorrectState,
  closeNodeValidReducesCorrectState,
];

const unitTestRouters = {
  title,
  tests,
  runTestsAsynchronously,
};

export { unitTestRouters };
