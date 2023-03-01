const { defineConfig } = require('cypress')
const {webPackConfig} = require('../frontend/node_modules/react-scripts/config/webpackDevServer.config')
module.exports = defineConfig({
  component: {
    devServer: {
      framework: 'create-react-app',
      webpackConfig: webPackConfig,
      bundler: 'webpack'

    },
  },
})