// brian taylor vann

import { StructureRender } from "../../type_flyweight/render";
import { SkeletonNodes, buildSkeleton } from "./build_skeleton";

type TextTextInterpolator = <A>(
  templateArray: TemplateStringsArray,
  ...injections: A[]
) => StructureRender<A>;

type CompareSkeletons = (
  source: SkeletonNodes,
  target: SkeletonNodes
) => boolean;

const title = "build_skeleton";
const runTestsAsynchronously = true;

const testTextInterpolator: TextTextInterpolator = (
  templateArray,
  ...injections
) => {
  return { templateArray, injections };
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
      node.vector.origin.arrayIndex !== targetNode.vector.origin.arrayIndex ||
      node.vector.origin.stringIndex !== targetNode.vector.origin.stringIndex ||
      node.vector.target.arrayIndex !== targetNode.vector.target.arrayIndex ||
      node.vector.target.stringIndex !== targetNode.vector.target.stringIndex
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
      vector: {
        target: { arrayIndex: 0, stringIndex: 20 },
        origin: { arrayIndex: 0, stringIndex: 0 },
      },
    },
  ];

  const testBlank = testTextInterpolator`no nodes to be found!`;
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
      vector: {
        target: { arrayIndex: 0, stringIndex: 2 },
        origin: { arrayIndex: 0, stringIndex: 0 },
      },
    },
  ];
  const testOpenNode = testTextInterpolator`<p>`;
  const testSkeleton = buildSkeleton(testOpenNode);
  if (!compareSkeletons(sourceSkeleton, testSkeleton)) {
    assertions.push("skeletons are not equal");
  }

  return assertions;
};

const findComplexFromPlainText = () => {
  const assertions: string[] = [];
  const sourceSkeleton: SkeletonNodes = [
    {
      nodeType: "CONTENT_NODE",
      vector: {
        target: { arrayIndex: 0, stringIndex: 4 },
        origin: { arrayIndex: 0, stringIndex: 0 },
      },
    },
    {
      nodeType: "OPEN_NODE_CONFIRMED",
      vector: {
        target: { arrayIndex: 0, stringIndex: 7 },
        origin: { arrayIndex: 0, stringIndex: 5 },
      },
    },
    {
      nodeType: "CONTENT_NODE",
      vector: {
        target: { arrayIndex: 0, stringIndex: 12 },
        origin: { arrayIndex: 0, stringIndex: 8 },
      },
    },
    {
      nodeType: "CLOSE_NODE_CONFIRMED",
      vector: {
        target: { arrayIndex: 0, stringIndex: 16 },
        origin: { arrayIndex: 0, stringIndex: 13 },
      },
    },
  ];
  const testComplexNode = testTextInterpolator`hello<p>world</p>`;
  const testSkeleton = buildSkeleton(testComplexNode);

  if (!compareSkeletons(sourceSkeleton, testSkeleton)) {
    assertions.push("skeletons are not equal");
  }

  return assertions;
};

const findCompoundFromPlainText = () => {
  const assertions: string[] = [];
  const sourceSkeleton: SkeletonNodes = [
    {
      nodeType: "OPEN_NODE_CONFIRMED",
      vector: {
        target: { arrayIndex: 0, stringIndex: 3 },
        origin: { arrayIndex: 0, stringIndex: 0 },
      },
    },
    {
      nodeType: "CONTENT_NODE",
      vector: {
        target: { arrayIndex: 0, stringIndex: 8 },
        origin: { arrayIndex: 0, stringIndex: 4 },
      },
    },
    {
      nodeType: "CLOSE_NODE_CONFIRMED",
      vector: {
        target: { arrayIndex: 0, stringIndex: 13 },
        origin: { arrayIndex: 0, stringIndex: 9 },
      },
    },
  ];
  const testComplexNode = testTextInterpolator`<h1>hello</h1>`;
  const testSkeleton = buildSkeleton(testComplexNode);

  if (!compareSkeletons(sourceSkeleton, testSkeleton)) {
    assertions.push("skeletons are not equal");
  }

  return assertions;
};

const findBrokenFromPlainText = () => {
  const assertions: string[] = [];
  const sourceSkeleton: SkeletonNodes = [
    {
      nodeType: "CONTENT_NODE",
      vector: {
        target: { arrayIndex: 1, stringIndex: 5 },
        origin: { arrayIndex: 0, stringIndex: 0 },
      },
    },
    {
      nodeType: "CLOSE_NODE_CONFIRMED",
      vector: {
        target: { arrayIndex: 1, stringIndex: 10 },
        origin: { arrayIndex: 1, stringIndex: 6 },
      },
    },
    {
      nodeType: "OPEN_NODE_CONFIRMED",
      vector: {
        target: { arrayIndex: 1, stringIndex: 13 },
        origin: { arrayIndex: 1, stringIndex: 11 },
      },
    },
    {
      nodeType: "CONTENT_NODE",
      vector: {
        target: { arrayIndex: 1, stringIndex: 18 },
        origin: { arrayIndex: 1, stringIndex: 14 },
      },
    },
    {
      nodeType: "CLOSE_NODE_CONFIRMED",
      vector: {
        target: { arrayIndex: 1, stringIndex: 22 },
        origin: { arrayIndex: 1, stringIndex: 19 },
      },
    },
  ];
  const testComplexNode = testTextInterpolator`<${"hello"}h2>hey</h2><p>howdy</p>`;
  const testSkeleton = buildSkeleton(testComplexNode);
  if (!compareSkeletons(sourceSkeleton, testSkeleton)) {
    assertions.push("skeletons are not equal");
  }

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
