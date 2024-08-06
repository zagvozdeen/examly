<template>
  <n-input
    v-model:value="value"
    :placeholder="placeholder"
  />
</template>

<script lang="ts" setup>
import { NInput } from 'naive-ui'
import { computed, watch } from 'vue'

const value = defineModel<string | undefined | null>('value')

const props = defineProps<{
  isLast: boolean
  isSecondToLast: boolean
}>()

const emit = defineEmits<{
  'lastUpdated': []
  'secondToLastUpdated': []
}>()

const placeholder = computed(() => {
  return props.isLast
    ? 'Добавить ответ'
    : 'Ответ'
})

watch(value, () => {
  if (props.isLast && typeof value.value === 'string' && value.value.length > 0) {
    emit('lastUpdated')
  }

  if (props.isSecondToLast && !value.value) {
    emit('secondToLastUpdated')
  }
})
</script>
