<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useToast } from '../composables/useToast' 
const router = useRouter()
const { showToast } = useToast()

const darkMode = ref(false)
const gestureLock = ref(true)

const showFeature = (feature) => {
  showToast(`正在开发${feature}功能...`, 'info')
}

const handleLogout = () => {
  showToast('退出登录功能开发中...', 'warning')
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
      <div class="w-10"></div>
    </div>

    <div class="flex-1 overflow-y-auto hide-scrollbar px-6 pt-2 relative z-10 pb-10">
      
      <!-- Personalization Section -->
      <p class="text-xs font-bold text-gray-400 mb-2 ml-2 uppercase tracking-wide animate-enter delay-100">个性化</p>
      <div class="glass-card rounded-2xl overflow-hidden mb-6 animate-enter delay-200">
        <div @click="darkMode = !darkMode" class="flex items-center justify-between p-4 border-b border-gray-100/50 hover:bg-white/50 transition cursor-pointer active-press">
          <span class="text-sm font-bold text-gray-700">深色模式</span>
          <div :class="darkMode ? 'bg-indigo-500' : 'bg-gray-200'" class="w-11 h-6 rounded-full relative transition-colors duration-200 ease-in-out cursor-pointer">
            <div :class="darkMode ? 'translate-x-5' : 'translate-x-0.5'" class="w-5 h-5 bg-white rounded-full shadow absolute top-0.5 transition-transform duration-200 ease-in-out"></div>
          </div>
        </div>
        <div @click="showFeature('主题皮肤')" class="flex items-center justify-between p-4 border-b border-gray-100/50 hover:bg-white/50 transition cursor-pointer active-press">
          <span class="text-sm font-bold text-gray-700">主题皮肤</span>
          <div class="flex items-center gap-2">
            <div class="w-4 h-4 rounded-full bg-gradient-to-tr from-blue-400 to-purple-400 border border-white shadow-sm"></div>
            <span class="text-xs text-gray-400">极光紫</span>
            <i class="fa-solid fa-chevron-right text-gray-300 text-xs ml-1"></i>
          </div>
        </div>
        <div @click="showFeature('货币单位')" class="flex items-center justify-between p-4 hover:bg-white/50 transition cursor-pointer active-press">
          <span class="text-sm font-bold text-gray-700">货币单位</span>
          <div class="flex items-center gap-2">
            <span class="text-xs text-gray-400">CNY (¥)</span>
            <i class="fa-solid fa-chevron-right text-gray-300 text-xs ml-1"></i>
          </div>
        </div>
      </div>

      <!-- Security & Data Section -->
      <p class="text-xs font-bold text-gray-400 mb-2 ml-2 uppercase tracking-wide animate-enter delay-300">安全与数据</p>
      <div class="glass-card rounded-2xl overflow-hidden mb-6 animate-enter delay-400">
        <div @click="gestureLock = !gestureLock" class="flex items-center justify-between p-4 border-b border-gray-100/50 hover:bg-white/50 transition cursor-pointer active-press">
          <span class="text-sm font-bold text-gray-700">手势密码锁定</span>
          <div :class="gestureLock ? 'bg-indigo-500' : 'bg-gray-200'" class="w-11 h-6 rounded-full relative transition-colors duration-200 ease-in-out cursor-pointer">
            <div :class="gestureLock ? 'translate-x-5' : 'translate-x-0.5'" class="w-5 h-5 bg-white rounded-full shadow absolute top-0.5 transition-transform duration-200 ease-in-out"></div>
          </div>
        </div>
        <div @click="showFeature('云端同步')" class="flex items-center justify-between p-4 border-b border-gray-100/50 hover:bg-white/50 transition cursor-pointer active-press">
          <span class="text-sm font-bold text-gray-700">云端同步</span>
          <div class="flex items-center gap-2">
            <span class="text-xs text-green-500 font-medium">已同步</span>
            <i class="fa-solid fa-rotate text-gray-300 text-xs ml-1"></i>
          </div>
        </div>
        <div @click="showFeature('清除缓存')" class="flex items-center justify-between p-4 hover:bg-white/50 transition cursor-pointer active-press">
          <span class="text-sm font-bold text-gray-700">清除缓存</span>
          <span class="text-xs text-gray-400">24.5 MB</span>
        </div>
      </div>

      <!-- About Section -->
      <div class="glass-card rounded-2xl overflow-hidden mb-8 animate-enter delay-500">
        <div @click="showFeature('关于极光记账')" class="flex items-center justify-between p-4 border-b border-gray-100/50 hover:bg-white/50 transition cursor-pointer active-press">
          <span class="text-sm font-bold text-gray-700">关于极光记账</span>
          <i class="fa-solid fa-chevron-right text-gray-300 text-xs"></i>
        </div>
        <div @click="showFeature('给个好评')" class="flex items-center justify-between p-4 hover:bg-white/50 transition cursor-pointer active-press">
          <span class="text-sm font-bold text-gray-700">给个好评</span>
          <i class="fa-solid fa-heart text-pink-400 text-xs"></i>
        </div>
      </div>

      <!-- Logout Button -->
      <button @click="handleLogout" class="w-full text-center text-red-500 text-sm font-bold py-3 hover:bg-red-50 rounded-xl transition active-press">
        退出登录
      </button>
    </div>
  </div>
</template>
