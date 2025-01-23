<template>
  <ul class="flex flex-wrap gap-1">
    <li
      v-for="tag in tags"
      :key="tag.id"
    >
      <button
        class="text-xs py-1 px-2 rounded-full"
        :class="{
          'bg-obscure-600': !ids.includes(tag.id),
          'bg-green-800': ids.includes(tag.id),
        }"
        type="button"
        @click="toggle(tag.id)"
      >
        {{ tag.name }}
      </button>
    </li>
    <li v-if="!all">
      <button
        class="text-xs bg-gray-300 text-obscure-600 font-semibold py-1 px-2 rounded-full"
        type="button"
        @click="all = true"
      >
        Показать все
      </button>
    </li>
  </ul>
</template>

<script lang="ts" setup>
import { tags as allTags, useTagStore } from '@/composables/useTagStore.ts'
import { computed, onMounted, ref } from 'vue'

const tagStore = useTagStore()
const all = ref(false)

const ids = defineModel<number[]>('values', { required: true })

const tags = computed(() => all.value ? allTags.value : allTags.value.slice(0, 10))

const toggle = (id: number) => {
  if (ids.value.includes(id)) {
    ids.value = ids.value.filter((value) => value !== id)
  } else {
    ids.value = [...ids.value, id]
  }
}

onMounted(() => {
  tagStore.getAllTags()
})
</script>