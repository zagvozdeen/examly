import { useKy } from '@/composables/useKy.ts'
import { User, UserExperience } from '@/types.ts'

export const useUserStore = () => {
  const ky = useKy()

  const getUsers = () => {
    return ky
      .get('users')
      .json<{data: User[]}>()
  }

  const getUserExperience = () => {
    return ky
      .get('users/experience')
      .json<{data: UserExperience | null}>()
  }

  const createUserExperience = (json: object) => {
    return ky
      .post('users/experience', { json })
      .json<{data: UserExperience}>()
  }

  return {
    getUsers,
    getUserExperience,
    createUserExperience,
  }
}