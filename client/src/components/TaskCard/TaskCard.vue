<template>
  <div
    class="w-full bg-white rounded-lg border-1 border-gray-300 my-2 p-2 shadow-sm relative"
    @mouseover="hover = true"
    @mouseleave="hover = false"
  >
    <div v-if="hover" class="absolute top-0 right-0 mr-2 mt-2 flex">
      <check-icon
        class="cursor-pointer items-center flex mx-0.3"
        :size="22"
        fill-color="#696969"
      />
      <delete-icon
        class="cursor-pointer items-center flex mx-0.3"
        :size="22"
        fill-color="#696969"
      />
    </div>

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
import CheckIcon from 'vue-material-design-icons/Check.vue'
import DeleteIcon from 'vue-material-design-icons/Delete.vue'

export default {
  name: 'TaskCard',
  components: {
    TaskTag,
    CheckIcon,
    DeleteIcon
  },
  data() {
    return {
      hover: false
    }
  }
}
</script>

<style scoped></style>
