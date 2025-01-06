<template>
  <div class="flex flex-col gap-4">
    <div
      class="flex flex-col items-center gap-1.5 rounded-md bg-obscure-500 bg-opacity-50 p-3"
    >
      <div class="bg-blue-400 rounded w-9 py-0.5 text-center">
        <i class="bi bi-info-lg text-2xl" />
      </div>
      <span class="text-lg">Подбор вопросов для подготовки</span>
      <span class="text-xs text-gray-100 text-center">
        Подбор вопросов для подготовки — это новая система подготовки к собеседованию.
        Ниже вы выбираете направления, к которым хотите подготовиться, и система предложит вам вопросы.
        Эти вопросы будут взяты из всех внутренних источников, объеденины в один тест и перемешаны.
      </span>
    </div>

    <n-form
      ref="formRef"
      :rules="formRules"
      :model="formValue"
      @submit.prevent="onSubmit"
    >
      <n-form-item
        label="Направления"
        path="tags_ids"
      >
        <AppDynamicTags
          v-model:values="formValue.tags_ids"
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
          Начать подготовку
        </n-button>
      </n-form-item>
    </n-form>
  </div>
</template>

<script lang="ts" setup>
import { FormInst, FormRules, NButton, NForm, NFormItem, useMessage } from 'naive-ui'
import { reactive, ref } from 'vue'
import { useRouter } from 'vue-router'
import { PageExpose, TestSessionType } from '@/types.ts'
import { useForm } from '@/composables/useForm.ts'
import AppDynamicTags from '@/components/AppDynamicTags.vue'
import { useTestSessionStore } from '@/composables/useTestSessionStore.ts'

const form = useForm()
const router = useRouter()
const message = useMessage()
const testSessionStore = useTestSessionStore()

defineExpose<PageExpose>({
  title: 'Система подбора',
  back: router.resolve({ name: 'me' }),
})

const formRef = ref<FormInst>()
const formValue = reactive({
  tags_ids: [] as number[],
})
const formRules: FormRules = {
  tags_ids: {
    required: true,
    type: 'array',
    message: 'Выберите хотя бы одно направление',
  },
}

const onSubmit = () => {
  form.submit(formRef.value, async () => {
    const payload = {
      type: TestSessionType.SelectionSystem,
      tags_ids: formValue.tags_ids,
      shuffle: true,
    }

    testSessionStore
      .createTestSession(payload)
      .then(data => {
        router.push({
          name: 'tests.show',
          params: { uuid: data.data.uuid },
        })

        message.success('Мы создали для вас уникальный тест')
      })

  })
}
</script>
