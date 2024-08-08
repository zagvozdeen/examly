import { useKy } from '@/composables/useKy.ts'
import { Course, CourseStats, UserCourse, UserQuestion } from '@/types.ts'

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

  const getAllCourses = () => {
    return ky
      .get('all-courses')
      .json<{ data: Course[] }>()
  }

  const getCourseByUuid = (uuid: string) => {
    return ky
      .get(`courses/${uuid}`)
      .json<{
        data: Course
        stats: CourseStats
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

  const checkAnswer = (uuid: string, json: object) => {
    return ky
      .patch(`user-questions/${uuid}`, { json })
      .json<{ data: UserQuestion }>()
  }

  return {
    getCourses,
    getMyCourses,
    getAllCourses,
    getCourseByUuid,
    getUserCourseByUuid,
    createCourse,
    createMarathon,
    checkAnswer,
  }
}