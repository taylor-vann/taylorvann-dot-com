// brian taylor vann

import { buildSkeleton } from "./build_skeleton";

type TextTextInterpolator = (
  brokenText: TemplateStringsArray,
  ...injections: string[]
) => TemplateStringsArray;

const getTemplateArray: TextTextInterpolator = (brokenText, ...injections) => {
  return brokenText;
};

// order of start, end aren't being respected
const title = "Bang XML Crawl";
const runTestsAsynchronously = true;

const findNothingWhenThereIsPlainText = () => {
  const testBlank = getTemplateArray`no nodes to be found!`;
  const testSkeleton = buildSkeleton(testBlank);
  // console.log(testSkeleton);

  const assertions: string[] = [];

  return assertions;
};

const findParagraphInPlainText = () => {
  const testOpenNode = getTemplateArray`<p>`;
  const testSkeleton = buildSkeleton(testOpenNode);

  // console.log(testOpenNode);

  const assertions: string[] = [];

  return assertions;
};

const findComplexFromPlainText = () => {
  const testComplexNode = getTemplateArray`hello<p>world</p>`;
  const testSkeleton = buildSkeleton(testComplexNode);

  // console.log(testSkeleton);

  const assertions: string[] = [];

  return assertions;
};

const findCompoundFromPlainText = () => {
  const testComplexNode = getTemplateArray`<h1>hello</h1><h2>world</h2><img/><p>howdy</p>`;
  const testSkeleton = buildSkeleton(testComplexNode);

  // console.log(testSkeleton);

  const assertions: string[] = [];

  return assertions;
};

const findBrokenFromPlainText = () => {
  const testComplexNode = getTemplateArray`<h1>hello</h1><${"hello"}h2>world</h2><p>howdy</p>`;
  const testSkeleton = buildSkeleton(testComplexNode);

  // console.log(testSkeleton);

  const assertions: string[] = [];

  return assertions;
};

const tests = [
  findNothingWhenThereIsPlainText,
  findParagraphInPlainText,
  findComplexFromPlainText,
  findCompoundFromPlainText,
  findBrokenFromPlainText,
];

const unitTestBuildSkeleton = {
  title,
  tests,
  runTestsAsynchronously,
};

export { unitTestBuildSkeleton };
