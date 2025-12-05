<script setup>
import { ref, onMounted, watch } from 'vue'
import { useRouter } from 'vue-router'
import { categoryAPI } from '@/api/category'
import { useToast } from '../composables/useToast'

const router = useRouter()
const { showToast } = useToast()

const type = ref('expense') // 'expense' or 'income'
const amount = ref('')
const description = ref('')
const date = ref(new Date().toISOString().split('T')[0])
const loading = ref(false)

const expenseCategories = ref([])
const incomeCategories = ref([])
const selectedCategory = ref(null)

onMounted(() => {
  loadCategories()
})

const loadCategories = async () => {
  loading.value = true
  try {
    const [expenseRes, incomeRes] = await Promise.all([
      categoryAPI.getCategories('expense'),
      categoryAPI.getCategories('income')
    ])
    expenseCategories.value = expenseRes.data || []
    incomeCategories.value = incomeRes.data || []
    
    // Set default category
    if (type.value === 'expense' && expenseCategories.value.length > 0) {
      selectedCategory.value = expenseCategories.value[0]
    } else if (type.value === 'income' && incomeCategories.value.length > 0) {
      selectedCategory.value = incomeCategories.value[0]
    }
  } catch (error) {
    console.error('Failed to load categories:', error)
    showToast('加载分类失败', 'error')
  } finally {
    loading.value = false
  }
}

const switchType = (newType) => {
  type.value = newType
  if (newType === 'income') {
    if (incomeCategories.value.length > 0) selectedCategory.value = incomeCategories.value[0]
  } else {
    if (expenseCategories.value.length > 0) selectedCategory.value = expenseCategories.value[0]
  }
}

const selectCategory = (cat) => {
  selectedCategory.value = cat
}

// Just for UI demo, no save functionality yet as requested
const handleSave = () => {
  if (!amount.value) {
    showToast('请输入金额', 'warning')
    return
  }
  showToast('功能开发中...', 'info')
}
</script>

<template>
  <div class="bg-white h-full flex flex-col">
    <!-- Header -->
    <div class="px-4 pt-10 pb-2 flex justify-between items-center bg-white z-10">
      <RouterLink to="/" class="w-8 h-8 flex items-center justify-center text-gray-400 hover:bg-gray-100 rounded-full active-press">
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

    <!-- Amount Input -->
    <div class="px-6 py-8 flex flex-col items-end border-b border-gray-50">
      <div class="flex items-center gap-2 mb-2" v-if="selectedCategory">
        <div 
          class="w-6 h-6 rounded-full flex items-center justify-center text-xs transition-colors"
          :class="`bg-${selectedCategory.color}-100 text-${selectedCategory.color}-500`"
        >
          <i class="fa-solid" :class="selectedCategory.icon"></i>
        </div>
        <span class="text-sm font-medium text-gray-500">{{ selectedCategory.name }}</span>
      </div>
      <div class="flex items-center gap-2 text-gray-800 w-full justify-end">
        <span class="text-3xl font-bold">¥</span>
        <input 
          v-model="amount"
          type="number" 
          placeholder="0.00"
          class="text-5xl font-bold text-right w-full bg-transparent border-none outline-none placeholder-gray-200"
          :class="type === 'income' ? 'text-green-600' : 'text-gray-800'"
          autofocus
        >
      </div>
    </div>

    <!-- Categories -->
    <div class="flex-1 overflow-y-auto p-4">
      <div v-if="loading" class="flex justify-center py-8">
        <i class="fa-solid fa-spinner fa-spin text-gray-400"></i>
      </div>

      <!-- Expense Categories -->
      <div v-else-if="type === 'expense'" class="grid grid-cols-4 gap-4 transition-opacity duration-300">
        <div 
          v-for="cat in expenseCategories" 
          :key="cat.id"
          class="flex flex-col items-center gap-2 cursor-pointer transition-opacity"
          :class="selectedCategory?.id === cat.id ? 'opacity-100' : 'opacity-50 hover:opacity-100'"
          @click="selectCategory(cat)"
        >
          <div 
            class="w-14 h-14 rounded-2xl flex items-center justify-center text-xl transition-all shadow-sm"
            :class="selectedCategory?.id === cat.id ? `bg-${cat.color}-100 text-${cat.color}-500 ring-2 ring-${cat.color}-500 ring-offset-2` : 'bg-gray-50 text-gray-400'"
          >
            <i class="fa-solid" :class="cat.icon"></i>
          </div>
          <span class="text-xs font-medium" :class="selectedCategory?.id === cat.id ? 'text-gray-800' : 'text-gray-400'">{{ cat.name }}</span>
        </div>
      </div>

      <!-- Income Categories -->
      <div v-else class="grid grid-cols-4 gap-4 transition-opacity duration-300">
        <div 
          v-for="cat in incomeCategories" 
          :key="cat.id"
          class="flex flex-col items-center gap-2 cursor-pointer transition-opacity"
          :class="selectedCategory?.id === cat.id ? 'opacity-100' : 'opacity-50 hover:opacity-100'"
          @click="selectCategory(cat)"
        >
          <div 
            class="w-14 h-14 rounded-2xl flex items-center justify-center text-xl transition-all shadow-sm"
            :class="selectedCategory?.id === cat.id ? `bg-${cat.color}-100 text-${cat.color}-500 ring-2 ring-${cat.color}-500 ring-offset-2` : 'bg-gray-50 text-gray-400'"
          >
            <i class="fa-solid" :class="cat.icon"></i>
          </div>
          <span class="text-xs font-medium" :class="selectedCategory?.id === cat.id ? 'text-gray-800' : 'text-gray-400'">{{ cat.name }}</span>
        </div>
      </div>
    </div>

    <!-- Action Bar -->
    <div class="p-4 bg-white border-t border-gray-50 pb-8">
      <div class="flex items-center gap-4">
        <div class="flex-1 bg-gray-50 rounded-xl px-4 py-3 flex items-center gap-2">
          <i class="fa-regular fa-calendar text-gray-400"></i>
          <span class="text-sm text-gray-600">今天</span>
        </div>
        <button 
          @click="handleSave"
          class="flex-1 py-3 rounded-xl text-white font-bold shadow-lg hover:opacity-90 active:scale-95 transition"
          :class="type === 'income' ? 'bg-green-600 shadow-green-200' : 'bg-gray-800 shadow-gray-200'"
        >
          完成
        </button>
      </div>
    </div>
  </div>
</template>

<style scoped>
/* Hide number input arrows */
input[type=number]::-webkit-inner-spin-button, 
input[type=number]::-webkit-outer-spin-button { 
  -webkit-appearance: none; 
  margin: 0; 
}
</style>
