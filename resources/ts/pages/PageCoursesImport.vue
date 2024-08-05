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
          :options="options"
          placeholder="Выберите курс"
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
          Импортировать
        </n-button>
      </n-form-item>
    </n-form>

    <!--    <span class="text-center">Ещё нет аккаунта? <router-link-->
    <!--      :to="{ name: 'register' }"-->
    <!--      class="underline"-->
    <!--    >-->
    <!--      Зарегистрироваться-->
    <!--    </router-link></span>-->
  </div>
</template>
<script lang="ts" setup>
import { NForm, NFormItem, NSelect, NButton, FormInst, FormRules, useMessage } from 'naive-ui'
import { computed, onMounted, reactive, ref } from 'vue'
import { HTTPError } from 'ky'
import { useRouter } from 'vue-router'
import { Course, PageExpose } from '@/types.ts'
import { useCourseStore } from '@/composables/useCourseStore.ts'
import { useQuestionStore } from '@/composables/useQuestionStore.ts'

const router = useRouter()
const message = useMessage()
const courseStore = useCourseStore()
const questionStore = useQuestionStore()

defineExpose<PageExpose>({
  title: 'Импортировать вопросы',
  back: router.resolve({ name: 'me' }),
})

const formRef = ref<FormInst>()
const formValue = reactive({
  // email: null as string | null,
  // password: null as string | null,
  course_id: null as number | null,
})
const formRules: FormRules = {
  course_id: {
    required: true,
    type: 'number',
    message: 'Выберите курс',
  },
  // password: {
  //   required: true,
  //   message: 'Введите пароль',
  // },
}
const courses = ref<Course[]>([])
const options = computed(() => courses.value.map(course => ({
  label: course.name,
  value: course.id,
})))

const clearForm = () => {
  formValue.course_id = null
}

const onSubmit = async () => {
  try {
    await formRef.value?.validate()

    await questionStore.importQuestions(formValue)

    clearForm()

    message.success('Импортирование вопросов прошло успешно')
  } catch (e) {
    if (e instanceof HTTPError) {
      message.error('При выполнении запроса прозошла ошибка')
      return
    }

    if (e instanceof Array) {
      message.error('Пожалуйста, исправьте ошибки в форме')
      return
    }

    message.error('При выполнении запроса произошла неизвестная ошибка')
    console.error(e)
  }
}

onMounted(() => {
  courseStore
    .getCourses()
    .then(data => {
      courses.value = data.data
    })
})
</script>