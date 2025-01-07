import { RouteRecordRaw } from 'vue-router'
import AppLayout from '@/components/AppLayout.vue'
import PageCourse from '@/pages/PageCourse.vue'
import PageLogin from '@/pages/PageLogin.vue'
import PageMain from '@/pages/PageMain.vue'
import PageRegister from '@/pages/PageRegister.vue'
import PageMe from '@/pages/PageMe.vue'
import PageMySettings from '@/pages/PageMySettings.vue'
import PageCoursesImport from '@/pages/PageCoursesImport.vue'
import PageMyCourses from '@/pages/PageCourses.vue'
import PageMyCoursesEdit from '@/pages/PageCoursesEdit.vue'
import PageMyModulesEdit from '@/pages/PageModulesEdit.vue'
import PageMyModules from '@/pages/PageModules.vue'
import PageMyQuestionsEdit from '@/pages/PageQuestionsEdit.vue'
import PageMyQuestions from '@/pages/PageQuestions.vue'
import PageStats from '@/pages/PageStats.vue'
// import PageCourseMarathon from '@/pages/PageCourseMarathon.vue'
import PageTest from '@/pages/PageTest.vue'
import PageSelectionSystem from '@/pages/PageSelectionSystem.vue'
import PageRecommendationSystem from '@/pages/PageRecommendationSystem.vue'

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
          path: '/settings',
          component: PageMySettings,
          name: 'settings',
        },
        {
          path: '/courses',
          component: PageMyCourses,
          name: 'courses',
        },
        {
          path: '/courses/create',
          component: PageMyCoursesEdit,
          name: 'courses.create',
        },
        {
          path: '/courses/:uuid/edit',
          component: PageMyCoursesEdit,
          name: 'courses.edit',
        },
        {
          path: '/modules',
          component: PageMyModules,
          name: 'modules',
        },
        {
          path: '/modules/create',
          component: PageMyModulesEdit,
          name: 'modules.create',
        },
        {
          path: '/modules/:uuid/edit',
          component: PageMyModulesEdit,
          name: 'modules.edit',
        },
        {
          path: '/questions',
          component: PageMyQuestions,
          name: 'questions',
        },
        {
          path: '/questions/create',
          component: PageMyQuestionsEdit,
          name: 'questions.create',
        },
        {
          path: '/questions/:uuid/edit',
          component: PageMyQuestionsEdit,
          name: 'questions.edit',
        },
        {
          path: '/questions/import',
          component: PageCoursesImport,
          name: 'questions.import',
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
        // {
        //   path: '/courses/:uuid/marathon',
        //   component: PageCourseMarathon,
        //   name: 'courses.show.marathon',
        // },
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
        {
          path: '/selection-system',
          component: PageSelectionSystem,
          name: 'selection-system',
        },
        {
          path: '/recommendation-system',
          component: PageRecommendationSystem,
          name: 'recommendation-system',
        },
        // {
        //   path: '/admin',
        //   children: [
        //     {
        //       path: 'courses',
        //       component: PageAdminCourses,
        //       name: 'admin.courses',
        //     },
        //   ],
        // },
      ],
    },
  ]
}