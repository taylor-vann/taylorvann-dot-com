// brian taylor vann

import { crawl } from "./crawl";

type TextTextInterpolator = (
  brokenText: TemplateStringsArray,
  ...injections: string[]
) => TemplateStringsArray;

const testTextInterpolator: TextTextInterpolator = (
  brokenText,
  ...injections
) => {
  return brokenText;
};

const testDebugTextInterpolator: TextTextInterpolator = (
  brokenText,
  ...injections
) => {
  console.log(brokenText);
  return brokenText;
};

const title = "Crawl";
const runTestsAsynchronously = true;

const findNothingWhenThereIsPlainText = () => {
  const testBlank = testTextInterpolator`no nodes to be found!`;
  const assertions: string[] = [];

  const result = crawl(testBlank);
  if (result === undefined) {
    assertions.push("undefined result");
  }

  if (result && result.nodeType !== "CONTENT_NODE") {
    assertions.push(`should return CONTENT_NODE instead of ${result.nodeType}`);
  }

  if (result && result.target.start.arrayIndex !== 0) {
    assertions.push(`should return start arrayIndex as 0`);
  }

  if (result && result.target.start.stringIndex !== 0) {
    assertions.push(`should return start stringIndex as 0`);
  }

  if (result && result.target.end.arrayIndex !== 0) {
    assertions.push(`should return end arrayIndex as 0`);
  }

  if (result && result.target.end.stringIndex !== 20) {
    assertions.push(`should return end stringIndex as 20`);
  }

  return assertions;
};

const findParagraphInPlainText = () => {
  const testOpenNode = testTextInterpolator`<p>`;
  const assertions: string[] = [];

  const result = crawl(testOpenNode);
  if (result === undefined) {
    assertions.push("undefined result");
  }
  if (result && result.nodeType !== "OPEN_NODE_CONFIRMED") {
    assertions.push(
      `should return OPEN_NODE_CONFIRMED instead of ${result.nodeType}`
    );
  }

  if (result && result.target.start.arrayIndex !== 0) {
    assertions.push(`should return start arrayIndex as 0`);
  }

  if (result && result.target.start.stringIndex !== 0) {
    assertions.push(`should return start stringIndex as 0`);
  }

  if (result && result.target.end.arrayIndex !== 0) {
    assertions.push(`should return end arrayIndex as 0`);
  }

  if (result && result.target.end.stringIndex !== 2) {
    assertions.push(`should return end stringIndex as 2`);
  }

  return assertions;
};

const findCloseParagraphInPlainText = () => {
  const testTextCloseNode = testTextInterpolator`</p>`;
  const assertions: string[] = [];

  const result = crawl(testTextCloseNode);
  if (result === undefined) {
    assertions.push("undefined result");
  }
  if (result && result.nodeType !== "CLOSE_NODE_CONFIRMED") {
    assertions.push(
      `should return CLOSE_NODE_CONFIRMED instead of ${result.nodeType}`
    );
  }

  if (result && result.target.start.arrayIndex !== 0) {
    assertions.push(`should return start arrayIndex as 0`);
  }

  if (result && result.target.start.stringIndex !== 0) {
    assertions.push(`should return start stringIndex as 2`);
  }

  if (result && result.target.end.arrayIndex !== 0) {
    assertions.push(`should return end arrayIndex as 0`);
  }

  if (result && result.target.end.stringIndex !== 3) {
    assertions.push(`should return end stringIndex as 3`);
  }

  return assertions;
};

const findIndependentParagraphInPlainText = () => {
  const testTextIndependentNode = testTextInterpolator`<p/>`;
  const assertions: string[] = [];

  const result = crawl(testTextIndependentNode);
  if (result === undefined) {
    assertions.push("undefined result");
  }
  if (result && result.nodeType !== "INDEPENDENT_NODE_CONFIRMED") {
    assertions.push(
      `should return INDEPENDENT_NODE_CONFIRMED instead of ${result.nodeType}`
    );
  }

  if (result && result.target.start.arrayIndex !== 0) {
    assertions.push(`should return start arrayIndex as 0`);
  }

  if (result && result.target.start.stringIndex !== 0) {
    assertions.push(`should return start stringIndex as 0`);
  }

  if (result && result.target.end.arrayIndex !== 0) {
    assertions.push(`should return end arrayIndex as 0`);
  }

  if (result && result.target.end.stringIndex !== 3) {
    assertions.push(`should return end stringIndex as 3`);
  }

  return assertions;
};

