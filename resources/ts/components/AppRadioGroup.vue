<template>
  <n-radio-group
    v-model:value="value"
    class="w-full sm:!flex !hidden"
    :disabled="disabled"
  >
    <n-radio-button
      v-for="type in props.options"
      :key="type.value"
      type="primary"
      :value="type.value"
      :label="type.label"
      class="flex-1 text-center"
    />
  </n-radio-group>
  <n-select
    v-model:value="value"
    class="sm:hidden block"
    :options="props.options"
    :disabled="disabled"
  />
</template>

<script lang="ts" setup>
import { NRadioButton, NRadioGroup, NSelect } from 'naive-ui'

const value = defineModel<number | undefined | null>('value', { required: true })
const props = defineProps<{
  options: {
    value: number
    label: string
  }[]
  disabled: boolean
}>()
</script>