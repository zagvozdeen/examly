import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import path from 'path'

// https://vitejs.dev/config/
export default defineConfig({
  server: {
    host: true,
    origin: 'http://localhost:5173',
  },
  base: 'https://examly.ru/dist/',
  plugins: [vue()],
  resolve: {
    alias: {
      '@': path.resolve(__dirname, 'ts'),
      '@images': path.resolve(__dirname, 'images'),
    },
  },
})
