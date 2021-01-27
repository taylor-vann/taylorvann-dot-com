// brian taylor vann
// test hooks

// make sure we can add and remove properties

const title = "test_hooks";
const runTestsAsynchronously = true;

const defaultFunc = () => {
  return ["fail automatically"];
};

const tests = [defaultFunc];

const unitTestBuildRender = {
  title,
  tests,
  runTestsAsynchronously,
};

export { unitTestBuildRender };
