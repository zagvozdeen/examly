<template>
  <div class="flex flex-col gap-4">
    <span class="text-2xl text-center py-8">Моя статистика</span>

    <div
      v-if="testSessions.length === 0"
      class="text-center text-base"
    >
      Пройдите хотя бы один тест, чтобы увидеть статистику
    </div>

    <ul
      v-else
      class="flex flex-col"
    >
      <li class="border-b border-gray-400 text-gray-400 text-xs">
        <div class="grid grid-cols-[1fr_40px_48px] p-2">
          <small>Название</small>
          <small class="text-right">Решено</small>
          <small class="text-right">Всего</small>
        </div>
      </li>
      <li
        v-for="stat in testSessions"
        :key="stat.id"
        class="text-xs"
      >
        <router-link
          :to="{
            name: 'tests.show',
            params: { uuid: stat.uuid },
          }"
          class="grid grid-cols-[1fr_40px_48px] gap-y-2 p-2 hover:bg-obscure-600"
        >
          <span>{{ TestSessionTypeTranslates[stat.type] }} от {{ format(parseISO(stat.created_at), 'dd.MM.yyyy HH:mm') }}</span>
          <span class="text-right">
            <span class="text-green-400">{{ stat.correct }}</span>
            <span> / </span>
            <span class="text-red-400">{{ stat.incorrect }}</span>
          </span>
          <span class="text-right text-gray-400">{{ stat.question_ids.length }}</span>
          <span class="relative h-1 col-span-full bg-gray-700">
            <span
              class="absolute h-1 block bg-blue-500"
              :style="{
                'width': `${Math.round((stat.correct + stat.incorrect) / stat.question_ids.length * 1000) / 10}%`,
              }"
            />
          </span>
        </router-link>
      </li>
    </ul>
  </div>
</template>

<script lang="ts" setup>
import { PageExpose, TestSession, TestSessionTypeTranslates } from '@/types.ts'
import { onMounted, ref } from 'vue'
import { format, parseISO } from 'date-fns'
import { useTestSessionStore } from '@/composables/useTestSessionStore.ts'

const testSessionStore = useTestSessionStore()
const testSessions = ref<TestSession[]>([])

defineExpose<PageExpose>({
  title: 'Статистика',
})

onMounted(() => {
  testSessionStore
    .getTestSessions({})
    .then(data => {
      testSessions.value = data.data || []
    })
})
</script>