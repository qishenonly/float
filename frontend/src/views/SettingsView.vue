<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useToast } from '../composables/useToast'
import { useAppUpdate } from '../composables/useAppUpdate'
import GlassCard from '../components/GlassCard.vue'

const router = useRouter()
const { showToast } = useToast()
const { currentVersion, checkUpdate, openUpdateDialog } = useAppUpdate()

const hasUpdate = ref(false)
const latestVersion = ref(null)
const checking = ref(false)

// Mock data for other settings
const darkMode = ref(false)
const gestureLock = ref(true)

onMounted(async () => {
  await performCheck()
})

const performCheck = async (manual = false) => {
  if (checking.value) return
  checking.value = true
  
  try {
    const latest = await checkUpdate(manual)
    
    if (latest) {
      hasUpdate.value = true
      latestVersion.value = latest
      if (manual) {
        openUpdateDialog(latest)
      }
    } else {
      hasUpdate.value = false
      if (manual) {
        showToast('当前已是最新版本', 'success')
      }
    }
  } catch (error) {
    console.error('Failed to check update:', error)
    if (manual) {
      showToast('检查更新失败，请稍后重试', 'error')
    }
  } finally {
    checking.value = false
  }
}

const handleUpdateClick = () => {
  if (hasUpdate.value && latestVersion.value) {
    openUpdateDialog(latestVersion.value)
  } else {
    performCheck(true)
  }
}
</script>

<template>
  <div>
    <!-- Header -->
    <div class="px-6 pt-12 pb-4 relative z-10 flex items-center justify-between animate-enter">
      <button @click="router.back()" class="w-10 h-10 rounded-full bg-white/50 flex items-center justify-center backdrop-blur-md text-gray-600 hover:bg-white transition active-press">
        <i class="fa-solid fa-arrow-left"></i>
      </button>
      <h1 class="text-lg font-bold text-gray-800">通用设置</h1>
      <div class="w-10"></div> <!-- Placeholder for balance -->
    </div>

    <div class="px-6 pb-32 space-y-6 animate-enter delay-100">
      
      <!-- Personalization -->
      <div>
        <h2 class="text-xs font-bold text-gray-400 uppercase tracking-wider mb-3 px-2">个性化</h2>
        <GlassCard class="rounded-2xl shadow-sm overflow-hidden">
          <div class="divide-y divide-gray-100/50">
             <div class="flex items-center justify-between p-4">
                <span class="text-sm font-bold text-gray-700">深色模式</span>
                <div class="relative inline-flex items-center cursor-pointer" @click="darkMode = !darkMode">
                  <input type="checkbox" v-model="darkMode" class="sr-only peer">
                  <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-indigo-600"></div>
                </div>
             </div>
             <div class="flex items-center justify-between p-4 hover:bg-white/50 cursor-pointer transition active-press">
                <span class="text-sm font-bold text-gray-700">主题皮肤</span>
                <div class="flex items-center gap-2">
                    <div class="w-3 h-3 rounded-full bg-indigo-500"></div>
                    <span class="text-xs text-gray-400">极光紫</span>
                    <i class="fa-solid fa-chevron-right text-gray-300 text-xs"></i>
                </div>
             </div>
             <div class="flex items-center justify-between p-4 hover:bg-white/50 cursor-pointer transition active-press">
                <span class="text-sm font-bold text-gray-700">货币单位</span>
                <div class="flex items-center gap-2">
                    <span class="text-xs text-gray-400">CNY (¥)</span>
                    <i class="fa-solid fa-chevron-right text-gray-300 text-xs"></i>
                </div>
             </div>
          </div>
        </GlassCard>
      </div>

      <!-- Security & Data -->
      <div>
        <h2 class="text-xs font-bold text-gray-400 uppercase tracking-wider mb-3 px-2">安全与数据</h2>
        <GlassCard class="rounded-2xl shadow-sm overflow-hidden">
          <div class="divide-y divide-gray-100/50">
             <div class="flex items-center justify-between p-4">
                <span class="text-sm font-bold text-gray-700">手势密码锁定</span>
                <div class="relative inline-flex items-center cursor-pointer" @click="gestureLock = !gestureLock">
                  <input type="checkbox" v-model="gestureLock" class="sr-only peer">
                  <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-indigo-600"></div>
                </div>
             </div>
             <div class="flex items-center justify-between p-4 hover:bg-white/50 cursor-pointer transition active-press">
                <span class="text-sm font-bold text-gray-700">云端同步</span>
                <div class="flex items-center gap-2">
                    <span class="text-xs text-green-500 font-medium">已同步</span>
                    <i class="fa-solid fa-rotate text-gray-300 text-xs"></i>
                </div>
             </div>
             <div class="flex items-center justify-between p-4 hover:bg-white/50 cursor-pointer transition active-press">
                <span class="text-sm font-bold text-gray-700">清除缓存</span>
                <span class="text-xs text-gray-400">24.5 MB</span>
             </div>
          </div>
        </GlassCard>
      </div>

      <!-- About -->
      <div>
        <GlassCard class="rounded-2xl shadow-sm overflow-hidden">
          <div class="divide-y divide-gray-100/50">
            <div class="flex items-center justify-between p-4 hover:bg-white/50 cursor-pointer transition active-press">
                <span class="text-sm font-bold text-gray-700">关于浮岛</span>
                <i class="fa-solid fa-chevron-right text-gray-300 text-xs"></i>
            </div>
            
            <!-- Check Update -->
            <div @click="handleUpdateClick" class="flex items-center justify-between p-4 hover:bg-white/50 cursor-pointer transition active-press">
                <span class="text-sm font-bold text-gray-700">检查更新</span>
                <div class="flex items-center gap-2">
                    <div v-if="hasUpdate" class="flex items-center gap-2">
                        <span class="text-xs text-red-500 font-medium">发现新版本</span>
                        <div class="w-2 h-2 bg-red-500 rounded-full animate-pulse"></div>
                    </div>
                    <span v-else class="text-xs text-gray-400">v{{ currentVersion.name }}</span>
                    <i class="fa-solid fa-chevron-right text-gray-300 text-xs"></i>
                </div>
            </div>

            <div @click="showToast('感谢您的支持！', 'success')" class="flex items-center justify-between p-4 hover:bg-white/50 cursor-pointer transition active-press">
                <span class="text-sm font-bold text-gray-700">给个好评</span>
                <i class="fa-solid fa-heart text-pink-400 text-sm"></i>
            </div>
          </div>
        </GlassCard>
      </div>
      
      <div class="text-center pt-4">
          <button @click="router.push('/login')" class="text-sm font-bold text-red-500 py-2 px-6">退出登录</button>
      </div>
    </div>
  </div>
</template>

<style scoped>
@keyframes fadeInUp {
  from { opacity: 0; transform: translateY(20px); }
  to { opacity: 1; transform: translateY(0); }
}

.animate-enter {
  animation: fadeInUp 0.5s ease-out forwards;
  opacity: 0;
}

.delay-100 { animation-delay: 0.1s; }

.active-press:active {
  transform: scale(0.98);
  transition: transform 0.1s;
}
</style>
