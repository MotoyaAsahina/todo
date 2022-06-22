<template>
  <div class="px-3 py-1.5 hover:bg-gray-200 cursor-pointer">
    <p class="text-base"><slot></slot></p>
  </div>
</template>

<script lang="ts">
export default {
  name: 'TaskPanelMenuItem'
}
</script>

<style scoped></style>
