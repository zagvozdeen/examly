import AppLayout from '@/components/AppLayout.vue'
import PageCourse from '@/pages/PageCourse.vue'
import PageLogin from '@/pages/PageLogin.vue'
import PageMain from '@/pages/PageMain.vue'
import PageRegister from '@/pages/PageRegister.vue'
import { RouteRecordRaw } from 'vue-router'
import PageMe from '@/pages/PageMe.vue'
import PageMeSettings from '@/pages/PageMeSettings.vue'
import PageCoursesImport from '@/pages/PageCoursesImport.vue'

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
          path: '/me',
          component: PageMe,
          name: 'me',
        },
        {
          path: '/me/settings',
          component: PageMeSettings,
          name: 'me.settings',
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
        {
          path: '/courses/import',
          component: PageCoursesImport,
          name: 'courses.import',
        },
      ],
    },
  ]
}