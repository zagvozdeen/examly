import { useKy } from '@/composables/useKy.ts'
import { UserAnswer } from '@/types.ts'

export const useUserAnswerStore = () => {
  const ky = useKy()

  const checkAnswer = (json: object) => {
    return ky
      .post('user-answers', { json })
      .json<{data: UserAnswer}>()
  }

  return {
    checkAnswer,
  }
}