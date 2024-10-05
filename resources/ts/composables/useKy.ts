import ky from 'ky'

export const useKy = () => {
  return ky.extend({
    prefixUrl: '/api/v1',
    hooks: {
      beforeRequest: [
        (request) => {
          request.headers.set('Authorization', `Bearer ${localStorage.getItem('token')}`)
        },
      ],
      afterResponse: [
        async (request, _, response) => {
          if (response.status === 401 || (request.url.endsWith('/api/v1/me') && response.status === 404)) {
            const token = await ky('/api/v1/auth/guest-token').json<{data: string}>()

            localStorage.setItem('token', token.data)
            request.headers.set('Authorization', `Bearer ${token.data}`)

            return ky(request)
          }
        },
      ],
    },
    retry: {
      limit: 0,
    },
  })
}