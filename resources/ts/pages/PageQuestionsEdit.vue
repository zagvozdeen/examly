<template>
  <div class="flex flex-col gap-4">
    <AppModerationForm
      v-if="isAdminMode && question"
      :status="question.status"
      :reason="question.moderation_reason"
      :callback="questionStore.moderateQuestion"
    />

    <div
      v-if="!isCreating"
      class="flex flex-col items-center gap-1.5 rounded-md bg-obscure-500 bg-opacity-50 p-3"
    >
      <div class="bg-orange-400 rounded w-9 py-0.5 text-center">
        <i class="bi bi-exclamation-lg text-2xl" />
      </div>
      <span class="text-lg">Изменения создадут новый вопрос</span>
      <span class="text-xs text-gray-100 text-center">
        Обратите внимание, что текущий вопрос не изменится, а будет создан новый.
        Когда вопрос пройдёт модерацию, то новый вопрос заменит старый.
        Так как кто-то может выполнять тест в данный момент, то мы рискуем сломать обратную совместимость.
        Поэтому в старых тестах сохранятся прошлые вопросы, но новые тесты будут обновлены.
        Не бойтесь что-то менять — после модерации изменения вступят в силу.
      </span>
    </div>

    <n-form
      ref="formRef"
      :rules="formRules"
      :model="formValue"
      @submit.prevent="onSubmit"
    >
      <n-form-item
        label="Курс"
        path="course_id"
      >
        <n-select
          v-model:value="formValue.course_id"
          placeholder="Выберите курс"
          :options="coursesOptions"
          @update:value="onChangeCourse"
        />
      </n-form-item>

      <n-form-item
        v-show="formValue.course_id"
        label="Модуль"
        path="module_id"
      >
        <n-select
          v-model:value="formValue.module_id"
          placeholder="Выберите модуль (необязательно)"
          :options="modulesOptions"
        />
      </n-form-item>

      <n-form-item
        label="Количество правильных ответов"
        path="type"
      >
        <n-radio-group
          v-model:value="formValue.type"
          class="w-full sm:!flex !hidden"
          @update:value="onChangeType"
        >
          <n-radio-button
            v-for="type in typesOptions"
            :key="type.value"
            type="primary"
            :value="type.value"
            :label="type.label"
            class="flex-1 text-center"
          />
        </n-radio-group>
        <n-select
          v-model:value="formValue.type"
          class="sm:hidden block"
          :options="typesOptions"
        />
      </n-form-item>

      <n-form-item
        label="Вопрос"
        path="title"
      >
        <n-input
          v-model:value="formValue.title"
          placeholder="Задайте вопрос"
          type="textarea"
          :rows="2"
        />
      </n-form-item>

      <n-form-item
        label="Содержание вопроса в формате Markdown (не обязательно)"
        path="content"
      >
        <AppTextEditor v-model="formValue.content" />
      </n-form-item>

      <n-form-item
        label="Варианты ответа"
        path="options"
      >
        <div class="flex flex-col gap-2 w-full">
          <template
            v-for="(answer, index) in formValue.options"
            :key="answer.id"
          >
            <div class="flex items-center gap-3">
              <n-checkbox
                v-if="formValue.type === QuestionType.MultipleChoice"
                v-model:checked="answer.is_correct"
                tabindex="-1"
              />
              <n-radio
                v-if="formValue.type === QuestionType.SingleChoice"
                :checked="answer.is_correct"
                name="questions-radio-button"
                :value="answer.id"
                tabindex="-1"
                @change="onChange(answer.id)"
              />
              <AppAnswerInput
                v-model:value="answer.content"
                :is-last="index === formValue.options.length - 1"
                :is-second-to-last="index === formValue.options.length - 2"
                @last-updated="onLastUpdated"
                @second-to-last-updated="onSecondToLastUpdated"
              />
            </div>
          </template>
        </div>
      </n-form-item>

      {{ formValue.options }}

      <n-form-item
        label="Объяснение"
        path="explanation"
      >
        <n-input
          v-model:value="formValue.explanation"
          placeholder="Введите объяснение (необязательно)"
          type="textarea"
          :rows="2"
        />
      </n-form-item>

      <n-form-item
        label="Теги"
        path="tags_ids"
      >
        <AppDynamicTags
          v-model:values="formValue.tags_ids"
        />
      </n-form-item>

      <n-form-item
        :show-feedback="false"
        :show-label="false"
      >
        <n-button
          attr-type="submit"
          type="primary"
          class="flex-1"
        >
          Сохранить
        </n-button>
      </n-form-item>
    </n-form>
  </div>
</template>

