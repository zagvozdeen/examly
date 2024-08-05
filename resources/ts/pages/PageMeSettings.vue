<template>
  <div class="flex flex-col gap-4">
    <n-form
      ref="formRef"
      :rules="formRules"
      :model="formValue"
      @submit.prevent="onSubmit"
    >
      <n-form-item
        label="Имя"
        path="first_name"
      >
        <n-input
          v-model:value="formValue.first_name"
          placeholder="Введите имя"
        />
      </n-form-item>

      <n-form-item
        label="Фамилия"
        path="last_name"
      >
        <n-input
          v-model:value="formValue.last_name"
          placeholder="Введите фамилию"
        />
      </n-form-item>

      <n-form-item
        label="Почта"
        path="email"
      >
        <n-input
          v-model:value="formValue.email"
          placeholder="Введите почту"
          :input-props="{
            type: 'email',
          }"
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
import { NForm, NFormItem, NInput, NButton, FormInst, FormRules, useMessage, useLoadingBar } from 'naive-ui'
import { onMounted, reactive, ref, watch } from 'vue'
import { me, useAuthStore } from '@/composables/useAuthStore.ts'
import { HTTPError } from 'ky'
import { useRouter } from 'vue-router'
import { PageExpose } from '@/types.ts'

const router = useRouter()
const message = useMessage()
const authStore = useAuthStore()
const loadingBar = useLoadingBar()

defineExpose<PageExpose>({
  title: 'Настройки',
  back: router.resolve({ name: 'me' }),
})

const formRef = ref<FormInst>()
const formValue = reactive({
  first_name: null as string | null,
  last_name: null as string | null,
  email: null as string | null,
})
const formRules: FormRules = {
  first_name: {
    required: true,
    type: 'string',
    message: 'Введите ваше имя',
  },
  last_name: {
    required: true,
    type: 'string',
    message: 'Введите вашу фамилию',
  },
  email: {
    required: true,
    type: 'email',
    message: 'Введите почту',
  },
}

const onSubmit = async () => {
  try {
    await formRef.value?.validate()

    loadingBar.start()

    await authStore.updateMe(formValue)

    message.success('Данные пользователя успешно обновлены')

    loadingBar.finish()
  } catch (e) {
    loadingBar.error()

    if (e instanceof HTTPError) {
      message.error('При выполнении запроса произошла ошибка')
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

const setForm = () => {
  formValue.first_name = me.value?.first_name || null
  formValue.last_name = me.value?.last_name || null
  formValue.email = me.value?.email || null
}

watch(me, setForm)

onMounted(() => {
  setForm()
})
</script>
