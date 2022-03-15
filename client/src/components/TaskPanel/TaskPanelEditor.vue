<template>
  <div class="w-sm max-w-full z-10 p-2">
    <div class="rounded-lg border-1 border-gray-300 bg-white shadow-md">
      <div
        class="mt-2.5 pb-1.5 px-3 flex items-center border-b-1 border-gray-300"
      >
        <span
          class="w-5 h-5 flex-initial rounded-lg bg-gray-200 text-sm text-center"
        >
          {{ groups?.length }}
        </span>
        <h3 class="flex-1 pl-2 font-semibold">Groups</h3>
        <add-icon />
      </div>

      <div class="h-lg px-2 overflow-scroll">
        <div
          v-for="group in groups"
          :key="group.id"
          class="my-2 pl-1 pr-2 flex items-center"
        >
          <up-arrow-icon :size="18" />
          <down-arrow-icon :size="18" class="ml-1" />
          <p class="ml-3">{{ group.name }}</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { PropType } from 'vue'
import { Groups } from '/@/lib/apis'
import AddIcon from '/@/components/UI/AddIcon.vue'
import UpArrowIcon from '/@/components/UI/UpArrowIcon.vue'
import DownArrowIcon from '/@/components/UI/DownArrowIcon.vue'

export default {
  name: 'TaskPanelEditor',
  components: {
    AddIcon,
    UpArrowIcon,
    DownArrowIcon
  },
  props: {
    groups: {
      type: Array as PropType<Groups>,
      required: true
    }
  }
}
</script>

<style scoped></style>
