<template>
  <header class="h-14 w-full bg-gray-100 border-b-1 border-gray-300">
    <div class="h-full flex items-center px-4">
      <div class="flex-1 text-xl">Todo</div>
      <add-icon />
      <order-icon class="ml-1" />
      <tag-icon class="ml-1" />
      <a href="login"><login-icon class="ml-2" /></a>
    </div>
  </header>
</template>

<script lang="ts">
import { ref } from 'vue'
import { apis, Tags } from '/@/lib/apis'
import AddIcon from '/@/components/UI/AddIcon.vue'
import OrderIcon from '/@/components/UI/ReorderHorizontalIcon.vue'
import TagIcon from '/@/components/UI/TagIcon.vue'
import LoginIcon from '/@/components/UI/LoginIcon.vue'

export default {
  name: 'PageHeader',
  components: {
    AddIcon,
    OrderIcon,
    TagIcon,
    LoginIcon
  },
  setup() {
    const tags = ref<Tags>()

    const fetchTags = async () => {
      const res = await apis.getTags()
      tags.value = res.data
    }

    try {
      fetchTags()
    } catch (e) {
      // eslint-disable-next-line no-console
      console.log(e)
    }

    return { tags }
  }
}
</script>

<style scoped></style>
