<template>
  <div class="w-18rem rounded-lg border-1 border-gray-200 bg-white shadow-md">
    <div class="p-2.4">
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
