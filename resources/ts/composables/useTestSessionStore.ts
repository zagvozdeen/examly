import { useKy } from '@/composables/useKy.ts'
import { TestSession, UserAnswer } from '@/types.ts'

interface GetParams {
  course_uuid?: string
}

export const useTestSessionStore = () => {
  const ky = useKy()

  const getTestSessions = (params: GetParams) => {
    return ky
      .get('test-sessions', { searchParams: params as Record<string, string> })
      .json<{data: TestSession[]}>()
  }

  const getTestSession = (uuid: string) => {
    return ky
      .get(`test-sessions/${uuid}`)
      .json<{
        data: TestSession
        answers: Record<number, UserAnswer>
      }>()
  }

  const createTestSession = (json: object) => {
    return ky
      .post('test-sessions', { json })
      .json<{data: TestSession}>()
  }

  return {
    getTestSessions,
    getTestSession,
    createTestSession,
  }
}