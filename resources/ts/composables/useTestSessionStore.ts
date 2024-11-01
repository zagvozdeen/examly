import { useKy } from '@/composables/useKy.ts'
import { TestSession } from '@/types.ts'

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
      .json<{data: TestSession}>()
  }

  const createTestSession = (json: object) => {
    return ky
      .post('test-sessions', { json })
      .json()
  }

  return {
    getTestSessions,
    getTestSession,
    createTestSession,
  }
}