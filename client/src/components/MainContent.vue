<template>
  <div
    style="height: calc(100% - 3.5rem)"
    class="w-full p-2 overflow-scroll flex"
  >
    <template v-for="group in groups" :key="group.id">
      <task-panel
        :tasks="tasksByGroup(group) ?? []"
        :group="group"
        :tags="tags ?? []"
      />
    </template>
  </div>
</template>

<script lang="ts">
import { ref } from 'vue'
import { apis, Groups, Group, Tasks, Tags } from '/@/lib/apis'
import TaskPanel from '/@/components/TaskPanel/TaskPanel.vue'

export default {
  name: 'MainContent',
  components: { TaskPanel },
  setup() {
    const groups = ref<Groups>()
    const tasks = ref<Tasks>()
    const tags = ref<Tags>()

    const fetchGroups = async () => {
      const res = await apis.getGroups()
      groups.value = res.data
    }
    const fetchTasks = async () => {
      const res = await apis.getTasks()
      tasks.value = res.data
    }
    const fetchTags = async () => {
      const res = await apis.getTags()
      tags.value = res.data
    }

    // TODO: エラーハンドリング
    try {
      fetchGroups()
      fetchTasks()
      fetchTags()
    } catch (e) {
      // eslint-disable-next-line no-console
      console.log(e)
    }

    const tasksByGroup = (group: Group) => {
      return tasks.value?.filter(task => task.group_id === group.id)
    }

    return { groups, tasksByGroup, tags }
  }
}
</script>

<style scoped></style>
