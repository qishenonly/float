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
      <h1 class="text-xl font-bold text-gray-800">订阅 & 账单</h1>
      <button @click="showFeature('添加账单')" class="w-10 h-10 rounded-full bg-indigo-500 text-white flex items-center justify-center shadow-lg shadow-indigo-200 active-press">
        <i class="fa-solid fa-plus"></i>
      </button>
    </div>

    <div class="px-6 relative z-10">
      
      <!-- Statistics -->
      <div class="flex gap-4 mb-8 animate-enter delay-100">
        <div class="glass-card p-4 rounded-3xl flex-1">
          <div class="w-8 h-8 rounded-full bg-purple-100 text-purple-600 flex items-center justify-center mb-2">
            <i class="fa-solid fa-calendar-check"></i>
          </div>
          <p class="text-[10px] text-gray-400">本月固定支出</p>
          <p class="text-lg font-bold text-gray-800">¥ 3,240</p>
        </div>
        <div class="glass-card p-4 rounded-3xl flex-1">
          <div class="w-8 h-8 rounded-full bg-red-100 text-red-500 flex items-center justify-center mb-2">
            <i class="fa-solid fa-clock"></i>
          </div>
          <p class="text-[10px] text-gray-400">即将到期</p>
          <p class="text-lg font-bold text-gray-800">3 笔</p>
        </div>
      </div>

      <!-- Upcoming Bills -->
      <h3 class="font-bold text-gray-800 mb-4 animate-enter delay-200">即将扣费</h3>
      <div class="space-y-3 pb-32 animate-enter delay-300">
        
        <div class="glass-card p-4 rounded-2xl flex items-center justify-between border-l-4 border-l-red-400 active-press">
          <div class="flex items-center gap-3">
            <div class="w-12 h-12 rounded-xl bg-black text-red-600 flex items-center justify-center text-xl overflow-hidden">
              <i class="fa-brands fa-netflix"></i>
            </div>
            <div>
              <h4 class="font-bold text-gray-800">Netflix 高级会员</h4>
              <p class="text-[10px] text-red-500 font-bold mt-0.5">明天到期</p>
            </div>
          </div>
          <div class="text-right">
            <p class="font-bold text-gray-800">¥ 39.00</p>
            <p class="text-[10px] text-gray-400">VISA • 1011</p>
          </div>
        </div>

        <div class="glass-card p-4 rounded-2xl flex items-center justify-between border-l-4 border-l-indigo-400 active-press">
          <div class="flex items-center gap-3">
            <div class="w-12 h-12 rounded-xl bg-green-100 text-green-600 flex items-center justify-center text-xl">
              <i class="fa-brands fa-spotify"></i>
            </div>
            <div>
              <h4 class="font-bold text-gray-800">Spotify 音乐</h4>
              <p class="text-[10px] text-gray-400 mt-0.5">3天后 • 自动续费</p>
            </div>
          </div>
          <div class="text-right">
            <p class="font-bold text-gray-800">¥ 15.00</p>
            <p class="text-[10px] text-gray-400">微信支付</p>
          </div>
        </div>
        
        <div class="glass-card p-4 rounded-2xl flex items-center justify-between border-l-4 border-l-indigo-400 active-press">
          <div class="flex items-center gap-3">
            <div class="w-12 h-12 rounded-xl bg-blue-100 text-blue-600 flex items-center justify-center text-xl">
              <i class="fa-solid fa-house"></i>
            </div>
            <div>
              <h4 class="font-bold text-gray-800">房租</h4>
              <p class="text-[10px] text-gray-400 mt-0.5">10天后 • 记得转账</p>
            </div>
          </div>
          <div class="text-right">
            <p class="font-bold text-gray-800">¥ 3,000</p>
            <p class="text-[10px] text-gray-400">招商银行</p>
          </div>
        </div>

      </div>
    </div>
  </div>
</template>
