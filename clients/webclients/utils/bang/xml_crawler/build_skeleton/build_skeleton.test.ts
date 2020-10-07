// brian taylor vann

import { buildSkeleton } from "./build_skeleton";

type TextTextInterpolator = (
  brokenText: TemplateStringsArray,
  ...injections: string[]
) => TemplateStringsArray;

// order of start, end aren't being respected
const title = "bang/xml_crawler/crawl";
const runTestsAsynchronously = true;

const findNothingWhenThereIsPlainText = () => {
  const testBlank = buildSkeleton`no nodes to be found!`;
  console.log(testBlank);

  const assertions: string[] = [];

  return assertions;
};

const findParagraphInPlainText = () => {
  const testOpenNode = buildSkeleton`<p>`;
  console.log(testOpenNode);

  const assertions: string[] = [];

  return assertions;
};

const findComplexFromPlainText = () => {
  const testComplexNode = buildSkeleton`hello<p>world</p>`;
  console.log(testComplexNode);

  const assertions: string[] = [];

  return assertions;
};

const findCompoundFromPlainText = () => {
  const testComplexNode = buildSkeleton`<h1>hello</h1><h2>world</h2><img/><p>howdy</p>`;
  console.log(testComplexNode);

  const assertions: string[] = [];

  return assertions;
};

// const findBrokenFromPlainText = () => {
//   const testComplexNode = buildSkeleton`<h1>hello</h1><${"hello"}h2>world</h2><p>howdy</p>`;
//   console.log(testComplexNode);

//   const assertions: string[] = [];

//   return assertions;
// };

// const findCloseParagraphInPlainText = () => {
//   const testTextCloseNode = testTextInterpolator`</p>`;
//   const assertions: string[] = [];

//   return assertions;
// };

// const findIndependentParagraphInPlainText = () => {
//   const testTextIndependentNode = testTextInterpolator`<p/>`;
//   const assertions: string[] = [];

//   return assertions;
// };

// const findOpenParagraphInTextWithArgs = () => {
//   const testTextWithArgs = testTextInterpolator`an ${"example"} <p>${"!"}</p>`;
//   const assertions: string[] = [];

//   return assertions;
// };

// const notFoundInUgglyMessText = () => {
//   const testInvalidUgglyMess = testTextInterpolator`an <${"invalid"}p> example${"!"}`;
//   const assertions: string[] = [];

//   return assertions;
// };

// const invalidCloseNodeWithArgs = () => {
//   const testInvlaidCloseNodeWithArgs = testTextInterpolator`closed </${"example"}p>`;
//   const assertions: string[] = [];

//   return assertions;
// };

// const validCloseNodeWithArgs = () => {
//   const testValidCloseNodeWithArgs = testTextInterpolator`closed </p ${"example"}>`;
//   const assertions: string[] = [];

//   return assertions;
// };

// const invalidIndependentNodeWithArgs = () => {
//   const testInvalidIndependentNode = testTextInterpolator`independent <${"example"}p/>`;
//   const assertions: string[] = [];

//   return assertions;
// };

// const validIndependentNodeWithArgs = () => {
//   const testValidIndependentNode = testTextInterpolator`independent <p ${"example"} / >`;
//   const assertions: string[] = [];

//   return assertions;
// };

// const invalidOpenNodeWithArgs = () => {
//   const testInvalidOpenNode = testTextInterpolator`open <${"example"}p>`;
//   const assertions: string[] = [];

//   return assertions;
// };

// const validOpenNodeWithArgs = () => {
//   const testValidOpenNode = testTextInterpolator`open <p ${"example"}>`;
//   const assertions: string[] = [];

//   return assertions;
// };

const tests = [
  findNothingWhenThereIsPlainText,
  findParagraphInPlainText,
  findComplexFromPlainText,
  findCompoundFromPlainText,
  // findBrokenFromPlainText,
  // findCloseParagraphInPlainText,
  // findIndependentParagraphInPlainText,
  // findOpenParagraphInTextWithArgs,
  // notFoundInUgglyMessText,
  // invalidCloseNodeWithArgs,
  // validCloseNodeWithArgs,
  // invalidIndependentNodeWithArgs,
  // validIndependentNodeWithArgs,
  // invalidOpenNodeWithArgs,
  // validOpenNodeWithArgs,
];

const unitTestBuildSkeleton = {
  title,
  tests,
  runTestsAsynchronously,
};

export { unitTestBuildSkeleton };
