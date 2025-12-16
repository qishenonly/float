import { createRouter, createWebHistory } from 'vue-router'
import AdminLayout from '../layout/AdminLayout.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/login',
      name: 'login',
      component: () => import('../views/LoginView.vue')
    },
    {
      path: '/',
      component: AdminLayout,
      children: [
        {
          path: '',
          name: 'dashboard',
          component: () => import('../views/DashboardView.vue'),
          meta: { title: '仪表盘 Dashboard' }
        },
        {
          path: 'users',
          name: 'users',
          component: () => import('../views/UsersView.vue'), // Placeholder
          meta: { title: '用户管理 Users' }
        },
        {
          path: 'transactions',
          name: 'transactions',
          component: () => import('../views/TransactionsView.vue'), // Placeholder
          meta: { title: '交易记录 Data' }
        },
        {
          path: 'versions',
          name: 'versions',
          component: () => import('../views/AppVersionsView.vue'),
          meta: { title: '版本发布 Versions' }
        },
        {
            path: 'settings',
            name: 'settings',
            component: () => import('../views/SettingsView.vue'), // Placeholder
            meta: {title: '系统设置 Settings'}
        }
      ]
    }
  ]
})

router.beforeEach((to, from, next) => {
  const token = localStorage.getItem('token')
  if (to.name !== 'login' && !token) {
    next({ name: 'login' })
  } else if (to.name === 'login' && token) {
    next({ name: 'dashboard' })
  } else {
    next()
  }
})

export default router
