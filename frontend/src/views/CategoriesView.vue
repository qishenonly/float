<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useToast } from '../composables/useToast'

const router = useRouter()
const { showToast } = useToast()

const currentTab = ref('expense')

const expenseCategories = [
  { icon: 'fa-utensils', name: '餐饮美食', color: 'orange' },
  { icon: 'fa-bag-shopping', name: '购物消费', color: 'purple' },
  { icon: 'fa-bus', name: '交通出行', color: 'blue' },
  { icon: 'fa-house', name: '住房物业', color: 'green' }
]

const incomeCategories = [
  { icon: 'fa-sack-dollar', name: '工资薪水', color: 'indigo' },
  { icon: 'fa-arrow-trend-up', name: '理财收益', color: 'red' },
  { icon: 'fa-briefcase', name: '兼职外快', color: 'green' },
  { icon: 'fa-gift', name: '礼金红包', color: 'pink' }
]

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
      <h1 class="text-lg font-bold text-gray-800">分类管理</h1>
      <button @click="showFeature('添加分类')" class="w-10 h-10 rounded-full bg-white/50 flex items-center justify-center backdrop-blur-md text-gray-600 hover:bg-white transition active-press">
        <i class="fa-solid fa-plus"></i>
      </button>
    </div>

    <!-- Tab Switcher -->
    <div class="px-6 mb-6 relative z-10 animate-enter delay-100">
      <div class="glass-card p-1 rounded-2xl flex relative transition-colors">
        <button 
          @click="currentTab = 'expense'" 
          :class="currentTab === 'expense' ? 'bg-white text-indigo-600 shadow-md' : 'text-gray-500'"
          class="flex-1 py-2.5 rounded-xl text-sm font-bold transition-all duration-300"
        >
          支出
        </button>
        <button 
          @click="currentTab = 'income'" 
          :class="currentTab === 'income' ? 'bg-white text-indigo-600 shadow-md' : 'text-gray-500'"
          class="flex-1 py-2.5 rounded-xl text-sm font-bold transition-all duration-300"
        >
          收入
        </button>
      </div>
    </div>

    <!-- Category List -->
    <div class="px-6 pb-24 relative z-10">
      <p class="text-xs text-gray-400 mb-4 px-2">长按可拖拽排序</p>
      
      <!-- Expense Categories -->
      <div v-show="currentTab === 'expense'" class="space-y-3 animate-enter delay-200">
        <div 
          v-for="(category, index) in expenseCategories" 
          :key="'expense-' + index"
          class="glass-card p-3 rounded-2xl flex items-center justify-between group active:scale-[0.98] transition cursor-pointer"
        >
          <div class="flex items-center gap-4">
            <div class="cursor-move text-gray-300 px-1">
              <i class="fa-solid fa-grip-lines"></i>
            </div>
            <div :class="`w-10 h-10 rounded-xl bg-${category.color}-100 text-${category.color}-500 flex items-center justify-center shadow-inner`">
              <i :class="`fa-solid ${category.icon}`"></i>
            </div>
            <span class="font-bold text-gray-700">{{ category.name }}</span>
          </div>
          <div class="flex gap-2 opacity-0 group-hover:opacity-100 transition">
            <button @click="showFeature('编辑分类')" class="w-8 h-8 rounded-full bg-white text-gray-400 hover:text-blue-500 flex items-center justify-center shadow-sm active-press">
              <i class="fa-solid fa-pen text-xs"></i>
            </button>
            <button @click="showFeature('删除分类')" class="w-8 h-8 rounded-full bg-white text-gray-400 hover:text-red-500 flex items-center justify-center shadow-sm active-press">
              <i class="fa-solid fa-trash text-xs"></i>
            </button>
          </div>
        </div>
      </div>

      <!-- Income Categories -->
      <div v-show="currentTab === 'income'" class="space-y-3 animate-enter delay-200">
        <div 
          v-for="(category, index) in incomeCategories" 
          :key="'income-' + index"
          class="glass-card p-3 rounded-2xl flex items-center justify-between group active:scale-[0.98] transition cursor-pointer"
        >
          <div class="flex items-center gap-4">
            <div class="cursor-move text-gray-300 px-1">
              <i class="fa-solid fa-grip-lines"></i>
            </div>
            <div :class="`w-10 h-10 rounded-xl bg-${category.color}-100 text-${category.color}-500 flex items-center justify-center shadow-inner`">
              <i :class="`fa-solid ${category.icon}`"></i>
            </div>
            <span class="font-bold text-gray-700">{{ category.name }}</span>
          </div>
          <div class="flex gap-2 opacity-0 group-hover:opacity-100 transition">
            <button @click="showFeature('编辑分类')" class="w-8 h-8 rounded-full bg-white text-gray-400 hover:text-blue-500 flex items-center justify-center shadow-sm active-press">
              <i class="fa-solid fa-pen text-xs"></i>
            </button>
            <button @click="showFeature('删除分类')" class="w-8 h-8 rounded-full bg-white text-gray-400 hover:text-red-500 flex items-center justify-center shadow-sm active-press">
              <i class="fa-solid fa-trash text-xs"></i>
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
