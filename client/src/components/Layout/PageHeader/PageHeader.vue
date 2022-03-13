<template>
  <header class="h-14 w-full bg-gray-100 shadow-md flex items-center px-4">
    <div class="flex-1 text-xl">Todo</div>
    <add-icon
      class="cursor-pointer items-center flex px-0.5"
      :size="22"
      fill-color="#2c3e50"
    />
    <order-icon
      class="cursor-pointer items-center flex px-0.5"
      :size="22"
      fill-color="#2c3e50"
    />
    <div class="w-2"></div>
    <tag-icon
      class="cursor-pointer items-center flex px-0.5"
      :size="22"
      fill-color="#2c3e50"
    />
  </header>
</template>

<script>
import AddIcon from 'vue-material-design-icons/Plus.vue'
import OrderIcon from 'vue-material-design-icons/ReorderHorizontal.vue'
import TagIcon from 'vue-material-design-icons/TagOutline.vue'

export default {
  name: 'PageHeader',
  components: {
    AddIcon,
    OrderIcon,
    TagIcon
  }
}
</script>

<style scoped></style>
