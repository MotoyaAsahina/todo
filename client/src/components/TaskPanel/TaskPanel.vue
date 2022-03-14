<template>
  <div class="w-sm h-full">
    <div
      style="width: calc(100% - 1rem); height: calc(100% - 1rem)"
      class="m-2 rounded-lg border-1 border-gray-300 bg-gray-50"
    >
      <div class="h-10 pt-3 pb-2 px-3 flex items-center">
        <span
          class="w-5 h-5 flex-initial rounded-lg bg-gray-200 text-sm text-center"
        >
          7
        </span>
        <h3 class="flex-1 pl-2 font-semibold">Panel Name</h3>
        <add-icon />
        <dots-icon class="ml-0.5" />
      </div>

      <div
        style="width: 100%; height: calc(100% - 2.5rem)"
        class="px-2 overflow-scroll"
      >
        <task-card></task-card>
        <task-card></task-card>
        <task-card></task-card>
        <task-card></task-card>
        <task-card></task-card>
        <task-card></task-card>
        <task-card></task-card>
        <task-card></task-card>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import TaskCard from '/@/components/TaskCard/TaskCard.vue'
import DotsIcon from '/@/components/UI/DotsHorizontalIcon.vue'
import AddIcon from '/@/components/UI/AddIcon.vue'

export default {
  name: 'TaskPanel',
  components: {
    TaskCard,
    DotsIcon,
    AddIcon
  }
}
</script>

<style scoped></style>
