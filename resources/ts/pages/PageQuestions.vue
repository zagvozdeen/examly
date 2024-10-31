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

    <n-table
      v-if="questions.length > 0"
      size="small"
    >
      <thead>
        <tr>
          <th>ID</th>
          <th>Название</th>
          <th>Статус</th>
          <th class="!text-right">
            Действия
          </th>
        </tr>
      </thead>
      <tbody>
        <tr
          v-for="question in questions"
          :key="question.id"
        >
          <td>{{ question.id }}</td>
          <td>{{ question.title }}</td>
          <td>
            <span
              class="rounded-full text-xs px-2 py-1 font-medium"
              :class="{
                [StatusBackgroundColors[question.status]]: true,
                [StatusTextColors[question.status]]: true,
              }"
            >{{ StatusTranslates[question.status] }}</span>
          </td>
          <td class="text-right">
            <router-link :to="{ name: 'questions.edit', params: {uuid: question.uuid} }">
              <n-button
                type="warning"
                size="tiny"
              >
                Редактировать
              </n-button>
            </router-link>
          </td>
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
import { StatusTranslates, PageExpose, Question, StatusBackgroundColors, StatusTextColors } from '@/types.ts'
import { useQuestionStore } from '@/composables/useQuestionStore.ts'
import { isAdminMode, me } from '@/composables/useAuthStore.ts'

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
      all: isAdminMode.value,
    })
    .then(data => {
      questions.value = data.data || []
    })
})
</script>
