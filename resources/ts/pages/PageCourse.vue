<template>
  <div
    v-if="course && stats"
    class="flex flex-col gap-4"
  >
    <span class="text-2xl text-center py-8">Курс «{{ course.name }}»</span>

    <router-link
      :to="{ name: 'stats', query: { course_uuid: course.uuid } }"
      class="bg-obscure-600 hover:bg-obscure-550 cursor-pointer rounded px-2 pt-2 pb-4"
    >
      <img
        class="mx-auto"
        src="../../images/Education.svg"
        alt="Education process"
        width="240"
        height="240"
      >
      <span class="grid grid-cols-[min-content_min-content_min-content] justify-center gap-10 text-xs px-10">
        <span
          v-for="stat in stats"
          :key="stat.name"
          class="flex flex-col gap-1 text-center"
        >
          <span class="flex gap-0.5 justify-center">
            <span class="text-gray-300">{{ stat.completed }}</span>
            <span class="text-gray-500">/</span>
            <span class="text-gray-500">{{ stat.total }}</span>
          </span>
          <span class="h-1 bg-obscure-500 group-hover:bg-gray-600 relative w-16">
            <span
              class="absolute block h-1 bg-blue-500"
              :style="{
                width: `${Math.round(stat.completed / stat.total * 1000) / 10}%`,
              }"
            />
          </span><span class="text-gray-300">{{ stat.name }}</span>
        </span>
      </span>
    </router-link>

    <div class="grid grid-cols-2 gap-4">
      <router-link
        :to="{
          name: 'courses.show.marathon',
          params: { uuid: course.uuid },
        }"
      >
        <n-button
          type="info"
          class="!w-full"
        >
          Марафон
        </n-button>
      </router-link>

      <n-button type="error">
        Экзамен
      </n-button>
      <n-button
        secondary
        type="success"
      >
        Модули
      </n-button>
      <n-button
        secondary
        type="success"
      >
        Ошибки
      </n-button>
      <n-button
        type="primary"
        class="col-span-2"
      >
        База
      </n-button>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { Course, CourseStats, PageExpose } from '@/types.ts'
import { useCourseStore } from '@/composables/useCourseStore.ts'
import { onMounted, ref } from 'vue'
import { useRoute } from 'vue-router'
import { NButton } from 'naive-ui'

defineExpose<PageExpose>({
  title: 'Курс',
})

const route = useRoute()
const courseStore = useCourseStore()

const course = ref<Course>()
const stats = ref<CourseStats>()

onMounted(() => {
  courseStore
    .getCourseByUuid(route.params.uuid as string)
    .then(data => {
      course.value = data.data
      stats.value = data.stats
    })
})
</script>