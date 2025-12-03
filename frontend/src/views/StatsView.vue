<script setup>
import GlassCard from '../components/GlassCard.vue'
</script>

<template>
  <div>
    <!-- Header -->
    <div class="px-6 pt-12 pb-2 flex justify-between items-center relative z-10 animate-enter">
      <div>
        <h1 class="text-2xl font-extrabold text-gray-800">收支统计</h1>
        <p class="text-xs text-gray-400 mt-1">每一笔花费都值得被记录</p>
      </div>
      <div class="glass-card px-3 py-1.5 rounded-full flex items-center gap-2 text-sm font-bold text-gray-700 shadow-sm cursor-pointer hover:bg-white/80 transition">
        <span>12月</span>
        <i class="fa-solid fa-chevron-down text-xs text-gray-400"></i>
      </div>
    </div>

    <div class="px-6 pt-6">
      <!-- Total Expenditure -->
      <div class="text-center mb-8 animate-enter delay-100">
        <p class="text-xs font-medium text-gray-400 tracking-wider uppercase mb-2">本月总支出</p>
        <h2 class="text-4xl font-extrabold text-gray-800">¥ 3,950.<span class="text-2xl text-gray-400 font-bold">00</span></h2>
        <div class="flex justify-center items-center gap-2 mt-2">
          <span class="bg-red-100 text-red-500 text-[10px] px-2 py-0.5 rounded-full font-bold">+12%</span>
          <span class="text-xs text-gray-400">对比上月</span>
        </div>
      </div>

      <!-- Weekly Trend Chart -->
      <GlassCard class="p-5 rounded-3xl shadow-sm mb-6 animate-enter delay-200">
        <div class="flex justify-between items-center mb-6">
          <h3 class="font-bold text-gray-800 text-sm">周消费趋势</h3>
          <i class="fa-solid fa-arrow-trend-up text-indigo-500 text-sm"></i>
        </div>
        
        <div class="flex justify-between items-end h-32 gap-3">
          <div class="flex flex-col items-center gap-2 w-full group cursor-pointer">
            <div class="w-full bg-indigo-100 rounded-t-lg relative h-full flex items-end overflow-hidden">
              <div class="w-full bg-indigo-300 h-[40%] rounded-t-lg bar-grow group-hover:bg-indigo-400 transition"></div>
            </div>
            <span class="text-[10px] text-gray-400 font-medium">Mon</span>
          </div>
          <div class="flex flex-col items-center gap-2 w-full group cursor-pointer">
             <div class="w-full bg-indigo-100 rounded-t-lg relative h-full flex items-end overflow-hidden">
              <div class="w-full bg-indigo-300 h-[60%] rounded-t-lg bar-grow group-hover:bg-indigo-400 transition" style="animation-delay: 0.1s"></div>
            </div>
            <span class="text-[10px] text-gray-400 font-medium">Tue</span>
          </div>
          <div class="flex flex-col items-center gap-2 w-full group cursor-pointer">
             <div class="w-full bg-indigo-100 rounded-t-lg relative h-full flex items-end overflow-hidden">
               <div class="w-full bg-gradient-to-t from-indigo-500 to-violet-500 h-[85%] rounded-t-lg bar-grow relative shadow-lg shadow-indigo-200" style="animation-delay: 0.2s">
                <div class="absolute -top-8 left-1/2 -translate-x-1/2 bg-gray-800 text-white text-[10px] px-2 py-1 rounded opacity-0 group-hover:opacity-100 transition whitespace-nowrap z-10">
                  ¥ 850
                </div>
              </div>
            </div>
            <span class="text-[10px] text-gray-800 font-bold">Wed</span>
          </div>
          <div class="flex flex-col items-center gap-2 w-full group cursor-pointer">
             <div class="w-full bg-indigo-100 rounded-t-lg relative h-full flex items-end overflow-hidden">
              <div class="w-full bg-indigo-300 h-[30%] rounded-t-lg bar-grow group-hover:bg-indigo-400 transition" style="animation-delay: 0.3s"></div>
            </div>
            <span class="text-[10px] text-gray-400 font-medium">Thu</span>
          </div>
          <div class="flex flex-col items-center gap-2 w-full group cursor-pointer">
             <div class="w-full bg-indigo-100 rounded-t-lg relative h-full flex items-end overflow-hidden">
              <div class="w-full bg-indigo-300 h-[50%] rounded-t-lg bar-grow group-hover:bg-indigo-400 transition" style="animation-delay: 0.4s"></div>
            </div>
            <span class="text-[10px] text-gray-400 font-medium">Fri</span>
          </div>
        </div>
      </GlassCard>

      <!-- Category Breakdown -->
      <div class="animate-enter delay-300">
        <h3 class="font-bold text-gray-800 text-sm mb-4">花哪里了?</h3>
        <div class="space-y-3">
          
          <GlassCard class="p-4 rounded-2xl flex items-center justify-between shadow-sm">
            <div class="flex items-center gap-3 w-full">
              <div class="w-10 h-10 rounded-xl bg-pink-100 text-pink-500 flex items-center justify-center shadow-inner">
                <i class="fa-solid fa-bag-shopping"></i>
              </div>
              <div class="flex-1">
                <div class="flex justify-between items-center mb-1">
                  <span class="text-sm font-bold text-gray-800">购物消费</span>
                  <span class="text-sm font-bold text-gray-800">¥ 1,200</span>
                </div>
                <div class="w-full bg-gray-200/50 h-1.5 rounded-full overflow-hidden">
                  <div class="bg-pink-500 h-full w-[45%] rounded-full"></div>
                </div>
              </div>
            </div>
          </GlassCard>

          <GlassCard class="p-4 rounded-2xl flex items-center justify-between shadow-sm">
            <div class="flex items-center gap-3 w-full">
              <div class="w-10 h-10 rounded-xl bg-orange-100 text-orange-500 flex items-center justify-center shadow-inner">
                <i class="fa-solid fa-utensils"></i>
              </div>
              <div class="flex-1">
                <div class="flex justify-between items-center mb-1">
                  <span class="text-sm font-bold text-gray-800">餐饮美食</span>
                  <span class="text-sm font-bold text-gray-800">¥ 800</span>
                </div>
                <div class="w-full bg-gray-200/50 h-1.5 rounded-full overflow-hidden">
                  <div class="bg-orange-500 h-full w-[30%] rounded-full"></div>
                </div>
              </div>
            </div>
          </GlassCard>

          <GlassCard class="p-4 rounded-2xl flex items-center justify-between shadow-sm">
            <div class="flex items-center gap-3 w-full">
              <div class="w-10 h-10 rounded-xl bg-blue-100 text-blue-500 flex items-center justify-center shadow-inner">
                <i class="fa-solid fa-bus"></i>
              </div>
              <div class="flex-1">
                <div class="flex justify-between items-center mb-1">
                  <span class="text-sm font-bold text-gray-800">交通出行</span>
                  <span class="text-sm font-bold text-gray-800">¥ 350</span>
                </div>
                <div class="w-full bg-gray-200/50 h-1.5 rounded-full overflow-hidden">
                  <div class="bg-blue-500 h-full w-[15%] rounded-full"></div>
                </div>
              </div>
            </div>
          </GlassCard>

        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
/* 柱状图生长动画 */
@keyframes growUp {
    from { height: 0; }
}
.bar-grow { animation: growUp 1s ease-out forwards; }
</style>
