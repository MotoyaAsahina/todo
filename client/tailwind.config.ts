import { defineConfig } from 'vite-plugin-windicss'
import forms from 'windicss/plugin/forms'
import lineClamp from 'windicss/plugin/line-clamp'

export default defineConfig({
  darkMode: 'class',
  plugins: [forms, lineClamp]
})
