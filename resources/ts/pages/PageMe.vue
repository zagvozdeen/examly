<template>
  <div class="flex flex-col gap-4">
    <span class="text-2xl text-center py-8">Личный кабинет</span>

    <router-link
      :to="{ name: 'my.settings' }"
      class="bg-obscure-700 rounded-md p-2 hover:bg-obscure-500 bg-opacity-50"
    >
      <div class="grid grid-cols-[min-content_1fr_min-content] items-center gap-2">
        <div class="w-10 h-10 rounded-full bg-gradient-to-tr from-gray-400 to-gray-500" />
        <div class="flex flex-col">
          <span class="text-sm font-medium">{{ me?.full_name }}</span>
          <span class="text-gray-400 text-xs">{{ me?.email }} • {{ UserRoleTranslates[me?.role || UserRole.GuestRole] }}</span>
        </div>
        <i class="bi bi-chevron-right" />
      </div>
    </router-link>

    <ul class="flex flex-col bg-obscure-700 rounded-md overflow-hidden">
      <li>
        <router-link
          class="grid grid-cols-[28px_1fr_min-content] items-center gap-2 hover:bg-obscure-500 bg-opacity-50 p-2"
          :to="{
            name: 'my.courses',
          }"
        >
          <div class="bg-blue-400 rounded w-full py-0.5 text-center">
            <i class="bi bi-card-checklist" />
          </div>
          <span>Мои курсы</span>
          <i class="bi bi-chevron-right" />
        </router-link>
      </li>
      <li class="h-px w-[calc(100%-28px-1rem)] ml-auto bg-obscure-500" />
      <li>
        <router-link
          class="grid grid-cols-[28px_1fr_min-content] items-center gap-2 hover:bg-obscure-500 bg-opacity-50 p-2"
          :to="{
            name: 'my.modules',
          }"
        >
          <div class="bg-red-400 rounded w-full py-0.5 text-center">
            <i class="bi bi-bar-chart-steps" />
          </div>
          <span>Добавленные модули</span>
          <i class="bi bi-chevron-right" />
        </router-link>
      </li>
      <li class="h-px w-[calc(100%-28px-1rem)] ml-auto bg-obscure-500" />
      <li>
        <router-link
          class="grid grid-cols-[28px_1fr_min-content] items-center gap-2 hover:bg-obscure-500 bg-opacity-50 p-2"
          :to="{
            name: 'my.questions',
          }"
        >
          <div class="bg-green-400 rounded w-full py-0.5 text-center">
            <i class="bi bi-question-circle-fill" />
          </div>
          <span>Созданные вопросы</span>
          <i class="bi bi-chevron-right" />
        </router-link>
      </li><li class="h-px w-[calc(100%-28px-1rem)] ml-auto bg-obscure-500" />
      <li>
        <router-link
          class="grid grid-cols-[28px_1fr_min-content] items-center gap-2 hover:bg-obscure-500 bg-opacity-50 p-2"
          :to="{
            name: 'my.questions.import',
          }"
        >
          <div class="bg-purple-400 rounded w-full py-0.5 text-center">
            <i class="bi bi-file-earmark-arrow-up-fill" />
          </div>
          <span>Импортировать вопросы</span>
          <i class="bi bi-chevron-right" />
        </router-link>
      </li>
    </ul>
  </div>
</template>

<script lang="ts" setup>
import { me } from '../composables/useAuthStore.ts'
import { PageExpose, UserRole, UserRoleTranslates } from '@/types.ts'

defineExpose<PageExpose>({
  title: 'Личный кабинет',
})
</script>
