// brian taylor vann

import { SkeletonNodes, buildSkeleton } from "./build_skeleton";

type TextTextInterpolator = (
  brokenText: TemplateStringsArray,
  ...injections: string[]
) => TemplateStringsArray;

type CompareSkeletons = (
  source: SkeletonNodes,
  target: SkeletonNodes
) => boolean;

const title = "build_skeleton";
const runTestsAsynchronously = true;

const getTemplateArray: TextTextInterpolator = (brokenText, ...injections) => {
  return brokenText;
};

const compareSkeletons: CompareSkeletons = (source, target) => {
  for (const sourceKey in source) {
    const node = source[sourceKey];
    const targetNode = target[sourceKey];

    if (targetNode === undefined) {
      return false;
    }

    if (node.nodeType !== targetNode.nodeType) {
      return false;
    }
    if (
      node.target.start.arrayIndex !== targetNode.target.start.arrayIndex ||
      node.target.start.stringIndex !== targetNode.target.start.stringIndex ||
      node.target.end.arrayIndex !== targetNode.target.end.arrayIndex ||
      node.target.end.stringIndex !== targetNode.target.end.stringIndex
    ) {
      return false;
    }
  }

  return true;
};

const findNothingWhenThereIsPlainText = () => {
  const assertions: string[] = [];
  const sourceSkeleton: SkeletonNodes = [
    {
      nodeType: "CONTENT_NODE",
      target: {
        end: { arrayIndex: 0, stringIndex: 20 },
        start: { arrayIndex: 0, stringIndex: 0 },
      },
    },
  ];

  const testBlank = getTemplateArray`no nodes to be found!`;
  const testSkeleton = buildSkeleton(testBlank);
  if (!compareSkeletons(sourceSkeleton, testSkeleton)) {
    assertions.push("skeletons are not equal");
  }

  return assertions;
};

const findParagraphInPlainText = () => {
  const assertions: string[] = [];
  const sourceSkeleton: SkeletonNodes = [
    {
      nodeType: "OPEN_NODE_CONFIRMED",
      target: {
        end: { arrayIndex: 0, stringIndex: 2 },
        start: { arrayIndex: 0, stringIndex: 0 },
      },
    },
  ];
  const testOpenNode = getTemplateArray`<p>`;
  const testSkeleton = buildSkeleton(testOpenNode);
  if (!compareSkeletons(sourceSkeleton, testSkeleton)) {
    assertions.push("skeletons are not equal");
  }

  return assertions;
};

const findComplexFromPlainText = () => {
  const testComplexNode = getTemplateArray`hello<p>world</p>`;
  const testSkeleton = buildSkeleton(testComplexNode);

  console.log(testSkeleton);

  const assertions: string[] = ["fail automatically"];

  return assertions;
};

const findCompoundFromPlainText = () => {
  const testComplexNode = getTemplateArray`<h1>hello</h1><p>howdy</p>`;
  const testSkeleton = buildSkeleton(testComplexNode);

  console.log(testSkeleton);

  const assertions: string[] = ["fail automatically"];

  return assertions;
};

const findBrokenFromPlainText = () => {
  const testComplexNode = getTemplateArray`<${"hello"}h2>hey</h2><p>howdy</p>`;
  const testSkeleton = buildSkeleton(testComplexNode);

  console.log(testSkeleton);

  const assertions: string[] = ["fail automatically"];

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
