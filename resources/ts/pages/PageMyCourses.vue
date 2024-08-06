<template>
  <div class="flex flex-col gap-4">
    <span class="text-gray-400 text-center text-xs">
      Здесь Вы можете посмотреть список созданных курсов и их статус.
      <br>
      Созданные курсы проходят модерацию, прежде чем попадут в общий доступ.
    </span>

    <router-link
      :to="{ name: 'my.courses.create' }"
      class="sm:self-center"
    >
      <n-button
        type="info"
        class="sm:!px-10 !w-full"
      >
        <div class="flex items-center gap-2">
          <i class="bi bi-plus-square-fill" />
          <span>Создать курс</span>
        </div>
      </n-button>
    </router-link>

    <n-table
      v-if="courses.length > 0"
      size="small"
    >
      <thead>
        <tr>
          <th>Название</th>
          <th>Статус</th>
        </tr>
      </thead>
      <tbody>
        <tr
          v-for="course in courses"
          :key="course.id"
        >
          <td>{{ course.name }}</td>
          <td>{{ CourseStatusTranslates[course.status] }}</td>
        </tr>
      </tbody>
    </n-table>
    <span
      v-else
      class="text-center text-gray-400"
    >Пока ничего</span>
  </div>
</template>

<script lang="ts" setup>
import { NTable, NButton } from 'naive-ui'
import { onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'
import { Course, CourseStatusTranslates, PageExpose } from '@/types.ts'
import { useCourseStore } from '@/composables/useCourseStore.ts'

const router = useRouter()
const courseStore = useCourseStore()

const courses = ref<Course[]>([])

defineExpose<PageExpose>({
  title: 'Мои курсы',
  back: router.resolve({ name: 'me' }),
})

onMounted(() => {
  courseStore
    .getMyCourses()
    .then(data => {
      courses.value = data.data
    })
})
</script>
