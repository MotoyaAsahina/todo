<template>
  <div class="w-xs rounded-lg border-1 border-gray-300 bg-white shadow-md">
    <div class="p-3">
      <div class="mb-2">
        <h3 class="font-semibold">New Task</h3>
      </div>
      <textarea class="w-full resize-none p-1 text-sm" rows="8"></textarea>
      <div class="flex items-center justify-end mt-2">
        <close-icon />
        <check-icon class="ml-0.5" />
      </div>
    </div>
  </div>
</template>

<script>
import CheckIcon from '/@/components/UI/CheckIcon.vue'
import CloseIcon from '/@/components/UI/CloseIcon.vue'

export default {
  name: 'TaskEditor',
  components: {
    CheckIcon,
    CloseIcon
  }
}
</script>

<style scoped></style>
