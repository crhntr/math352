const path = require("path");
const webpack = require('webpack');

module.exports = {
    context: __dirname,

    entry: {
        main: './frontend/main',
    },

    output: {
        path: path.resolve('../static/bundles/'),
        filename: "[name].js",
    },

    module: {
        loaders: [
            {
                test: /\.vue$/,
                exclude: /node_modules/,
                loader: 'vue-loader'
            },
            {
                test: /\.jsx?$/,
                exclude: /node_modules/,
                loader: 'babel-loader'
            }
        ]
    },

    resolve: {
        alias: {'vue$': 'vue/dist/vue.esm.js'},
        // modulesDirectories: ['node_modules', 'bower_components'],
        extensions: ['.vue', '.js', '.jsx']
    }
};
