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
  resolve: {
    alias: {
      '@blueskyfish/pierflow/app': resolve('src/app'),
      '@blueskyfish/pierflow/components': resolve('src/components'),
      '@blueskyfish/pierflow/pages': resolve('src/pages'),
      // '@blueskyfish/pierflow/hooks': './src/hooks',
      // '@blueskyfish/pierflow/utils': './src/utils',
      // '@blueskyfish/pierflow/store': './src/store',
    },
  },
  plugins: [react(), tailwindcss()],
});
