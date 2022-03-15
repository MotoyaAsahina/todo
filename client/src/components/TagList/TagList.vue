<template>
  <header-list :title="'Tags'" :item-length="tags.length">
    <div v-for="tag in tags" :key="tag.id" class="my-2 px-2">
      <div class="flex items-center">
        <task-tag class="w-auto" :tag="tag" />
        <p class="ml-2">{{ tag.color }}</p>
      </div>
      <p>{{ tag.name }}</p>
    </div>
  </header-list>
</template>

<script lang="ts">
import { PropType } from 'vue'
import { Tags } from '/@/lib/apis'
import TaskTag from '/@/components/TaskTag/TaskTag.vue'
import HeaderList from '/@/components/Layout/PageHeader/HeaderList.vue'

export default {
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
}
</script>

<style scoped></style>
