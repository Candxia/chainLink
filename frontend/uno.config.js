import { defineConfig, presetUno, presetAttributify } from 'unocss'
import transformerVariantGroup from '@unocss/transformer-variant-group'

export default defineConfig({
  presets: [
    presetUno(),
    presetAttributify(),
  ],
  transformers: [
    transformerVariantGroup(),
  ],
  shortcuts: [
    { 'flex-center': 'flex items-center justify-center' },
    { 'flex-between': 'flex items-center justify-between' },
  ],
})
