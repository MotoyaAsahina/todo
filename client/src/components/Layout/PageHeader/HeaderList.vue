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
        <a><add-icon /></a>

        <div class="absolute top-6.2 right-1">
          <div
            class="w-60 p-3 bg-white rounded-lg border-1 border-gray-300 shadow-md"
          >
            <div class="mb-2">
              <h3 class="font-semibold">New</h3>
            </div>
            <textarea
              class="w-full resize-none p-1 text-sm"
              rows="3"
            ></textarea>
            <div class="flex items-center justify-end mt-2">
              <a><close-icon /></a>
              <a><check-icon class="ml-0.5" /></a>
            </div>
          </div>
        </div>
      </div>

      <div class="h-sm py-2 overflow-scroll">
        <slot></slot>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import AddIcon from '/@/components/UI/AddIcon.vue'
import CheckIcon from '/@/components/UI/CheckIcon.vue'
import CloseIcon from '/@/components/UI/CloseIcon.vue'

export default {
  name: 'HeaderList',
  components: {
    AddIcon,
    CheckIcon,
    CloseIcon
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
