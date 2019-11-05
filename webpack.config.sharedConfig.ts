import { TypedCssModulesPlugin } from "typed-css-modules-webpack-plugin";
import * as HtmlWebPackPlugin from "html-webpack-plugin";
import * as MiniCSSExtractPlugin from "mini-css-extract-plugin";

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

// create Typed CSS Plugin for typescript
const typedCssModulesPlugin = new TypedCssModulesPlugin({
  globPattern: "src/**/*.css",
});

const sharedModuleConfig = {
  rules: [
    {
      exclude: /node_modules/,
      test: /\.tsx?$/,
      use: "ts-loader",
    },
    {
      test: /\.css$/,
      use: [
        MiniCSSExtractPlugin.loader,
        {
          loader: "css-loader",
          options: {
            modules: {
              localIdentName: "[name]__[local]___[hash:base64:5]",
            },
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
            name: "fonts/[name]__[hash].[ext]",
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
};

const sharedPlugins = [htmlPlugin, cssExtractPlugin, typedCssModulesPlugin];

export { sharedModuleConfig, sharedPlugins };
