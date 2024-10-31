<template>
  <n-form
    ref="formRef"
    class="relative border-2 rounded-xl border-red-500 border-dashed p-4"
    :rules="formRules"
    :model="formValue"
    @submit.prevent="onSubmit"
  >
    <div class="text-red-500 select-none uppercase text-xs bg-obscure-800 font-bold absolute -top-[10px] px-1 left-1/2 -translate-x-1/2">
      Раздел администратора
    </div>
    <n-form-item
      label="Статус"
      path="status"
    >
      <n-select
        v-model:value="formValue.status"
        placeholder="Выберите статус"
        :options="statuses"
      />
    </n-form-item>

    <n-form-item
      v-show="formValue.status === Status.Inactive"
      label="Причина отказа"
      path="moderation_reason"
    >
      <n-input
        v-model:value="formValue.moderation_reason"
        placeholder="Введите причину отказа"
        type="textarea"
      />
    </n-form-item>

    <n-form-item
      :show-feedback="false"
      :show-label="false"
    >
      <n-button
        attr-type="submit"
        type="error"
        class="flex-1"
        size="small"
      >
        Сохранить
      </n-button>
    </n-form-item>
  </n-form>
</template>

<script lang="ts" setup>
import { FormInst, FormRules, NButton, NForm, NFormItem, NInput, NSelect, useLoadingBar, useMessage } from 'naive-ui'
import { reactive, ref } from 'vue'
import { Status, StatusTranslates } from '@/types.ts'
import { useRoute } from 'vue-router'

const props = defineProps<{
  status: Status
  reason: string | null
  callback: (uuid: string, json: object) => Promise<unknown>
}>()


const message = useMessage()
const loadingBar = useLoadingBar()
const route = useRoute()

const formRef = ref<FormInst>()
const formValue = reactive({
  status: props.status as string,
  moderation_reason: props.reason as string | null,
})
const formRules: FormRules = {
  status: {
    required: true,
    type: 'string',
    message: 'Выберите статус',
  },
}

const statuses = Object.values(Status).map(status => ({
  label: StatusTranslates[status],
  value: status,
}))

const onSubmit = async () => {
  try {
    await formRef.value?.validate()

    loadingBar.start()

    await props.callback(route.params.uuid as string, formValue)

    message.success(`Статус успешно изменён на «${StatusTranslates[formValue.status as Status]}»`)

    loadingBar.finish()
  } catch (e) {
    console.error(e)
    loadingBar.error()
  }
}
</script>