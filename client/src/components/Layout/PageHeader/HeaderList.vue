<template>
  <div class="w-21rem max-w-full z-20 p-2">
    <div class="rounded-lg border-1 border-gray-200 bg-white shadow-md">
      <div
        class="mt-2.5 pb-1.5 px-2.4 flex items-center border-b-1 border-gray-200 relative"
      >
        <span
          class="w-auto px-1.4 h-4.4 flex-initial rounded-lg bg-gray-200 text-xs text-center leading-1.1rem"
        >
          {{ itemLength }}
        </span>
        <h3 class="flex-1 pl-1.6 text-base font-semibold">{{ title }}</h3>
        <a @click="$emit('editorNew', null)"><add-icon /></a>

        <slot name="editor"></slot>
      </div>

      <div class="h-sm py-1.6 overflow-scroll">
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
