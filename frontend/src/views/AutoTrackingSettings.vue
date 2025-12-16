<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import AutoBookkeeping from '../plugins/autoBookkeeping'
import GlassCard from '../components/GlassCard.vue'

const router = useRouter()
const permissions = ref({
  overlay: false,
  accessibility: false,
  notification: false,
  battery: false
})

const monitorEnabled = ref(localStorage.getItem('notification_monitor_enabled') === 'true')

const toggleMonitor = () => {
    monitorEnabled.value = !monitorEnabled.value
    localStorage.setItem('notification_monitor_enabled', monitorEnabled.value)
    AutoBookkeeping.setNotificationMonitorEnabled({ enabled: monitorEnabled.value })
}


const checkPermissions = async () => {
  try {
    const result = await AutoBookkeeping.checkPermissions()
    permissions.value = result
    // Sync monitor switch with actual native state
    if (result.monitorEnabled !== undefined) {
        monitorEnabled.value = result.monitorEnabled
        localStorage.setItem('notification_monitor_enabled', result.monitorEnabled)
    }
  } catch (e) {
    console.error('Failed to check permissions', e)
  }
}

onMounted(() => {
  checkPermissions()
  // Resume check when coming back from settings
  document.addEventListener('resume', checkPermissions)
})

const goBack = () => router.back()

const requestOverlay = async () => {
  await AutoBookkeeping.requestOverlayPermission()
}

const requestAccessibility = async () => {
  await AutoBookkeeping.requestAccessibilityPermission()
}

const requestNotification = async () => {
  await AutoBookkeeping.requestNotificationPermission()
}

const requestBattery = async () => {
  await AutoBookkeeping.requestIgnoreBatteryOptimizations()
}
</script>