const findOpenParagraphInTextWithArgs = () => {
  const testTextWithArgs = testTextInterpolator`an ${"example"} <p>${"!"}</p>`;
  const assertions: string[] = [];

  const result = crawl(testTextWithArgs);
  if (result === undefined) {
    assertions.push("undefined result");
  }
  if (result && result.nodeType !== "OPEN_NODE_CONFIRMED") {
    assertions.push(
      `should return OPEN_NODE_CONFIRMED instead of ${result.nodeType}`
    );
  }

  if (result && result.target.start.arrayIndex !== 1) {
    assertions.push(`should return start arrayIndex as 1`);
  }

  if (result && result.target.start.stringIndex !== 1) {
    assertions.push(`should return start stringIndex as 1`);
  }

  if (result && result.target.end.arrayIndex !== 1) {
    assertions.push(`should return end arrayIndex as 1`);
  }

  if (result && result.target.end.stringIndex !== 3) {
    assertions.push(`should return end stringIndex as 3`);
  }

  return assertions;
};

const notFoundInUgglyMessText = () => {
  const testInvalidUgglyMess = testTextInterpolator`an <${"invalid"}p> example${"!"}`;
  const assertions: string[] = [];

  const result = crawl(testInvalidUgglyMess);
  if (result === undefined) {
    assertions.push("undefined result");
  }
  if (result && result.nodeType !== "CONTENT_NODE") {
    assertions.push(`should return CONTENT_NODE instead of ${result.nodeType}`);
  }

  if (result && result.target.start.arrayIndex !== 0) {
    assertions.push(`should return start arrayIndex as 0`);
  }

  if (result && result.target.start.stringIndex !== 0) {
    assertions.push(`should return start stringIndex as 0`);
  }

  if (result && result.target.end.arrayIndex !== 2) {
    assertions.push(`should return end arrayIndex as 2`);
  }

  if (result && result.target.end.stringIndex !== -1) {
    assertions.push(`should return end stringIndex as -1`);
  }

  return assertions;
};

const invalidCloseNodeWithArgs = () => {
  const testInvlaidCloseNodeWithArgs = testTextInterpolator`closed </${"example"}p>`;
  const assertions: string[] = [];

  const result = crawl(testInvlaidCloseNodeWithArgs);
  if (result === undefined) {
    assertions.push("undefined result");
  }
  if (result && result.nodeType !== "CONTENT_NODE") {
    assertions.push(`should return CONTENT_NODE instead of ${result.nodeType}`);
  }

  if (result && result.target.start.arrayIndex !== 0) {
    assertions.push(`should return start arrayIndex as 0`);
  }

  if (result && result.target.start.stringIndex !== 0) {
    assertions.push(`should return start stringIndex as 0`);
  }

  if (result && result.target.end.arrayIndex !== 1) {
    assertions.push(`should return end arrayIndex as 1`);
  }

  if (result && result.target.end.stringIndex !== 1) {
    assertions.push(`should return end stringIndex as 1`);
  }

  return assertions;
};

const validCloseNodeWithArgs = () => {
  const testValidCloseNodeWithArgs = testTextInterpolator`closed </p ${"example"}>`;
  const assertions: string[] = [];

  const result = crawl(testValidCloseNodeWithArgs);
  if (result === undefined) {
    assertions.push("undefined result");
  }

  if (result && result.nodeType !== "CLOSE_NODE_CONFIRMED") {
    assertions.push(
      `should return CLOSE_NODE_CONFIRMED instead of ${result.nodeType}`
    );
  }

  if (result && result.target.start.arrayIndex !== 0) {
    assertions.push(`should return start arrayIndex as 0`);
  }

  if (result && result.target.start.stringIndex !== 7) {
    assertions.push(`should return start stringIndex as 7`);
  }

  if (result && result.target.end.arrayIndex !== 1) {
    assertions.push(`should return end arrayIndex as 1`);
  }

  if (result && result.target.end.stringIndex !== 0) {
    assertions.push(`should return end stringIndex as 0`);
  }

  return assertions;
};

const invalidIndependentNodeWithArgs = () => {
  const testInvalidIndependentNode = testTextInterpolator`independent <${"example"}p/>`;
  const assertions: string[] = [];

  const result = crawl(testInvalidIndependentNode);
  if (result === undefined) {
    assertions.push("undefined result");
  }
  if (result && result.nodeType !== "CONTENT_NODE") {
    assertions.push(`should return CONTENT_NODE instead of ${result.nodeType}`);
  }

  if (result && result.target.start.arrayIndex !== 0) {
    assertions.push(`should return start arrayIndex as 0`);
  }

  if (result && result.target.start.stringIndex !== 0) {
    assertions.push(`should return start stringIndex as 0`);
  }

  if (result && result.target.end.arrayIndex !== 1) {
    assertions.push(`should return end arrayIndex as 1`);
  }

  if (result && result.target.end.stringIndex !== 2) {
    assertions.push(`should return end stringIndex as 2`);
  }

  return assertions;
};

