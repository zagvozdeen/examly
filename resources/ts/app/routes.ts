import AppLayout from '@/components/AppLayout.vue'
import PageCourse from '@/pages/PageCourse.vue'
import PageLogin from '@/pages/PageLogin.vue'
import PageMain from '@/pages/PageMain.vue'
import PageRegister from '@/pages/PageRegister.vue'
import { RouteRecordRaw } from 'vue-router'

export const createRoutes = (): Array<RouteRecordRaw> => {
  return [
    {
      path: '/',
      component: AppLayout,
      children: [
        {
          path: '/',
          component: PageMain,
          name: 'main',
        },
        {
          path: '/login',
          component: PageLogin,
          name: 'login',
        },
        {
          path: '/register',
          component: PageRegister,
          name: 'register',
        },
        {
          path: '/courses/:uuid',
          component: PageCourse,
          name: 'courses.show',
        },
      ],
    },
  ]
}