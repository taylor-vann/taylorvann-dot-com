// brian taylor vann
// build structure

const title = "build_structure";
const runTestsAsynchronously = true;

const defaultFunc = () => {
  return ["fail automatically"];
};

const tests = [defaultFunc];

const unitTestBuildStructure = {
  title,
  tests,
  runTestsAsynchronously,
};

export { unitTestBuildStructure };
