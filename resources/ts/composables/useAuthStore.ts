import { useKy } from '@/composables/useKy.ts'
import { User, UserRole } from '@/types.ts'
import { computed, ref } from 'vue'

export const me = ref<User>()
export const isAuthenticated = ref<boolean>(false)
export const isModerator = ref<boolean>(false)

const adminMode = ref<boolean>(localStorage.getItem('moderation') === 'enable')

export const isAdminMode = computed({
  get: () => isModerator.value && adminMode.value,
  set: (value: boolean) => {
    if (!isModerator.value) {
      return
    }

    adminMode.value = value

    if (value) {
      localStorage.setItem('moderation', 'enable')
    } else {
      localStorage.removeItem('moderation')
    }
  },
})

export const useAuthStore = () => {
  const ky = useKy()

  const getMe = async () => {
    const data = await ky
      .get('me')
      .json<{ data: User }>()

    me.value = data.data
    isAuthenticated.value = true
    isModerator.value = data.data.role === UserRole.Moderator || data.data.role === UserRole.Admin
  }

  const updateMe = async (json: object) => {
    const data = await ky
      .patch('me', { json })
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

  const logout = async () => {
    localStorage.removeItem('token')
    me.value = undefined
    isAuthenticated.value = false
    isModerator.value = false
  }

  return {
    getMe,
    updateMe,
    login,
    register,
    logout,
  }
}