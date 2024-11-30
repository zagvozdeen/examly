import { defineConfig, loadEnv } from 'vite'
import vue from '@vitejs/plugin-vue'
import path from 'path'

// https://vitejs.dev/config/
export default () => {
  const mode = process.env.NODE_ENV || 'development'
  process.env = { ...process.env, ...loadEnv(mode, process.cwd()) }

  return defineConfig({
    server: {
      host: true,
      origin: process.env.VITE_APP_URL,
    },
    // base: process.env.VITE_APP_URL,
    plugins: [vue()],
    resolve: {
      alias: {
        '@': path.resolve(__dirname, 'ts'),
        '@images': path.resolve(__dirname, 'images'),
      },
    },
  })
}
