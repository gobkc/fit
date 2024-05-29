import {createRouter, createWebHistory} from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/setting',
      name: 'setting',
      // @ts-ignore
      component: () => import('../views/Setting.vue')
    },
    {
      path: '/',
      name: 'note',
      // @ts-ignore
      component: () => import('../views/Note.vue')
    },
  ]
})

export default router
