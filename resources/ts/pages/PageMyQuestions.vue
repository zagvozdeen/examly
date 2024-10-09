<template>
  <div class="flex flex-col gap-4">
    <span class="text-gray-400 text-center text-xs">
      Вопросы и ответы тоже проходят проверку модерацией
      <br>
      перед тем как попасть в общий доступ.
    </span>

    <router-link
      :to="{ name: 'questions.create' }"
      class="sm:self-center"
    >
      <n-button
        type="info"
        class="sm:!px-10 !w-full"
      >
        <div class="flex items-center gap-2">
          <i class="bi bi-plus-square-fill" />
          <span>Создать вопрос</span>
        </div>
      </n-button>
    </router-link>

    <n-table v-if="questions.length > 0">
      <thead>
        <tr>
          <th>Название</th>
          <th>Статус</th>
        </tr>
      </thead>
      <tbody>
        <tr
          v-for="question in questions"
          :key="question.id"
        >
          <td>{{ question.title }}</td>
          <td>{{ StatusTranslates[question.status] }}</td>
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
import { StatusTranslates, PageExpose, Question } from '@/types.ts'
import { useQuestionStore } from '@/composables/useQuestionStore.ts'
import { me } from '@/composables/useAuthStore.ts'

const router = useRouter()
const questionStore = useQuestionStore()

const questions = ref<Question[]>([])

defineExpose<PageExpose>({
  title: 'Созданные вопросы',
  back: router.resolve({ name: 'me' }),
})

onMounted(() => {
  questionStore
    .getQuestions({
      created_by: me.value?.id,
    })
    .then(data => {
      if (data.data) {
        questions.value = data.data
      }
    })
})
</script>
