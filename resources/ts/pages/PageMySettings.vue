<template>
  <div class="flex flex-col gap-4">
    <n-form
      ref="formRef"
      :rules="formRules"
      :model="formValue"
      @submit.prevent="onSubmit"
    >
      <n-form-item
        label="Роль"
        path="role"
      >
        <n-radio-group
          v-model:value="formValue.role"
          class="w-full sm:!flex !hidden"
          size="small"
          :disabled="moreThanMember"
        >
          <n-radio-button
            v-for="type in rolesOptions"
            :key="type.value"
            type="primary"
            :value="type.value"
            :label="type.label"
            class="flex-1 text-center"
          />
        </n-radio-group>
        <n-select
          v-model:value="formValue.role"
          class="sm:hidden block"
          :options="rolesOptions"
          :disabled="moreThanMember"
        />
      </n-form-item>

      <n-form-item
        v-if="formValue.role === UserRole.Referral"
        :show-feedback="false"
        :show-label="false"
      >
        <div class="flex flex-col items-center gap-1.5 rounded-md bg-obscure-500 bg-opacity-50 p-3 mb-4">
          <div class="bg-blue-400 rounded w-9 py-0.5 text-center">
            <i class="bi bi-info-lg text-2xl" />
          </div>
          <span class="text-lg">Реферальная программа</span>
          <span class="text-xs text-gray-100 text-center">
            Вы можете стать внести свой вклад в сообщество и оставить контакты своего работодателя, если у вас есть открытые вакансии.
            Пользователи, которые ищут работу, увидят название компании, краткое описание и контакт, по которому смогут связаться с вами или напрямую с HR.
          </span>
        </div>
      </n-form-item>

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

      <template v-if="[UserRole.Referral, UserRole.Company].includes(formValue.role)">
        <n-form-item
          label="Название компании"
          path="company_name"
        >
          <n-input
            v-model:value="formValue.company_name"
            placeholder="Введите название компании"
          />
        </n-form-item>

        <n-form-item
          label="Небольшое описание"
          path="description"
        >
          <n-input
            v-model:value="formValue.description"
            placeholder="Введите описание компании"
            type="textarea"
          />
        </n-form-item>

        <n-form-item
          label="Контакт для связи"
          path="contact"
        >
          <n-input
            v-model:value="formValue.contact"
            placeholder="Введите контакт для связи"
          />
        </n-form-item>
      </template>

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
  NRadioButton,
  NRadioGroup,
  NSelect,
  useLoadingBar,
  useMessage,
} from 'naive-ui'
import { onMounted, reactive, ref, watch } from 'vue'
import { me, useAuthStore } from '@/composables/useAuthStore.ts'
import { HTTPError } from 'ky'
import { useRouter } from 'vue-router'
import { PageExpose, UserRole, UserRoleTranslates } from '@/types.ts'

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
  role: UserRole.Member as string | null,
  first_name: null as string | null,
  last_name: null as string | null,
  email: null as string | null,
  description: null as string | null,
  company_name: null as string | null,
  contact: null as string | null,
})
const formRules: FormRules = {
  role: {
    required: true,
    type: 'string',
    message: 'Выберите вашу роль',
  },
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

const rolesOptions = [
  { label: UserRoleTranslates[UserRole.Member], value: UserRole.Member },
  { label: UserRoleTranslates[UserRole.Referral], value: UserRole.Referral },
  { label: UserRoleTranslates[UserRole.Company], value: UserRole.Company },
]
const moreThanMember = [UserRole.Moderator, UserRole.Admin].includes(me.value?.role || UserRole.Guest)
const defaultRole = moreThanMember ? (me.value?.role || UserRole.Member) : UserRole.Member
if (moreThanMember) {
  rolesOptions.push({ label: UserRoleTranslates[defaultRole], value: defaultRole })
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
      message.error(`При выполнении запроса произошла ошибка: ${e.response}`)
      console.log(e.message, e.response.body)
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
  formValue.role = me.value?.role || defaultRole
  formValue.first_name = me.value?.first_name || null
  formValue.last_name = me.value?.last_name || null
  formValue.email = me.value?.email || null
  formValue.description = me.value?.description || null
  formValue.company_name = me.value?.company_name || null
  formValue.contact = me.value?.contact || null
}

watch(me, setForm)

onMounted(() => {
  setForm()
})
</script>
