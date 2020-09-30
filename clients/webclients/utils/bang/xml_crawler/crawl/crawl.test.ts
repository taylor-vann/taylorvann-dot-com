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

const title = "bang/xml_crawler/crawl";
const runTestsAsynchronously = true;

const findNothingWhenThereIsPlainText = () => {
  const testBlank = testTextInterpolator`no nodes to be found!`;
  const assertions: string[] = [];

  const result = crawl({ brokenText: testBlank });
  if (result && result.nodeType !== "NOT_FOUND") {
    assertions.push(`should return NOT_FOUND instead of ${result.nodeType}`);
  }

  if (result.target.startPosition.stringArrayIndex !== 0) {
    assertions.push(`should return startPosition stringArrayIndex as 0`);
  }

  if (result.target.startPosition.stringIndex !== 0) {
    assertions.push(`should return startPosition stringIndex as 0`);
  }

  if (result.target.endPosition.stringArrayIndex !== 0) {
    assertions.push(`should return endPosition stringArrayIndex as 0`);
  }

  if (result.target.endPosition.stringIndex !== 0) {
    assertions.push(`should return endPosition stringIndex as 0`);
  }

  return assertions;
};

const findParagraphInPlainText = () => {
  const testOpenNode = testTextInterpolator`<p>`;
  const assertions: string[] = [];

  const result = crawl({ brokenText: testOpenNode });
  if (result.nodeType !== "OPEN_NODE_CONFIRMED") {
    assertions.push(
      `should return OPEN_NODE_CONFIRMED instead of ${result.nodeType}`
    );
  }

  if (result.target.startPosition.stringArrayIndex !== 0) {
    assertions.push(`should return startPosition stringArrayIndex as 0`);
  }

  if (result.target.startPosition.stringIndex !== 0) {
    assertions.push(`should return startPosition stringIndex as 0`);
  }

  if (result.target.endPosition.stringArrayIndex !== 0) {
    assertions.push(`should return endPosition stringArrayIndex as 0`);
  }

  if (result.target.endPosition.stringIndex !== 2) {
    assertions.push(`should return endPosition stringIndex as 2`);
  }

  return assertions;
};

const findCloseParagraphInPlainText = () => {
  const testTextCloseNode = testTextInterpolator`</p>`;
  const assertions: string[] = [];

  const result = crawl({ brokenText: testTextCloseNode });
  if (result && result.nodeType !== "CLOSE_NODE_CONFIRMED") {
    assertions.push(
      `should return CLOSE_NODE_CONFIRMED instead of ${result.nodeType}`
    );
  }

  if (result.target.startPosition.stringArrayIndex !== 0) {
    assertions.push(`should return startPosition stringArrayIndex as 0`);
  }

  if (result.target.startPosition.stringIndex !== 0) {
    assertions.push(`should return startPosition stringIndex as 2`);
  }

  if (result.target.endPosition.stringArrayIndex !== 0) {
    assertions.push(`should return endPosition stringArrayIndex as 0`);
  }

  if (result.target.endPosition.stringIndex !== 3) {
    assertions.push(`should return endPosition stringIndex as 3`);
  }

  return assertions;
};

const findIndependentParagraphInPlainText = () => {
  const testTextIndependentNode = testTextInterpolator`<p/>`;
  const assertions: string[] = [];

  const result = crawl({ brokenText: testTextIndependentNode });
  if (result && result.nodeType !== "INDEPENDENT_NODE_CONFIRMED") {
    assertions.push(
      `should return INDEPENDENT_NODE_CONFIRMED instead of ${result.nodeType}`
    );
  }

  if (result.target.startPosition.stringArrayIndex !== 0) {
    assertions.push(`should return startPosition stringArrayIndex as 0`);
  }

  if (result.target.startPosition.stringIndex !== 0) {
    assertions.push(`should return startPosition stringIndex as 0`);
  }

  if (result.target.endPosition.stringArrayIndex !== 0) {
    assertions.push(`should return endPosition stringArrayIndex as 0`);
  }

  if (result.target.endPosition.stringIndex !== 3) {
    assertions.push(`should return endPosition stringIndex as 3`);
  }

  return assertions;
};

