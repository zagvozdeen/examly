<template>
  <div class="flex flex-col gap-4">
    <AppLogo />

    <router-link to="/login">
      Вход
    </router-link>

    <router-link to="/register">
      Регистрация
    </router-link>
  </div>
</template>

<script lang="ts" setup>
import { useAuthStore } from '@/composables/useAuthStore.ts'
import AppLogo from '@/components/AppLogo.vue'

const authStore = useAuthStore()

authStore.getMe()
</script>