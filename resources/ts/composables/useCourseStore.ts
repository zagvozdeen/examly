import { useKy } from '@/composables/useKy.ts'
import { Course } from '@/types.ts'

export const useCourseStore = () => {
  const ky = useKy()

  const getCourses = () => {
    return ky
      .get('courses')
      .json<{ data: Course[] }>()
  }

  const getCourse = (id: number) => {
    return ky
      .get(`/courses/${id}`)
      .json<{ data: Course }>()
  }

  return {
    getCourses,
    getCourse,
  }
}