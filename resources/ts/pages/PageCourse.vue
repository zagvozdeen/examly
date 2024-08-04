<template>
  <div class="flex flex-col gap-4">
    <span class="text-2xl text-center py-8">Курс «{{ course?.name }}»</span>

    <router-link :to="{ name: 'main' }">
      BACK
    </router-link>
  </div>
</template>

<script lang="ts" setup>
import { Course, PageExpose } from '@/types.ts'
import { useCourseStore } from '@/composables/useCourseStore.ts'
import { onMounted, ref } from 'vue'
import { useRoute } from 'vue-router'

defineExpose<PageExpose>({
  title: 'Курс',
})

const route = useRoute()
const courseStore = useCourseStore()

const course = ref<Course>()

onMounted(() => {
  courseStore
    .getCourseByUuid(route.params.uuid as string)
    .then(data => {
      course.value = data.data
    })
})
</script>