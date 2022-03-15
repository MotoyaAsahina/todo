<template>
  <div class="w-xs rounded-lg border-1 border-gray-300 bg-white shadow-md">
    <div class="p-3">
      <slot></slot>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
  name: 'TaskEditor'
})
</script>

<style scoped></style>
