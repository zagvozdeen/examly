<template>
  <div class="sm:max-w-xl mx-auto sm:p-4">
    <router-view v-slot="{ Component }">
      <div class="sm:block sm:max-h-none sm:min-h-0 sm:overflow-auto grid min-h-dvh max-h-dvh overflow-hidden grid-rows-[1fr_min-content]">
        <div class="overflow-auto sm:p-0 p-4 flex flex-col gap-4">
          <AppHeaderWithBack
            v-if="currentComponent?.back"
            :title="currentComponent.title"
            :back="isRef(currentComponent.back) ? currentComponent.back.value : currentComponent.back"
          />
          <template v-else>
            <AppHeader class="sm:block hidden" />
            <router-link
              class="sm:hidden"
              :to="{ name: 'main' }"
            >
              <AppLogo />
            </router-link>
          </template>
          <main>
            <div
              v-if="loading"
              class="text-base mt-4 text-center"
            >
              <n-spin size="small" />
            </div>

            <component
              :is="Component"
              v-else
              ref="currentComponent"
            />
          </main>
        </div>
        <AppFooter
          v-if="!currentComponent?.back"
          class="sm:hidden"
        />
      </div>
    </router-view>
  </div>
</template>

<script lang="ts" setup>
import { isRef, onMounted, ref, watch } from 'vue'
import { PageExpose } from '@/types.js'
import { useAuthStore } from '@/composables/useAuthStore.ts'
import { NSpin } from 'naive-ui'
import AppFooter from '@/components/AppFooter.vue'
import AppHeader from '@/components/AppHeader.vue'
import AppLogo from '@/components/AppLogo.vue'
import AppHeaderWithBack from '@/components/AppHeaderWithBack.vue'

const authStore = useAuthStore()

const currentComponent = ref<PageExpose>()
const loading = ref<boolean>(true)

watch(currentComponent, (component) => {
  document.title = component?.title || import.meta.env.VITE_APP_NAME
})

onMounted(async () => {
  await authStore.getMe()
  loading.value = false
})
</script>