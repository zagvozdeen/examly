<template>
  <div class="flex flex-col gap-4">
    <span
      v-if="!isStatsByCourse"
      class="text-2xl text-center py-8"
    >Моя статистика</span>

    <ul class="flex flex-col">
      <li class="border-b border-gray-400 text-gray-400 text-xs">
        <div class="grid grid-cols-[1fr_40px_48px] p-2">
          <small>Название</small>
          <small class="text-right">Решено</small>
          <small class="text-right">Всего</small>
        </div>
      </li>
      <li
        v-for="stat in stats"
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
          <span>{{ UserCourseTypeTranslates[stat.type] }} от {{ format(stat.created_at, 'dd.MM.yyyy HH:mm') }}</span>
          <span class="text-right">
            <span class="text-green-400">{{ stat.correct }}</span>
            <span> / </span>
            <span class="text-red-400">{{ stat.incorrect }}</span>
          </span>
          <span class="text-right text-gray-400">{{ stat.total }}</span>
          <span class="relative h-1 col-span-full bg-gray-700">
            <span
              class="absolute h-1 block bg-blue-500"
              :style="{
                'width': `${Math.round((stat.correct + stat.incorrect) / stat.total * 1000) / 10}%`,
              }"
            />
          </span>
        </router-link>
      </li>
    </ul>
  </div>
</template>

<script lang="ts" setup>
import { FullCourseStats, PageExpose, UserCourseTypeTranslates } from '@/types.ts'
import { onMounted, ref } from 'vue'
import { useStatsStore } from '@/composables/useStatsStore.ts'
import { useRoute, useRouter } from 'vue-router'
import { format } from 'date-fns'

const route = useRoute()
const router = useRouter()
const statsStore = useStatsStore()

const isStatsByCourse = route.name === 'courses.stats'

const stats = ref<FullCourseStats[]>()

defineExpose<PageExpose>({
  title: 'Статистика',
  back: isStatsByCourse
    ? router.resolve({
      name: 'courses.show',
      params: { uuid: route.params.uuid },
    })
    : undefined,
})

onMounted(() => {
  if (isStatsByCourse) {
    statsStore
      .getCourseStats(route.params.uuid as string)
      .then(data => {
        stats.value = data.data
      })
  } else {
    statsStore
      .getStats()
      .then(data => {
        stats.value = data.data
      })
  }
})
</script>