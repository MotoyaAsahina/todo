<template>
  <div class="bg-red-200 h-5 cursor-pointer px-2 rounded-xl">
    <p class="text-sm font-medium">{{ tagName }}</p>
  </div>
</template>

<script>
export default {
  name: 'TaskTag',
  props: {
    tagName: {
      type: String,
      required: true
    }
  }
}
</script>

<style scoped></style>
