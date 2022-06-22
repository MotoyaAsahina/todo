<template>
  <header class="h-3rem w-full bg-gray-100 border-b-1 border-gray-200 relative">
    <div class="h-full flex items-center px-4">
      <div class="flex-1 text-lg">Todo</div>
      <a @click="operateGroupEditor"><order-icon /></a>
      <a @click="operateTagEditor"><tag-icon class="ml-1.5" /></a>
    </div>

    <tag-list
      v-show="editingTags"
      class="absolute right-0 top-8"
      :tags="tags ?? []"
    />
    <group-list
      v-show="editingGroups"
      class="absolute right-0 top-8"
      :groups="groups ?? []"
    />
  </header>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'
import { apis, Groups, Tags } from '/@/lib/apis'
import { addRefreshListener } from '/@/lib/refresh'
import TagList from '/@/components/TagList/TagList.vue'
import GroupList from '/@/components/GroupList/GroupList.vue'
import OrderIcon from '/@/components/UI/ReorderHorizontalIcon.vue'
import TagIcon from '/@/components/UI/TagIcon.vue'

export default defineComponent({
  name: 'PageHeader',
  components: {
    TagList,
    GroupList,
    OrderIcon,
    TagIcon
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

    addRefreshListener(fetchGroups, fetchTags)

    try {
      fetchGroups()
      fetchTags()
    } catch (e) {
      // eslint-disable-next-line no-console
      console.log(e)
    }

    const editingTags = ref(false)
    const editingGroups = ref(false)

    const operateTagEditor = () => {
      const temp = editingTags.value
      closeEditors()
      if (!temp) editingTags.value = true
    }
    const operateGroupEditor = () => {
      const temp = editingGroups.value
      closeEditors()
      if (!temp) editingGroups.value = true
    }
    const closeEditors = () => {
      editingTags.value = false
      editingGroups.value = false
    }

    return {
      groups,
      tags,
      editingTags,
      editingGroups,
      operateTagEditor,
      operateGroupEditor
    }
  }
})
</script>

<style scoped></style>
