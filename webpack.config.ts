import * as webpack from "webpack";
import * as HtmlWebPackPlugin from "html-webpack-plugin";
import * as MiniCSSExtractPlugin from "mini-css-extract-plugin";
import * as path from "path";

// html loader
const htmlPlugin = new HtmlWebPackPlugin({
  template: "./src/index.html",
});

// create css plugin
const cssExtractPlugin = new MiniCSSExtractPlugin({
  filename: "styles/[name].css",
  chunkFilename: "styles/[id].css",
  ignoreOrder: false, // Enable to remove warnings about conflicting order
});

// webpack config for a react app in typescript
const webpackConfig: webpack.Configuration = {
  mode: "production",
  entry: "./src/app_root.tsx",
  output: {
    filename: "js/bundled_typescript.js",
    path: path.resolve(__dirname, "dist"),
  },
  devtool: "inline-source-map",
  resolve: {
    extensions: [".js", ".ts", ".tsx"],
  },
  module: {
    rules: [
      {
        exclude: /node_modules/,
        test: /\.tsx?$/,
        use: "ts-loader",
      },
      {
        exclude: /node_modules/,
        test: /\.css$/,
        use: [
          { loader: "style-loader" },
          {
            loader: "css-loader",
            options: {
              modules: true,
            },
          },
        ],
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
              sourceMap: true,
            },
          },
        ],
      },
      {
        test: /\.(woff|woff2|eot|ttf|otf)$/,
        use: [
          {
            loader: "url-loader",
            options: {
              limit: 10240,
              name: "styles/fonts/[name]__[hash].[ext]",
            },
          },
        ],
      },
      {
        test: /\.(png|jp(e*)g|svg)$/,
        use: [
          {
            loader: "url-loader",
            options: {
              limit: 30720,
              name: "images/[name]__[hash].[ext]",
            },
          },
        ],
      },
    ],
  },
  plugins: [htmlPlugin, cssExtractPlugin],
};

module.exports = webpackConfig;
