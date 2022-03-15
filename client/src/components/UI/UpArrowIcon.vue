<template>
  <icon
    class="cursor-pointer items-center flex"
    :size="size"
    fill-color="#696969"
  />
</template>

<script lang="ts">
import Icon from 'vue-material-design-icons/ArrowUp.vue'

export default {
  name: 'ApArrowIcon',
  components: {
    Icon
  },
  props: {
    size: {
      type: Number,
      default: 22
    }
  }
}
</script>

<style scoped></style>
