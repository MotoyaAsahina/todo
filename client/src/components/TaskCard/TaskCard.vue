<template>
  <div
    class="w-full bg-white rounded-lg border-1 border-gray-300 my-2 p-2 shadow-sm relative"
    @mouseover="cardHover = true"
    @mouseleave="cardHover = false"
  >
    <div v-show="cardHover" class="absolute top-0 right-0 mr-2 mt-2 flex">
      <check-icon />
      <delete-icon class="ml-0.5" />
    </div>

    <a class="cursor-pointer font-bold">Task Title</a>
    <div class="flex flex-wrap gap-1 mt-0.6">
      <p class="mr-1 leading-5">2/24(木) 23:59</p>
      <task-tag tag-name="priority:HIGH"></task-tag>
    </div>
    <p
      class="mt-1.2 text-sm"
      :class="{
        'whitespace-nowrap overflow-hidden overflow-ellipsis': !cardClick
      }"
      @click="cardClick = !cardClick"
    >
      コメントおおおおおおおおおおおおおおおおおおおおおおおおおおおおおおおおおおおおおお
    </p>
  </div>
</template>

<script lang="ts">
import TaskTag from '/@/components/TaskTag/TaskTag.vue'
import CheckIcon from '/@/components/UI/CheckIcon.vue'
import DeleteIcon from '/@/components/UI/DeleteIcon.vue'

export default {
  name: 'TaskCard',
  components: {
    TaskTag,
    CheckIcon,
    DeleteIcon
  },
  data() {
    return {
      cardHover: false,
      cardClick: false
    }
  }
}
</script>

<style scoped></style>
