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
        label="Курс"
        path="module_id"
      >
        <n-select
          v-model:value="formValue.module_id"
          placeholder="Выберите модуль"
          :options="modulesOptions"
        />
      </n-form-item>

      <n-form-item
        label="Вопрос"
        path="content"
      >
        <n-input
          v-model:value="formValue.content"
          placeholder="Введите вопрос"
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
import { NForm, NFormItem, NInput, NSelect, NButton, FormInst, FormRules, useMessage } from 'naive-ui'
import { computed, onMounted, reactive, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { Course, Module, PageExpose } from '@/types.ts'
import { useForm } from '@/composables/useForm.ts'
import { useCourseStore } from '@/composables/useCourseStore.ts'
import { useModuleStore } from '@/composables/useModuleStore.ts'
import { useQuestionStore } from '@/composables/useQuestionStore.ts'

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

const formRef = ref<FormInst>()
const formValue = reactive({
  course_id: null as string | null,
  module_id: null as string | null,
  content: null as string | null,
})
const formRules: FormRules = {
  course_id: {
    required: true,
    type: 'number',
    message: 'Выберите курс',
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
const modulesOptions = computed(() => modules.value.map(module => ({
  label: module.name,
  value: module.id,
})))

const clearForm = () => {
  formValue.course_id = null
  formValue.module_id = null
  formValue.content = null
}

const onSubmit = () => {
  form.handle(formRef.value, isCreating, async () => {
    await questionStore.createQuestion(formValue)

    message.success('Вопрос успешно создан')
    clearForm()
    // await router.push({ name: 'my.modules' })
  }, async () => {
    console.log('NOT IMPLEMENTED')
  })
}


onMounted(() => {
  courseStore
    .getCourses()
    .then(data => {
      courses.value = data.data
    })

  moduleStore
    .getMyModules()
    .then(data => {
      modules.value = data.data
    })
})
</script>
