import AppLayout from '@/components/AppLayout.vue'
import PageCourse from '@/pages/PageCourse.vue'
import PageLogin from '@/pages/PageLogin.vue'
import PageMain from '@/pages/PageMain.vue'
import PageRegister from '@/pages/PageRegister.vue'
import { RouteRecordRaw } from 'vue-router'
import PageMe from '@/pages/PageMe.vue'
import PageMySettings from '@/pages/PageMySettings.vue'
import PageCoursesImport from '@/pages/PageCoursesImport.vue'
import PageMyCourses from '@/pages/PageMyCourses.vue'
import PageMyCoursesEdit from '@/pages/PageMyCoursesEdit.vue'

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
          path: '/my/settings',
          component: PageMySettings,
          name: 'my.settings',
        },
        {
          path: '/my/courses',
          component: PageMyCourses,
          name: 'my.courses',
        },
        {
          path: '/my/courses/create',
          component: PageMyCoursesEdit,
          name: 'my.courses.create',
        },
        {
          path: '/my/courses/:uuid/edit',
          component: PageMyCoursesEdit,
          name: 'my.courses.edit',
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