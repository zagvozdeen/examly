<template>
  <div class="flex flex-col gap-4">
    <span class="text-2xl text-center py-16">Какому курсу ты хочешь научиться?</span>

    <ul class="flex flex-col bg-obscure-700 rounded-md overflow-hidden">
      <template
        v-for="(course, index) in courses"
        :key="course.id"
      >
        <li>
          <router-link
            class="grid grid-cols-[28px_1fr_min-content] items-center gap-2 hover:bg-obscure-500 bg-opacity-50 p-2"
            :to="{
              name: 'courses.show',
              params: { uuid: course.uuid },
            }"
          >
            <div
              class="rounded w-full py-0.5 text-center"
              :class="{
                [course.color]: true,
              }"
            >
              <i
                class="bi"
                :class="{
                  [course.icon]: true,
                }"
              />
            </div>
            <div class="flex flex-col">
              <span>{{ course.name }}</span>
              <span
                v-if="course.description"
                class="text-xs text-gray-400"
              >{{ course.description }}</span>
            </div>
            <i class="bi bi-chevron-right" />
          </router-link>
        </li>
        <li
          v-if="index + 1 !== courses.length"
          class="h-px w-[calc(100%-28px-1rem)] ml-auto bg-obscure-500"
        />
      </template>
    </ul>
  </div>
</template>

<script lang="ts" setup>
import { Course, PageExpose } from '@/types.ts'
import { useCourseStore } from '@/composables/useCourseStore.ts'
import { onMounted, ref } from 'vue'

defineExpose<PageExpose>({
  title: 'Главная',
})

const courseStore = useCourseStore()

const courses = ref<Course[]>([])

onMounted(() => {
  courseStore
    .getCourses()
    .then(data => {
      courses.value = data.data
    })
})
</script>