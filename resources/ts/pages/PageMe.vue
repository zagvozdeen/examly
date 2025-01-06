<template>
  <div class="flex flex-col gap-4">
    <span class="text-2xl text-center py-8">
      {{ me?.role === UserRole.Company ? 'Личный кабинет работодателя' : 'Личный кабинет' }}
    </span>

    <div
      v-if="!me || me.role === UserRole.Guest"
      class="grid grid-cols-[28px_1fr_min-content] items-center gap-2 rounded-md bg-obscure-500 bg-opacity-50 p-2"
    >
      <div class="bg-orange-400 rounded w-full py-0.5 text-center">
        <i class="bi bi-exclamation-lg" />
      </div>
      <div class="flex flex-col">
        <span>Вы вошли как гость</span>
        <span class="text-xs text-gray-100"><router-link
          class="underline"
          :to="{ name: 'login' }"
        >Войдите</router-link> или <router-link
          class="underline"
          :to="{ name: 'register' }"
        >зарегистрируйтесь</router-link> чтобы получить полный доступ к разделам</span>
      </div>
    </div>

    <template v-else>
      <router-link
        :to="{ name: 'settings' }"
        class="bg-obscure-700 rounded-md p-2 hover:bg-obscure-500 bg-opacity-50"
      >
        <div class="grid grid-cols-[min-content_1fr_min-content] items-center gap-2">
          <div class="w-10 h-10 rounded-full bg-gradient-to-tr from-gray-400 to-gray-500" />
          <div class="flex flex-col">
            <span class="text-sm font-medium">{{ me.full_name }}</span>
            <span class="text-gray-400 text-xs">{{ me.email }} • {{ UserRoleTranslates[me.role] }}</span>
          </div>
          <i class="bi bi-chevron-right" />
        </div>
      </router-link>

      <ul class="flex flex-col bg-obscure-700 rounded-md overflow-hidden">
        <li>
          <router-link
            class="grid grid-cols-[28px_1fr_min-content] items-center gap-2 hover:bg-obscure-500 bg-opacity-50 p-2"
            :to="{
              name: 'courses',
            }"
          >
            <div class="bg-blue-400 rounded w-full py-0.5 text-center">
              <i class="bi bi-card-checklist" />
            </div>
            <span>{{ isAdminMode ? 'Список всех курсов' : 'Мои курсы' }}</span>
            <i class="bi bi-chevron-right" />
          </router-link>
        </li>
        <li class="h-px w-[calc(100%-28px-1rem)] ml-auto bg-obscure-500" />
        <li>
          <router-link
            class="grid grid-cols-[28px_1fr_min-content] items-center gap-2 hover:bg-obscure-500 bg-opacity-50 p-2"
            :to="{
              name: 'modules',
            }"
          >
            <div class="bg-red-400 rounded w-full py-0.5 text-center">
              <i class="bi bi-bar-chart-steps" />
            </div>
            <span>{{ isAdminMode ? 'Список модулей' : 'Добавленные модули' }}</span>
            <i class="bi bi-chevron-right" />
          </router-link>
        </li>
        <li class="h-px w-[calc(100%-28px-1rem)] ml-auto bg-obscure-500" />
        <li>
          <router-link
            class="grid grid-cols-[28px_1fr_min-content] items-center gap-2 hover:bg-obscure-500 bg-opacity-50 p-2"
            :to="{
              name: 'questions',
            }"
          >
            <div class="bg-green-400 rounded w-full py-0.5 text-center">
              <i class="bi bi-question-circle-fill" />
            </div>
            <span>{{ isAdminMode ? 'Список всех вопросов' : 'Созданные вопросы' }}</span>
            <i class="bi bi-chevron-right" />
          </router-link>
        </li>
        <li class="h-px w-[calc(100%-28px-1rem)] ml-auto bg-obscure-500" />
        <li>
          <router-link
            class="grid grid-cols-[28px_1fr_min-content] items-center gap-2 hover:bg-obscure-500 bg-opacity-50 p-2"
            :to="{
              name: 'selection-system',
            }"
          >
            <div class="bg-zinc-400 rounded w-full py-0.5 text-center">
              <i class="bi bi-list-columns" />
            </div>
            <span>Система подбора</span>
            <i class="bi bi-chevron-right" />
          </router-link>
        </li>
        <template v-if="isModerator">
          <li class="h-px w-[calc(100%-28px-1rem)] ml-auto bg-obscure-500" />
          <li>
            <div
              class="grid grid-cols-[28px_1fr_min-content] items-center gap-2 bg-opacity-50 p-2"
            >
              <div class="bg-orange-400 rounded w-full py-0.5 text-center">
                <i class="bi bi-arrow-left-right" />
              </div>
              <span class="select-none">Режим администратора</span>
              <AppAdminMode />
            </div>
          </li>
        </template>
      </ul>

      <button
        class="block w-full gap-2 bg-obscure-700 rounded-md hover:bg-obscure-500 bg-opacity-50 p-2 text-red-400"
        type="button"
        @click="logout"
      >
        <span>Выйти</span>
      </button>
    </template>
  </div>
</template>

<script lang="ts" setup>
import { isAdminMode, isModerator, me, useAuthStore } from '../composables/useAuthStore.ts'
import { PageExpose, UserRole, UserRoleTranslates } from '@/types.ts'
import AppAdminMode from '@/components/AppAdminMode.vue'
import { useRouter } from 'vue-router'
import { useMessage } from 'naive-ui'

defineExpose<PageExpose>({
  title: 'Личный кабинет',
})

const router = useRouter()
const message = useMessage()
const authStore = useAuthStore()

const logout = () => {
  authStore.logout()
  router.push({ name: 'login' })
  message.info('Вы вышли из аккаунта')
}
</script>
