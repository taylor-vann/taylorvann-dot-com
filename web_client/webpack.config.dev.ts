import * as path from "path";
import * as webpack from "webpack";

import {
  sharedModuleConfig,
  sharedPlugins,
} from "./webpack.config.sharedConfig";

// webpack config for a react app in typescript
const webpackConfig: webpack.Configuration = {
  mode: "development",
  devtool: "inline-source-map",
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
};

module.exports = webpackConfig;
