import { useKy } from '@/composables/useKy.ts'
import {  } from '@/types.ts'

export const useTestSessionStore = () => {
  const ky = useKy()

  const createTestSession = (json: object) => {
    return ky
      .post('test-sessions', { json })
      .json()
  }

  return {
    createTestSession,
  }
}