import { useKy } from '@/composables/useKy.ts'
import { Course } from '@/types.ts'

export const useCourseStore = () => {
  const ky = useKy()

  const getCourses = () => {
    return ky
      .get('courses')
      .json<{ data: Course[] }>()
  }

  const getCourseByUuid = (uuid: string) => {
    return ky
      .get(`courses/${uuid}`)
      .json<{ data: Course }>()
  }

  return {
    getCourses,
    getCourseByUuid,
  }
}