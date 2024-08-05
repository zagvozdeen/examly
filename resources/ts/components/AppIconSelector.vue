<template>
  <div class="flex flex-col gap-4">
    <ul class="flex gap-3 px-1 flex-wrap">
      <li
        v-for="icon in icons"
        :key="icon"
      >
        <input
          :id="`app-icon-${icon}-selector`"
          v-model="value"
          :value="icon"
          type="radio"
          name="app-icon-selector"
          class="hidden peer"
        >
        <label
          :for="`app-icon-${icon}-selector`"
          class="flex items-center justify-center peer-checked:outline rounded w-8 h-8 cursor-pointer hover:outline outline-offset-2 outline-2 outline-orange-400d outline-white"
        >
          <i
            :class="{
              'text-lg bi': true,
              [icon]: true,
            }"
          />
        </label>
      </li>
    </ul>

    <n-button
      v-if="available"
      @click="iteration++"
    >
      Загрузить ещё
    </n-button>
  </div>
</template>

<script lang="ts" setup>
import { BOOTSTRAP_ICONS } from '@/composables/useIcons.ts'
import { NButton } from 'naive-ui'
import { computed, ref } from 'vue'

const PER_PAGE = 48
const iteration = ref(1)
const value = defineModel('value')

const available = computed(() => {
  return BOOTSTRAP_ICONS.length > PER_PAGE * iteration.value
})

const icons = computed(() => {
  return BOOTSTRAP_ICONS
    .slice(0, PER_PAGE * iteration.value)
    .map(icon => `bi-${icon}`)
})
</script>
