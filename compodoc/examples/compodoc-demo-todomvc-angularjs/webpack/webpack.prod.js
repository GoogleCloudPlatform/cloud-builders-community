var loaders = require("./loaders"),
    HtmlWebpackPlugin = require('html-webpack-plugin'),
    webpack = require('webpack'),
    sass = require('node-sass'),
    fs = require('fs-extra');

sass.render({
    file: 'src/main.scss',
    outputStyle: 'compressed'
}, function(error, result) {
    if (error) {
        console.log(error.status);
        console.log(error.column);
        console.log(error.message);
        console.log(error.line);
    } else {
        fs.outputFile('dist/style.css', result.css.toString(), function (err) {
            if (err) return console.log(err);
            console.log('Sass compiled');
        });
    }
});

module.exports = {
    entry: ['./src/main.ts'],
    output: {
        filename: 'build.js',
        path: 'dist'
    },
    devtool: 'source-map',
    resolve: {
        root: __dirname,
        extensions: ['', '.ts', '.js', '.json']
    },
    resolveLoader: {
        modulesDirectories: ["node_modules"]
    },
    plugins: [
        new webpack.optimize.UglifyJsPlugin(
            {
                warning: false,
                mangle: true,
                comments: false
            }
        ),
        new HtmlWebpackPlugin({
            template: './src/index.html',
            inject: 'body',
            hash: true
        })
    ],
    module:{
        loaders: loaders
    }
};
