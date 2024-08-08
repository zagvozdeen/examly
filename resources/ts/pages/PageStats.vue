<template>
  <div class="flex flex-col gap-4">
    <span class="text-2xl text-center py-8">Статистика</span>
  </div>
</template>

<script lang="ts" setup>
import { PageExpose } from '@/types.ts'
import { onMounted } from 'vue'

defineExpose<PageExpose>({
  title: 'Статистика',
})

onMounted(() => {
})
</script>