<script setup>
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()
const type = ref('expense') // 'expense' or 'income'
const amount = ref('48.00')

const switchType = (newType) => {
  type.value = newType
  // Reset selection logic here if needed, for now just UI toggle
  if (newType === 'income') {
    selectedCategory.value = incomeCategories[0]
  } else {
    selectedCategory.value = expenseCategories[0]
  }
}

const expenseCategories = [
  { name: '餐饮', icon: 'fa-utensils', colorClass: 'text-orange-500', bgClass: 'bg-orange-100', borderClass: 'border-orange-500' },
  { name: '购物', icon: 'fa-bag-shopping', colorClass: 'text-gray-600', bgClass: 'bg-gray-50', borderClass: 'border-gray-600' }, // Simplified for inactive state
  { name: '交通', icon: 'fa-bus', colorClass: 'text-gray-600', bgClass: 'bg-gray-50', borderClass: 'border-gray-600' },
  { name: '居住', icon: 'fa-house', colorClass: 'text-gray-600', bgClass: 'bg-gray-50', borderClass: 'border-gray-600' },
]

const incomeCategories = [
  { name: '工资', icon: 'fa-sack-dollar', colorClass: 'text-green-600', bgClass: 'bg-green-100', borderClass: 'border-green-600' },
  { name: '理财', icon: 'fa-hand-holding-dollar', colorClass: 'text-gray-600', bgClass: 'bg-gray-50', borderClass: 'border-gray-600' },
  { name: '红包', icon: 'fa-gift', colorClass: 'text-gray-600', bgClass: 'bg-gray-50', borderClass: 'border-gray-600' },
]

const selectedCategory = ref(expenseCategories[0])

const selectCategory = (cat) => {
  selectedCategory.value = cat
}

// Helper to get active style for category
const getCategoryStyle = (cat) => {
  if (selectedCategory.value.name === cat.name) {
    // Active state reconstruction based on type
    if (type.value === 'expense') {
        return {
            bg: 'bg-orange-100',
            text: 'text-orange-500',
            border: 'border-2 border-orange-500'
        }
    } else {
         return {
            bg: 'bg-green-100',
            text: 'text-green-600',
            border: 'border-2 border-green-600'
        }
    }
  }
  return {
    bg: 'bg-gray-50',
    text: 'text-gray-600',
    border: ''
  }
}

</script>

