import * as MiniCSSExtractPlugin from "mini-css-extract-plugin";
import * as path from "path";
import { TypedCssModulesPlugin } from "typed-css-modules-webpack-plugin";

const WEB_CLIENT = "web_client";

// // html loader
// const htmlPluginWebClient: HtmlWebPackPlugin = new HtmlWebPackPlugin({
//   filename: `./${WEB_CLIENT}/dist/index.html`,
//   template: "./index_template.ejs",
//   title: "taylor vann",
// });

// create css plugin
const cssExtractPluginWebClient = new MiniCSSExtractPlugin({
  filename: `./${WEB_CLIENT}/dist/styles/${WEB_CLIENT}.css`,
  ignoreOrder: false, // Enable to remove warnings about conflicting order
});

// create Typed CSS Plugin for typescript
const typedCssModulesPlugin = new TypedCssModulesPlugin({
  globPattern: "src/components/**.css",
});

const sharedEntry = {
  web_client: `./${WEB_CLIENT}/src/app_root.tsx`,
};

const sharedModules = {
  rules: [
    {
      exclude: /node_modules/,
      test: /\.tsx?$/,
      use: "ts-loader",
    },
    {
      test: /\.worker\.ts$/,
      exclude: /node_modules/,
      use: [{ loader: "worker-loader" }, { loader: "ts-loader" }],
    },
    {
      test: /\.*.css$/,
      exclude: /node_modules/,
      include: /web_client\/src\/components/,
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

const sharedOutput = {
  filename: "./[name]/dist/scripts/[name].js",
  path: path.resolve(__dirname),
  publicPath: "/",
};

const sharedPlugins = [
  // htmlPluginWebClient,
  cssExtractPluginWebClient,
  typedCssModulesPlugin,
];

const sharedResolve = {
  extensions: [".js", ".jsx", ".ts", ".tsx"],
};

export {
  sharedEntry,
  sharedModules,
  sharedOutput,
  sharedPlugins,
  sharedResolve,
};
