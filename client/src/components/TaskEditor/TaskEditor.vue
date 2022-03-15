<template>
  <div class="w-xs rounded-lg border-1 border-gray-300 bg-white shadow-md">
    <div class="p-3">
      <div class="mb-2">
        <h3 class="font-semibold">New Task</h3>
      </div>
      <textarea
        v-model="rawTaskData"
        class="w-full resize-none p-1 text-sm"
        rows="6"
      ></textarea>
      <div class="flex flex-wrap items-end gap-1 relative">
        <a @click="openingTagList = !openingTagList"><tag-icon /></a>
        <task-tag v-for="tag in selectingTags" :key="tag.id" :tag="tag" />

        <div
          v-if="openingTagList"
          class="top-0 left-6 w-60 absolute bg-white rounded-lg border-1 border-gray-300 shadow-md z-10"
        >
          <div class="m-2 flex flex-wrap gap-1">
            <task-tag
              v-for="tag in tags"
              :key="tag.id"
              :tag="tag"
              @click="selectTag(tag)"
            />
          </div>
        </div>
      </div>
      <div class="flex items-center justify-end mt-2">
        <a @click="$emit('close')"><close-icon /></a>
        <a @click="postTask"><check-icon class="ml-0.5" /></a>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, PropType } from 'vue'
import { apis, Group, Tag, Tags } from '/@/lib/apis'
import TaskTag from '/@/components/TaskTag/TaskTag.vue'
import CheckIcon from '/@/components/UI/CheckIcon.vue'
import CloseIcon from '/@/components/UI/CloseIcon.vue'
import TagIcon from '/@/components/UI/TagIcon.vue'

export default defineComponent({
  name: 'TaskEditor',
  components: {
    TaskTag,
    CheckIcon,
    CloseIcon,
    TagIcon
  },
  props: {
    group: {
      type: Object as PropType<Group>,
      required: true
    },
    tags: {
      type: Array as PropType<Tags>,
      required: true
    }
  },
  emits: ['close'],
  setup(props) {
    const selectingTags: Set<Tag> = new Set()

    const selectTag = (tag: Tag) => {
      if (selectingTags.has(tag)) {
        selectingTags.delete(tag)
      } else {
        selectingTags.add(tag)
      }
    }

    const rawTaskData = ref('')

    const postTask = async () => {
      let rawData = rawTaskData.value.split('\n')
      if (rawData.length < 2 || rawData[0] === '' || rawData[1] === '') return
      await apis.postTask({
        group_id: props.group.id,
        title: rawData[0] ?? '',
        description: rawData[2] ?? '',
        due_date: rawData[1] ?? '',
        tags: [...selectingTags].map(tag => tag.id)
      })
    }

    return { selectingTags, selectTag, rawTaskData, postTask }
  },
  data() {
    return {
      openingTagList: false
    }
  }
})
</script>

<style scoped></style>
