<template>
  <div class="flex flex-col gap-4">
    <n-form
      ref="formRef"
      :rules="formRules"
      :model="formValue"
      @submit.prevent="onSubmit"
    >
      <n-form-item
        label="Почта"
        path="email"
      >
        <n-input
          v-model:value="formValue.email"
          placeholder="Введите почту"
          :input-props="{type: 'email'}"
        />
      </n-form-item>

      <n-form-item
        label="Пароль"
        path="password"
      >
        <n-input
          v-model:value="formValue.password"
          placeholder="Введите пароль"
          :input-props="{type: 'password'}"
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
          Войти
        </n-button>
      </n-form-item>
    </n-form>

    <span class="text-center">Ещё нет аккаунта? <router-link
      :to="{ name: 'register' }"
      class="underline"
    >
      Зарегистрироваться
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
  title: 'Вход',
})

const router = useRouter()
const message = useMessage()
const authStore = useAuthStore()

const formRef = ref<FormInst>()
const formValue = reactive({
  email: null as string | null,
  password: null as string | null,
})
const formRules: FormRules = {
  email: {
    required: true,
    type: 'email',
    message: 'Введите почту',
  },
  password: {
    required: true,
    message: 'Введите пароль',
  },
}

const clearForm = () => {
  formValue.email = null
  formValue.password = null
}

const onSubmit = async () => {
  try {
    await formRef.value?.validate()

    const data = await authStore.login(formValue)

    localStorage.setItem('token', data.data)

    clearForm()

    message.success('Вы успешно вошли в систему')

    await router.push({ name: 'main' })
  } catch (e) {
    if (e instanceof HTTPError) {
      message.error('Пользователь не найден или введенные данные не верны')
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