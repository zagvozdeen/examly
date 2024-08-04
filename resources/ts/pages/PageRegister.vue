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
          :input-props="{
            autocomplete: 'given-name',
          }"
        />
      </n-form-item>

      <n-form-item
        label="Фамилия"
        path="last_name"
      >
        <n-input
          v-model:value="formValue.last_name"
          placeholder="Введите фамилию"
          :input-props="{
            autocomplete: 'family-name',
          }"
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
            autocomplete: 'email',
          }"
        />
      </n-form-item>

      <n-form-item
        label="Пароль"
        path="password"
      >
        <n-input
          v-model:value="formValue.password"
          placeholder="Введите пароль"
          :input-props="{
            type: 'password',
            autocomplete: 'new-password',
          }"
        />
      </n-form-item>

      <n-form-item
        label="Повторите пароль"
        path="password_confirmation"
      >
        <n-input
          v-model:value="formValue.password_confirmation"
          placeholder="Введите повторный пароль"
          :input-props="{
            type: 'password',
            autocomplete: 'new-password',
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
          Создать аккаунт
        </n-button>
      </n-form-item>
    </n-form>

    <span class="text-center">Уже есть аккаунт? <router-link
      :to="{ name: 'login' }"
      class="underline"
    >
      Войти в аккаунт
    </router-link></span>
  </div>
</template>
<script lang="ts" setup>
import { NForm, NFormItem, NInput, NButton, FormInst, FormRules, useMessage } from 'naive-ui'
import { reactive, ref } from 'vue'
import { useAuthStore } from '@/composables/useAuthStore.ts'
import { HTTPError } from 'ky'
import { useRouter } from 'vue-router'
import { PageExpose } from '@/types.ts'

defineExpose<PageExpose>({
  title: 'Регистрация',
})

const router = useRouter()
const message = useMessage()
const authStore = useAuthStore()

const formRef = ref<FormInst>()
const formValue = reactive({
  first_name: null as string | null,
  last_name: null as string | null,
  email: null as string | null,
  password: null as string | null,
  password_confirmation: null as string | null,
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
  password: {
    required: true,
    message: 'Введите пароль',
  },
  password_confirmation: [{
    required: true,
    message: 'Введите повторный пароль',
  }, {
    validator: (_, value) => value === formValue.password,
    message: 'Пароли не совпадают',
  }],
}

const clearForm = () => {
  formValue.first_name = null
  formValue.last_name = null
  formValue.email = null
  formValue.password = null
  formValue.password_confirmation = null
}

const onSubmit = async () => {
  try {
    await formRef.value?.validate()

    await authStore.register(formValue)

    clearForm()

    message.success('Вы успешно зарегистрировались, теперь вы можете войти в свой аккаунт')

    await router.push({ name: 'login' })
  } catch (e) {
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
</script>