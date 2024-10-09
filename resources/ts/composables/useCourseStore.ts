import { useKy } from '@/composables/useKy.ts'
import { Course, CourseStats, UserCourse, UserQuestion } from '@/types.ts'

export const useCourseStore = () => {
  const ky = useKy()

  const getCourses = (params: object) => {
    return ky
      .get('courses', { searchParams: params as Record<string, string> })
      .json<{ data: Course[] }>()
  }

  const getCourseByUuid = (uuid: string) => {
    return ky
      .get(`courses/${uuid}`)
      .json<{
        data: Course
        stats: CourseStats
        errors: Course
      }>()
  }

  const getUserCourseByUuid = (uuid: string) => {
    return ky
      .get(`user-courses/${uuid}`)
      .json<{
        data: UserCourse
      }>()
  }

  const createCourse = (json: object) => {
    return ky
      .post('courses', { json })
      .json()
  }

  const createMarathon = (uuid: string) => {
    return ky
      .post(`courses/${uuid}/marathon`)
      .json<{ data: string }>()
  }

  const createExam = (uuid: string) => {
    return ky
      .post(`courses/${uuid}/exam`)
      .json<{ data: string }>()
  }

  const checkAnswer = (uuid: string, json: object) => {
    return ky
      .patch(`user-questions/${uuid}`, { json })
      .json<{ data: UserQuestion }>()
  }

  const exportCourses = () => {
    return ky
      .post('courses/export')
      .json<{ data: string }>()
  }

  return {
    getCourses,
    getCourseByUuid,
    getUserCourseByUuid,
    createCourse,
    createMarathon,
    createExam,
    checkAnswer,
    exportCourses,
  }
}