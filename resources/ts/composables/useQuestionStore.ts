import { useKy } from '@/composables/useKy.ts'
import { Question } from '@/types.ts'

interface GetParams {
  created_by?: number
  all?: boolean
}

export const useQuestionStore = () => {
  const ky = useKy()

  const getQuestions = (params: GetParams) => {
    return ky
      .get('questions', { searchParams: params as Record<string, string> })
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

  const moderateQuestion = (uuid: string, json: object) => {
    return ky
      .patch(`questions/${uuid}/moderate`, { json })
      .json()
  }

  const updateQuestion = (uuid: string, json: object) => {
    return ky
      .patch(`questions/${uuid}`, { json })
      .json()
  }

  const importQuestions = (json: object) => {
    return ky
      .post('questions/import', { json })
      .json()
  }

  return {
    getQuestions,
    createQuestion,
    getQuestionByUuid,
    moderateQuestion,
    updateQuestion,
    importQuestions,
  }
}