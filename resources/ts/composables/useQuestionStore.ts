import { useKy } from '@/composables/useKy.ts'
import { Question } from '@/types.ts'

export const useQuestionStore = () => {
  const ky = useKy()

  const getMyQuestions = () => {
    return ky
      .get('my-questions')
      .json<{data: Question[]}>()
  }

  const createQuestion = (json: object) => {
    return ky
      .post('questions', { json })
      .json<{data: Question}>()
  }

  const getQuestionByUuid = (uuid: string) => {
    return ky
      .get(`questions/${uuid}`)
      .json<{ data: Question }>()
  }

  const importQuestions = (json: object) => {
    return ky
      .post('questions/import', { json })
      .json()
  }

  return {
    getMyQuestions,
    createQuestion,
    getQuestionByUuid,
    importQuestions,
  }
}