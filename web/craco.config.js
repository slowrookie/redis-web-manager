const MonacoWebpackPlugin = require('monaco-editor-webpack-plugin');

module.exports = {
  devServer: {
    devMiddleware: {
      writeToDisk: false,
    },
  },
  webpack: {
    plugins: {
      add: [new MonacoWebpackPlugin({languages: ['json', 'lua']})]
    }
  }
};