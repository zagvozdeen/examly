<template>
  <div class="flex flex-col gap-4">
    <div class="flex flex-col items-center gap-1.5 rounded-md bg-obscure-500 bg-opacity-50 p-3">
      <div class="bg-blue-400 rounded w-9 py-0.5 text-center">
        <i class="bi bi-info-lg text-2xl" />
      </div>
      <span class="text-lg">Сбор обратной связи</span>
      <span class="text-xs text-gray-100 text-center">
        Заполнив форму, вы поможете стать нам лучше и сделать наше приложение лучше!
        После полного заполнения формы вы получите 200 баллов на счёт.
      </span>
      <span
        v-if="filled"
        class="text-xs text-green-400 font-semibold"
      >Спасибо за обратную связь!</span>
    </div>

    <n-form
      ref="formRef"
      :rules="formRules"
      :model="form"
      @submit.prevent="onSubmit"
    >
      <n-form-item
        label="Насколько вы удовлетворены общим опытом использования нашего ресурса?"
        path="one"
      >
        <AppRadioGroup
          v-model:value="form.one"
          :options="fiveOptions"
          :disabled="filled"
        />
      </n-form-item>

      <n-form-item
        label="Какова вероятность, что вы порекомендуете наш ресурс друзьям или коллегам?"
        path="two"
      >
        <AppRadioGroup
          v-model:value="form.two"
          :options="tenOptions"
          :disabled="filled"
        />
      </n-form-item>

      <n-form-item
        label="Насколько информация на нашем ресурсе помогла вам подготовиться к собеседованию?"
        path="three"
      >
        <AppRadioGroup
          v-model:value="form.three"
          :options="fiveOptions"
          :disabled="filled"
        />
      </n-form-item>

      <n-form-item
        label="Какие конкретные разделы сайта были для вас наиболее полезными и почему?"
        path="four"
      >
        <n-input
          v-model:value="form.four"
          placeholder="Введите"
          type="textarea"
          :disabled="filled"
        />
      </n-form-item>

      <n-form-item
        label="Насколько наш ресурс помог вам в поиске работы?"
        path="five"
      >
        <AppRadioGroup
          v-model:value="form.five"
          :options="fiveOptions"
          :disabled="filled"
        />
      </n-form-item>

      <n-form-item
        label="Сколько предложений о работе вы получили после использования нашего ресурса?"
        path="six"
      >
        <n-input-number
          v-model:value="form.six"
          class="w-full"
          placeholder="Введите число"
          :min="0"
          :max="999"
          :disabled="filled"
        />
      </n-form-item>

      <n-form-item
        label="Какие функции нашего ресурса, по вашему мнению, нуждаются в улучшениях?"
        path="seven"
      >
        <n-input
          v-model:value="form.seven"
          placeholder="Введите"
          type="textarea"
          :disabled="filled"
        />
      </n-form-item>

      <n-form-item
        label="Испытывали ли вы какие-либо проблемы с навигацией или использованием ресурса? Если да, то какие?"
        path="eight"
      >
        <n-input
          v-model:value="form.eight"
          placeholder="Введите"
          type="textarea"
          :disabled="filled"
        />
      </n-form-item>

      <n-form-item
        label="Как вы оцениваете качество материалов и информации на сайте?"
        path="nine"
      >
        <AppRadioGroup
          v-model:value="form.nine"
          :options="fiveOptions"
          :disabled="filled"
        />
      </n-form-item>

      <n-form-item
        label="Есть ли темы, о которых вы хотели бы видеть больше информации на нашем ресурсе?"
        path="ten"
      >
        <n-input
          v-model:value="form.ten"
          placeholder="Введите"
          type="textarea"
          :disabled="filled"
        />
      </n-form-item>

      <n-form-item
        label="Насколько легко было найти нужную информацию на нашем ресурсе?"
        path="eleven"
      >
        <AppRadioGroup
          v-model:value="form.eleven"
          :options="fiveOptions"
          :disabled="filled"
        />
      </n-form-item>

      <n-form-item
        label="Какие аспекты нашего ресурса вы бы хотели изменить для улучшения пользовательского опыта?"
        path="twelve"
      >
        <n-input
          v-model:value="form.twelve"
          placeholder="Введите"
          type="textarea"
          :disabled="filled"
        />
      </n-form-item>

      <n-form-item
        label="Смогли ли вы получить работу с помощью нашего ресурса? Если да, каков был ваш процесс?"
        path="thirteen"
      >
        <n-input
          v-model:value="form.thirteen"
          placeholder="Введите"
          type="textarea"
          :disabled="filled"
        />
      </n-form-item>

      <n-form-item
        v-if="!filled"
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
  NForm,
  NFormItem,
  NInput,
  NInputNumber,
  NButton,
  FormInst,
  FormRules,
  useMessage,
} from 'naive-ui'
import { onMounted, reactive, ref } from 'vue'
import { useRouter } from 'vue-router'
import { PageExpose } from '@/types.ts'
import { useForm } from '@/composables/useForm.ts'
import AppRadioGroup from '@/components/AppRadioGroup.vue'
import { useUserStore } from '@/composables/useUserStore.ts'

