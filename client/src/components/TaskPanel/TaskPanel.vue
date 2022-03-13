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
        <h3 class="flex-1 pl-2">Panel Name</h3>
        <add-icon
          class="cursor-pointer items-center flex px-0.25"
          :size="22"
          fill-color="#2c3e50"
        />
        <dots-icon
          class="cursor-pointer items-center flex px-0.25"
          :size="22"
          fill-color="#2c3e50"
        />
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

<script>
import TaskCard from '../TaskCard/TaskCard.vue'
import DotsIcon from 'vue-material-design-icons/DotsHorizontal.vue'
import AddIcon from 'vue-material-design-icons/Plus.vue'

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
