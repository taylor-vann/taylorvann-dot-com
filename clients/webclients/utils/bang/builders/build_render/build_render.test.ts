// brian taylor vann
// build structure

const title = "build_render";
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
