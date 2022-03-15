<template>
  <div class="w-sm max-w-full z-10 p-2">
    <div class="rounded-lg border-1 border-gray-300 bg-white">
      <div
        class="mt-2.5 pb-1.5 px-3 flex items-center border-b-1 border-gray-300"
      >
        <span
          class="w-5 h-5 flex-initial rounded-lg bg-gray-200 text-sm text-center"
        >
          {{ tags?.length }}
        </span>
        <h3 class="flex-1 pl-2 font-semibold">Tags</h3>
        <add-icon />
      </div>

      <div class="h-lg px-2 overflow-scroll">
        <div v-for="tag in tags" :key="tag.id" class="my-2 px-2">
          <div class="flex items-center">
            <task-tag class="w-auto" :tag="tag" />
            <p class="ml-2">{{ tag.color }}</p>
          </div>
          <p>{{ tag.name }}</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { PropType } from 'vue'
import { Tags } from '/@/lib/apis'
import TaskTag from '/@/components/TaskTag/TaskTag.vue'
import AddIcon from '/@/components/UI/AddIcon.vue'

export default {
  name: 'TaskTagEditor',
  components: {
    TaskTag,
    AddIcon
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
