<template>
  <div class="max-w-md mx-auto p-4">
    <router-view v-slot="{ Component }">
      <component
        :is="Component"
        ref="currentComponent"
      />
    </router-view>
  </div>
</template>

<script lang="ts" setup>
import { onMounted, ref, watch } from 'vue'
import { PageExpose } from '@/types.js'
import { useAuthStore } from '@/composables/useAuthStore.ts'

const authStore = useAuthStore()

const currentComponent = ref<PageExpose>()

watch(currentComponent, (component) => {
  document.title = component?.title || import.meta.env.VITE_APP_NAME
})

onMounted(() => {
  authStore.getMe()
})
</script>