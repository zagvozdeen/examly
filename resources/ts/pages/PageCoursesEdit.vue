<template>
  <div class="flex flex-col gap-4">
    <AppModerationForm
      v-if="isAdminMode && course"
      :status="course.status"
      :reason="course.moderation_reason"
      :callback="courseStore.moderateCourse"
    />

    <n-form
      ref="formRef"
      :rules="formRules"
      :model="formValue"
      @submit.prevent="onSubmit"
    >
      <n-form-item
        label="Название"
        path="name"
      >
        <n-input
          v-model:value="formValue.name"
          placeholder="Введите название"
        />
      </n-form-item>

      <n-form-item
        label="Описание"
        path="description"
      >
        <n-input
          v-model:value="formValue.description"
          placeholder="Введите короткое описание курса"
          type="textarea"
        />
      </n-form-item>

      <n-form-item
        label="Цвет курса"
        path="color"
      >
        <AppColorSelector v-model:value="formValue.color" />
      </n-form-item>

      <n-form-item
        label="Логотип курса"
        path="icon"
      >
        <AppIconSelector v-model:value="formValue.icon" />
      </n-form-item>

      <n-form-item label="Предпросмотр карточки">
        <ul class="w-full flex flex-col bg-obscure-700 rounded-md overflow-hidden">
          <li>
            <div
              class="grid grid-cols-[28px_1fr_min-content] items-center gap-2 hover:bg-obscure-500 bg-opacity-50 p-2"
            >
              <div
                class="rounded w-full py-0.5 text-center"
                :class="{[formValue.color || 'bg-orange-400']: true}"
              >
                <i
                  class="bi"
                  :class="{[formValue.icon || 'bi-1-circle-fill']: true}"
                />
              </div>
              <div class="flex flex-col">
                <span>{{ formValue.name || 'Здесь будет ваше название' }}</span>
                <span
                  v-show="formValue.description"
                  class="text-xs text-gray-400"
                >{{ formValue.description }}</span>
              </div>
              <i class="bi bi-chevron-right" />
            </div>
          </li>
        </ul>
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
import { NForm, NFormItem, NInput, NButton, FormInst, FormRules, useMessage } from 'naive-ui'
import { onMounted, reactive, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { PageExpose, Course } from '@/types.ts'
import { useForm } from '@/composables/useForm.ts'
import { useCourseStore } from '@/composables/useCourseStore.ts'
import AppColorSelector from '@/components/AppColorSelector.vue'
import AppIconSelector from '@/components/AppIconSelector.vue'
import { isAdminMode } from '@/composables/useAuthStore.ts'
import AppModerationForm from '@/components/AppModerationForm.vue'

const form = useForm()
const route = useRoute()
const router = useRouter()
const courseStore = useCourseStore()
const message = useMessage()
const course = ref<Course>()

const isCreating = String(route.name).endsWith('create')

defineExpose<PageExpose>({
  title: isCreating ? 'Создание курса' : 'Редактирование курса',
  back: router.resolve({ name: 'courses' }),
})

const formRef = ref<FormInst>()
const formValue = reactive({
  name: null as string | null,
  description: null as string | null,
  color: null as string | null,
  icon: null as string | null,
})
const formRules: FormRules = {
  name: {
    required: true,
    type: 'string',
    message: 'Введите название курса',
  },
  description: {
    required: true,
    type: 'string',
    message: 'Введите название курса',
  },
  color: {
    required: true,
    type: 'string',
    message: 'Выберите цвет курса',
  },
  icon: {
    required: true,
    type: 'string',
    message: 'Выберите логотип курса',
  },
}

const clearForm = () => {
  formValue.name = null
  formValue.color = null
  formValue.icon = null
}

const onSubmit = () => {
  form.handle(formRef.value, isCreating, async () => {
    await courseStore.createCourse(formValue)

    message.success('Курс успешно создан')
    clearForm()
    await router.push({ name: 'courses' })
  }, async () => {
    console.log('NOT IMPLEMENTED')
  })
}

onMounted(() => {
  if (!isCreating) {
    courseStore
      .getCourseByUuid(route.params.uuid as string)
      .then(data => {
        course.value = data.data
        formValue.name = data.data.name
        formValue.description = data.data.description
        formValue.color = data.data.color
        formValue.icon = data.data.icon
      })
  }
})
</script>
