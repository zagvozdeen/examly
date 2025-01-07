import { useKy } from '@/composables/useKy.ts'
import { User } from '@/types.ts'

export const useUserStore = () => {
  const ky = useKy()

  const getUsers = () => {
    return ky
      .get('users')
      .json<{data: User[]}>()
  }

  return {
    getUsers,
  }
}