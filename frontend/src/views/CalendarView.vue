<script setup>
import { useRouter } from 'vue-router'
import GlassCard from '../components/GlassCard.vue'

const router = useRouter()
</script>

<template>
  <div>
    <!-- Header -->
    <div class="px-6 pt-12 pb-4 relative z-10 flex items-center justify-between animate-enter">
      <button @click="router.back()" class="w-10 h-10 rounded-full bg-white/50 flex items-center justify-center backdrop-blur-md text-gray-600 hover:bg-white transition active-press">
        <i class="fa-solid fa-arrow-left"></i>
      </button>
      <div class="flex flex-col items-center">
        <span class="text-xs text-gray-500 font-medium">2025</span>
        <div class="flex items-center gap-2">
          <h1 class="text-lg font-bold text-gray-800">December</h1>
          <i class="fa-solid fa-caret-down text-xs text-gray-400"></i>
        </div>
      </div>
      <div class="w-10"></div>
    </div>

    <div class="px-6 relative z-10">
      
      <!-- Calendar -->
      <div class="glass-card p-4 rounded-3xl mb-6 shadow-sm animate-enter delay-100">
        <div class="grid grid-cols-7 mb-4 text-center">
          <span class="text-[10px] text-gray-400 font-bold">S</span>
          <span class="text-[10px] text-gray-400 font-bold">M</span>
          <span class="text-[10px] text-gray-400 font-bold">T</span>
          <span class="text-[10px] text-gray-400 font-bold">W</span>
          <span class="text-[10px] text-gray-400 font-bold">T</span>
          <span class="text-[10px] text-gray-400 font-bold">F</span>
          <span class="text-[10px] text-gray-400 font-bold">S</span>
        </div>
        
        <div class="grid grid-cols-7 gap-y-4 text-center">
          <span class="text-sm text-gray-300">29</span>
          <span class="text-sm text-gray-300">30</span>
          
          <div class="flex flex-col items-center gap-1 cursor-pointer active-press">
            <span class="text-sm font-medium text-gray-600">1</span>
            <div class="w-1 h-1 rounded-full bg-gray-300"></div>
          </div>
          <div class="flex flex-col items-center gap-1 cursor-pointer active-press">
            <span class="text-sm font-medium text-gray-600">2</span>
            <div class="w-1 h-1 rounded-full bg-blue-500"></div>
          </div>
          <div class="flex flex-col items-center gap-1 cursor-pointer relative active-press">
            <div class="w-8 h-8 rounded-full bg-blue-600 absolute -top-1.5 -z-10 shadow-lg shadow-blue-300"></div>
            <span class="text-sm font-bold text-white">3</span>
            <div class="w-1 h-1 rounded-full bg-white"></div>
          </div>
          <div class="flex flex-col items-center gap-1 cursor-pointer active-press">
            <span class="text-sm font-medium text-gray-600">4</span>
          </div>
          <div class="flex flex-col items-center gap-1 cursor-pointer active-press">
            <span class="text-sm text-gray-600">5</span>
            <div class="w-1 h-1 rounded-full bg-blue-500"></div>
          </div>
          <div class="flex flex-col items-center gap-1 cursor-pointer active-press">
            <span class="text-sm text-gray-600">6</span>
            <div class="w-1 h-1 rounded-full bg-red-400"></div>
          </div>
        </div>
      </div>

      <!-- Daily Transactions -->
      <div class="pb-32 animate-enter delay-200">
        <div class="flex justify-between items-end mb-4 px-2">
          <h3 class="font-bold text-gray-800">12月3日</h3>
          <span class="text-sm font-bold text-gray-800">- ¥ 129.00</span>
        </div>

        <div class="space-y-3">
          <GlassCard class="p-4 rounded-2xl flex items-center justify-between">
            <div class="flex items-center gap-3">
              <div class="w-10 h-10 rounded-xl bg-orange-100 text-orange-500 flex items-center justify-center">
                <i class="fa-solid fa-burger"></i>
              </div>
              <div>
                <h4 class="font-bold text-gray-800 text-sm">Shake Shack</h4>
                <p class="text-[10px] text-gray-400">18:30</p>
              </div>
            </div>
            <span class="font-bold text-gray-800">- 85.00</span>
          </GlassCard>
          
          <GlassCard class="p-4 rounded-2xl flex items-center justify-between">
            <div class="flex items-center gap-3">
              <div class="w-10 h-10 rounded-xl bg-purple-100 text-purple-500 flex items-center justify-center">
                <i class="fa-solid fa-mug-hot"></i>
              </div>
              <div>
                <h4 class="font-bold text-gray-800 text-sm">Starbucks</h4>
                <p class="text-[10px] text-gray-400">09:15</p>
              </div>
            </div>
            <span class="font-bold text-gray-800">- 44.00</span>
          </GlassCard>
        </div>
      </div>

    </div>
  </div>
</template>
