<template>
  <div class="w-sm max-w-full z-10 p-2">
    <div class="rounded-lg border-1 border-gray-300 bg-white shadow-md">
      <div
        class="mt-2.5 pb-1.5 px-3 flex items-center border-b-1 border-gray-300"
      >
        <span
          class="w-5 h-5 flex-initial rounded-lg bg-gray-200 text-sm text-center"
        >
          {{ itemLength }}
        </span>
        <h3 class="flex-1 pl-2 font-semibold">{{ title }}</h3>
        <add-icon />
      </div>

      <div class="h-lg px-2 overflow-scroll">
        <slot></slot>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import AddIcon from '/@/components/UI/AddIcon.vue'

export default {
  name: 'HeadEditor',
  components: {
    AddIcon
  },
  props: {
    title: {
      type: String,
      required: true
    },
    itemLength: {
      type: Number,
      required: true
    }
  }
}
</script>

<style scoped></style>