const findOpenParagraphInTextWithArgs = () => {
  const testTextWithArgs = testTextInterpolator`an ${"example"} <p>${"!"}</p>`;
  const assertions: string[] = [];

  const result = crawl({ brokenText: testTextWithArgs });
  if (result && result.nodeType !== "OPEN_NODE_CONFIRMED") {
    assertions.push(
      `should return OPEN_NODE_CONFIRMED instead of ${result.nodeType}`
    );
  }

  if (result.target.startPosition.stringArrayIndex !== 1) {
    assertions.push(`should return startPosition stringArrayIndex as 1`);
  }

  if (result.target.startPosition.stringIndex !== 1) {
    assertions.push(`should return startPosition stringIndex as 1`);
  }

  if (result.target.endPosition.stringArrayIndex !== 1) {
    assertions.push(`should return endPosition stringArrayIndex as 1`);
  }

  if (result.target.endPosition.stringIndex !== 3) {
    assertions.push(`should return endPosition stringIndex as 3`);
  }

  return assertions;
};

const notFoundInUgglyMessText = () => {
  const testInvalidUgglyMess = testTextInterpolator`an <${"invalid"}p> example${"!"}`;
  const assertions: string[] = [];

  const result = crawl({ brokenText: testInvalidUgglyMess });
  if (result && result.nodeType !== "NOT_FOUND") {
    assertions.push(`should return NOT_FOUND instead of ${result.nodeType}`);
  }

  if (result.target.startPosition.stringArrayIndex !== 0) {
    assertions.push(`should return startPosition stringArrayIndex as 0`);
  }

  if (result.target.startPosition.stringIndex !== 0) {
    assertions.push(`should return startPosition stringIndex as 0`);
  }

  if (result.target.endPosition.stringArrayIndex !== 0) {
    assertions.push(`should return endPosition stringArrayIndex as 0`);
  }

  if (result.target.endPosition.stringIndex !== 0) {
    assertions.push(`should return endPosition stringIndex as 0`);
  }

  return assertions;
};

const invalidCloseNodeWithArgs = () => {
  const testInvlaidCloseNodeWithArgs = testTextInterpolator`closed </${"example"}p>`;
  const assertions: string[] = [];

  const result = crawl({ brokenText: testInvlaidCloseNodeWithArgs });
  if (result && result.nodeType !== "NOT_FOUND") {
    assertions.push(`should return NOT_FOUND instead of ${result.nodeType}`);
  }

  if (result.target.startPosition.stringArrayIndex !== 0) {
    assertions.push(`should return startPosition stringArrayIndex as 0`);
  }

  if (result.target.startPosition.stringIndex !== 0) {
    assertions.push(`should return startPosition stringIndex as 0`);
  }

  if (result.target.endPosition.stringArrayIndex !== 0) {
    assertions.push(`should return endPosition stringArrayIndex as 0`);
  }

  if (result.target.endPosition.stringIndex !== 0) {
    assertions.push(`should return endPosition stringIndex as 0`);
  }

  return assertions;
};

const validCloseNodeWithArgs = () => {
  const testValidCloseNodeWithArgs = testTextInterpolator`closed </p ${"example"}>`;
  const assertions: string[] = [];

  const result = crawl({ brokenText: testValidCloseNodeWithArgs });
  if (result && result.nodeType !== "CLOSE_NODE_CONFIRMED") {
    assertions.push(
      `should return CLOSE_NODE_CONFIRMED instead of ${result.nodeType}`
    );
  }

  if (result.target.startPosition.stringArrayIndex !== 0) {
    assertions.push(`should return startPosition stringArrayIndex as 0`);
  }

  if (result.target.startPosition.stringIndex !== 7) {
    assertions.push(`should return startPosition stringIndex as 7`);
  }

  if (result.target.endPosition.stringArrayIndex !== 1) {
    assertions.push(`should return endPosition stringArrayIndex as 1`);
  }

  if (result.target.endPosition.stringIndex !== 0) {
    assertions.push(`should return endPosition stringIndex as 0`);
  }

  return assertions;
};

