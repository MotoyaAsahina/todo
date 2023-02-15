import { defineConfig } from 'windicss/helpers'
import scrollSnapPlugin from 'windicss/plugin/scroll-snap'

export default defineConfig({
  plugins: [scrollSnapPlugin],
  theme: {
    extend: {
      fontSize: {
        base: '0.9rem',
        sm: '0.8rem'
      }
    }
  }
})
