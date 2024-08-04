import { createApp } from 'vue'
import '../css/index.css'
import App from '@/app/App.vue'
import { createRouter, createWebHashHistory } from 'vue-router'
import { createRoutes } from '@/app/routes.ts'

const router = createRouter({
  history: createWebHashHistory(),
  routes: createRoutes(),
})

createApp(App)
  .use(router)
  .mount('#app')
