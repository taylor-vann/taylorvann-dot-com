import * as OptimizeCSSAssetsPlugin from "optimize-css-assets-webpack-plugin";
import * as path from "path";
import * as TerserJSPlugin from "terser-webpack-plugin";
import * as webpack from "webpack";

import {
  sharedModuleConfig,
  sharedPlugins,
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
  mode: "production",
  devtool: "source-map",
  resolve: {
    extensions: [".js", ".jsx", ".ts", ".tsx"],
  },
  entry: "./src/app_root.tsx",
  output: {
    filename: "js/bundled_typescript.js",
    path: path.resolve(__dirname, "dist"),
  },
  plugins: sharedPlugins,
  module: sharedModuleConfig,
  optimization: {
    minimize: true,
    minimizer: [terserPlugin, optimizeCssPlugin],
  },
};

module.exports = webpackConfig;
