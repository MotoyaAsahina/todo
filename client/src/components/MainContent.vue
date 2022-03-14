<template>
  <div
    style="height: calc(100% - 3.5rem)"
    class="p-2 overflow-scroll grid grid-rows-1 grid-flow-col"
  >
    <task-panel></task-panel>
    <task-panel></task-panel>
    <task-panel></task-panel>
    <task-panel></task-panel>
  </div>
</template>

<script lang="ts">
import TaskPanel from '/@/components/TaskPanel/TaskPanel.vue'

export default {
  name: 'MainContent',
  components: { TaskPanel }
}
</script>

<style scoped></style>
