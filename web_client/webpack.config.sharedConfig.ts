import { TypedCssModulesPlugin } from "typed-css-modules-webpack-plugin";
import * as HtmlWebPackPlugin from "html-webpack-plugin";
import * as MiniCSSExtractPlugin from "mini-css-extract-plugin";
import * as CompressionWebpackPlugin from "compression-webpack-plugin";

// html loader
const htmlPlugin: HtmlWebPackPlugin = new HtmlWebPackPlugin({
  filname: "index.html",
  template: "./src/index_template.ejs",
  title: "Three Demo",
});

// create css plugin
const cssExtractPlugin = new MiniCSSExtractPlugin({
  filename: "styles/[name].css",
  chunkFilename: "styles/[id].css",
  ignoreOrder: false, // Enable to remove warnings about conflicting order
});

// create Typed CSS Plugin for typescript
const typedCssModulesPlugin = new TypedCssModulesPlugin({
  globPattern: "src/components/*.css",
});

const compressionPlugin = new CompressionWebpackPlugin({
  filename: "compressed/bundled_app.gzip",
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
      exclude: /node_modules/,
      include: /src\/components/,
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
      exclude: /node_modules/,
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
      exclude: /node_modules/,
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

const sharedPlugins = [
  htmlPlugin,
  cssExtractPlugin,
  typedCssModulesPlugin,
  compressionPlugin,
];

export { sharedModuleConfig, sharedPlugins };
