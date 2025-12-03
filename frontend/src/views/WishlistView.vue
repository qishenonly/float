<script setup>
import { useRouter } from 'vue-router'
import { useToast } from '../composables/useToast'

const router = useRouter()
const { showToast } = useToast()

const showFeature = (feature) => {
  showToast(`正在开发${feature}功能...`, 'info')
}
</script>

<template>
  <div>
    <!-- Header -->
    <div class="px-6 pt-12 pb-4 relative z-10 flex items-center justify-between animate-enter">
      <button @click="router.back()" class="w-10 h-10 rounded-full bg-white/50 flex items-center justify-center backdrop-blur-md text-gray-600 hover:bg-white transition active-press">
        <i class="fa-solid fa-arrow-left"></i>
      </button>
      <h1 class="text-xl font-bold text-gray-800">我的心愿单</h1>
      <button @click="showFeature('添加心愿')" class="w-10 h-10 rounded-full bg-pink-500 text-white flex items-center justify-center shadow-lg shadow-pink-200 active-press">
        <i class="fa-solid fa-plus"></i>
      </button>
    </div>

    <div class="px-6 pb-8 relative z-10 space-y-6">
      
      <!-- Summary -->
      <div class="text-center py-4 animate-enter delay-100">
        <p class="text-xs text-gray-400 mb-1">距离实现所有梦想还差</p>
        <h2 class="text-3xl font-extrabold text-pink-500">¥ 12,400</h2>
      </div>

      <!-- Wishlist Items -->
      <div class="glass-card p-5 rounded-3xl relative overflow-hidden group animate-enter delay-200 active-press">
        <div class="flex justify-between items-start mb-4">
          <div class="flex gap-4">
            <div class="w-16 h-16 rounded-2xl bg-gray-100 flex items-center justify-center text-3xl">🎮</div>
            <div>
              <h3 class="font-bold text-gray-800">Nintendo Switch</h3>
              <p class="text-xs text-gray-400 mt-1">给自己的生日礼物</p>
            </div>
          </div>
          <span class="bg-pink-100 text-pink-500 text-[10px] font-bold px-2 py-1 rounded-full">进行中</span>
        </div>
        
        <div class="mb-2 flex justify-between text-xs font-medium">
          <span class="text-pink-500">已存 ¥1,500</span>
          <span class="text-gray-400">目标 ¥2,200</span>
        </div>
        <div class="w-full bg-gray-200 h-2 rounded-full overflow-hidden">
          <div class="bg-gradient-to-r from-pink-400 to-rose-500 h-full w-[68%] rounded-full shadow-[0_0_10px_rgba(244,63,94,0.6)]"></div>
        </div>
      </div>

      <div class="glass-card p-5 rounded-3xl relative overflow-hidden group animate-enter delay-300 active-press">
        <div class="flex justify-between items-start mb-4">
          <div class="flex gap-4">
            <div class="w-16 h-16 rounded-2xl bg-gray-100 flex items-center justify-center text-3xl">✈️</div>
            <div>
              <h3 class="font-bold text-gray-800">日本京都之旅</h3>
              <p class="text-xs text-gray-400 mt-1">预计明年樱花季</p>
            </div>
          </div>
          <span class="bg-blue-100 text-blue-500 text-[10px] font-bold px-2 py-1 rounded-full">计划中</span>
        </div>
        
        <div class="mb-2 flex justify-between text-xs font-medium">
          <span class="text-blue-500">已存 ¥3,000</span>
          <span class="text-gray-400">目标 ¥15,000</span>
        </div>
        <div class="w-full bg-gray-200 h-2 rounded-full overflow-hidden">
          <div class="bg-gradient-to-r from-blue-400 to-cyan-400 h-full w-[20%] rounded-full"></div>
        </div>
      </div>

      <div class="glass-card p-5 rounded-3xl relative overflow-hidden opacity-80 grayscale-[0.5] animate-enter delay-400 active-press">
        <div class="flex justify-between items-start mb-4">
          <div class="flex gap-4">
            <div class="w-16 h-16 rounded-2xl bg-gray-100 flex items-center justify-center text-3xl">🎧</div>
            <div>
              <h3 class="font-bold text-gray-800">AirPods Max</h3>
              <p class="text-xs text-gray-400 mt-1">已达成 ✨</p>
            </div>
          </div>
          <span class="bg-green-100 text-green-500 text-[10px] font-bold px-2 py-1 rounded-full">已完成</span>
        </div>
        <div class="w-full bg-green-500 h-2 rounded-full overflow-hidden"></div>
      </div>

    </div>
  </div>
</template>