<template>
  <div>
    <!-- Header -->
    <div class="px-6 pt-12 pb-4 relative z-10 flex items-center justify-between animate-enter">
      <button @click="goBack" class="w-10 h-10 rounded-full bg-white/50 flex items-center justify-center backdrop-blur-md text-gray-600 hover:bg-white transition active-press">
        <i class="fa-solid fa-arrow-left"></i>
      </button>
      <h1 class="text-lg font-bold text-gray-800">自动记账</h1>
      <div class="w-10"></div> <!-- Placeholder for balance -->
    </div>

    <div class="px-6 pb-32 space-y-6 animate-enter delay-100">
      
      <!-- Intro Card -->
      <GlassCard class="p-5 rounded-2xl">
        <div class="flex items-center gap-4 mb-3">
          <div class="w-12 h-12 bg-gradient-to-br from-indigo-500 to-violet-500 rounded-2xl flex items-center justify-center text-white shadow-lg shadow-indigo-200">
             <i class="fa-solid fa-wand-magic-sparkles text-xl"></i>
          </div>
          <div>
            <h2 class="font-bold text-gray-800">智能记账助手</h2>
            <p class="text-xs text-gray-500 mt-0.5">支付成功后自动弹出记账窗口</p>
          </div>
        </div>
        <p class="text-sm text-gray-600 leading-relaxed">
          开启后，当检测到支付宝或微信支付成功页面时，Float 会自动弹出一个悬浮窗，帮您快速记录这笔支出。
        </p>
      </GlassCard>

      <!-- Required Permissions -->
      <div>
        <h2 class="text-xs font-bold text-gray-400 uppercase tracking-wider mb-3 px-2">必要权限</h2>
        <GlassCard class="rounded-2xl shadow-sm overflow-hidden">
          <div class="divide-y divide-gray-100/50">
            <!-- Overlay Permission -->
            <div @click="requestOverlay" class="flex items-center justify-between p-4 hover:bg-white/50 cursor-pointer transition active-press">
              <div class="flex items-center gap-3">
                 <div class="w-8 h-8 rounded-full flex items-center justify-center" :class="permissions.overlay ? 'bg-green-100 text-green-500' : 'bg-gray-100 text-gray-400'">
                    <i class="fa-solid fa-layer-group text-sm"></i>
                 </div>
                 <div>
                   <span class="text-sm font-bold text-gray-700">悬浮窗权限</span>
                   <p class="text-xs text-gray-400">用于显示快速记账窗口</p>
                 </div>
              </div>
              <div class="flex items-center gap-2">
                <span class="text-xs font-medium" :class="permissions.overlay ? 'text-green-500' : 'text-orange-500'">
                  {{ permissions.overlay ? '已开启' : '去开启' }}
                </span>
                <i class="fa-solid fa-chevron-right text-gray-300 text-xs"></i>
              </div>
            </div>

            <!-- Accessibility Permission -->
            <div @click="requestAccessibility" class="flex items-center justify-between p-4 hover:bg-white/50 cursor-pointer transition active-press">
               <div class="flex items-center gap-3">
                 <div class="w-8 h-8 rounded-full flex items-center justify-center" :class="permissions.accessibility ? 'bg-green-100 text-green-500' : 'bg-gray-100 text-gray-400'">
                    <i class="fa-solid fa-universal-access text-sm"></i>
                 </div>
                 <div>
                   <span class="text-sm font-bold text-gray-700">无障碍服务</span>
                   <p class="text-xs text-gray-400">用于识别支付成功页面</p>
                 </div>
              </div>
              <div class="flex items-center gap-2">
                <span class="text-xs font-medium" :class="permissions.accessibility ? 'text-green-500' : 'text-orange-500'">
                  {{ permissions.accessibility ? '已开启' : '去开启' }}
                </span>
                 <i class="fa-solid fa-chevron-right text-gray-300 text-xs"></i>
              </div>
            </div>
          </div>
        </GlassCard>
      </div>
      
      <!-- Background Keep Alive -->
      <div>
        <h2 class="text-xs font-bold text-gray-400 uppercase tracking-wider mb-3 px-2">后台运行</h2>
        <GlassCard class="rounded-2xl shadow-sm overflow-hidden">
          <div class="divide-y divide-gray-100/50">
             <div @click="requestBattery" class="flex items-center justify-between p-4 hover:bg-white/50 cursor-pointer transition active-press">
               <div class="flex items-center gap-3">
                 <div class="w-8 h-8 rounded-full flex items-center justify-center" :class="permissions.battery ? 'bg-green-100 text-green-500' : 'bg-gray-100 text-gray-400'">
                    <i class="fa-solid fa-battery-full text-sm"></i>
                 </div>
                 <div>
                   <span class="text-sm font-bold text-gray-700">忽略电池优化</span>
                   <p class="text-xs text-gray-400">防止应用在后台被系统清理</p>
                 </div>
              </div>
              <div class="flex items-center gap-2">
                <span class="text-xs font-medium" :class="permissions.battery ? 'text-green-500' : 'text-orange-500'">
                  {{ permissions.battery ? '已开启' : '去设置' }}
                </span>
                 <i class="fa-solid fa-chevron-right text-gray-300 text-xs"></i>
              </div>
            </div>
          </div>
        </GlassCard>
      </div>

      <!-- Alternative Options -->
      <div>
        <h2 class="text-xs font-bold text-gray-400 uppercase tracking-wider mb-3 px-2">备用方案</h2>
        <GlassCard class="rounded-2xl shadow-sm overflow-hidden">
          <div class="divide-y divide-gray-100/50">
            <!-- Notification Permission -->
            <div @click="requestNotification" class="flex items-center justify-between p-4 hover:bg-white/50 cursor-pointer transition active-press">
               <div class="flex items-center gap-3">
                 <div class="w-8 h-8 rounded-full flex items-center justify-center" :class="permissions.notification ? 'bg-green-100 text-green-500' : 'bg-gray-100 text-gray-400'">
                    <i class="fa-solid fa-bell text-sm"></i>
                 </div>
                 <div>
                   <span class="text-sm font-bold text-gray-700">通知监听权限</span>
                   <p class="text-xs text-gray-400">无法使用无障碍服务时的备选方案</p>
                 </div>
              </div>
              <div class="flex items-center gap-2">
                <span class="text-xs font-medium" :class="permissions.notification ? 'text-green-500' : 'text-gray-400'">
                  {{ permissions.notification ? '已开启' : '未开启' }}
                </span>
                 <i class="fa-solid fa-chevron-right text-gray-300 text-xs"></i>
              </div>
            </div>
          </div>
        </GlassCard>
      </div>

       <!-- Test Monitor Toggle -->
       <div>
        <h2 class="text-xs font-bold text-gray-400 uppercase tracking-wider mb-3 px-2">开发者选项</h2>
        <GlassCard class="rounded-2xl shadow-sm overflow-hidden">
           <div class="flex items-center justify-between p-4">
               <div class="flex items-center gap-3">
                 <div class="w-8 h-8 rounded-full bg-blue-100 flex items-center justify-center text-blue-500">
                    <i class="fa-solid fa-bug text-sm"></i>
                 </div>
                 <div>
                   <span class="text-sm font-bold text-gray-700">测试模式</span>
                   <p class="text-xs text-gray-400">监听所有通知并弹窗显示</p>
                 </div>
              </div>
               <div class="relative inline-flex items-center cursor-pointer" @click="toggleMonitor">
                  <input type="checkbox" v-model="monitorEnabled" class="sr-only peer">
                  <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-indigo-600"></div>
               </div>
            </div>
        </GlassCard>
       </div>

       <!-- Tips -->
       <div class="bg-indigo-50/80 backdrop-blur-sm text-indigo-600 p-4 rounded-2xl text-xs leading-relaxed border border-indigo-100">
         <i class="fa-solid fa-circle-info mr-1"></i>
         如果开启了权限但无法自动识别，请检查应用是否被后台清理。建议在系统设置中允许应用"后台运行"或"自启动"。
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
