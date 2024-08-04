import { useKy } from '@/composables/useKy.ts'
import { Course } from '@/types.ts'

export const useCourseStore = () => {
  const ky = useKy()

  const getCourse = (id: number) => {
    return ky.get(`/courses/${id}`).json<Course>()
  }

  return {
    getCourse,
  }
}