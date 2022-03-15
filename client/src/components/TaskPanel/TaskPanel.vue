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
          @close="closeEditors"
        />
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
          <task-card :task="task" :tags="tags" />
        </template>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, PropType } from 'vue'
import { Group, Tags, Tasks } from '/@/lib/apis'
import TaskCard from '/@/components/TaskCard/TaskCard.vue'
import DotsIcon from '/@/components/UI/DotsHorizontalIcon.vue'
import AddIcon from '/@/components/UI/AddIcon.vue'
import TaskEditor from '/@/components/TaskEditor/TaskEditor.vue'
import TaskPanelMenu from '/@/components/TaskPanelMenu/TaskPanelMenu.vue'

export default defineComponent({
  name: 'TaskPanel',
  components: {
    TaskPanelMenu,
    TaskEditor,
    TaskCard,
    DotsIcon,
    AddIcon
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
  data() {
    return {
      openingMenu: false,
      editingTask: false
    }
  },
  methods: {
    operateTaskEditor() {
      const temp = this.editingTask
      this.closeEditors()
      if (!temp) this.editingTask = true
    },
    operateMenu() {
      const temp = this.openingMenu
      this.closeEditors()
      if (!temp) this.openingMenu = true
    },
    closeEditors() {
      this.editingTask = false
      this.openingMenu = false
    }
  }
})
</script>

<style scoped></style>
