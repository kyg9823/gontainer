import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vite.dev/config/
export default defineConfig({
  base: '/gontainer',
  build: {
    outDir: '../static'
  },
  plugins: [react()],
  server: {
    proxy: {
      '/gontainer/api': {
        target: 'http://localhost:8080',
        changeOrigin: true,
      }
    }
  }
})
