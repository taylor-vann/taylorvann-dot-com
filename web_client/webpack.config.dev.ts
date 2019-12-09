import * as webpack from "webpack";

import {
  sharedModules,
  sharedPlugins,
  sharedOutput,
  sharedEntry,
  sharedResolve,
} from "./webpack.config.sharedConfig";

// webpack config for a react app in typescript
const webpackConfig: webpack.Configuration = {
  devtool: "inline-source-map",
  entry: sharedEntry,
  mode: "development",
  module: sharedModules,
  output: sharedOutput,
  plugins: sharedPlugins,
  resolve: sharedResolve,
};

module.exports = webpackConfig;
