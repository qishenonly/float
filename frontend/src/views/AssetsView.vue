<script setup>
import GlassCard from '../components/GlassCard.vue'
</script>

<template>
  <div>
    <!-- Header -->
    <div class="px-6 pt-12 pb-4 relative z-10 animate-enter">
      <h1 class="text-2xl font-extrabold text-gray-800">我的资产</h1>
      <p class="text-xs text-gray-400 mt-1">财富积累是一个美好的过程</p>
    </div>

    <div class="px-6">
      <!-- Net Worth Card -->
      <div class="relative overflow-hidden rounded-[2rem] p-6 text-white shadow-xl shadow-indigo-200/50 mb-8 animate-enter delay-100 group">
        <div class="absolute inset-0 bg-gradient-to-br from-[#4e54c8] to-[#8f94fb] z-0"></div>
        <div class="absolute right-0 top-0 w-32 h-32 bg-white opacity-10 rounded-full blur-2xl transform translate-x-1/2 -translate-y-1/2"></div>
        
        <div class="relative z-10">
          <p class="text-indigo-100 text-xs font-medium mb-1 tracking-wide opacity-80">当前净资产 (CNY)</p>
          <h2 class="text-3xl font-bold mb-6 tracking-tight">¥ 124,500.00</h2>
          
          <div class="flex items-center gap-4">
            <div class="flex-1">
              <div class="w-full bg-black/20 h-1.5 rounded-full mb-2 overflow-hidden backdrop-blur-sm">
                <div class="bg-cyan-300 h-full w-[80%] shadow-[0_0_10px_rgba(103,232,249,0.5)]"></div>
              </div>
              <div class="flex justify-between text-[10px]">
                <span class="text-indigo-100/70">总资产</span>
                <span class="font-bold text-cyan-300">14.5万</span>
              </div>
            </div>
            <div class="flex-1">
              <div class="w-full bg-black/20 h-1.5 rounded-full mb-2 overflow-hidden backdrop-blur-sm">
                <div class="bg-pink-400 h-full w-[20%] shadow-[0_0_10px_rgba(244,114,182,0.5)]"></div>
              </div>
              <div class="flex justify-between text-[10px]">
                <span class="text-indigo-100/70">总负债</span>
                <span class="font-bold text-pink-300">2.05万</span>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Fund Accounts -->
      <div class="mb-6 animate-enter delay-200">
        <div class="flex justify-between items-center mb-4">
          <h3 class="text-sm font-bold text-gray-800">资金账户</h3>
          <button class="w-7 h-7 rounded-full bg-white shadow-sm flex items-center justify-center text-gray-500 hover:text-indigo-600 transition">
            <i class="fa-solid fa-plus text-xs"></i>
          </button>
        </div>
        
        <div class="space-y-3">
          <GlassCard class="p-4 rounded-2xl flex items-center justify-between shadow-sm hover:shadow-md transition cursor-pointer">
            <div class="flex items-center gap-4">
              <div class="w-10 h-10 rounded-xl bg-red-50 flex items-center justify-center shadow-inner">
                <i class="fa-solid fa-building-columns text-red-500"></i>
              </div>
              <div>
                <h4 class="font-bold text-gray-800 text-sm">招商银行</h4>
                <p class="text-[10px] text-gray-400 mt-0.5">储蓄卡 • 8899</p>
              </div>
            </div>
            <span class="font-bold text-gray-800">¥ 85,200</span>
          </GlassCard>

          <GlassCard class="p-4 rounded-2xl flex items-center justify-between shadow-sm hover:shadow-md transition cursor-pointer">
            <div class="flex items-center gap-4">
              <div class="w-10 h-10 rounded-xl bg-blue-50 flex items-center justify-center shadow-inner">
                <i class="fa-brands fa-alipay text-blue-500 text-lg"></i>
              </div>
              <div>
                <h4 class="font-bold text-gray-800 text-sm">支付宝</h4>
                <p class="text-[10px] text-gray-400 mt-0.5">余额宝</p>
              </div>
            </div>
            <span class="font-bold text-gray-800">¥ 42,300</span>
          </GlassCard>

           <GlassCard class="p-4 rounded-2xl flex items-center justify-between shadow-sm hover:shadow-md transition cursor-pointer">
            <div class="flex items-center gap-4">
              <div class="w-10 h-10 rounded-xl bg-green-50 flex items-center justify-center shadow-inner">
                <i class="fa-brands fa-weixin text-green-500 text-lg"></i>
              </div>
              <div>
                <h4 class="font-bold text-gray-800 text-sm">微信钱包</h4>
                <p class="text-[10px] text-gray-400 mt-0.5">零钱</p>
              </div>
            </div>
            <span class="font-bold text-gray-800">¥ 1,540</span>
          </GlassCard>
        </div>
      </div>

      <!-- Credit Accounts -->
      <div class="mb-6 animate-enter delay-300">
        <div class="flex justify-between items-center mb-4">
          <h3 class="text-sm font-bold text-gray-800">信用账户</h3>
        </div>
        
        <div class="space-y-3">
          <GlassCard class="p-4 rounded-2xl flex items-center justify-between shadow-sm hover:shadow-md transition cursor-pointer">
            <div class="flex items-center gap-4">
              <div class="w-10 h-10 rounded-xl bg-cyan-50 flex items-center justify-center shadow-inner">
                <i class="fa-solid fa-face-smile text-cyan-500"></i>
              </div>
              <div>
                <h4 class="font-bold text-gray-800 text-sm">蚂蚁花呗</h4>
                <p class="text-[10px] text-gray-400 mt-0.5">免息期 20天</p>
              </div>
            </div>
            <span class="font-bold text-gray-800">- 2,300</span>
          </GlassCard>

          <GlassCard class="p-4 rounded-2xl flex items-center justify-between shadow-sm hover:shadow-md transition cursor-pointer">
            <div class="flex items-center gap-4">
              <div class="w-10 h-10 rounded-xl bg-indigo-50 flex items-center justify-center shadow-inner">
                <i class="fa-brands fa-cc-visa text-indigo-500 text-lg"></i>
              </div>
              <div>
                <h4 class="font-bold text-gray-800 text-sm">Visa 信用卡</h4>
                <p class="text-[10px] text-gray-400 mt-0.5">账单日 5号</p>
              </div>
            </div>
            <span class="font-bold text-gray-800">- 18,200</span>
          </GlassCard>
        </div>
      </div>
    </div>
  </div>
</template>
