<template>
  <header-list
    :title="'Groups'"
    :item-length="groups.length"
    @editor-new="setEditGroup"
  >
    <!-- TODO: TagList と GroupList の共通部分をまとめる -->
    <template #editor>
      <div v-show="editing" class="absolute top-6.2 right-1">
        <div
          class="w-60 p-3 bg-white rounded-lg border-1 border-gray-300 shadow-md"
        >
          <div class="mb-2">
            <h3 class="font-semibold">{{ newOrEdit() }}</h3>
          </div>
          <textarea
            v-model="rawData"
            class="w-full resize-none p-1 text-sm"
            rows="3"
          ></textarea>
          <div class="flex items-center justify-end mt-2">
            <a @click="editing = false"><close-icon /></a>
            <a @click="postGroup()"><check-icon class="ml-0.5" /></a>
          </div>
        </div>
      </div>
    </template>

    <template #list>
      <div
        v-for="group in groups"
        :key="group.id"
        class="my-2 px-3 flex items-center"
      >
        <up-arrow-icon :size="18" />
        <down-arrow-icon :size="18" class="ml-1" />
        <p class="ml-3 cursor-pointer" @click="setEditGroup(group)">
          {{ group.name }}
        </p>
      </div>
    </template>
  </header-list>
</template>

<script lang="ts">
import { defineComponent, ref, PropType } from 'vue'
import { apis, Group, Groups } from '/@/lib/apis'
import { refresh } from '/@/lib/refresh'
import HeaderList from '/@/components/Layout/PageHeader/HeaderList.vue'
import UpArrowIcon from '/@/components/UI/UpArrowIcon.vue'
import DownArrowIcon from '/@/components/UI/DownArrowIcon.vue'
import CheckIcon from '/@/components/UI/CheckIcon.vue'
import CloseIcon from '/@/components/UI/CloseIcon.vue'

export default defineComponent({
  name: 'GroupList',
  components: {
    HeaderList,
    UpArrowIcon,
    DownArrowIcon,
    CheckIcon,
    CloseIcon
  },
  props: {
    groups: {
      type: Array as PropType<Groups>,
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

    let editingGroup: Group | null = null
    const setEditGroup = (group: Group | null) => {
      if (group) {
        editingGroup = group
        isNew.value = false
        rawData.value = `${group.name}`
      } else {
        if (editing.value && !editingGroup) {
          editing.value = false
          return
        }
        editingGroup = null
        isNew.value = true
        rawData.value = ''
      }
      editing.value = true
    }

    const postGroup = async () => {
      let name = rawData.value.split('\n')[0]
      if (name?.length === 0) return

      let reqGroup = {
        name: name ?? ''
      }

      if (isNew.value) {
        await apis.postGroup(reqGroup).then(() => {
          refresh()
          closeEditor()
        })
      } else {
        await apis.putGroup(editingGroup?.id ?? '', reqGroup).then(() => {
          refresh()
          closeEditor()
        })
      }
    }

    const closeEditor = () => {
      editing.value = false
      rawData.value = ''
    }

    return { rawData, editing, isNew, newOrEdit, setEditGroup, postGroup }
  }
})
</script>

<style scoped></style>
