module.exports = {
    devServer: {
        open: true,
        port: 9527,
        proxy: {
            '/': {
              target: 'http://127.0.0.1:8080/',
              changeOrigin: true, // 是否跨域
              pathRewrite: {
                '^/': '' // 需要rewrite的,
              }
            },
        }
    },
    assetsDir: 'static',

}
