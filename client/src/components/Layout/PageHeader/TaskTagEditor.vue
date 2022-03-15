<template>
  <head-editor :title="'Tags'" :item-length="tags.length">
    <div v-for="tag in tags" :key="tag.id" class="my-2 px-2">
      <div class="flex items-center">
        <task-tag class="w-auto" :tag="tag" />
        <p class="ml-2">{{ tag.color }}</p>
      </div>
      <p>{{ tag.name }}</p>
    </div>
  </head-editor>
</template>

<script lang="ts">
import { PropType } from 'vue'
import { Tags } from '/@/lib/apis'
import TaskTag from '/@/components/TaskTag/TaskTag.vue'
import HeadEditor from '/@/components/Layout/PageHeader/HeadEditor.vue'

export default {
  name: 'TaskTagEditor',
  components: {
    HeadEditor,
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
