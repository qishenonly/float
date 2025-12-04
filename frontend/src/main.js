import { createApp } from 'vue'
import { createPinia } from 'pinia'
import App from './App.vue'
import router from './router'
import './style.css'
import { App as CapacitorApp } from '@capacitor/app'

const app = createApp(App)

app.use(createPinia())
app.use(router)

app.mount('#app')

// 配置硬件返回键监听器
router.isReady().then(() => {
  CapacitorApp.addListener('backButton', () => {
    // 获取当前路由
    const currentRoute = router.currentRoute.value
    
    // 如果当前在首页,直接退出应用
    if (currentRoute.name === 'home' || currentRoute.path === '/') {
      CapacitorApp.exitApp()
    } else {
      // 其他页面则返回上一页
      router.back()
    }
  })
})
