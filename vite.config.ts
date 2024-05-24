import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite'
import {ElementPlusResolver} from 'unplugin-vue-components/resolvers'
import viteCompression from 'vite-plugin-compression'
import IconsResolver from 'unplugin-icons/resolver'
import Icons from 'unplugin-icons/vite'

const INVALID_CHAR_REGEX = /[\x00-\x1F\x7F<>*#"{}|^[\]`;?:&_=+$,]/g;
const DRIVE_LETTER_REGEX = /^[a-z]:/i;

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    AutoImport({
      resolvers: [
        ElementPlusResolver(),
        IconsResolver({
          prefix: 'Icon'
        })
      ],
    }),
    Components({
      resolvers: [
        ElementPlusResolver(),
        IconsResolver({
          enabledCollections: ['ep']
        }),
      ],
    }),
    viteCompression(),
    Icons({
      autoInstall: true,
    }),
  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  },
  build: {
    target: ["es2015", "chrome63"],
    rollupOptions: {
      output: {
        sanitizeFileName(name) {
          const match = DRIVE_LETTER_REGEX.exec(name);
          let driveLetter = match ? match[0] : '';
          name = name.replaceAll('-','')
          name = name.replaceAll('_','')
          driveLetter = driveLetter.replace('-','')
          driveLetter = driveLetter.replace('_','')
          console.log('================',name)
          return (
              driveLetter +
              name.slice(driveLetter.length).replace('-','').replace('_','').replace(INVALID_CHAR_REGEX, "")
          );
        },
      },
    },
  },
})
