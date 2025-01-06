<template>
  <div
    v-if="testSession && currentQuestion"
    class="flex flex-col gap-4"
  >
    <ul class="flex gap-2 overflow-auto pb-4">
      <li
        v-for="(question, index) in testSession.questions"
        :key="question.id"
      >
        <button
          class="w-8 h-8 rounded relative"
          :class="{
            'bg-obscure-500': !answers[question.id] && currentQuestion.id === question.id,
            'bg-obscure-600 hover:bg-obscure-500': !answers[question.id] && currentQuestion.id !== question.id,
            'bg-green-500 hover:bg-green-600': answers[question.id] && answers[question.id].is_correct,
            'bg-red-500 hover:bg-red-600': answers[question.id] && !answers[question.id].is_correct,
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
      v-if="!currentQuestion.content"
      class="grid grid-cols-[min-content_1fr] items-center gap-2 -mt-4"
    >
      <span class="text-gray-400 text-xs whitespace-nowrap select-none">Вопрос без содержания</span>
      <span class="h-px bg-obscure-50d0 bg-gray-400" />
    </div>
    <div
      v-else
      class="grid grid-cols-[min-content_1fr] items-center gap-2 -mt-4"
    >
      <span class="text-gray-400 text-xs whitespace-nowrap select-none">Вопрос с доп. информацией</span>
      <span class="h-px bg-obscure-50d0 bg-gray-400" />
    </div>

    <span class="text-base">{{ currentQuestion.title }}</span>
    <span class="text-base">{{ currentQuestion.content }}</span>

    <n-input
      v-if="currentQuestion.type === QuestionType.Plaintext"
      v-model:value="form.input"
      size="small"
      placeholder="Введите ответ"
      :disabled="!!answers[currentQuestion.id]"
      @keydown.enter="checkAnswer"
    />

    <ul
      v-if="currentQuestion.type === QuestionType.SingleChoice"
      class="flex flex-col gap-2"
    >
      <li
        v-for="(answer, index) in currentQuestion.options"
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
      v-if="currentQuestion.type === QuestionType.MultipleChoice"
      class="flex flex-col gap-2"
    >
      <li
        v-for="(answer, index) in currentQuestion.options"
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

    <template v-if="answers[currentQuestion.id]">
      <span>Правильный ответ: {{ currentQuestion.options.filter(a => a.is_correct).map(a => a.content).join(', ') }}</span>
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
      v-show="(form.answers_ids.length > 0 || form.answer_id || form.input.length > 0) && !answers[currentQuestion.id]"
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
import { PageExpose, Question, QuestionType, TestSession, UserAnswer } from '@/types.ts'
import { computed, onMounted, reactive, ref, watch } from 'vue'
import { RouteLocationResolved, useRoute, useRouter } from 'vue-router'
import { NButton, NInput, useLoadingBar, useMessage } from 'naive-ui'
import { useTestSessionStore } from '@/composables/useTestSessionStore.ts'
import { useUserAnswerStore } from '@/composables/useUserAnswerStore.ts'

const route = useRoute()
const router = useRouter()
const loadingBar = useLoadingBar()
const message = useMessage()
const testSessionStore = useTestSessionStore()
const userAnswerStore = useUserAnswerStore()
const back = ref<RouteLocationResolved>(router.resolve({ name: 'main' }))

defineExpose<PageExpose>({
  title: 'Прохождение теста',
  back: back,
})

const loading = ref<boolean>(false)
const testSession = ref<TestSession>()
const currentQuestion = ref<Question>()
const answers = ref<Record<number, UserAnswer>>({})
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

const onClickQuestion = (question: Question) => {
  currentQuestion.value = question
}

const checkAnswer = () => {
  if (!currentQuestion.value || !testSession.value) {
    message.error('Вопрос не найден')
    return
  }

  loading.value = true
  loadingBar.start()

  userAnswerStore
    .checkAnswer({
      course_id: testSession.value.id,
      question_id: currentQuestion.value.id,
      answer_id: form.answer_id,
      answers_ids: form.answers_ids,
      plaintext: form.input,
    })
    .then(data => {
      answers.value[data.data.question_id] = data.data
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
  if (!testSession.value || !testSession.value.questions) return -1

  const index = testSession.value.questions.findIndex(question => question.id === currentQuestion.value?.id)

  if (index === -1) return -1

  return testSession.value.questions[index + 1] ? index + 1 : -1
})

const nextQuestion = () => {
  if (nextQuestionIndex.value !== -1 && testSession.value && testSession.value.questions) {
    currentQuestion.value = testSession.value.questions[nextQuestionIndex.value]
  }
}

watch(currentQuestion, (question: Question | undefined) => {
  clearForm()

  const answer = answers.value[question?.id || 0]
  if (answer) {
    form.answer_id = answer.answer_data.answer_id
    form.answers_ids = answer.answer_data.answers_ids
    form.input = answer.answer_data.plaintext
  }
})

onMounted(() => {
  testSessionStore
    .getTestSession(route.params.uuid as string)
    .then(data => {
      if (data.data && data.data.questions) {
        testSession.value = data.data
        answers.value = data.answers
        back.value = router.resolve({
          name: 'courses.show',
          params: {
            uuid: data.data.course_uuid,
          },
        })

        if (data.data.last_question_id) {
          const question = data.data.questions.find(question => question.id === data.data.last_question_id)
          currentQuestion.value = question || data.data.questions[0]
        } else {
          currentQuestion.value = data.data.questions[0]
        }
      }
    })
})
</script>