<template>
  main

  <router-link to="/login">
    Вход
  </router-link>

  <router-link to="/register">
    Регистрация
  </router-link>
</template>

<script lang="ts" setup>
import { useAuthStore } from '@/composables/useAuthStore.ts'

const authStore = useAuthStore()

authStore.getMe()
</script>