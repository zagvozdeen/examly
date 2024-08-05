import { useKy } from '@/composables/useKy.ts'
import { FileModel } from '@/types.ts'

export const useFileStore = () => {
  const ky = useKy()

  const uploadFile = (body: FormData) => {
    return ky
      .post('files', { body })
      .json<{ data: FileModel }>()
  }

  return {
    uploadFile,
  }
}