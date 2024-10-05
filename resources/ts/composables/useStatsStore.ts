import { useKy } from '@/composables/useKy.ts'
import { FullCourseStats } from '@/types.ts'

export const useStatsStore = () => {
  const ky = useKy()

  const getStats = () => {
    return ky
      .get('test-sessions/stats')
      .json<{ data: FullCourseStats[] }>()
  }

  const getCourseStats = (uuid: string) => {
    return ky
      .get(`courses/${uuid}/stats`)
      .json<{ data: FullCourseStats[] }>()
  }

  return {
    getStats,
    getCourseStats,
  }
}