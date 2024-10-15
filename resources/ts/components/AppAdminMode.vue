<template>
  <AppleCheckbox v-model="isAdminMode" />
</template>

<script lang="ts" setup>
import AppleCheckbox from '@/components/AppleCheckbox.vue'
import { useMessage } from 'naive-ui'
import { watch } from 'vue'
import { isAdminMode } from '@/composables/useAuthStore.ts'

const message = useMessage()

watch(isAdminMode, (value: boolean) => {
  const text = value
    ? 'Режим администратора включён'
    : 'Режим администратора выключен'

  message.success(text)
})
</script>