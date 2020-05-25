import * as webpack from "webpack";
import * as WebClientConfig from "./web_client.webpack.sharedConfig";
import * as WebClientLoginConfig from "./web_client_login.webpack.sharedConfig";

const webpackConfig: webpack.Configuration = {
  name: "web_client",
  devtool: "source-map",
  entry: WebClientConfig.sharedEntry,
  mode: "development",
  module: WebClientConfig.sharedModules,
  output: WebClientConfig.sharedOutput,
  plugins: WebClientConfig.sharedPlugins,
  resolve: WebClientConfig.sharedResolve,
};

const webpackWebClientLoginConfig: webpack.Configuration = {
  name: "web_client_login",
  devtool: "source-map",
  entry: WebClientLoginConfig.sharedEntry,
  mode: "development",
  module: WebClientLoginConfig.sharedModules,
  output: WebClientLoginConfig.sharedOutput,
  plugins: WebClientLoginConfig.sharedPlugins,
  resolve: WebClientLoginConfig.sharedResolve,
};

module.exports = [webpackConfig, webpackWebClientLoginConfig];