const invalidIndependentNodeWithArgs = () => {
  const testInvalidIndependentNode = testTextInterpolator`independent <${"example"}p/>`;
  const assertions: string[] = [];

  const result = crawl({ brokenText: testInvalidIndependentNode });
  if (result && result.nodeType !== "NOT_FOUND") {
    assertions.push(`should return NOT_FOUND instead of ${result.nodeType}`);
  }

  if (result.target.startPosition.stringArrayIndex !== 0) {
    assertions.push(`should return startPosition stringArrayIndex as 0`);
  }

  if (result.target.startPosition.stringIndex !== 0) {
    assertions.push(`should return startPosition stringIndex as 0`);
  }

  if (result.target.endPosition.stringArrayIndex !== 0) {
    assertions.push(`should return endPosition stringArrayIndex as 0`);
  }

  if (result.target.endPosition.stringIndex !== 0) {
    assertions.push(`should return endPosition stringIndex as 0`);
  }

  return assertions;
};

const validIndependentNodeWithArgs = () => {
  const testValidIndependentNode = testTextInterpolator`independent <p ${"example"} / >`;
  const assertions: string[] = [];

  const result = crawl({ brokenText: testValidIndependentNode });
  if (result && result.nodeType !== "INDEPENDENT_NODE_CONFIRMED") {
    assertions.push(
      `should return INDEPENDENT_NODE_CONFIRMED instead of ${result.nodeType}`
    );
  }

  if (result.target.startPosition.stringArrayIndex !== 0) {
    assertions.push(`should return startPosition stringArrayIndex as 0`);
  }

  if (result.target.startPosition.stringIndex !== 12) {
    assertions.push(`should return startPosition stringIndex as 12`);
  }

  if (result.target.endPosition.stringArrayIndex !== 1) {
    assertions.push(`should return endPosition stringArrayIndex as 1`);
  }

  if (result.target.endPosition.stringIndex !== 3) {
    assertions.push(`should return endPosition stringIndex as 3`);
  }

  return assertions;
};

const invalidOpenNodeWithArgs = () => {
  const testInvalidOpenNode = testTextInterpolator`open <${"example"}p>`;
  const assertions: string[] = [];

  const result = crawl({ brokenText: testInvalidOpenNode });
  if (result && result.nodeType !== "NOT_FOUND") {
    assertions.push(`should return NOT_FOUND instead of ${result.nodeType}`);
  }

  if (result.target.startPosition.stringArrayIndex !== 0) {
    assertions.push(`should return startPosition stringArrayIndex as 0`);
  }

  if (result.target.startPosition.stringIndex !== 0) {
    assertions.push(`should return startPosition stringIndex as 0`);
  }

  if (result.target.endPosition.stringArrayIndex !== 0) {
    assertions.push(`should return endPosition stringArrayIndex as 0`);
  }

  if (result.target.endPosition.stringIndex !== 0) {
    assertions.push(`should return endPosition stringIndex as 0`);
  }

  return assertions;
};

const validOpenNodeWithArgs = () => {
  const testValidOpenNode = testTextInterpolator`open <p ${"example"}>`;
  const assertions: string[] = [];

  const result = crawl({ brokenText: testValidOpenNode });
  if (result && result.nodeType !== "OPEN_NODE_CONFIRMED") {
    assertions.push(
      `should return OPEN_NODE_CONFIRMED instead of ${result.nodeType}`
    );
  }

  if (result.target.startPosition.stringArrayIndex !== 0) {
    assertions.push(`should return startPosition stringArrayIndex as 0`);
  }

  if (result.target.startPosition.stringIndex !== 5) {
    assertions.push(`should return startPosition stringIndex as 5`);
  }

  if (result.target.endPosition.stringArrayIndex !== 1) {
    assertions.push(`should return endPosition stringArrayIndex as 1`);
  }

  if (result.target.endPosition.stringIndex !== 0) {
    assertions.push(`should return endPosition stringIndex as 0`);
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
];

const unitTestCrawl = {
  title,
  tests,
  runTestsAsynchronously,
};

export { unitTestCrawl };
