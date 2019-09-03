const webpack = require('webpack');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const meteorExternals = require('webpack-meteor-externals');
const path = require('path');
const ExtractTextPlugin = require("extract-text-webpack-plugin");
const clientConfig = {
    entry: './client/main.js',
    module: {
      rules: [
        {
          test: /\.(js|jsx)$/,
          exclude: /node_modules/,
          loader: 'babel-loader',
          options: {
            presets: [
            '@babel/preset-env',
            '@babel/react',
            {
              'plugins': ['@babel/plugin-proposal-class-properties']
            }]
          }
        },
        {
          test: /\.css$/,
          use: [ 'style-loader', 'css-loader' ]
        },
        {
          test: /\.(png|jpg|gif|svg|eot|ttf|woff|woff2)$/,
          use: ['url-loader']
        }
      ]
    },
    plugins: [
      new HtmlWebpackPlugin({
          template: './client/main.html'
      }),
      new webpack.HotModuleReplacementPlugin()
    ],
    resolve: {
      extensions: ['*', '.js', '.jsx'],
      modules: [
        path.resolve(__dirname, 'node_modules'),
        path.resolve(__dirname, './')
      ],
      alias: {
        imports: path.resolve(__dirname, './imports'),
        ui: path.resolve(__dirname, './imports/ui'),
        api: path.resolve(__dirname, './imports/api'),
        client: path.resolve(__dirname, './client'),
        assets: path.resolve(__dirname, './assets'),
        store: path.resolve(__dirname, './imports/store')
      }
    },
    output: {
      path: __dirname + '/dist',
      publicPath: '/',
      filename: 'bundle.js'
    },
    devServer: {
      contentBase: './dist',
      hot: true
    },
    externals: [
      meteorExternals()
    ]
};
const serverConfig = {
    entry: [
        './server/main.js'
    ],
    target: 'node',
    externals: [
      meteorExternals()
    ],
    resolve: {
      extensions: ['*', '.js'],
      modules: [
        path.resolve(__dirname, 'node_modules'),
        path.resolve(__dirname, './')
      ],
      alias: {
        imports: path.resolve(__dirname, './imports'),
        ui: path.resolve(__dirname, './imports/ui'),
        api: path.resolve(__dirname, './imports/api'),
      }
    },
};
module.exports = [clientConfig, serverConfig];
