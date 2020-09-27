// brian taylor vann

import { crawl } from "./crawl";

// xml-crawl tests

const testTextInterpolator = (
  brokenText: TemplateStringsArray,
  ...injections: string[]
) => {
  console.log(brokenText);

  return brokenText;
};

const testText = testTextInterpolator`<p>a modest start</p>`;
const testTextInjected = testTextInterpolator`<p>an interesting ${"example"}</p>`;
const testTextReally = testTextInterpolator`a most <p>interesting example${"!"}</p>`;

const title = "bang/xml_crawler/crawl";

const defaultFailTest = () => {
  return ["fail crawl immediately"];
};

// frist testText
// {0, 0}, {0, 2}

const unitTestCrawl = {
  title,
  tests: [defaultFailTest],
};

export { unitTestCrawl };
