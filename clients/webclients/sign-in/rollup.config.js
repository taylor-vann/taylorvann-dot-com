export default [
  {
    input: "./build/main.js",
    output: {
      file: "./sign-in/scripts/bundled.js",
    },
  },
  {
    input: "./build/tests.js",
    output: {
      file: "./sign-in/tests/bundled.tests.js",
    },
  },
];
