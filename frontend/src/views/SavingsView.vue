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
      <h1 class="text-xl font-bold text-gray-800">我的小金库</h1>
      <div class="w-10"></div>
    </div>

    <div class="px-6 relative z-10">
      
      <!-- Savings Summary Card -->
      <div class="relative overflow-hidden rounded-[2.5rem] p-8 text-center shadow-xl shadow-orange-200/50 mb-8 mt-4 animate-enter delay-100">
        <div class="absolute inset-0 bg-gradient-to-tr from-orange-400 to-amber-300 z-0"></div>
        <div class="absolute top-0 left-0 w-full h-full bg-[url('data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iNjAiIGhlaWdodD0iNjAiIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyI+PHBhdGggZD0iTTAgMGg2MHY2MEgweiIgZmlsbD0ibm9uZSIvPjxwYXRoIGQ9Ik0zMCAwdjYwTTAgMzBoNjAiIHN0cm9rZT0iIzAwMCIgc3Ryb2tlLXdpZHRoPSIuNSIgb3BhY2l0eT0iLjEiLz48L3N2Zz4=')] opacity-10"></div>
        
        <div class="relative z-10 active-press">
          <div class="w-20 h-20 bg-white/20 rounded-full mx-auto mb-4 flex items-center justify-center text-4xl backdrop-blur-sm shadow-inner text-white">
            <i class="fa-solid fa-piggy-bank"></i>
          </div>
          <p class="text-orange-50 font-medium mb-2 tracking-wide">当前累计存款</p>
          <h2 class="text-4xl font-extrabold text-white mb-6">¥ 35,200</h2>
          
          <button @click="showFeature('存入')" class="w-full bg-white text-orange-500 font-bold py-3.5 rounded-xl shadow-lg hover:scale-[1.02] transition active-press">
            <i class="fa-solid fa-coins mr-2"></i> 存入一笔
          </button>
        </div>
      </div>

      <!-- Active Plans -->
      <div class="space-y-4 pb-8 animate-enter delay-200">
        <h3 class="font-bold text-gray-800 px-2">正在进行</h3>
        
        <div class="glass-card p-5 rounded-3xl flex items-center gap-4 active-press">
          <div class="relative w-14 h-14">
            <svg class="w-full h-full transform -rotate-90">
              <circle cx="28" cy="28" r="24" stroke="currentColor" stroke-width="4" fill="transparent" class="text-gray-200" />
              <circle cx="28" cy="28" r="24" stroke="currentColor" stroke-width="4" fill="transparent" stroke-dasharray="150" stroke-dashoffset="30" class="text-orange-400" />
            </svg>
            <div class="absolute inset-0 flex items-center justify-center text-[10px] font-bold text-gray-600">80%</div>
          </div>
          <div class="flex-1">
            <h4 class="font-bold text-gray-800">52周存钱法</h4>
            <p class="text-xs text-gray-400 mt-0.5">第 32 周 • 坚持就是胜利</p>
          </div>
          <button @click="showFeature('存钱计划详情')" class="w-8 h-8 rounded-full bg-orange-100 text-orange-500 flex items-center justify-center active-press">
            <i class="fa-solid fa-chevron-right text-xs"></i>
          </button>
        </div>

        <div class="glass-card p-5 rounded-3xl flex items-center gap-4 active-press">
          <div class="relative w-14 h-14">
            <svg class="w-full h-full transform -rotate-90">
              <circle cx="28" cy="28" r="24" stroke="currentColor" stroke-width="4" fill="transparent" class="text-gray-200" />
              <circle cx="28" cy="28" r="24" stroke="currentColor" stroke-width="4" fill="transparent" stroke-dasharray="150" stroke-dashoffset="120" class="text-yellow-400" />
            </svg>
            <div class="absolute inset-0 flex items-center justify-center text-[10px] font-bold text-gray-600">20%</div>
          </div>
          <div class="flex-1">
            <h4 class="font-bold text-gray-800">买房基金</h4>
            <p class="text-xs text-gray-400 mt-0.5">积少成多</p>
          </div>
          <button @click="showFeature('存钱计划详情')" class="w-8 h-8 rounded-full bg-orange-100 text-orange-500 flex items-center justify-center active-press">
            <i class="fa-solid fa-chevron-right text-xs"></i>
          </button>
        </div>
      </div>

    </div>
  </div>
</template>
