<template>
  <div
    class="h-1.1rem cursor-pointer px-2 rounded-xl inline-block"
    :style="{
      'background-color': tag?.color,
      color: pickBlackOrWhite(tag?.color)
    }"
  >
    <p class="text-xs leading-1.1rem font-medium">{{ tag?.name }}</p>
  </div>
</template>

<script lang="ts">
import { PropType } from 'vue'
import { Tag } from '/@/lib/apis'

export default {
  name: 'TaskTag',
  props: {
    tag: {
      type: Object as PropType<Tag>,
      required: true
    }
  },
  setup() {
    const pickBlackOrWhite = (color: string) => {
      if (!color) return '#ffffff'
      let colorStr = color.slice(-6)
      let r = parseInt(colorStr.slice(0, 2), 16)
      let g = parseInt(colorStr.slice(2, 4), 16)
      let b = parseInt(colorStr.slice(4, 6), 16)
      let yiq = (r * 299 + g * 587 + b * 114) / 1000
      return yiq >= 128 ? '#000000' : '#ffffff'
    }

    return { pickBlackOrWhite }
  }
}
</script>

<style scoped></style>
