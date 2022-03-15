<template>
  <header-list :title="'Groups'" :item-length="groups.length">
    <div
      v-for="group in groups"
      :key="group.id"
      class="my-2 pl-1 pr-2 flex items-center"
    >
      <up-arrow-icon :size="18" />
      <down-arrow-icon :size="18" class="ml-1" />
      <p class="ml-3">{{ group.name }}</p>
    </div>
  </header-list>
</template>

<script lang="ts">
import { PropType } from 'vue'
import { Groups } from '/@/lib/apis'
import HeaderList from '/@/components/Layout/PageHeader/HeaderList.vue'
import UpArrowIcon from '/@/components/UI/UpArrowIcon.vue'
import DownArrowIcon from '/@/components/UI/DownArrowIcon.vue'

export default {
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
}
</script>

<style scoped></style>
