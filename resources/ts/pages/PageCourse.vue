<template>
  <div
    v-if="course"
    class="flex flex-col gap-4"
  >
    <span class="text-2xl text-center py-8">Курс «{{ course.name }}»</span>

    <ul class="flex flex-col bg-obscure-700 rounded-md overflow-hidden">
      <!--      <li>-->
      <!--        <router-link-->
      <!--          class="grid grid-cols-[28px_1fr_min-content] w-full text-left items-center gap-2 hover:bg-obscure-500 bg-opacity-50 p-2"-->
      <!--          :to="{-->
      <!--            name: 'courses.stats',-->
      <!--            params: { uuid: course.uuid },-->
      <!--          }"-->
      <!--        >-->
      <!--          <div class="bg-orange-400 rounded w-full py-0.5 text-center">-->
      <!--            <i class="bi bi-bar-chart-line-fill" />-->
      <!--          </div>-->
      <!--          <span>Статистика по курсу</span>-->
      <!--          <i class="bi bi-chevron-right" />-->
      <!--        </router-link>-->
      <!--      </li>-->
      <!--      <li class="h-px w-[calc(100%-28px-1rem)] ml-auto bg-obscure-500" />-->
      <li>
        <button
          class="grid grid-cols-[28px_1fr_min-content] w-full text-left items-center gap-2 hover:bg-obscure-500 bg-opacity-50 p-2"
          @click="handleCreateMarathon"
        >
          <span class="bg-blue-400 rounded w-full py-0.5 text-center">
            <i class="bi bi-window-fullscreen" />
          </span>
          <span>Начать марафон</span>
          <i class="bi bi-chevron-right" />
        </button>
      </li>
      <li class="h-px w-[calc(100%-28px-1rem)] ml-auto bg-obscure-500" />
      <li>
        <button
          class="grid grid-cols-[28px_1fr_min-content] w-full text-left items-center gap-2 hover:bg-obscure-500 bg-opacity-50 p-2"
          @click="handleCreateExam"
        >
          <span class="bg-red-400 rounded w-full py-0.5 text-center">
            <i class="bi bi-patch-check-fill" />
          </span>
          <span>Начать экзамен</span>
          <i class="bi bi-chevron-right" />
        </button>
      </li>
      <li class="h-px w-[calc(100%-28px-1rem)] ml-auto bg-obscure-500" />
      <li>
        <router-link
          class="grid grid-cols-[28px_1fr_min-content] w-full text-left items-center gap-2 hover:bg-obscure-500 bg-opacity-50 p-2"
          :to="{
            name: 'tests.show',
            params: { uuid: mistakes?.uuid }
          }"
        >
          <div class="bg-slate-400 rounded w-full py-0.5 text-center">
            <i class="bi bi-x-octagon" />
          </div>
          <span>Ошибки</span>
          <i class="bi bi-chevron-right" />
        </router-link>
      </li>
    </ul>

    <span class="text-xl text-center pt-4">Статистика</span>

    <ul
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
import { Course, PageExpose, TestSession, TestSessionType, TestSessionTypeTranslates } from '@/types.ts'
import { useCourseStore } from '@/composables/useCourseStore.ts'
import { onMounted, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useMessage } from 'naive-ui'
import { useTestSessionStore } from '@/composables/useTestSessionStore.ts'
import { format, parseISO } from 'date-fns'

defineExpose<PageExpose>({
  title: 'Курс',
})

const route = useRoute()
const router = useRouter()
const message = useMessage()
const courseStore = useCourseStore()
const testSessionStore = useTestSessionStore()

const course = ref<Course>()
const mistakes = ref<TestSession>()
const testSessions = ref<TestSession[]>([])

const handleCreateExam = () => {
  const payload = {
    course_uuid: course.value?.uuid,
    type: TestSessionType.Exam,
    shuffle: true,
  }

  testSessionStore
    .createTestSession(payload)
    .then(data => {
      router.push({
        name: 'tests.show',
        params: { uuid: data.data.uuid },
      })

      message.success('Экзамен начат')
    })
}

const handleCreateMarathon = () => {
  const payload = {
    course_uuid: course.value?.uuid,
    type: TestSessionType.Marathon,
    shuffle: true,
  }

  testSessionStore
    .createTestSession(payload)
    .then(data => {
      router.push({
        name: 'tests.show',
        params: { uuid: data.data.uuid },
      })

      message.success('Марафон начат')
    })
}

onMounted(() => {
  courseStore
    .getCourseByUuid(route.params.uuid as string)
    .then(data => {
      course.value = data.data
      mistakes.value = data.mistakes
    })

  testSessionStore
    .getTestSessions({
      course_uuid: route.params.uuid as (string | undefined),
    })
    .then(data => {
      if (data.data) {
        testSessions.value = data.data
      }
    })
})
</script>