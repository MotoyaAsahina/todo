<template>
  <header class="h-14 w-full bg-gray-100 border-b-1 border-gray-300 relative">
    <div class="h-full flex items-center px-4">
      <div class="flex-1 text-xl">Todo</div>
      <order-icon />
      <tag-icon class="ml-1" />
      <a href="login"><login-icon class="ml-2" /></a>
    </div>

    <task-tag-editor
      v-show="editingTags"
      class="absolute right-0 top-10"
      :tags="tags ?? []"
    />
    <task-panel-editor
      v-show="editingPanels"
      class="absolute right-0 top-10"
      :groups="groups ?? []"
    />
  </header>
</template>

<script lang="ts">
import { ref } from 'vue'
import { apis, Groups, Tags } from '/@/lib/apis'
import TaskTagEditor from '/@/components/TaskTag/TaskTagEditor.vue'
import TaskPanelEditor from '/@/components/TaskPanel/TaskPanelEditor.vue'
import OrderIcon from '/@/components/UI/ReorderHorizontalIcon.vue'
import TagIcon from '/@/components/UI/TagIcon.vue'
import LoginIcon from '/@/components/UI/LoginIcon.vue'

export default {
  name: 'PageHeader',
  components: {
    TaskTagEditor,
    TaskPanelEditor,
    OrderIcon,
    TagIcon,
    LoginIcon
  },
  setup() {
    const groups = ref<Groups>()
    const tags = ref<Tags>()

    const fetchGroups = async () => {
      const res = await apis.getGroups()
      groups.value = res.data
    }
    const fetchTags = async () => {
      const res = await apis.getTags()
      tags.value = res.data
    }

    try {
      fetchGroups()
      fetchTags()
    } catch (e) {
      // eslint-disable-next-line no-console
      console.log(e)
    }

    return { groups, tags }
  },
  data() {
    return {
      editingTags: false,
      editingPanels: false
    }
  }
}
</script>

<style scoped></style>
