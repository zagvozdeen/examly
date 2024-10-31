<template>
  <div class="flex flex-col gap-2 w-full">
    <div class="flex gap-2">
      <label
        for="text-editor-image-uploader"
        class="bg-obscure-500 flex items-center justify-center cursor-pointer hover:bg-obscure-500 bg-opacity-50 rounded w-8 h-8"
        type="button"
      >
        <i class="bi bi-image" />
      </label>
    </div>
    <input
      id="text-editor-image-uploader"
      class="hidden"
      type="file"
      accept="image/*"
      @change="onChangeImageUploader"
    >
    <n-input
      v-model:value="modelValue"
      type="textarea"
      placeholder="Введите описание"
    />
  </div>
</template>

<script lang="ts" setup>
import { NInput, useMessage } from 'naive-ui'
import { useFileStore } from '@/composables/useFileStore.ts'

const fileStore = useFileStore()
const message = useMessage()
const modelValue = defineModel<string | null, string>('modelValue', { default: '' })

const onChangeImageUploader = async (event: Event) => {
  const target = event.target
  if (!(target instanceof HTMLInputElement)) {
    message.error('При загрузке изображения что-то пошло не так')
    return
  }
  const file = target.files?.[0]

  if (!(file instanceof File)) {
    message.error('При загрузке изображения что-то пошло не так')
    return
  }

  const form = new FormData()
  form.append('file', file)

  try {
    const response = await fileStore.uploadFile(form)

    target.value = ''

    modelValue.value += `![${response.data.origin_name}](${response.data.content})`
  } catch (e) {
    message.error('При отправке файла на сервер произошла ошибка')
    console.error(e)
  }
}
</script>