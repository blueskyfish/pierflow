import tailwindcss from '@tailwindcss/vite';
import react from '@vitejs/plugin-react';
import { resolve } from 'node:path';
import { defineConfig } from 'vite';

// https://vite.dev/config/
export default defineConfig({
  build: {
    outDir: '../cmd/serve/web',
    emptyOutDir: true,
  },
  server: {
    proxy: {
      '/api': {
        target: 'http://localhost:58080',
        changeOrigin: true,
        secure: false,
      },
    },
  },
  resolve: {
    alias: {
      '@blueskyfish/pierflow/app': resolve('src/app'),
      '@blueskyfish/pierflow/components': resolve('src/components'),
      '@blueskyfish/pierflow/pages': resolve('src/pages'),
      '@blueskyfish/pierflow/stores': resolve('src/stores'),
      '@blueskyfish/pierflow/utils': resolve('./src/utils'),
      // '@blueskyfish/pierflow/hooks': './src/hooks',
    },
  },
  plugins: [react(), tailwindcss()],
});
