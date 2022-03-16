<template>
  <div class="w-sm h-full">
    <div
      style="width: 23rem; height: calc(100% - 1rem)"
      class="m-2 rounded-lg border-1 border-gray-300 bg-gray-50"
    >
      <div class="h-10 pt-3 pb-2 px-3 flex items-center relative">
        <span
          class="w-5 h-5 flex-initial rounded-lg bg-gray-200 text-sm text-center"
        >
          {{ tasks?.length }}
        </span>
        <h3 class="flex-1 pl-2 font-semibold">{{ group.name }}</h3>
        <a @click="operateTaskEditor"><add-icon /></a>
        <a @click="operateMenu"><dots-icon class="ml-0.5" /></a>

        <task-editor
          v-show="editingTask"
          class="right-1 top-9 absolute z-8"
          :group="group"
          :tags="tags ?? []"
          @keydown.meta.enter="postTask"
          @keydown.esc="closeEditors"
        >
          <div class="mb-2">
            <h3 class="font-semibold">{{ newOrEdit() }} Task</h3>
          </div>
          <textarea
            :id="`task-editor-input-${group.id}`"
            v-model="rawTaskData"
            class="w-full resize-none p-1 text-sm"
            rows="6"
          ></textarea>
          <div class="flex flex-wrap items-end gap-1 relative">
            <a @click="openingTagList = !openingTagList"><tag-icon /></a>
            <task-tag v-for="tag in selectingTags" :key="tag.id" :tag="tag" />

            <div
              v-if="openingTagList"
              class="top-0 left-6 w-60 absolute bg-white rounded-lg border-1 border-gray-300 shadow-md z-10"
            >
              <div class="m-2 flex flex-wrap gap-1">
                <task-tag
                  v-for="tag in tags"
                  :key="tag.id"
                  :tag="tag"
                  @click="selectTag(tag)"
                />
              </div>
            </div>
          </div>
          <div class="flex items-center justify-end mt-2">
            <a @click="closeEditors"><close-icon /></a>
            <a @click="postTask"><check-icon class="ml-0.5" /></a>
          </div>
        </task-editor>

        <task-panel-menu
          v-show="openingMenu"
          class="right-1 top-9 absolute z-8"
        />
      </div>

      <div
        style="width: 100%; height: calc(100% - 2.5rem)"
        class="px-2 overflow-scroll"
      >
        <template v-for="task in tasks" :key="task.id">
          <task-card :task="task" :tags="tags" @edit-task="setEditTask" />
        </template>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, PropType } from 'vue'
import { apis, Group, Tag, Tags, Task, Tasks } from '/@/lib/apis'
import { refresh } from '/@/lib/refresh'
import TaskCard from '/@/components/TaskCard/TaskCard.vue'
import DotsIcon from '/@/components/UI/DotsHorizontalIcon.vue'
import AddIcon from '/@/components/UI/AddIcon.vue'
import TaskEditor from '/@/components/TaskEditor/TaskEditor.vue'
import TaskPanelMenu from '/@/components/TaskPanelMenu/TaskPanelMenu.vue'
import TaskTag from '/@/components/TaskTag/TaskTag.vue'
import CheckIcon from '/@/components/UI/CheckIcon.vue'
import CloseIcon from '/@/components/UI/CloseIcon.vue'
import TagIcon from '/@/components/UI/TagIcon.vue'

export default defineComponent({
  name: 'TaskPanel',
  components: {
    TaskPanelMenu,
    TaskEditor,
    TaskCard,
    DotsIcon,
    AddIcon,
    TaskTag,
    CheckIcon,
    CloseIcon,
    TagIcon
  },
  props: {
    group: {
      type: Object as PropType<Group>,
      required: true
    },
    tasks: {
      type: Array as PropType<Tasks>,
      required: true
    },
    tags: {
      type: Array as PropType<Tags>,
      required: true
    }
  },
  setup(props) {
    const openingMenu = ref(false)
    const editing = ref(false)

    const editingTask = ref<Task | null>(null)
    const isNew = ref(false)

    const operateTaskEditor = () => {
      const temp = editing.value
      if (openingMenu.value || (temp && !editingTask.value)) closeEditors()
      if (!temp) {
        editing.value = true
      }
      isNew.value = true
      editingTask.value = null
      rawTaskData.value = ''
      selectingTags.clear()
      window.setTimeout(function () {
        document.getElementById(`task-editor-input-${props.group.id}`)?.focus()
      }, 10)
    }
    const operateMenu = () => {
      const temp = openingMenu.value
      closeEditors()
      if (!temp) openingMenu.value = true
    }
    const closeEditors = () => {
      editing.value = false
      openingMenu.value = false
    }

    const newOrEdit = () => (isNew.value ? 'New' : 'Edit')

    const setEditTask = (task: Task) => {
      closeEditors()
      isNew.value = false
      editing.value = true
      editingTask.value = task
      rawTaskData.value = `${task.title}\n${task.due_date}\n${task.description}`
      selectingTags.clear()
      if (task.tags?.length > 0) {
        for (const tagID of task.tags) {
          selectingTags.add(
            props.tags.find(t => t.id === tagID) ?? {
              id: '',
              name: '',
              color: ''
            }
          )
        }
      }
      window.setTimeout(function () {
        document.getElementById(`task-editor-input-${props.group.id}`)?.focus()
      }, 10)
    }

    let selectingTags: Set<Tag> = new Set()
    const openingTagList = ref(false)

    const selectTag = (tag: Tag) => {
      if (selectingTags.has(tag)) {
        selectingTags.delete(tag)
      } else {
        selectingTags.add(tag)
      }
    }

    const rawTaskData = ref('')

    const postTask = async () => {
      let rawData = rawTaskData.value.split('\n')
      if (rawData.length < 2 || rawData[0] === '' || rawData[1] === '') return

      const reqTask = {
        group_id: props.group.id,
        title: rawData[0] ?? '',
        description: rawData.slice(2).join('\n') ?? '',
        due_date: rawData[1] ?? '',
        tags: [...selectingTags].map(tag => tag.id)
      }

      if (isNew.value) {
        await apis.postTask(reqTask).then(() => {
          refresh()
          closeEditors()
        })
      } else {
        if (!editingTask.value) return
        await apis.putTask(editingTask.value.id, reqTask).then(() => {
          refresh()
          closeEditors()
        })
      }
    }

    return {
      openingMenu,
      editingTask: editing,
      operateTaskEditor,
      operateMenu,
      closeEditors,
      newOrEdit,
      setEditTask,
      openingTagList,
      selectingTags,
      selectTag,
      rawTaskData,
      postTask
    }
  }
})
</script>

<style scoped></style>
