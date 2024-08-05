import { useKy } from '@/composables/useKy.ts'
import { Module } from '@/types.ts'

export const useModuleStore = () => {
  const ky = useKy()

  const getMyModules = () => {
    return ky
      .get('my/modules')
      .json<{data: Module[]}>()
  }

  const createModule = (json: object) => {
    return ky
      .post('modules', { json })
      .json<{data: Module}>()
  }

  const getModuleByUuid = (uuid: string) => {
    return ky
      .get(`modules/${uuid}`)
      .json<{ data: Module }>()
  }

  return {
    getMyModules,
    createModule,
    getModuleByUuid,
  }
}