<template>
  <div class="flex flex-col gap-4">
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
        label="Иллюстрация к вопросу"
        path="file_id"
      >
        <AppUploadFile
          v-model:value="formValue.file_id"
          accept="image/*"
        >
          <span class="text-xs">
            Выберите изображение <small>(необязательно)</small>
          </span>
        </AppUploadFile>
      </n-form-item>

      <n-form-item
        label="Вопрос"
        path="content"
      >
        <n-input
          v-model:value="formValue.content"
          placeholder="Задайте вопрос"
          type="textarea"
          :rows="2"
        />
      </n-form-item>

      <n-form-item
        label="Варианты ответа"
        path="content"
      >
        <div class="flex flex-col gap-2 w-full">
          <template
            v-for="(answer, index) in formValue.answers"
            :key="answer.id"
          >
            <div class="flex items-center gap-3">
              <n-checkbox
                v-if="formValue.type === QuestionType.MultiplyAnswersType"
                v-model:checked="answer.is_true"
                tabindex="-1"
              />
              <n-radio
                v-if="formValue.type === QuestionType.OneAnswerType"
                :checked="answer.is_true"
                name="questions-radio-button"
                :value="answer.id"
                tabindex="-1"
                @change="onChange(answer.id)"
              />
              <AppAnswerInput
                v-model:value="answer.content"
                :is-last="index === formValue.answers.length - 1"
                :is-second-to-last="index === formValue.answers.length - 2"
                @last-updated="onLastUpdated"
                @second-to-last-updated="onSecondToLastUpdated"
              />
            </div>
          </template>
        </div>
      </n-form-item>

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
import { computed, onMounted, reactive, ref, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { Course, Module, PageExpose, QuestionType, QuestionTypeTranslates } from '@/types.ts'
import { useForm } from '@/composables/useForm.ts'
import { useCourseStore } from '@/composables/useCourseStore.ts'
import { useModuleStore } from '@/composables/useModuleStore.ts'
import { useQuestionStore } from '@/composables/useQuestionStore.ts'
import AppAnswerInput from '@/components/AppAnswerInput.vue'
import AppUploadFile from '@/components/AppUploadFile.vue'

const form = useForm()
const route = useRoute()
const router = useRouter()
const courseStore = useCourseStore()
const moduleStore = useModuleStore()
const questionStore = useQuestionStore()
const message = useMessage()

const isCreating = String(route.name).endsWith('create')

defineExpose<PageExpose>({
  title: isCreating ? 'Создание вопроса' : 'Редактирование вопроса',
  back: router.resolve({ name: 'my.questions' }),
})

const answerCreator = (start: number = 1) => {
  let id = start

  return () => ({
    id: id++,
    content: null,
    is_true: false,
  })
}

const createAnswer = answerCreator()

const formRef = ref<FormInst>()
const formValue = reactive({
  course_id: null as number | null,
  module_id: null as number | null,
  file_id: null as number | null,
  content: null as string | null,
  type: QuestionType.OneAnswerType,
  answers: [createAnswer(), createAnswer()] as Array<{
    id: number
    content: string | null
    is_true: boolean
  }>,
  explanation: null as string | null,
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
  content: {
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
  formValue.content = null
  formValue.file_id = null
  formValue.explanation = null
  formValue.type = QuestionType.OneAnswerType
  formValue.answers = [
    createAnswer(),
    createAnswer(),
  ]
}

const onSubmit = () => {
  const prepareData = {
    course_id: formValue.course_id,
    module_id: formValue.module_id,
    file_id: formValue.file_id,
    content: formValue.content,
    type: formValue.type,
    answers: formValue.answers.filter(answer => answer.content),
    explanation: formValue.explanation,
  }

  form.handle(formRef.value, isCreating, async () => {
    await questionStore.createQuestion(prepareData)

    message.success('Вопрос успешно создан')
    clearForm()
  }, async () => {
    console.log('NOT IMPLEMENTED')
  })
}

const onLastUpdated = () => {
  formValue.answers.push(createAnswer())
}

const onSecondToLastUpdated = () => {
  formValue.answers.pop()
}

const onChange = (id: number) => {
  formValue.answers
    .forEach(answer => {
      answer.is_true = answer.id === id
    })
}

watch(() => formValue.course_id, () => {
  formValue.module_id = null
})

watch(() => formValue.type, () => {
  formValue.answers
    .forEach(answer => {
      answer.is_true = false
    })
})

onMounted(() => {
  courseStore
    .getAllCourses()
    .then(data => {
      courses.value = data.data
    })

  moduleStore
    .getAllModules()
    .then(data => {
      modules.value = data.data
    })
})
</script>
