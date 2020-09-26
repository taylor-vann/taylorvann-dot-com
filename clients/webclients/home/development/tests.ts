// brian taylor vann
// home - tests

import { runTests } from "../../utils/little_test_runner/little_test_runner";

const title = "first unit test";

const littleUnitTest = () => {
  return ["fail immediately"];
};

const firstUnitTests = {
  title,
  tests: [littleUnitTest],
};

runTests({ testCollection: [firstUnitTests] })
  .then((results) => console.log("results: ", results))
  .catch((errors) => console.log("errors: ", errors));
