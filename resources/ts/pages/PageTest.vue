<template>
  <div
    v-if="course && currentQuestion"
    class="flex flex-col gap-4"
  >
    <ul class="flex gap-2 overflow-auto pb-4">
      <li
        v-for="(question, index) in course.questions"
        :key="question.id"
      >
        <button
          class="w-8 h-8 rounded relative"
          :class="{
            'bg-obscure-500': question.is_true === null && currentQuestion.id === question.id,
            'bg-obscure-600 hover:bg-obscure-500': question.is_true === null && currentQuestion.id !== question.id,
            'bg-green-500 hover:bg-green-600': question.is_true !== null && question.is_true,
            'bg-red-500 hover:bg-red-600': question.is_true !== null && !question.is_true,
          }"
          @click="onClickQuestion(question)"
        >
          {{ index + 1 }}
          <span
            class="absolute w-2 h-2 top-[calc(100%-0.25rem)] left-[calc(50%-0.25rem)] rotate-45 bg-obscure-800"
            :class="{
              'hidden': currentQuestion.id !== question.id,
            }"
          />
        </button>
      </li>
    </ul>

    <div
      v-if="!currentQuestion.file_id"
      class="grid grid-cols-[min-content_1fr] items-center gap-2 -mt-4"
    >
      <span class="text-gray-400 text-xs whitespace-nowrap select-none">Вопрос без изображения</span>
      <span class="h-px bg-obscure-50d0 bg-gray-400" />
    </div>

    <span class="text-base">{{ currentQuestion.content }}</span>

    <n-input
      v-if="currentQuestion.type === QuestionType.InputType"
      v-model:value="form.input"
      size="small"
      placeholder="Введите ответ"
    />

    <ul
      v-if="currentQuestion.type === QuestionType.OneAnswerType"
      class="flex flex-col gap-2"
    >
      <li
        v-for="(answer, index) in currentQuestion.answers"
        :key="answer.id"
      >
        <label
          :for="`answer-${answer.id}`"
          class="grid grid-cols-[min-content_1fr] rounded gap-2 select-none p-2 text-xs cursor-pointer"
          :class="{
            'bg-obscure-600 hover:bg-obscure-500': form.answer_id !== answer.id,
            'bg-obscure-500': form.answer_id === answer.id,
          }"
        >
          <input
            :id="`answer-${answer.id}`"
            v-model="form.answer_id"
            type="radio"
            name="answer-radio-button"
            class="hidden"
            :value="answer.id"
          >
          <span>{{ index + 1 }}.</span>
          <span>{{ answer.content }}</span>
        </label>
      </li>
    </ul>

    <ul
      v-if="currentQuestion.type === QuestionType.MultiplyAnswersType"
      class="flex flex-col gap-2"
    >
      <li
        v-for="(answer, index) in currentQuestion.answers"
        :key="answer.id"
      >
        <label
          :for="`answer-${answer.id}`"
          class="grid grid-cols-[min-content_1fr] rounded gap-2 select-none p-2 text-xs cursor-pointer"
          :class="{
            'bg-obscure-600 hover:bg-obscure-500': !form.answers_ids.includes(answer.id),
            'bg-obscure-500': form.answers_ids.includes(answer.id),
          }"
        >
          <input
            :id="`answer-${answer.id}`"
            v-model="form.answers_ids"
            type="checkbox"
            name="answer-checkbox"
            class="hidden"
            :value="answer.id"
          >
          <span>{{ index + 1 }}.</span>
          <span>{{ answer.content }}</span>
        </label>
      </li>
    </ul>

    <template v-if="currentQuestion.is_true !== null">
      <span>Правильный ответ: {{ currentQuestion.answers.filter(a => a.is_true).map(a => a.content).join(', ') }}</span>
      <n-button
        type="primary"
        size="small"
        :disabled="nextQuestionIndex === -1"
        @click="nextQuestion"
      >
        Следующий вопрос
      </n-button>
    </template>

    <n-button
      v-show="(form.answer_id || form.answers_ids.length > 0 || form.input.length > 0) && currentQuestion.is_true === null"
      type="primary"
      size="small"
      :loading="loading"
      :disabled="loading"
      @click="checkAnswer"
    >
      Подтвердить ответ
    </n-button>
  </div>
</template>

<script lang="ts" setup>
import { PageExpose, QuestionType, UserCourse, UserQuestion } from '@/types.ts'
import { useCourseStore } from '@/composables/useCourseStore.ts'
import { computed, onMounted, reactive, ref, watch } from 'vue'
import { useRoute } from 'vue-router'
import { NButton, NInput, useLoadingBar, useMessage } from 'naive-ui'

defineExpose<PageExpose>({
  title: 'Прохождение теста',
})

const route = useRoute()
const courseStore = useCourseStore()
const loadingBar = useLoadingBar()
const message = useMessage()

const loading = ref<boolean>(false)
const course = ref<UserCourse>()
const currentQuestion = ref<UserQuestion>()
const form = reactive({
  answer_id: null as null | number,
  answers_ids: [] as number[],
  input: '',
})

const clearForm = () => {
  form.answer_id = null
  form.answers_ids = []
  form.input = ''
}

const onClickQuestion = (question: UserQuestion) => {
  currentQuestion.value = question
}

const checkAnswer = () => {
  if (!currentQuestion.value) {
    message.error('Вопрос не найден')
    return
  }

  loading.value = true
  loadingBar.start()

  courseStore
    .checkAnswer(currentQuestion.value.uuid, form)
    .then(data => {
      if (course.value) {
        const index = course.value.questions.findIndex(question => question.id === currentQuestion.value?.id)
        course.value.questions[index] = data.data
        currentQuestion.value = data.data

        if (data.data.is_true) {
          nextQuestion()
        }
      }

      loadingBar.finish()
    })
    .catch(() => {
      loadingBar.error()
    })
    .finally(() => {
      loading.value = false
    })
}

const nextQuestionIndex = computed(() => {
  if (!course.value) return -1

  const index = course.value.questions.findIndex(question => question.id === currentQuestion.value?.id)

  if (index === -1) return -1

  return course.value.questions[index + 1] ? index + 1 : -1
})

const nextQuestion = () => {
  if (nextQuestionIndex.value !== -1 && course.value) {
    currentQuestion.value = course.value.questions[nextQuestionIndex.value]
  }
}

watch(currentQuestion, (question: UserQuestion | undefined) => {
  clearForm()

  if (question && question.is_true !== null) {
    switch (question.type) {
    case QuestionType.OneAnswerType:
      form.answer_id = question.answers.find(answer => answer.is_chosen)?.id || null
      break
    case QuestionType.MultiplyAnswersType:
      form.answers_ids = question.answers.filter(answer => answer.is_chosen).map(answer => answer.id)
      break
    }
  }
})

onMounted(() => {
  courseStore
    .getUserCourseByUuid(route.params.uuid as string)
    .then(data => {
      course.value = data.data

      if (data.data.last_question_id) {
        const question = data.data.questions.find(question => question.id === data.data.last_question_id)
        currentQuestion.value = question || data.data.questions[0]
      } else {
        currentQuestion.value = data.data.questions[0]
      }
    })
})
</script>