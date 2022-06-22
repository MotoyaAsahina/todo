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
          class="w-60 p-2.4 bg-white rounded-lg border-1 border-gray-200 shadow-md"
        >
          <div class="mb-2">
            <h3 class="text-base font-semibold">{{ newOrEdit() }}</h3>
          </div>
          <textarea
            id="group-list-editor-input"
            v-model="rawData"
            class="w-full resize-none p-1 text-sm border-1 border-gray-400"
            rows="3"
            @keydown.meta.enter="postGroup"
            @keydown.esc="closeEditor"
          ></textarea>
          <div class="flex items-center justify-end mt-1.6">
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
        class="my-1.6 px-2.4 flex items-center"
      >
        <a @click="moveGroup(group.id, 'up')"><up-arrow-icon :size="18" /></a>
        <a @click="moveGroup(group.id, 'down')"
          ><down-arrow-icon :size="18" class="ml-1"
        /></a>
        <p class="ml-3 cursor-pointer text-base" @click="setEditGroup(group)">
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
      document
        .getElementById('group-list-editor-input')
        ?.removeAttribute('style')
      window.setTimeout(function () {
        document.getElementById('group-list-editor-input')?.focus()
      }, 10)
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

    const moveGroup = async (id: string, direction: 'up' | 'down') => {
      if (direction === 'up') {
        await apis.putGroupUp(id).then(() => {
          refresh()
        })
      } else {
        await apis.putGroupDown(id).then(() => {
          refresh()
        })
      }
    }

    return {
      rawData,
      editing,
      isNew,
      newOrEdit,
      setEditGroup,
      postGroup,
      closeEditor,
      moveGroup
    }
  }
})
</script>

<style scoped></style>
