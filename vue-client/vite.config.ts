import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

export default defineConfig({
  plugins: [vue()],
  optimizeDeps: {
    include: ['grpc-web', 'google-protobuf']
  },
  build: {
    commonjsOptions: {
      include: [/grpc-web/, /google-protobuf/, /src\/gen/]
    }
  }
})