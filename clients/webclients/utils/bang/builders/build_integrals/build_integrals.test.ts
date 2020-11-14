// brian taylor vann
// build structure

const title = "build_structure";
const runTestsAsynchronously = true;

const defaultFunc = () => {
  return ["fail automatically"];
};

const tests = [defaultFunc];

const unitTestBuildIntegrals = {
  title,
  tests,
  runTestsAsynchronously,
};

export { unitTestBuildIntegrals };
