<template>
  <header class="h-14 w-full bg-gray-100 shadow-md flex items-center px-4">
    <div class="flex-1 text-xl">Todo</div>
    <add-icon />
    <order-icon class="ml-1" />
    <tag-icon class="ml-1" />
    <a href="login"><key-icon class="ml-2" /></a>
  </header>
</template>

<script lang="ts">
import AddIcon from '/@/components/UI/AddIcon.vue'
import OrderIcon from '/@/components/UI/ReorderHorizontalIcon.vue'
import TagIcon from '/@/components/UI/TagIcon.vue'
import KeyIcon from '/@/components/UI/KeyIcon.vue'

export default {
  name: 'PageHeader',
  components: {
    AddIcon,
    OrderIcon,
    TagIcon,
    KeyIcon
  }
}
</script>

<style scoped></style>
