<template>
  <header-list :title="'Tags'" :item-length="tags.length">
    <div
      v-for="(tag, index) in tags"
      :key="tag.id"
      class="py-1 px-3"
      :class="{ 'border-t-1 border-gray-200': index > 0 }"
    >
      <div class="flex items-center">
        <task-tag class="w-auto" :tag="tag" />
        <p class="ml-2">{{ tag.color }}</p>
      </div>
    </div>
  </header-list>
</template>

<script lang="ts">
import { defineComponent, PropType } from 'vue'
import { Tags } from '/@/lib/apis'
import TaskTag from '/@/components/TaskTag/TaskTag.vue'
import HeaderList from '/@/components/Layout/PageHeader/HeaderList.vue'

export default defineComponent({
  name: 'TagList',
  components: {
    HeaderList,
    TaskTag
  },
  props: {
    tags: {
      type: Array as PropType<Tags>,
      required: true
    }
  }
})
</script>

<style scoped></style>
