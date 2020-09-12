export default [{
  input: "./build/main.js",
  output: {
    file: "./home/scripts/bundled.js",
  },
}, {
  input: "./build/test.js",
  output: {
    file: "./home/tests/bundled.test.js",
  },
}];
