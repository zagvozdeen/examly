import { useKy } from '@/composables/useKy.ts'
import { User } from '@/types.ts'
import { ref } from 'vue'

export const me = ref<User>()

export const useAuthStore = () => {
  const ky = useKy()

  const getMe = async () => {
    const data = await ky
      .get('me')
      .json<{ data: User }>()

    me.value = data.data
  }

  const login = async (json: object) => {
    return await ky
      .post('auth/login', { json })
      .json<{data: string}>()
  }

  const register = async (json: object) => {
    return await ky
      .post('auth/register', { json })
      .json<{data: number}>()
  }

  return {
    getMe,
    login,
    register,
  }
}