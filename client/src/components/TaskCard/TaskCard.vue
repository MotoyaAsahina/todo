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

    <a class="cursor-pointer font-bold">{{ task.title }}</a>
    <div class="flex flex-wrap gap-1 mt-0.6">
      <p class="mr-1 leading-5">{{ formatDueDate(task.due_date) }}</p>
      <task-tag v-for="tagID in task.tags" :key="tagID" :tag="findTag(tagID)" />
    </div>
    <p
      class="mt-1.2 text-sm"
      :class="{
        'whitespace-nowrap overflow-hidden overflow-ellipsis': !cardClick
      }"
      @click="cardClick = !cardClick"
    >
      {{ task.description }}
    </p>
  </div>
</template>

<script lang="ts">
import { defineComponent, PropType } from 'vue'
import { Task, Tags } from '/@/lib/apis'
import TaskTag from '/@/components/TaskTag/TaskTag.vue'
import CheckIcon from '/@/components/UI/CheckIcon.vue'
import DeleteIcon from '/@/components/UI/DeleteIcon.vue'

export default defineComponent({
  name: 'TaskCard',
  components: {
    TaskTag,
    CheckIcon,
    DeleteIcon
  },
  props: {
    task: {
      type: Object as PropType<Task>,
      required: true
    },
    tags: {
      type: Array as PropType<Tags>,
      required: true
    }
  },
  setup(props) {
    const formatDueDate = (d: string) => {
      let date = new Date(d)
      let year =
        date.getFullYear() !== new Date().getFullYear()
          ? `${date.getFullYear()}/`
          : ''
      let month = date.getMonth() + 1
      let day = date.getDate()
      let weekDay = ['日', '月', '火', '水', '木', '金', '土'][date.getDay()]
      let hour = ('00' + date.getHours()).slice(-2)
      let minute = ('00' + date.getMinutes()).slice(-2)
      return `${year}${month}/${day}(${weekDay}) ${hour}:${minute}`
    }

    const findTag = (id: string) => {
      return props.tags?.find(t => t.id === id)
    }

    return { formatDueDate, findTag }
  },
  data() {
    return {
      cardHover: false,
      cardClick: false
    }
  }
})
</script>

<style scoped></style>
