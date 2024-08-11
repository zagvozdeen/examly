import { RouteRecordRaw } from 'vue-router'
import AppLayout from '@/components/AppLayout.vue'
import PageCourse from '@/pages/PageCourse.vue'
import PageLogin from '@/pages/PageLogin.vue'
import PageMain from '@/pages/PageMain.vue'
import PageRegister from '@/pages/PageRegister.vue'
import PageMe from '@/pages/PageMe.vue'
import PageMySettings from '@/pages/PageMySettings.vue'
import PageCoursesImport from '@/pages/PageCoursesImport.vue'
import PageMyCourses from '@/pages/PageMyCourses.vue'
import PageMyCoursesEdit from '@/pages/PageMyCoursesEdit.vue'
import PageMyModulesEdit from '@/pages/PageMyModulesEdit.vue'
import PageMyModules from '@/pages/PageMyModules.vue'
import PageMyQuestionsEdit from '@/pages/PageMyQuestionsEdit.vue'
import PageMyQuestions from '@/pages/PageMyQuestions.vue'
import PageStats from '@/pages/PageStats.vue'
import PageCourseMarathon from '@/pages/PageCourseMarathon.vue'
import PageTest from '@/pages/PageTest.vue'

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
          path: '/my/modules',
          component: PageMyModules,
          name: 'my.modules',
        },
        {
          path: '/my/modules/create',
          component: PageMyModulesEdit,
          name: 'my.modules.create',
        },
        {
          path: '/my/modules/:uuid/edit',
          component: PageMyModulesEdit,
          name: 'my.modules.edit',
        },
        {
          path: '/my/questions',
          component: PageMyQuestions,
          name: 'my.questions',
        },
        {
          path: '/my/questions/create',
          component: PageMyQuestionsEdit,
          name: 'my.questions.create',
        },
        {
          path: '/my/questions/:uuid/edit',
          component: PageMyQuestionsEdit,
          name: 'my.questions.edit',
        },
        {
          path: '/my/questions/import',
          component: PageCoursesImport,
          name: 'my.questions.import',
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
          path: '/courses/:uuid/marathon',
          component: PageCourseMarathon,
          name: 'courses.show.marathon',
        },
        {
          path: '/courses/:uuid/stats',
          component: PageStats,
          name: 'courses.stats',
        },
        {
          path: '/stats',
          component: PageStats,
          name: 'stats',
        },
        {
          path: '/tests/:uuid',
          component: PageTest,
          name: 'tests.show',
        },
      ],
    },
  ]
}