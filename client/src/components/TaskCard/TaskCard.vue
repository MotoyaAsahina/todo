<template>
  <div
    class="w-full bg-white rounded-lg border-1 border-gray-200 my-1.6 px-1.6 pt-1.2 pb-1.6 shadow-sm relative"
    :class="isPending(task) ? 'bg-gray-100 shadow-none' : ''"
    @mouseover="cardHover = true"
    @mouseleave="cardHover = false"
  >
    <div
      v-show="isPending(task)"
      class="absolute rounded-lg bg-gray-100 w-full h-full left-0 top-0 bg-opacity-40 pointer-events-none"
    ></div>
    <div v-show="cardHover" class="absolute top-0 right-0 mr-1.6 mt-1.2 flex">
      <a @click="putTaskDone"><check-icon class="mr-0.5" /></a>
      <a @click="deleteTask"><delete-icon /></a>
    </div>

    <div>
      <span class="mr-1.2 text-base">{{ stamp(task.due_date) }}</span>
      <a
        class="cursor-pointer text-base"
        :class="!isPending(task) ? 'font-bold' : ''"
        @click="$emit('editTask', task)"
      >
        {{ task.title }}
      </a>
    </div>
    <div class="flex flex-wrap gap-1 mt-0.2">
      <p class="mr-1 text-base leading-1.1rem" @click="cardClick = !cardClick">
        {{ formatDueDate(task.due_date) }}
      </p>
      <task-tag v-for="tagID in task.tags" :key="tagID" :tag="findTag(tagID)" />
    </div>
    <div
      v-if="removeAnnotations(task.description)?.length > 0"
      class="pt-0.6"
      @click="cardClick = !cardClick"
    >
      <p
        v-if="!cardClick"
        class="text-sm leading-snug whitespace-nowrap overflow-hidden overflow-ellipsis"
        v-html="makeURL(removeAnnotations(task.description))"
      ></p>
      <p
        v-if="cardClick"
        class="text-sm leading-snug"
        v-html="makeBR(makeURL(removeAnnotations(task.description)))"
      ></p>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, PropType } from 'vue'
import { apis, Task, Tags, Tag } from '/@/lib/apis'
import { refresh } from '/@/lib/refresh'
import { selectStamp } from '/@/lib/stamp'
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
  emits: ['editTask'],
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

    const findTag = (id: string): Tag => {
      return (
        props.tags?.find(t => t.id === id) ?? { id: '', name: '', color: '' }
      )
    }

    const putTaskDone = async () => {
      await apis.putTaskDone(props.task.id).then(() => {
        refresh()
      })
    }

    const deleteTask = async () => {
      await apis.deleteTask(props.task.id).then(() => {
        refresh()
      })
    }

    const stamp = (d: string): string => selectStamp(d)

    const surroundURL = (url: string) => {
      const style = 'overflow-wrap: break-word; color: #135fab'
      const rel = 'noopener noreferrer'
      return `<a style="${style}" href="${url}" target="_blank" rel="${rel}">${url}</a>`
    }

    const makeURL = (text: string) => {
      return text.replace(
        /(https?:\/\/[-_.!~*'()a-zA-Z0-9;/?:@&=+$,%#]+)/g,
        surroundURL
      )
    }

    const isPending = (task: Task) => {
      return task.description.match(/!pending\[.*]/)
    }

    const removeAnnotations = (text: string) => {
      return text
        .replace(/\n?!notice\[.*]/g, '')
        .replace(/\n?!pending\[.*]/g, '')
    }

    const makeBR = (text: string) => {
      return text.replace(/\n/g, '<br>')
    }

    return {
      formatDueDate,
      findTag,
      putTaskDone,
      deleteTask,
      stamp,
      isPending,
      removeAnnotations,
      makeURL,
      makeBR
    }
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
