import { useKy } from '@/composables/useKy.ts'

export const useUserAnswerStore = () => {
  const ky = useKy()

  const checkAnswer = (json: object) => {
    return ky
      .post('user-answers', { json })
      .json()
  }

  return {
    checkAnswer,
  }
}