const validIndependentNodeWithArgs = () => {
  const testValidIndependentNode = testTextInterpolator`independent <p ${"example"} / >`;
  const assertions: string[] = [];

  const result = crawl(testValidIndependentNode);
  if (result === undefined) {
    assertions.push("undefined result");
  }
  if (result && result.nodeType !== "INDEPENDENT_NODE_CONFIRMED") {
    assertions.push(
      `should return INDEPENDENT_NODE_CONFIRMED instead of ${result.nodeType}`
    );
  }

  if (result && result.target.start.arrayIndex !== 0) {
    assertions.push(`should return start arrayIndex as 0`);
  }

  if (result && result.target.start.stringIndex !== 12) {
    assertions.push(`should return start stringIndex as 12`);
  }

  if (result && result.target.end.arrayIndex !== 1) {
    assertions.push(`should return end arrayIndex as 1`);
  }

  if (result && result.target.end.stringIndex !== 3) {
    assertions.push(`should return end stringIndex as 3`);
  }

  return assertions;
};

const invalidOpenNodeWithArgs = () => {
  const testInvalidOpenNode = testTextInterpolator`open <${"example"}p>`;
  const assertions: string[] = [];

  const result = crawl(testInvalidOpenNode);
  if (result === undefined) {
    assertions.push("undefined result");
  }
  if (result && result.nodeType !== "CONTENT_NODE") {
    assertions.push(`should return CONTENT_NODE instead of ${result.nodeType}`);
  }

  if (result && result.target.start.arrayIndex !== 0) {
    assertions.push(`should return start arrayIndex as 0`);
  }

  if (result && result.target.start.stringIndex !== 0) {
    assertions.push(`should return start stringIndex as 0`);
  }

  if (result && result.target.end.arrayIndex !== 1) {
    assertions.push(`should return end arrayIndex as 1`);
  }

  if (result && result.target.end.stringIndex !== 1) {
    assertions.push(`should return end stringIndex as 1`);
  }

  return assertions;
};

const validOpenNodeWithArgs = () => {
  const testValidOpenNode = testTextInterpolator`open <p ${"example"}>`;
  const assertions: string[] = [];

  const result = crawl(testValidOpenNode);
  if (result === undefined) {
    assertions.push("undefined result");
  }
  if (result && result.nodeType !== "OPEN_NODE_CONFIRMED") {
    assertions.push(
      `should return OPEN_NODE_CONFIRMED instead of ${result.nodeType}`
    );
  }

  if (result && result.target.start.arrayIndex !== 0) {
    assertions.push(`should return start arrayIndex as 0`);
  }

  if (result && result.target.start.stringIndex !== 5) {
    assertions.push(`should return start stringIndex as 5`);
  }

  if (result && result.target.end.arrayIndex !== 1) {
    assertions.push(`should return end arrayIndex as 1`);
  }

  if (result && result.target.end.stringIndex !== 0) {
    assertions.push(`should return end stringIndex as 0`);
  }
  return assertions;
};

const validSecondaryIndependentNodeWithArgs = () => {
  const testValidOpenNode = testTextInterpolator`<p ${"small"}/>${"example"}<p/>`;
  const assertions: string[] = [];

  const previousCrawl = crawl(testValidOpenNode);
  const result = crawl(testValidOpenNode, previousCrawl);
  if (result === undefined) {
    assertions.push("undefined result");
  }

  if (result && result.nodeType !== "INDEPENDENT_NODE_CONFIRMED") {
    assertions.push(
      `should return INDEPENDENT_NODE_CONFIRMED instead of ${result.nodeType}`
    );
  }

  if (result && result.target.start.arrayIndex !== 2) {
    assertions.push(`should return start arrayIndex as 2`);
  }

  if (result && result.target.start.stringIndex !== 0) {
    assertions.push(`should return start stringIndex as 0`);
  }

  if (result && result.target.end.arrayIndex !== 2) {
    assertions.push(`should return end arrayIndex as 1`);
  }

  if (result && result.target.end.stringIndex !== 3) {
    assertions.push(`should return end stringIndex as 3`);
  }
  return assertions;
};

const tests = [
  findNothingWhenThereIsPlainText,
  findParagraphInPlainText,
  findCloseParagraphInPlainText,
  findIndependentParagraphInPlainText,
  findOpenParagraphInTextWithArgs,
  notFoundInUgglyMessText,
  invalidCloseNodeWithArgs,
  validCloseNodeWithArgs,
  invalidIndependentNodeWithArgs,
  validIndependentNodeWithArgs,
  invalidOpenNodeWithArgs,
  validOpenNodeWithArgs,
  validSecondaryIndependentNodeWithArgs,
];

const unitTestCrawl = {
  title,
  tests,
  runTestsAsynchronously,
};

export { unitTestCrawl };
