import { useKy } from '@/composables/useKy.ts'

export const useQuestionStore = () => {
  const ky = useKy()

  const importQuestions = (json: object) => {
    return ky
      .post('questions/import', { json })
      .json()
  }

  return {
    importQuestions,
  }
}