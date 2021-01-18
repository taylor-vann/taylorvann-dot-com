// brian taylor vann
// build integrals

import { samestuff } from "../../../little_test_runner/samestuff/samestuff";
import { Template } from "../../type_flyweight/template";
import { create, incrementTarget } from "../../text_vector/text_vector";
import { crawlForContent } from "./content_crawl";

type TextTextInterpolator = <A>(
  templateArray: TemplateStringsArray,
  ...injections: A[]
) => Template<A>;

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

  const expectedResults = {
    action: "CREATE_CONTENT",
    params: {
      contentVector: {
        origin: {
          arrayIndex: 0,
          stringIndex: 0,
        },
        target: {
          arrayIndex: 0,
          stringIndex: 0,
        },
      },
    },
  };

  const template = testTextInterpolator``;
  const vector = create();

  const results = crawlForContent(template, vector);

  if (!samestuff(expectedResults, results)) {
    assertions.push("unexpected results found.");
  }

  return assertions;
};

const testLargeEmptyString = () => {
  const assertions = [];

  const expectedResults = {
    action: "CREATE_CONTENT",
    params: {
      contentVector: {
        origin: {
          arrayIndex: 0,
          stringIndex: 0,
        },
        target: {
          arrayIndex: 0,
          stringIndex: 4,
        },
      },
    },
  };

  const template = testTextInterpolator`     `;
  const vector = create();

  let safety = 0;
  while (incrementTarget(template, vector) && safety < RECURSION_SAFETY) {
    safety += 1;
  }

  const results = crawlForContent(template, vector);

  if (!samestuff(expectedResults, results)) {
    assertions.push("unexpected results found.");
  }

  if (results === undefined) {
    assertions.push("this should not have failed");
  }

  return assertions;
};

const testInjectionString = () => {
  const assertions = [];

  const expectedResults = {
    action: "CREATE_CONTENT",
    params: {
      contentVector: {
        origin: {
          arrayIndex: 0,
          stringIndex: 0,
        },
        target: {
          arrayIndex: 1,
          stringIndex: 0,
        },
      },
    },
  };

  const template = testTextInterpolator`${"hello, world!"}`;
  const vector = create();

  let safety = 0;
  while (incrementTarget(template, vector) && safety < RECURSION_SAFETY) {
    safety += 1;
  }

  const results = crawlForContent(template, vector);

  if (!samestuff(expectedResults, results)) {
    assertions.push("unexpected results found.");
  }

  return assertions;
};

const testLargeInjectionString = () => {
  const assertions = [];

  const expectedResults = {
    action: "CREATE_CONTENT",
    params: {
      contentVector: {
        origin: {
          arrayIndex: 0,
          stringIndex: 0,
        },
        target: {
          arrayIndex: 1,
          stringIndex: 4,
        },
      },
    },
  };

  const template = testTextInterpolator`     ${"hello, world!"}     `;
  const vector = create();
  let safety = 0;
  while (incrementTarget(template, vector) && safety < RECURSION_SAFETY) {
    safety += 1;
  }

  const results = crawlForContent(template, vector);

  if (!samestuff(expectedResults, results)) {
    assertions.push("unexpected results found.");
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
