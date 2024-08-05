import { useKy } from '@/composables/useKy.ts'
import { Course } from '@/types.ts'

export const useCourseStore = () => {
  const ky = useKy()

  const getCourses = () => {
    return ky
      .get('courses')
      .json<{ data: Course[] }>()
  }

  const getMyCourses = () => {
    return ky
      .get('my-courses')
      .json<{ data: Course[] }>()
  }

  const getCourseByUuid = (uuid: string) => {
    return ky
      .get(`courses/${uuid}`)
      .json<{ data: Course }>()
  }

  const createCourse = (json: object) => {
    return ky
      .post('courses', { json })
      .json()
  }

  return {
    getCourses,
    getMyCourses,
    getCourseByUuid,
    createCourse,
  }
}