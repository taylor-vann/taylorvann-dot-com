// brian taylor vann
// build integrals

import { StructureRender } from "../../../type_flyweight/structure";
import { create, incrementTarget } from "../../../text_vector/text_vector";

import { crawlForContent } from "./content_crawl";

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

const title = "content_crawl";
const runTestsAsynchronously = true;

const testEmptyString = () => {
  const assertions = [];

  const template = testTextInterpolator``;
  const vector = create();
  const results = crawlForContent(template, vector);

  if (results === undefined) {
    assertions.push("this should not have failed");
  }
  if (results !== undefined && results.action === "CREATE_CONTENT") {
    if (results.params.contentVector.origin.arrayIndex !== 0) {
      assertions.push("contentVector origin.arrayIndex should be 0.");
    }
    if (results.params.contentVector.origin.stringIndex !== 0) {
      assertions.push("contentVector origin.stringIndex should be 0.");
    }

    if (results.params.contentVector.target.arrayIndex !== 0) {
      assertions.push("contentVector target.arrayIndex should be 0.");
    }
    if (results.params.contentVector.target.stringIndex !== 0) {
      assertions.push("contentVector target.stringIndex should be 0.");
    }
  }

  return assertions;
};

const testLargeEmptyString = () => {
  const assertions = [];

  const template = testTextInterpolator`     `;
  const vector = create();
  let safety = 0;
  while (incrementTarget(template, vector) && safety < RECURSION_SAFETY) {
    safety += 1;
  }

  const results = crawlForContent(template, vector);

  if (results === undefined) {
    assertions.push("this should not have failed");
  }
  if (results !== undefined && results.action === "CREATE_CONTENT") {
    if (results.params.contentVector.origin.arrayIndex !== 0) {
      assertions.push("contentVector origin.arrayIndex should be 0.");
    }
    if (results.params.contentVector.origin.stringIndex !== 0) {
      assertions.push("contentVector origin.stringIndex should be 0.");
    }

    if (results.params.contentVector.target.arrayIndex !== 0) {
      assertions.push("contentVector target.arrayIndex should be 0.");
    }
    if (results.params.contentVector.target.stringIndex !== 4) {
      assertions.push("contentVector target.stringIndex should be 4.");
    }
  }

  return assertions;
};

const testInjectionString = () => {
  const assertions = [];

  const template = testTextInterpolator`${"hello, world!"}`;
  const vector = create();
  let safety = 0;
  while (incrementTarget(template, vector) && safety < RECURSION_SAFETY) {
    safety += 1;
  }
  const results = crawlForContent(template, vector);

  if (results === undefined) {
    assertions.push("this should not have failed");
  }
  if (results !== undefined && results.action === "CREATE_CONTENT") {
    if (results.params.contentVector.origin.arrayIndex !== 0) {
      assertions.push("contentVector origin.arrayIndex should be 0.");
    }
    if (results.params.contentVector.origin.stringIndex !== 0) {
      assertions.push("contentVector origin.stringIndex should be 0.");
    }

    if (results.params.contentVector.target.arrayIndex !== 1) {
      assertions.push("contentVector target.arrayIndex should be 1.");
    }
    if (results.params.contentVector.target.stringIndex !== 0) {
      assertions.push("contentVector target.stringIndex should be 0.");
    }
  }

  return assertions;
};

const testLargeInjectionString = () => {
  const assertions = [];

  const template = testTextInterpolator`     ${"hello, world!"}     `;
  const vector = create();
  let safety = 0;
  while (incrementTarget(template, vector) && safety < RECURSION_SAFETY) {
    safety += 1;
  }
  const results = crawlForContent(template, vector);

  if (results === undefined) {
    assertions.push("this should not have failed");
  }
  if (results !== undefined && results.action === "CREATE_CONTENT") {
    if (results.params.contentVector.origin.arrayIndex !== 0) {
      assertions.push("contentVector origin.arrayIndex should be 0.");
    }
    if (results.params.contentVector.origin.stringIndex !== 0) {
      assertions.push("contentVector origin.stringIndex should be 0.");
    }

    if (results.params.contentVector.target.arrayIndex !== 1) {
      assertions.push("contentVector target.arrayIndex should be 1.");
    }
    if (results.params.contentVector.target.stringIndex !== 4) {
      assertions.push("contentVector target.stringIndex should be 0.");
    }
  }

  return assertions;
};

const tests = [
  testEmptyString,
  testLargeEmptyString,
  testInjectionString,
  testLargeInjectionString,
];

const unitTestContentCrawl = {
  title,
  tests,
  runTestsAsynchronously,
};

export { unitTestContentCrawl };
