const path = require("path");
const MiniCSSExtractPlugin = require("mini-css-extract-plugin");
const HtmlWebPackPlugin = require('html-webpack-plugin');

// html loader
const html_plugin = new HtmlWebPackPlugin({
  template: "./src/index.html"
});

// create css plugin
const css_extract_plugin = new MiniCSSExtractPlugin({
  filename: "[name].css",
  chunkFilename: "[id].css",
  ignoreOrder: false // Enable to remove warnings about conflicting order
});

// webpack config for a react app in typescript
const webpack_config = {
  mode: "production",
  entry: "./src/app_root.tsx",
  devtool: "inline-source-map",
  resolve: {
    extensions: [".js", ".ts", ".tsx"]
  },
  module: {
    rules: [
      {
        exclude: /node_modules/,
        test: /\.tsx?$/,
        use: "ts-loader"
      },
      {
        exclude: /node_modules/,
        test: /\.css$/,
        use: [
          { loader: "style-loader" },
          {
            loader: "css-loader",
            options: {
              modules: true
            }
          }
        ]
      },
      {
        test: /\.css$/,
        use: [
          MiniCSSExtractPlugin.loader,
          {
            loader: "css-loader",
            options: {
              localIdentName: "[name]__[local]___[hash:base64:5]",
              modules: true,
              sourceMap: true
            }
          }
        ]
      },
      {
        test: /\.(woff|woff2|eot|ttf|otf)$/,
        use: [
          {
            loader: "url-loader",
            options: {
              limit: 10240,
              name: "styles/fonts/[name]__[hash].[ext]"
            }
          }
        ]
      },
      {
        test: /\.(png|jp(e*)g|svg)$/,
        use: [
          {
            loader: "url-loader",
            options: {
              limit: 30720,
              name: "images/[name]__[hash].[ext]"
            }
          }
        ]
      }
    ]
  },
  output: {
    filename: "bundled_typescript.js",
    path: path.resolve(__dirname, "dist/js")
  },
  plugins: [html_plugin, css_extract_plugin]
};

module.exports = webpack_config;
