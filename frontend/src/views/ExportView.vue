<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useToast } from '../composables/useToast'

const router = useRouter()
const { showToast } = useToast()

const selectedFormat = ref('excel')
const email = ref('alex@example.com')

const handleExport = () => {
  showToast('导出功能正在开发中...', 'info')
}
</script>

<template>
  <div>
    <!-- Header -->
    <div class="px-6 pt-12 pb-4 relative z-10 flex items-center justify-between animate-enter">
      <button @click="router.back()" class="w-10 h-10 rounded-full bg-white/50 flex items-center justify-center backdrop-blur-md text-gray-600 hover:bg-white transition active-press">
        <i class="fa-solid fa-arrow-left"></i>
      </button>
      <h1 class="text-lg font-bold text-gray-800">导出数据</h1>
      <div class="w-10"></div>
    </div>

    <div class="flex-1 overflow-y-auto px-6 pt-4 relative z-10 pb-32">
      
      <!-- Time Range -->
      <h3 class="font-bold text-gray-800 mb-4 px-1 animate-enter delay-100">时间范围</h3>
      <div class="glass-card p-4 rounded-2xl mb-8 flex items-center justify-between animate-enter delay-200">
        <div class="text-center flex-1">
          <p class="text-xs text-gray-400 mb-1">开始日期</p>
          <p class="font-bold text-gray-800">2023-11-01</p>
        </div>
        <div class="text-gray-300">
          <i class="fa-solid fa-arrow-right"></i>
        </div>
        <div class="text-center flex-1">
          <p class="text-xs text-gray-400 mb-1">结束日期</p>
          <p class="font-bold text-gray-800">2023-12-03</p>
        </div>
      </div>

      <!-- Export Format -->
      <h3 class="font-bold text-gray-800 mb-4 px-1 animate-enter delay-300">导出格式</h3>
      <div class="grid grid-cols-2 gap-4 mb-8 animate-enter delay-400">
        <label 
          @click="selectedFormat = 'excel'"
          :class="selectedFormat === 'excel' ? 'border-indigo-600 bg-indigo-50' : 'border-transparent'"
          class="cursor-pointer relative group glass-card p-4 rounded-2xl flex flex-col items-center gap-3 border-2 transition-all hover:bg-white/80 h-32 justify-center active-press"
        >
          <div class="w-12 h-12 bg-green-100 rounded-full flex items-center justify-center text-green-600 text-xl">
            <i class="fa-solid fa-file-excel"></i>
          </div>
          <span class="font-bold text-gray-700 text-sm">Excel 表格</span>
          <div v-if="selectedFormat === 'excel'" class="absolute top-3 right-3 text-indigo-600">
            <i class="fa-solid fa-circle-check"></i>
          </div>
        </label>

        <label 
          @click="selectedFormat = 'pdf'"
          :class="selectedFormat === 'pdf' ? 'border-indigo-600 bg-indigo-50' : 'border-transparent'"
          class="cursor-pointer relative group glass-card p-4 rounded-2xl flex flex-col items-center gap-3 border-2 transition-all hover:bg-white/80 h-32 justify-center active-press"
        >
          <div class="w-12 h-12 bg-red-100 rounded-full flex items-center justify-center text-red-600 text-xl">
            <i class="fa-solid fa-file-pdf"></i>
          </div>
          <span class="font-bold text-gray-700 text-sm">PDF 文档</span>
          <div v-if="selectedFormat === 'pdf'" class="absolute top-3 right-3 text-indigo-600">
            <i class="fa-solid fa-circle-check"></i>
          </div>
        </label>
      </div>

      <!-- Email Input -->
      <h3 class="font-bold text-gray-800 mb-4 px-1 animate-enter delay-500">接收邮箱</h3>
      <div class="glass-card px-4 py-3 rounded-2xl flex items-center gap-3 mb-8 animate-enter delay-600">
        <i class="fa-regular fa-envelope text-gray-400"></i>
        <input 
          v-model="email"
          type="email" 
          class="bg-transparent w-full outline-none text-gray-700 font-medium placeholder-gray-400" 
          placeholder="输入您的邮箱"
        >
      </div>
    </div>

    <!-- Export Button -->
    <div class="px-6 pb-8 pt-4 bg-white/50 backdrop-blur-md relative z-20">
      <button @click="handleExport" class="w-full bg-gradient-to-r from-gray-800 to-gray-900 text-white font-bold py-4 rounded-2xl shadow-lg hover:shadow-xl hover:scale-[1.02] transition active:scale-[0.98]">
        确认导出
      </button>
    </div>
  </div>
</template>
