import * as OptimizeCSSAssetsPlugin from "optimize-css-assets-webpack-plugin";
import * as TerserJSPlugin from "terser-webpack-plugin";
import * as webpack from "webpack";
import * as WebClientConfig from "./web_client.webpack.sharedConfig";
import * as WebClientLoginConfig from "./web_client_login.webpack.sharedConfig";

// general optimization
const terserPlugin = new TerserJSPlugin({
  sourceMap: true,
  extractComments: false,
});

// css optimization
const optimizeCssPlugin = new OptimizeCSSAssetsPlugin({});

const webpackConfig: webpack.Configuration = {
  name: "web_client",
  devtool: "source-map",
  entry: WebClientConfig.sharedEntry,
  mode: "production",
  module: WebClientConfig.sharedModules,
  optimization: {
    minimize: true,
    minimizer: [terserPlugin, optimizeCssPlugin],
  },
  output: WebClientConfig.sharedOutput,
  plugins: WebClientConfig.sharedPlugins,
  resolve: WebClientConfig.sharedResolve,
};

const webpackWebClientLoginConfig: webpack.Configuration = {
  name: "web_client_login",
  devtool: "source-map",
  entry: WebClientLoginConfig.sharedEntry,
  mode: "production",
  module: WebClientLoginConfig.sharedModules,
  optimization: {
    minimize: true,
    minimizer: [terserPlugin, optimizeCssPlugin],
  },
  output: WebClientLoginConfig.sharedOutput,
  plugins: WebClientLoginConfig.sharedPlugins,
  resolve: WebClientLoginConfig.sharedResolve,
};

module.exports = [webpackConfig, webpackWebClientLoginConfig];
