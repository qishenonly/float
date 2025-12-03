<script setup>
import GlassCard from '../components/GlassCard.vue'
</script>

<template>
  <div>
    <!-- Header -->
    <div class="px-6 pt-10 pb-2 flex justify-between items-center animate-enter">
      <div>
        <div class="flex items-center gap-2 text-gray-500 mb-1">
          <span class="text-xs font-medium tracking-wider uppercase">Tuesday, Dec 03</span>
          <i class="fa-solid fa-cloud-sun text-yellow-400"></i>
        </div>
        <h1 class="text-2xl font-extrabold text-gray-800 tracking-tight">Hi, Alex 👋</h1>
        <p class="text-xs text-gray-400 mt-1">今天也是充满希望的一天 ✨</p>
      </div>
      <RouterLink to="/profile" class="w-11 h-11 rounded-full p-1 bg-white shadow-sm cursor-pointer hover:scale-105 transition">
        <img src="https://api.dicebear.com/7.x/avataaars/svg?seed=Alex&backgroundColor=ffdfbf" class="w-full h-full rounded-full" alt="avatar">
      </RouterLink>
    </div>

    <!-- Balance Card -->
    <div class="px-6 mt-6 animate-enter delay-100">
      <div class="relative overflow-hidden rounded-[2rem] p-6 text-white shadow-xl shadow-indigo-200/50 group transition-transform hover:scale-[1.02] duration-300">
        <div class="absolute inset-0 bg-gradient-to-tr from-violet-600 to-indigo-500 z-0"></div>
        <div class="absolute -right-10 -top-10 w-40 h-40 bg-white opacity-10 rounded-full blur-2xl"></div>
        <div class="absolute -left-10 -bottom-10 w-32 h-32 bg-purple-400 opacity-20 rounded-full blur-xl"></div>

        <div class="relative z-10">
          <div class="flex justify-between items-start mb-6">
            <div>
              <p class="text-indigo-100 text-xs font-medium mb-1 tracking-wide opacity-80">本月结余</p>
              <div class="flex items-baseline gap-2">
                <h2 class="text-3xl font-bold tracking-tight">¥ 4,250.<span class="text-lg opacity-70">00</span></h2>
                <button class="text-indigo-200 hover:text-white transition"><i class="fa-solid fa-eye-slash text-sm"></i></button>
              </div>
            </div>
            <div class="bg-white/20 backdrop-blur-md p-2 rounded-xl">
              <i class="fa-solid fa-chart-simple text-white"></i>
            </div>
          </div>

          <div class="mb-4">
            <div class="flex justify-between text-[10px] text-indigo-100 mb-1.5 px-0.5">
              <span>预算进度</span>
              <span>已用 45%</span>
            </div>
            <div class="w-full bg-black/20 h-1.5 rounded-full overflow-hidden backdrop-blur-sm">
              <div class="bg-white h-full w-[45%] rounded-full shadow-[0_0_10px_rgba(255,255,255,0.5)]"></div>
            </div>
          </div>

          <div class="flex gap-3 mt-4">
            <div class="bg-indigo-900/30 backdrop-blur-md rounded-xl px-3 py-2 flex-1 border border-white/10">
              <div class="flex items-center gap-2 mb-1">
                <div class="w-1.5 h-1.5 rounded-full bg-green-400"></div>
                <span class="text-[10px] text-indigo-100">收入</span>
              </div>
              <span class="text-sm font-semibold">8,200</span>
            </div>
            <div class="bg-indigo-900/30 backdrop-blur-md rounded-xl px-3 py-2 flex-1 border border-white/10">
              <div class="flex items-center gap-2 mb-1">
                <div class="w-1.5 h-1.5 rounded-full bg-pink-400"></div>
                <span class="text-[10px] text-indigo-100">支出</span>
              </div>
              <span class="text-sm font-semibold">3,950</span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Quick Actions -->
    <div class="px-6 mt-6 grid grid-cols-4 gap-4 animate-enter delay-200">
      <div class="flex flex-col items-center gap-2 cursor-pointer group">
        <div class="w-14 h-14 rounded-2xl bg-white shadow-sm flex items-center justify-center text-xl text-pink-500 group-hover:shadow-md group-hover:-translate-y-1 transition duration-300">
          <i class="fa-solid fa-gift"></i>
        </div>
        <span class="text-[10px] font-medium text-gray-500">心愿单</span>
      </div>
      <div class="flex flex-col items-center gap-2 cursor-pointer group">
        <div class="w-14 h-14 rounded-2xl bg-white shadow-sm flex items-center justify-center text-xl text-blue-500 group-hover:shadow-md group-hover:-translate-y-1 transition duration-300">
          <i class="fa-solid fa-calendar-days"></i>
        </div>
        <span class="text-[10px] font-medium text-gray-500">日历</span>
      </div>
      <div class="flex flex-col items-center gap-2 cursor-pointer group">
        <div class="w-14 h-14 rounded-2xl bg-white shadow-sm flex items-center justify-center text-xl text-orange-500 group-hover:shadow-md group-hover:-translate-y-1 transition duration-300">
          <i class="fa-solid fa-piggy-bank"></i>
        </div>
        <span class="text-[10px] font-medium text-gray-500">存钱</span>
      </div>
      <div class="flex flex-col items-center gap-2 cursor-pointer group">
        <div class="w-14 h-14 rounded-2xl bg-white shadow-sm flex items-center justify-center text-xl text-indigo-500 group-hover:shadow-md group-hover:-translate-y-1 transition duration-300">
          <i class="fa-solid fa-file-invoice"></i>
        </div>
        <span class="text-[10px] font-medium text-gray-500">账单</span>
      </div>
    </div>

    <!-- Recent Transactions -->
    <div class="px-6 mt-8 animate-enter delay-300">
      <div class="flex justify-between items-center mb-4">
        <h3 class="text-lg font-bold text-gray-800">今日动态</h3>
        <button class="w-8 h-8 rounded-full bg-white flex items-center justify-center text-gray-400 shadow-sm hover:text-gray-800">
          <i class="fa-solid fa-ellipsis"></i>
        </button>
      </div>

      <div class="space-y-4">
        <GlassCard class="p-4 flex items-center justify-between cursor-pointer">
          <div class="flex items-center gap-4">
            <div class="w-12 h-12 rounded-[1rem] bg-orange-50 flex items-center justify-center text-orange-500 text-lg shadow-inner">
              <i class="fa-solid fa-burger"></i>
            </div>
            <div>
              <h4 class="font-bold text-gray-800 text-sm">Shake Shack</h4>
              <p class="text-[10px] text-gray-400 mt-0.5">晚餐 • 幸福感满满</p>
            </div>
          </div>
          <div class="text-right">
            <span class="block font-bold text-gray-800">- 85.00</span>
            <span class="text-[10px] text-gray-400">招商银行</span>
          </div>
        </GlassCard>

        <GlassCard class="p-4 flex items-center justify-between cursor-pointer">
          <div class="flex items-center gap-4">
            <div class="w-12 h-12 rounded-[1rem] bg-purple-50 flex items-center justify-center text-purple-500 text-lg shadow-inner">
              <i class="fa-solid fa-film"></i>
            </div>
            <div>
              <h4 class="font-bold text-gray-800 text-sm">Netflix 订阅</h4>
              <p class="text-[10px] text-gray-400 mt-0.5">娱乐 • 月付</p>
            </div>
          </div>
          <div class="text-right">
            <span class="block font-bold text-gray-800">- 39.00</span>
            <span class="text-[10px] text-gray-400">VISA</span>
          </div>
        </GlassCard>
        
        <GlassCard class="p-4 flex items-center justify-between cursor-pointer">
          <div class="flex items-center gap-4">
            <div class="w-12 h-12 rounded-[1rem] bg-blue-50 flex items-center justify-center text-blue-500 text-lg shadow-inner">
              <i class="fa-solid fa-train-subway"></i>
            </div>
            <div>
              <h4 class="font-bold text-gray-800 text-sm">地铁通勤</h4>
              <p class="text-[10px] text-gray-400 mt-0.5">交通 • 08:30</p>
            </div>
          </div>
          <div class="text-right">
            <span class="block font-bold text-gray-800">- 5.00</span>
            <span class="text-[10px] text-gray-400">交通卡</span>
          </div>
        </GlassCard>
      </div>
    </div>
  </div>
</template>
