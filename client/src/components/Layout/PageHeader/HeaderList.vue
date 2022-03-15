<template>
  <div class="w-sm max-w-full z-20 p-2">
    <div class="rounded-lg border-1 border-gray-300 bg-white shadow-md">
      <div
        class="mt-2.5 pb-1.5 px-3 flex items-center border-b-1 border-gray-300 relative"
      >
        <span
          class="w-5 h-5 flex-initial rounded-lg bg-gray-200 text-sm text-center"
        >
          {{ itemLength }}
        </span>
        <h3 class="flex-1 pl-2 font-semibold">{{ title }}</h3>
        <a @click="$emit('editorNew', null)"><add-icon /></a>

        <slot name="editor"></slot>
      </div>

      <div class="h-sm py-2 overflow-scroll">
        <slot name="list"></slot>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import AddIcon from '/@/components/UI/AddIcon.vue'

export default {
  name: 'HeaderList',
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
  },
  emits: ['editorNew']
}
</script>

<style scoped></style>