const { submit } = useForm()
const router = useRouter()
const message = useMessage()
const userStore = useUserStore()
const filled = ref(true)

defineExpose<PageExpose>({
  title: 'Опрос пользы приложения',
  back: router.resolve({ name: 'me' }),
})

const formRef = ref<FormInst>()
const form = reactive({
  one: null as number | null,
  two: null as number | null,
  three: null as number | null,
  four: null as string | null,
  five: null as number | null,
  six: null as number | null,
  seven: null as string | null,
  eight: null as string | null,
  nine: null as number | null,
  ten: null as string | null,
  eleven: null as number | null,
  twelve: null as string | null,
  thirteen: null as string | null,
})
const formRules: FormRules = {
  one: {
    required: true,
    type: 'integer',
    message: 'Необходимо сделать выбор',
  },
  two: {
    required: true,
    type: 'integer',
    message: 'Необходимо сделать выбор',
  },
  three: {
    required: true,
    type: 'integer',
    message: 'Необходимо сделать выбор',
  },
  four: {
    required: true,
    type: 'string',
    message: 'Необходимо сделать выбор',
  },
  five: {
    required: true,
    type: 'integer',
    message: 'Необходимо сделать выбор',
  },
  six: {
    required: true,
    type: 'integer',
    message: 'Необходимо сделать выбор',
  },
  seven: {
    required: true,
    type: 'string',
    message: 'Необходимо сделать выбор',
  },
  eight: {
    required: true,
    type: 'string',
    message: 'Необходимо сделать выбор',
  },
  nine: {
    required: true,
    type: 'integer',
    message: 'Необходимо сделать выбор',
  },
  ten: {
    required: true,
    type: 'string',
    message: 'Необходимо сделать выбор',
  },
  eleven: {
    required: true,
    type: 'integer',
    message: 'Необходимо сделать выбор',
  },
  twelve: {
    required: true,
    type: 'string',
    message: 'Необходимо сделать выбор',
  },
  thirteen: {
    required: true,
    type: 'string',
    message: 'Необходимо сделать выбор',
  },
}

const onSubmit = () => {
  submit(formRef.value, async () => {
    return userStore
      .createUserExperience(form)
      .then(() => {
        filled.value = true
        message.success('Спасибо за обратную связь!')
      })
  })
}

const fiveOptions = Array.from({ length: 5 }, (_, i) => ({
  value: i + 1,
  label: String(i + 1),
}))

const tenOptions = Array.from({ length: 10 }, (_, i) => ({
  value: i + 1,
  label: String(i + 1),
}))

onMounted(() => {
  userStore
    .getUserExperience()
    .then(({ data: ue }) => {
      if (ue != null) {
        form.one = ue.one
        form.two = ue.two
        form.three = ue.three
        form.four = ue.four
        form.five = ue.five
        form.six = ue.six
        form.seven = ue.seven
        form.eight = ue.eight
        form.nine = ue.nine
        form.ten = ue.ten
        form.eleven = ue.eleven
        form.twelve = ue.twelve
        form.thirteen = ue.thirteen
      } else {
        filled.value = false
      }
    })
})
</script>
