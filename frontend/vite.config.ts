import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import styleImport from 'vite-plugin-style-import'
import {resolve} from 'path'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    styleImport({
      libs: [
        {
          libraryName: 'element-plus',
          esModule: true,
          ensureStyleFile: true,
          resolveStyle: (name) => {
            return `element-plus/lib/theme-chalk/${name}.css`;
          },
          resolveComponent: (name) => {
            return `element-plus/lib/${name}`;
          },
        }
      ]
    })
  ],

  /**
   * 在生产中服务时的基本公共路径。
   * @default '/'
   */
  base: './',
  /**
  * 与“根”相关的目录，构建输出将放在其中。如果目录存在，它将在构建之前被删除。
  * @default 'dist'
  */
  // outDir: 'dist',
  server: {
    // hostname: '0.0.0.0',
    host: "localhost",
    port: 3001,
    // // 是否自动在浏览器打开
    // open: true,
    // // 是否开启 https
    // https: false,
    // // 服务端渲染
    // ssr: false,
    proxy: {
      '/api': {
        target: 'http://localhost:8888',
        changeOrigin: true,
        // ws: true,
        rewrite: (pathStr) => pathStr.replace('/api', '/api')
      },
    },
  },
  resolve: {
    // 导入文件夹别名
    alias: {
      '@': resolve(__dirname, './src'),
      views: resolve(__dirname, './src/views'),
      components: resolve(__dirname, './src/components'),
      utils: resolve(__dirname, './src/utils'),
      less: resolve(__dirname, "./src/less"),
      assets: resolve(__dirname, "./src/assets"),
      com: resolve(__dirname, "./src/components"),
      store: resolve(__dirname, "./src/store"),
      mixins: resolve(__dirname, "./src/mixins")
    },
  }
})
