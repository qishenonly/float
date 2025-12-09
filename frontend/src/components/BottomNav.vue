<script setup>
import { RouterLink } from 'vue-router';
</script>

<template>
  <div class="absolute bottom-0 w-full bg-white/80 backdrop-blur-xl border-t border-white/50 px-6 py-4 pb-6 z-20">
    <div class="flex justify-between items-center relative">
      <RouterLink to="/" class="flex flex-col items-center gap-1 text-gray-400 hover:text-gray-900 transition active-press" active-class="!text-gray-900">
        <i class="fa-solid fa-house text-xl drop-shadow-sm"></i>
        <span class="text-[10px] font-bold">首页</span>
      </RouterLink>
      
      <RouterLink to="/stats" class="flex flex-col items-center gap-1 text-gray-400 hover:text-gray-600 transition active-press" active-class="!text-gray-900">
        <i class="fa-solid fa-chart-pie text-xl"></i>
        <span class="text-[10px] font-medium">统计</span>
      </RouterLink>
      
      <RouterLink to="/add" class="relative -top-10 group active-press">
        <div class="absolute inset-0 bg-indigo-500 rounded-full blur opacity-40 group-hover:opacity-60 transition animate-pulse"></div>
        <div class="w-16 h-16 bg-gradient-to-tr from-indigo-600 to-violet-600 rounded-full flex items-center justify-center shadow-2xl shadow-indigo-300 text-white border-[6px] border-[#fbfdff] relative z-10 transform group-hover:scale-110 transition duration-300">
          <i class="fa-solid fa-plus text-2xl"></i>
        </div>
      </RouterLink>

      <RouterLink to="/assets" class="flex flex-col items-center gap-1 text-gray-400 hover:text-gray-600 transition active-press" active-class="!text-gray-900">
        <i class="fa-solid fa-wallet text-xl"></i>
        <span class="text-[10px] font-medium">资产</span>
      </RouterLink>
      
      <RouterLink to="/profile" class="flex flex-col items-center gap-1 text-gray-400 hover:text-gray-600 transition active-press" active-class="!text-gray-900">
        <i class="fa-solid fa-user text-xl"></i>
        <span class="text-[10px] font-medium">我的</span>
      </RouterLink>
    </div>
  </div>
</template>
