import {createRouter, createWebHistory} from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'setting',
      // @ts-ignore
      component: () => import('../views/Setting.vue')
    },
    {
      path: '/console',
      name: 'Console',
      // @ts-ignore
      component: () => import('../views/Console.vue'),
      redirect: '/console/user/user-admin',
      children: [
        {
          path: '/console/user',
          name: '用户管理',
          // @ts-ignore
          redirect: '/console/user/user-admin',
          meta:{icon: "i-ep-user-filled"},
          children: [
            {
              path: '/console/user/user-admin',
              name: '用户管理',
              // @ts-ignore
              component: () => import('../views/UserAdmin.vue'),
            },
          ]
        },
        {
          path: '/console/exit',
          name: 'Login out',
          meta:{icon: "i-ep-arrow-left-bold"},
          // @ts-ignore
          component: () => import('../views/Exit.vue'),
        },
      ]
    }
  ]
})

export default router