<script lang="ts" setup>
import {
  FormInst,
  FormRules,
  NButton,
  NForm,
  NFormItem,
  NInput,
  NRadio,
  NRadioButton,
  NRadioGroup,
  NSelect,
  NCheckbox,
  useMessage,
} from 'naive-ui'
import { computed, onMounted, reactive, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { Course, Module, Option, PageExpose, Question, QuestionType, QuestionTypeTranslates } from '@/types.ts'
import { useForm } from '@/composables/useForm.ts'
import { useCourseStore } from '@/composables/useCourseStore.ts'
import { useModuleStore } from '@/composables/useModuleStore.ts'
import { useQuestionStore } from '@/composables/useQuestionStore.ts'
import AppAnswerInput from '@/components/AppAnswerInput.vue'
import { isAdminMode, me } from '@/composables/useAuthStore.ts'
import AppTextEditor from '@/components/AppTextEditor.vue'
import AppModerationForm from '@/components/AppModerationForm.vue'
import AppDynamicTags from '@/components/AppDynamicTags.vue'

const form = useForm()
const route = useRoute()
const router = useRouter()
const courseStore = useCourseStore()
const moduleStore = useModuleStore()
const questionStore = useQuestionStore()
const message = useMessage()
const question = ref<Question>()

const isCreating = String(route.name).endsWith('create')

defineExpose<PageExpose>({
  title: isCreating ? 'Создание вопроса' : 'Редактирование вопроса',
  back: router.resolve({ name: 'questions' }),
})

const answerCreator = (start: number = 0) => {
  let id = start

  return () => ({
    id: ++id,
    content: '',
    is_correct: false,
  })
}

let createAnswer = answerCreator()

const formRef = ref<FormInst>()
const formValue = reactive({
  course_id: null as number | null,
  module_id: null as number | null,
  title: null as string | null,
  moderation_reason: null as string | null,
  content: null as string | null,
  type: QuestionType.SingleChoice,
  options: [createAnswer(), createAnswer()] as Array<Option>,
  explanation: null as string | null,
  tags_ids: [] as number[],
})
const formRules: FormRules = {
  course_id: {
    required: true,
    type: 'number',
    message: 'Выберите курс',
  },
  type: {
    required: true,
    type: 'string',
    message: 'Выберите количество правильных ответов',
  },
  title: {
    required: true,
    type: 'string',
    message: 'Введите вопрос',
  },
}
const courses = ref<Course[]>([])
const coursesOptions = computed(() => courses.value.map(course => ({
  label: course.name,
  value: course.id,
})))
const modules = ref<Module[]>([])
const modulesOptions = computed(() => {
  return modules.value
    .filter(module => module.course_id === formValue.course_id)
    .map(module => ({
      label: module.name,
      value: module.id,
    }))
})
const typesOptions = Object
  .values(QuestionType)
  .map(type => ({
    label: QuestionTypeTranslates[type],
    value: type,
  }))

const clearForm = () => {
  formValue.title = null
  formValue.content = null
  formValue.explanation = null
  formValue.type = QuestionType.SingleChoice
  formValue.options = [
    createAnswer(),
    createAnswer(),
  ]
  formValue.tags_ids = []
}

const onSubmit = () => {
  const prepareData = {
    course_id: formValue.course_id,
    module_id: formValue.module_id,
    title: formValue.title,
    content: formValue.content,
    type: formValue.type,
    options: formValue.options.filter(answer => answer.content),
    explanation: formValue.explanation,
    tags_ids: formValue.tags_ids,
  }

  form.handle(formRef.value, isCreating, async () => {
    await questionStore.createQuestion(prepareData)

    message.success('Вопрос успешно создан')
    clearForm()
  }, async () => {
    await questionStore.updateQuestion(route.params.uuid as string, prepareData)

    message.success('Новый вопрос создан, этот будет удален после модерации')
  })
}

const onLastUpdated = () => {
  formValue.options.push(createAnswer())
}

const onSecondToLastUpdated = () => {
  formValue.options.pop()
}

const onChange = (id: number) => {
  formValue.options
    .forEach(answer => {
      answer.is_correct = answer.id === id
    })
}

const onChangeCourse = () => {
  formValue.module_id = null
}

const onChangeType = () => {
  formValue
    .options
    .forEach(answer => {
      answer.is_correct = false
    })
}

onMounted(() => {
  if (!isCreating) {
    questionStore
      .getQuestionByUuid(route.params.uuid as string)
      .then(data => {
        question.value = data.data
        formValue.type = data.data.type
        formValue.title = data.data.title
        formValue.content = data.data.content
        formValue.options = data.data.options
        formValue.explanation = data.data.explanation
        formValue.course_id = data.data.course_id
        formValue.module_id = data.data.module_id
        formValue.moderation_reason = data.data.moderation_reason
        formValue.tags_ids = data.data.tags_ids || []

        createAnswer = answerCreator(data.data.options.at(-1)?.id || 1)

        onLastUpdated()
      })
  }

  courseStore
    .getCourses({
      or_created_by: me.value?.id,
    })
    .then(data => {
      courses.value = data.data
    })

  moduleStore
    .getModules({
      or_created_by: me.value?.id,
    })
    .then(data => {
      modules.value = data.data
    })
})
</script>
