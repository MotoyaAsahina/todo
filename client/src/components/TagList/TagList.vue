<template>
  <header-list
    :title="'Tags'"
    :item-length="tags.length"
    @editor-new="setEditTag"
  >
    <!-- TODO: TagList と GroupList の共通部分をまとめる -->
    <template #editor>
      <div v-show="editing" class="absolute top-6.2 right-1">
        <div
          class="w-60 p-2.4 bg-white rounded-lg border-1 border-gray-200 shadow-md"
        >
          <div class="mb-2">
            <h3 class="text-base font-semibold">{{ newOrEdit() }}</h3>
          </div>
          <textarea
            id="tag-list-editor-input"
            v-model="rawData"
            class="w-full resize-none p-1 text-sm border-1 border-gray-400"
            rows="3"
            @keydown.meta.enter="postTag"
            @keydown.esc="closeEditor"
          ></textarea>
          <div class="flex items-center justify-end mt-1.6">
            <a @click="editing = false"><close-icon /></a>
            <a @click="postTag()"><check-icon class="ml-0.5" /></a>
          </div>
        </div>
      </div>
    </template>

    <template #list>
      <div
        v-for="(tag, index) in tags"
        :key="tag.id"
        class="py-0.8 px-2.4"
        :class="{ 'border-t-1 border-gray-200': index > 0 }"
      >
        <div class="flex items-center">
          <task-tag class="w-auto" :tag="tag" @click="setEditTag(tag)" />
          <p class="ml-2 text-base">{{ tag.color }}</p>
        </div>
      </div>
    </template>
  </header-list>
</template>

<script lang="ts">
import { defineComponent, PropType, ref } from 'vue'
import { apis, Tag, Tags } from '/@/lib/apis'
import TaskTag from '/@/components/TaskTag/TaskTag.vue'
import HeaderList from '/@/components/Layout/PageHeader/HeaderList.vue'
import CheckIcon from '/@/components/UI/CheckIcon.vue'
import CloseIcon from '/@/components/UI/CloseIcon.vue'
import { refresh } from '/@/lib/refresh'

export default defineComponent({
  name: 'TagList',
  components: {
    HeaderList,
    TaskTag,
    CheckIcon,
    CloseIcon
  },
  props: {
    tags: {
      type: Array as PropType<Tags>,
      required: true
    }
  },
  setup() {
    const rawData = ref('')
    const editing = ref(false)
    const isNew = ref(false)

    const newOrEdit = (): string => {
      return isNew.value ? 'New' : 'Edit'
    }

    let editingTag: Tag | null = null
    const setEditTag = (tag: Tag | null) => {
      if (tag) {
        editingTag = tag
        isNew.value = false
        rawData.value = `${tag.name}\n${tag.color}`
      } else {
        if (editing.value && !editingTag) {
          editing.value = false
          return
        }
        editingTag = null
        isNew.value = true
        rawData.value = ''
      }
      editing.value = true
      document.getElementById('tag-list-editor-input')?.removeAttribute('style')
      window.setTimeout(function () {
        document.getElementById('tag-list-editor-input')?.focus()
      }, 10)
    }

    const postTag = async () => {
      let raw = rawData.value.split('\n')
      if (raw.length < 2 || raw[0]?.length === 0 || raw[1]?.length === 0) return

      let reqTag = {
        name: raw[0] ?? '',
        color: raw[1] ?? ''
      }

      if (isNew.value) {
        await apis.postTag(reqTag).then(() => {
          refresh()
          closeEditor()
        })
      } else {
        await apis.putTag(editingTag?.id ?? '', reqTag).then(() => {
          refresh()
          closeEditor()
        })
      }
    }

    const closeEditor = () => {
      editing.value = false
      rawData.value = ''
    }

    return {
      rawData,
      editing,
      isNew,
      newOrEdit,
      setEditTag,
      postTag,
      closeEditor
    }
  }
})
</script>

<style scoped></style>
