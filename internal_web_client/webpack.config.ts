import * as OptimizeCSSAssetsPlugin from "optimize-css-assets-webpack-plugin";
import * as TerserJSPlugin from "terser-webpack-plugin";
import * as webpack from "webpack";

import {
  sharedModules,
  sharedPlugins,
  sharedOutput,
  sharedEntry,
  sharedResolve,
} from "./webpack.config.sharedConfig";

// general optimization
const terserPlugin = new TerserJSPlugin({
  sourceMap: true,
  extractComments: false,
});

// css optimization
const optimizeCssPlugin = new OptimizeCSSAssetsPlugin({});

// webpack config for production
const webpackConfig: webpack.Configuration = {
  devtool: "source-map",
  entry: sharedEntry,
  mode: "production",
  module: sharedModules,
  optimization: {
    minimize: true,
    minimizer: [terserPlugin, optimizeCssPlugin],
  },
  output: sharedOutput,
  plugins: sharedPlugins,
  resolve: sharedResolve,
};

module.exports = webpackConfig;
