import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import vueJsx from '@vitejs/plugin-vue-jsx'

// export default is used to export the configuration object
export default defineConfig({
  // Use the Vue and Vue JSX plugins
  plugins: [
    vue(),
    vueJsx(),
  ],
  // Set the base path to the current directory
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  },
  // Set the port to 3000
  server: {
    port: 3000,

    // Enable strict port checking, which will block the server from starting if the port is already in use
    strictPort: true
  }
})