<template>
  <div class="bg-white h-full flex flex-col">
    <!-- Header -->
    <div class="px-4 pt-10 pb-2 flex justify-between items-center bg-white z-10">
      <RouterLink to="/" class="w-8 h-8 flex items-center justify-center text-gray-400 hover:bg-gray-100 rounded-full">
        <i class="fa-solid fa-xmark text-lg"></i>
      </RouterLink>
      
      <div class="flex bg-gray-100 p-1 rounded-lg relative">
        <button 
          @click="switchType('expense')" 
          class="px-6 py-1.5 rounded text-sm transition-all duration-200"
          :class="type === 'expense' ? 'bg-white shadow-sm text-gray-800 font-bold' : 'text-gray-500 font-medium'"
        >
          支出
        </button>
        <button 
          @click="switchType('income')" 
          class="px-6 py-1.5 rounded text-sm transition-all duration-200"
          :class="type === 'income' ? 'bg-white shadow-sm text-gray-800 font-bold' : 'text-gray-500 font-medium'"
        >
          收入
        </button>
      </div>
      
      <div class="w-8"></div>
    </div>

    <!-- Amount Display -->
    <div class="px-6 py-8 flex flex-col items-end border-b border-gray-50">
      <div class="flex items-center gap-2 mb-2">
        <div 
          class="w-6 h-6 rounded-full flex items-center justify-center text-xs transition-colors"
          :class="type === 'expense' ? 'bg-orange-100 text-orange-500' : 'bg-green-100 text-green-600'"
        >
          <i class="fa-solid" :class="selectedCategory.icon"></i>
        </div>
        <span class="text-sm font-medium text-gray-500">{{ selectedCategory.name }}</span>
      </div>
      <div class="flex items-baseline gap-2 text-gray-800">
        <span class="text-2xl font-bold">¥</span>
        <h1 
          class="text-5xl font-bold transition-colors"
          :class="type === 'income' ? 'text-green-600' : 'text-gray-800'"
        >
          {{ amount }}<span class="text-blue-500 animate-pulse">|</span>
        </h1>
      </div>
    </div>

    <!-- Categories -->
    <div class="flex-1 overflow-y-auto p-4">
      <!-- Expense Categories -->
      <div v-if="type === 'expense'" class="grid grid-cols-4 gap-4 transition-opacity duration-300">
        <div 
          v-for="cat in expenseCategories" 
          :key="cat.name"
          class="flex flex-col items-center gap-2 cursor-pointer transition-opacity"
          :class="selectedCategory.name === cat.name ? 'opacity-100' : 'opacity-50 hover:opacity-100'"
          @click="selectCategory(cat)"
        >
          <div 
            class="w-14 h-14 rounded-2xl flex items-center justify-center text-xl transition-all"
            :class="selectedCategory.name === cat.name ? 'bg-orange-100 text-orange-500 border-2 border-orange-500' : 'bg-gray-50 text-gray-600'"
          >
            <i class="fa-solid" :class="cat.icon"></i>
          </div>
          <span class="text-xs font-medium" :class="selectedCategory.name === cat.name ? 'text-gray-800' : 'text-gray-500'">{{ cat.name }}</span>
        </div>
      </div>

      <!-- Income Categories -->
      <div v-else class="grid grid-cols-4 gap-4 transition-opacity duration-300">
        <div 
          v-for="cat in incomeCategories" 
          :key="cat.name"
          class="flex flex-col items-center gap-2 cursor-pointer transition-opacity"
          :class="selectedCategory.name === cat.name ? 'opacity-100' : 'opacity-50 hover:opacity-100'"
          @click="selectCategory(cat)"
        >
          <div 
            class="w-14 h-14 rounded-2xl flex items-center justify-center text-xl transition-all"
            :class="selectedCategory.name === cat.name ? 'bg-green-100 text-green-600 border-2 border-green-600' : 'bg-gray-50 text-gray-600'"
          >
            <i class="fa-solid" :class="cat.icon"></i>
          </div>
          <span class="text-xs font-medium" :class="selectedCategory.name === cat.name ? 'text-gray-800' : 'text-gray-500'">{{ cat.name }}</span>
        </div>
      </div>
    </div>

    <!-- Keypad -->
    <div class="bg-gray-50 p-2 pb-8">
      <div class="grid grid-cols-4 gap-2 h-64">
        <div class="col-span-3 grid grid-cols-3 gap-2">
          <button class="bg-white rounded-xl shadow-sm text-xl font-bold text-gray-700 hover:bg-gray-100 active:scale-95 transition">1</button>
          <button class="bg-white rounded-xl shadow-sm text-xl font-bold text-gray-700 hover:bg-gray-100 active:scale-95 transition">2</button>
          <button class="bg-white rounded-xl shadow-sm text-xl font-bold text-gray-700 hover:bg-gray-100 active:scale-95 transition">3</button>
          <button class="bg-white rounded-xl shadow-sm text-xl font-bold text-gray-700 hover:bg-gray-100 active:scale-95 transition">4</button>
          <button class="bg-white rounded-xl shadow-sm text-xl font-bold text-gray-700 hover:bg-gray-100 active:scale-95 transition">5</button>
          <button class="bg-white rounded-xl shadow-sm text-xl font-bold text-gray-700 hover:bg-gray-100 active:scale-95 transition">6</button>
          <button class="bg-white rounded-xl shadow-sm text-xl font-bold text-gray-700 hover:bg-gray-100 active:scale-95 transition">7</button>
          <button class="bg-white rounded-xl shadow-sm text-xl font-bold text-gray-700 hover:bg-gray-100 active:scale-95 transition">8</button>
          <button class="bg-white rounded-xl shadow-sm text-xl font-bold text-gray-700 hover:bg-gray-100 active:scale-95 transition">9</button>
          <button class="bg-white rounded-xl shadow-sm text-xl font-bold text-gray-700 hover:bg-gray-100 active:scale-95 transition">.</button>
          <button class="bg-white rounded-xl shadow-sm text-xl font-bold text-gray-700 hover:bg-gray-100 active:scale-95 transition">0</button>
          <button class="bg-white rounded-xl shadow-sm flex items-center justify-center text-gray-700 hover:bg-gray-100 active:scale-95 transition">
            <i class="fa-solid fa-delete-left"></i>
          </button>
        </div>
        <div class="col-span-1 grid grid-rows-2 gap-2">
          <button class="bg-gray-200 rounded-xl flex flex-col items-center justify-center text-gray-600 text-xs font-bold gap-1 hover:bg-gray-300 transition">
            <i class="fa-regular fa-calendar"></i>
            今天
          </button>
          <RouterLink 
            to="/" 
            class="rounded-xl flex items-center justify-center text-white text-lg font-bold shadow-lg hover:opacity-90 active:scale-95 transition"
            :class="type === 'income' ? 'bg-green-600' : 'bg-gray-800'"
          >
            完成
          </RouterLink>
        </div>
      </div>
    </div>
  </div>
</template>
