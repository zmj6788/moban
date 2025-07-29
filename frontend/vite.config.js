import {fileURLToPath, URL} from 'node:url'

import {defineConfig, loadEnv} from 'vite'
import vue from '@vitejs/plugin-vue'


export default ({mode}) => {
    const env = loadEnv(mode, process.cwd())
    const baseUrl = env.VITE_API
    const wsUrl = env.VITE_WEBSOCKET
    return defineConfig({
        envPrefix: ["VITE_"],  // 需要使用的前缀
        plugins: [vue()],
        resolve: {
            alias: {
                '@': fileURLToPath(new URL('./src', import.meta.url))
            },
        },
        server: {
            host: "127.0.0.1",
            port: 80,
            proxy: {
                "/uploads": {
                    target: baseUrl
                },
                "/api": {
                    target: baseUrl,
                    changeOrigin: true,
                },
                "/ws": {
                    target: wsUrl,
                    changeOrigin: true,
                    ws: true, //websocket代理设置
                    rewrite: (path) => path.replace(/^\/ws/, "")
                }
            }
        },
    })
}