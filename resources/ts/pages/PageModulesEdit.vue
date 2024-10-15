<template>
  <div class="flex flex-col gap-4">
    <AppModerationForm
      v-if="isAdminMode && module"
      :status="module.status"
      :reason="module.moderation_reason"
      :callback="moduleStore.moderateModule"
    />

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
import { Course, Module, PageExpose } from '@/types.ts'
import { useForm } from '@/composables/useForm.ts'
import { useCourseStore } from '@/composables/useCourseStore.ts'
import { useModuleStore } from '@/composables/useModuleStore.ts'
import { isAdminMode, me } from '@/composables/useAuthStore.ts'
import AppModerationForm from '@/components/AppModerationForm.vue'

const form = useForm()
const route = useRoute()
const router = useRouter()
const courseStore = useCourseStore()
const moduleStore = useModuleStore()
const message = useMessage()
const module = ref<Module>()

const isCreating = String(route.name).endsWith('create')

defineExpose<PageExpose>({
  title: isCreating ? 'Создание модуля' : 'Редактирование модуля',
  back: router.resolve({ name: 'modules' }),
})

const formRef = ref<FormInst>()
const formValue = reactive({
  course_id: null as number | null,
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
    await router.push({ name: 'modules' })
  }, async () => {
    console.log('NOT IMPLEMENTED')
  })
}


onMounted(() => {
  if (!isCreating) {
    moduleStore
      .getModuleByUuid(route.params.uuid as string)
      .then(data => {
        module.value = data.data
        formValue.name = data.data.name
        formValue.course_id = data.data.course_id
      })
  }

  courseStore
    .getCourses({
      or_created_by: me.value?.id,
    })
    .then(data => {
      courses.value = data.data
    })
})
</script>
