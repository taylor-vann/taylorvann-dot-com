// brian taylor vann
// tag name crawl

import { StructureRender } from "../../../type_flyweight/structure";
import { create, incrementTarget } from "../../../text_vector/text_vector";

import { crawlForTagName } from "./tag_name_crawl";

type TextTextInterpolator = <A>(
  templateArray: TemplateStringsArray,
  ...injections: A[]
) => StructureRender<A>;

const RECURSION_SAFETY = 256;

const testTextInterpolator: TextTextInterpolator = (
  templateArray,
  ...injections
) => {
  return { templateArray, injections };
};

const title = "tag_name_crawl";
const runTestsAsynchronously = true;

const testEmptyString = () => {
  const assertions = [];

  const template = testTextInterpolator``;
  const vector = create();
  const results = crawlForTagName(template, vector);

  if (results !== undefined) {
    assertions.push("this should have failed");
  }

  return assertions;
};

const testEmptySpaceString = () => {
  const assertions = [];

  const template = testTextInterpolator` `;
  const vector = create();
  const results = crawlForTagName(template, vector);

  if (results !== undefined) {
    assertions.push("this should have failed");
  }

  return assertions;
};

const testSingleCharacterString = () => {
  const assertions = [];

  const template = testTextInterpolator`a`;
  const vector = create();
  const results = crawlForTagName(template, vector);

  if (results === undefined) {
    assertions.push("this should have returned a vector");
  }

  if (results !== undefined && results.origin.arrayIndex !== 0) {
    assertions.push("incorrect origin array index");
  }
  if (results !== undefined && results.origin.stringIndex !== 0) {
    assertions.push("incorrect origin string index");
  }
  if (results !== undefined && results.target.arrayIndex !== 0) {
    assertions.push("incorrect target array index");
  }
  if (results !== undefined && results.target.stringIndex !== 0) {
    assertions.push("incorrect target string index");
  }

  return assertions;
};

const testCharaceterString = () => {
  const assertions = [];

  const template = testTextInterpolator`a `;
  const vector = create();

  let safety = 0;
  while (incrementTarget(template, vector) && safety < RECURSION_SAFETY) {
    safety += 1;
  }

  const results = crawlForTagName(template, vector);

  if (results !== undefined) {
    assertions.push("this should have returned a vector");
  }

  if (results !== undefined && results.origin.arrayIndex !== 0) {
    assertions.push("incorrect origin array index");
  }
  if (results !== undefined && results.origin.stringIndex !== 0) {
    assertions.push("incorrect origin string index");
  }
  if (results !== undefined && results.target.arrayIndex !== 0) {
    assertions.push("incorrect target array index");
  }
  if (results !== undefined && results.target.stringIndex !== 0) {
    assertions.push("incorrect target string index");
  }

  return assertions;
};

const testMultiCharaceterString = () => {
  const assertions = [];

  const template = testTextInterpolator`aaa `;
  const vector = create();

  let safety = 0;
  while (incrementTarget(template, vector) && safety < RECURSION_SAFETY) {
    safety += 1;
  }

  const results = crawlForTagName(template, vector);

  if (results !== undefined) {
    assertions.push("this should have returned a vector");
  }

  if (results !== undefined && results.origin.arrayIndex !== 0) {
    assertions.push("incorrect origin array index");
  }
  if (results !== undefined && results.origin.stringIndex !== 0) {
    assertions.push("incorrect origin string index");
  }
  if (results !== undefined && results.target.arrayIndex !== 0) {
    assertions.push("incorrect target array index");
  }
  if (results !== undefined && results.target.stringIndex !== 2) {
    assertions.push("incorrect target string index");
  }

  return assertions;
};

const testMultiCharaceterStringWithTrailingSpaces = () => {
  const assertions = [];

  const template = testTextInterpolator`aaa     `;
  const vector = create();

  let safety = 0;
  while (incrementTarget(template, vector) && safety < RECURSION_SAFETY) {
    safety += 1;
  }

  const results = crawlForTagName(template, vector);

  if (results !== undefined) {
    assertions.push("this should have returned a vector");
  }

  if (results !== undefined && results.origin.arrayIndex !== 0) {
    assertions.push("incorrect origin array index");
  }
  if (results !== undefined && results.origin.stringIndex !== 0) {
    assertions.push("incorrect origin string index");
  }
  if (results !== undefined && results.target.arrayIndex !== 0) {
    assertions.push("incorrect target array index");
  }
  if (results !== undefined && results.target.stringIndex !== 2) {
    assertions.push("incorrect target string index");
  }

  return assertions;
};

const tests = [
  testEmptyString,
  testEmptySpaceString,
  testSingleCharacterString,
  testCharaceterString,
  testMultiCharaceterString,
  testMultiCharaceterStringWithTrailingSpaces,
];

const unitTestTagNameCrawl = {
  title,
  tests,
  runTestsAsynchronously,
};

export { unitTestTagNameCrawl };
