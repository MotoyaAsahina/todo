<template>
  <div
    class="w-full bg-white rounded-lg border-1 border-gray-300 my-2 p-2 shadow-sm"
  >
    <a class="cursor-pointer font-bold">Task Title</a>
    <p>2/24(木) 23:59</p>
    <div class="h-1"></div>
    <div class="flex flex-wrap gap-1">
      <task-tag tag-name="priority:HIGH"></task-tag>
    </div>
    <div class="h-1.2"></div>
    <p class="text-sm">コメント</p>
  </div>
</template>

<script>
import TaskTag from '../TaskTag/TaskTag.vue'
export default {
  name: 'TaskCard',
  components: { TaskTag }
}
</script>

<style scoped></style>
