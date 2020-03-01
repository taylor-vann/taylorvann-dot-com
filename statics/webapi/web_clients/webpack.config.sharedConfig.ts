import * as HtmlWebPackPlugin from "html-webpack-plugin";
import * as MiniCSSExtractPlugin from "mini-css-extract-plugin";
import * as path from "path";
import { TypedCssModulesPlugin } from "typed-css-modules-webpack-plugin";

// html loader
const htmlPlugin: HtmlWebPackPlugin = new HtmlWebPackPlugin({
  filname: "index.html",
  template: "./index_template.ejs",
  title: "Three Demo",
});

// create css plugin
const cssExtractPlugin = new MiniCSSExtractPlugin({
  filename: "./[name]/dist/styles/bundled_styles.css",
  chunkFilename: "./[name]/dist/styles/chunks/[id].css",
  ignoreOrder: false, // Enable to remove warnings about conflicting order
});

// create Typed CSS Plugin for typescript
const typedCssModulesPlugin = new TypedCssModulesPlugin({
  globPattern: "src/components/**.css",
});

const sharedEntry = {
  web_client: "./web_client/src/app_root.tsx",
  internal_web_client: "./internal_web_client/src/app_root.tsx",
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

const sharedOutput = {
  filename: "./[name]/dist/js/bundled_typesecript.js",
  path: path.resolve(__dirname),
};

const sharedPlugins = [htmlPlugin, cssExtractPlugin, typedCssModulesPlugin];

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
