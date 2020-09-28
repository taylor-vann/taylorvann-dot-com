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

const testBlank = testTextInterpolator`no nodes to be found!`;
const testOpenNode = testTextInterpolator`<p>a modest start</p>`;
const testTextInjected = testTextInterpolator`<p>an interesting</p>`;
const testTextReally = testTextInterpolator`a most ${"example"} <p>interesting example${"!"}</p>`;
const testUgglyMess = testTextInterpolator`a most <${"interesting"}p>invalid example example${"!"}</p>`;
const testValidUgglyMess = testTextInterpolator`a most <p ç>interesting example${"!"}</p>`;
const testCloseNode = testTextInterpolator`an interesting closed example</p>`;
const testInvlaidCloseNode = testTextInterpolator`an interesting closed example</${"example"}p>`;
const testValidCloseNode = testTextInterpolator`an interesting closed example</p ${"example"}>`;
const testInvalidIndependentNode = testTextInterpolator`an interesting independentd example</${"example"}p>`;
const testValidIndependentNode = testTextInterpolator`an interesting independent example</p ${"example"}>`;

const findNothingWhenThereIsPlainText = () => {
  const assertions: string[] = [];

  const result = crawl({ brokenText: testBlank });
  if (result && result.nodeType !== "NOT_FOUND") {
    assertions.push(`should return NOT_FOUND instead of ${result.nodeType}`);
  }

  return assertions;
};

const tests = [findNothingWhenThereIsPlainText];

const unitTestCrawl = {
  title,
  tests,
  runTestsAsynchronously,
};

export { unitTestCrawl };
