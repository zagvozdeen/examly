<template>
  <n-upload
    v-model:file-list="fileList"
    list-type="image-card"
    :max="1"
    @before-upload="onBeforeUpload"
    @remove="onRemove"
  >
    <slot>
      Выберите файл
    </slot>
  </n-upload>
</template>

<script lang="ts" setup>
import { NUpload, UploadFileInfo, useMessage } from 'naive-ui'
import { ref, watch } from 'vue'
import { useFileStore } from '@/composables/useFileStore.ts'

const fileStore = useFileStore()
const message = useMessage()

const value = defineModel<number | null>('value')

const fileList = ref<UploadFileInfo[]>([])

const onBeforeUpload = async (data: { file: UploadFileInfo }) => {
  if (!(data.file.file instanceof File))  {
    return false
  }

  try {
    const formData = new FormData()
    formData.set('file', data.file.file)

    const response = await fileStore.uploadFile(formData)

    fileList.value.push({
      id: response.data.uuid,
      name: response.data.origin_name,
      status: 'finished',
      url: response.data.content,
      file: data.file.file,
      type: response.data.mime_type,
    })

    value.value = response.data.id
  } catch (e) {
    message.error('При загрузке файла произошла ошибка')
    console.error(e)
  }

  return false
}

const onRemove = () => {
  value.value = null

  return false
}

watch(value, (newValue) => {
  if (!newValue) {
    fileList.value = []
  }
})
</script>
