import { HTTPError } from 'ky'
import { FormInst, useLoadingBar, useMessage } from 'naive-ui'

export const useForm = () => {
  const message = useMessage()
  const loadingBar = useLoadingBar()
  
  const handle = async (formRef: FormInst | undefined, isCreating: boolean, creatingCallback: () => Promise<void>, editingCallback: () => Promise<void>) => {
    return submit(formRef, isCreating ? creatingCallback : editingCallback)
  }

  const submit = async (formRef: FormInst | undefined, callback: () => Promise<void>) => {
    try {
      await formRef?.validate()

      loadingBar.start()

      await callback()

      loadingBar.finish()
    } catch (e) {
      if (e instanceof Array) {
        message.error('Пожалуйста, исправьте ошибки в форме')
        return
      }

      loadingBar.error()

      if (e instanceof HTTPError) {
        message.error('При выполнении запроса произошла ошибка')
        return
      }

      message.error('При выполнении запроса произошла неизвестная ошибка')
      console.error(e)
    }
  }

  return {
    handle,
    submit,
  }
}