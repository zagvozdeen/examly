<template>
  <header>
    <div class="grid grid-cols-[60px_1fr_60px] gap-2 items-center">
      <router-link
        :to="back"
        class="text-blue-400 hover:text-blue-500 font-medium inline-flex items-center gap-0.5"
      >
        <i class="bi bi-chevron-left" />
        <span>Назад</span>
      </router-link>
      <span class="text-center text-base select-none">
        {{ title }}
      </span>
    </div>
  </header>
</template>

<script lang="ts" setup>
import { RouteLocationResolved } from 'vue-router'

defineProps<{
  title: string
  back: RouteLocationResolved
}>()
</script>
