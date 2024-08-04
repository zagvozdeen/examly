<template>
  <div class="flex flex-col gap-4">
    <span>Acting as: {{ me?.email || me?.role }}</span>

    <span>Courses:</span>
    <ul class="list-inside list-disc">
      <li
        v-for="course in courses"
        :key="course.id"
      >
        {{ course.name }}
      </li>
    </ul>

    <router-link to="/login">
      Вход
    </router-link>

    <router-link to="/register">
      Регистрация
    </router-link>
  </div>
</template>

<script lang="ts" setup>
import { Course, PageExpose } from '@/types.ts'
import { me } from '@/composables/useAuthStore.ts'
import { useCourseStore } from '@/composables/useCourseStore.ts'
import { onMounted, ref } from 'vue'

defineExpose<PageExpose>({
  title: 'TEST',
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