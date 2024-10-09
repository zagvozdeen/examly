import { useKy } from '@/composables/useKy.ts'
import { Module } from '@/types.ts'

interface GetParams {
  created_by?: number
  or_created_by?: number
}

export const useModuleStore = () => {
  const ky = useKy()

  const getModules = (params: GetParams) => {
    return ky
      .get('modules', { searchParams: params as Record<string, string> })
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
    getModules,
    createModule,
    getModuleByUuid,
  }
}