// brian taylor vann
// home - tests

import { runTests } from "../../utils/little_test_runner/little_test_runner";

import { tests } from "../../utils/bang/bang.test";

const testCollection = [...tests];

runTests({ testCollection })
  .then((results) => console.log("results: ", results))
  .catch((errors) => console.log("errors: ", errors));
