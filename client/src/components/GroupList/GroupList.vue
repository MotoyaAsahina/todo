<template>
  <header-list :title="'Groups'" :item-length="groups.length">
    <div
      v-for="group in groups"
      :key="group.id"
      class="my-2 px-3 flex items-center"
    >
      <up-arrow-icon :size="18" />
      <down-arrow-icon :size="18" class="ml-1" />
      <p class="ml-3 cursor-pointer">{{ group.name }}</p>
    </div>
  </header-list>
</template>

<script lang="ts">
import { defineComponent, PropType } from 'vue'
import { Groups } from '/@/lib/apis'
import HeaderList from '/@/components/Layout/PageHeader/HeaderList.vue'
import UpArrowIcon from '/@/components/UI/UpArrowIcon.vue'
import DownArrowIcon from '/@/components/UI/DownArrowIcon.vue'

export default defineComponent({
  name: 'GroupList',
  components: {
    HeaderList,
    UpArrowIcon,
    DownArrowIcon
  },
  props: {
    groups: {
      type: Array as PropType<Groups>,
      required: true
    }
  }
})
</script>

<style scoped></style>
