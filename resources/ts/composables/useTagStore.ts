import { useKy } from '@/composables/useKy.ts'
import { Tag } from '@/types.ts'
import { ref } from 'vue'

export const tags = ref<Tag[]>([])

export const useTagStore = () => {
  const ky = useKy()

  const getAllTags = () => {
    if (tags.value.length > 0) {
      return
    }

    ky.get('tags')
      .json<{ data: Tag[] }>()
      .then(data => {
        tags.value = data.data
      })
  }

  return {
    getAllTags,
  }
}