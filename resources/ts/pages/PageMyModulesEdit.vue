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
          :options="options"
        />
      </n-form-item>

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
import { Course, PageExpose } from '@/types.ts'
import { useForm } from '@/composables/useForm.ts'
import { useCourseStore } from '@/composables/useCourseStore.ts'
import { useModuleStore } from '@/composables/useModuleStore.ts'

const form = useForm()
const route = useRoute()
const router = useRouter()
const courseStore = useCourseStore()
const moduleStore = useModuleStore()
const message = useMessage()

const isCreating = String(route.name).endsWith('create')

defineExpose<PageExpose>({
  title: isCreating ? 'Создание модуля' : 'Редактирование модуля',
  back: router.resolve({ name: 'my.modules' }),
})

const formRef = ref<FormInst>()
const formValue = reactive({
  course_id: null as string | null,
  name: null as string | null,
})
const formRules: FormRules = {
  course_id: {
    required: true,
    type: 'number',
    message: 'Выберите курс',
  },
  name: {
    required: true,
    type: 'string',
    message: 'Введите название модуля',
  },
}
const courses = ref<Course[]>([])
const options = computed(() => courses.value.map(course => ({
  label: course.name,
  value: course.id,
})))

const clearForm = () => {
  formValue.course_id = null
  formValue.name = null
}

const onSubmit = () => {
  form.handle(formRef.value, isCreating, async () => {
    await moduleStore.createModule(formValue)

    message.success('Модуль успешно создан')
    clearForm()
    await router.push({ name: 'my.modules' })
  }, async () => {
    console.log('NOT IMPLEMENTED')
  })
}


onMounted(() => {
  courseStore
    .getAllCourses()
    .then(data => {
      courses.value = data.data
    })
})
</script>
