<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { transactionAPI } from '@/api/transaction'
import { accountAPI } from '@/api/account'
import { useToast } from '../composables/useToast'
import GlassCard from '../components/GlassCard.vue'

const router = useRouter()
const { showToast } = useToast()

const transactions = ref([])
const accounts = ref([])
const loading = ref(false)
const showFilterDrawer = ref(false)

// 筛选状态
const activeQuickFilter = ref('all') // 'all', 'month', 'expense', 'income', 'large'
const searchKeyword = ref('')

// 高级筛选状态
const advancedFilters = ref({
  startDate: '',
  endDate: '',
  type: 'all', // 'all', 'expense', 'income', 'transfer'
  accountId: null
})

// 统计数据
const statistics = ref({
  totalIncome: 0,
  totalExpense: 0
})

onMounted(() => {
  loadAccounts()
  loadTransactions()
})

const loadAccounts = async () => {
  try {
    const res = await accountAPI.getAccounts()
    accounts.value = res.data || []
  } catch (error) {
    console.error('Failed to load accounts:', error)
  }
}

const loadTransactions = async () => {
  loading.value = true
  try {
    const params = buildFilterParams()
    
    const res = await transactionAPI.getTransactions(params)
    transactions.value = res.data?.items || []
    
    // 计算统计
    calculateStatistics()
  } catch (error) {
    console.error('Failed to load transactions:', error)
    showToast('加载交易记录失败', 'error')
  } finally {
    loading.value = false
  }
}

const buildFilterParams = () => {
  const params = {
    page: 1,
    page_size: 1000 // 加载所有数据用于前端分组
  }

  // 快速筛选
  if (activeQuickFilter.value === 'month') {
    const firstDay = new Date()
    firstDay.setDate(1)
    params.start_date = firstDay.toISOString().split('T')[0]
    params.end_date = new Date().toISOString().split('T')[0]
  } else if (activeQuickFilter.value === 'expense') {
    params.type = 'expense'
  } else if (activeQuickFilter.value === 'income') {
    params.type = 'income'
  } else if (activeQuickFilter.value === 'large') {
    // 大额交易逻辑：前端筛选金额 > 1000
    // 先加载所有数据
  }

  // 高级筛选
  if (advancedFilters.value.startDate) {
    params.start_date = advancedFilters.value.startDate
  }
  if (advancedFilters.value.endDate) {
    params.end_date = advancedFilters.value.endDate
  }
  if (advancedFilters.value.type !== 'all') {
    params.type = advancedFilters.value.type
  }
  if (advancedFilters.value.accountId) {
    params.account_id = advancedFilters.value.accountId
  }

  // 搜索关键词
  if (searchKeyword.value.trim()) {
    params.search_keyword = searchKeyword.value.trim()
  }

  return params
}

const calculateStatistics = () => {
  let income = 0
  let expense = 0
  
  transactions.value.forEach(t => {
    if (t.type === 'income') income += t.amount
    else if (t.type === 'expense') expense += t.amount
  })
  
  statistics.value = {
    totalIncome: income,
    totalExpense: expense
  }
}

const setQuickFilter = (filter) => {
  activeQuickFilter.value = filter
  loadTransactions()
}

const handleSearch = () => {
  loadTransactions()
}

const toggleFilterDrawer = () => {
  showFilterDrawer.value = !showFilterDrawer.value
}

const applyAdvancedFilters = () => {
  showFilterDrawer.value = false
  loadTransactions()
}

const resetFilters = () => {
  advancedFilters.value = {
    startDate: '',
    endDate: '',
    type: 'all',
    accountId: null
  }
}

// 按日期分组
const groupedTransactions = computed(() => {
  let filtered = [...transactions.value]
  
  // 大额交易筛选
  if (activeQuickFilter.value === 'large') {
    filtered = filtered.filter(t => t.amount > 1000)
  }
  
  const groups = {}
  filtered.forEach(t => {
    // 从 transaction_date 或 created_at 提取日期部分（YYYY-MM-DD）
    let dateStr = t.transaction_date || t.created_at
    if (!dateStr) return
    
    // 提取日期部分
    const date = dateStr.split('T')[0]
    
    if (!groups[date]) {
      groups[date] = []
    }
    groups[date].push(t)
  })
  
  return Object.entries(groups).sort((a, b) => new Date(b[0]) - new Date(a[0]))
})

const formatDate = (dateStr) => {
  const date = new Date(dateStr + 'T00:00:00')
  const today = new Date()
  const yesterday = new Date(today)
  yesterday.setDate(yesterday.getDate() - 1)
  
  if (dateStr === today.toISOString().split('T')[0]) {
    return '今天'
  } else if (dateStr === yesterday.toISOString().split('T')[0]) {
    return '昨天'
  } else {
    return date.toLocaleDateString('zh-CN', { month: '2-digit', day: '2-digit' }).replace(/\//g, '月') + '日'
  }
}

const getDailySummary = (transactions) => {
  let total = 0
  transactions.forEach(t => {
    if (t.type === 'income') total += t.amount
    else if (t.type === 'expense') total -= t.amount
  })
  return total
}

const formatAmount = (amount) => {
  return amount.toFixed(2)
}

const getCategoryIcon = (category) => {
  if (!category) return 'fa-receipt'
  return category.icon || 'fa-receipt'
}

const getCategoryColor = (category) => {
  if (!category) return 'gray'
  return category.color || 'gray'
}

const formatTime = (transaction) => {
  if (transaction.transaction_time) {
    return transaction.transaction_time.substring(0, 5)
  }
  if (transaction.created_at) {
    const date = new Date(transaction.created_at)
    return date.toLocaleTimeString('zh-CN', { hour: '2-digit', minute: '2-digit', hour12: false })
  }
  return ''
}

const goBack = () => {
  router.push('/')
}

const goToExport = () => {
  router.push('/export')
}

const getIconClasses = (account) => {
  const brandAccounts = ['alipay', 'wechat']
  const isBrand = brandAccounts.includes(account?.account_type)
  return isBrand ? ['fa-brands', account.icon] : ['fa-solid', account.icon]
}

const setAdvancedFilterType = (type) => {
  advancedFilters.value.type = type
}

const setAdvancedFilterAccount = (accountId) => {
  advancedFilters.value.accountId = accountId
}
</script>

<template>
  <div class="flex-1 h-full bg-gradient-to-br from-[#eff6ff] via-[#fff1f2] to-[#f0f9ff] relative overflow-hidden flex flex-col">
    <!-- Background Blobs -->
    <div class="blob w-64 h-64 bg-indigo-200 rounded-full -top-10 -left-10 mix-blend-multiply"></div>
    <div class="blob w-64 h-64 bg-cyan-100 rounded-full bottom-20 -right-10 mix-blend-multiply"></div>

    <!-- Header -->
    <div class="px-6 pt-12 pb-2 relative z-10 flex flex-col gap-4 bg-white/30 backdrop-blur-md border-b border-white/50">
      <div class="flex items-center justify-between">
        <button @click="goBack" class="w-10 h-10 rounded-full bg-white/60 flex items-center justify-center text-gray-600 hover:bg-white transition shadow-sm">
          <i class="fa-solid fa-arrow-left"></i>
        </button>
        <h1 class="text-lg font-bold text-gray-800">全部账单</h1>
        <div class="w-10"></div>
      </div>

      <!-- Search and Filter -->
      <div class="flex gap-3 pb-4">
        <div class="flex-1 glass-card px-4 py-2.5 rounded-2xl flex items-center gap-2 text-gray-500 shadow-sm focus-within:bg-white transition">
          <i class="fa-solid fa-magnifying-glass text-sm"></i>
          <input 
            v-model="searchKeyword"
            @input="handleSearch"
            type="text" 
            class="bg-transparent w-full outline-none text-sm text-gray-700 placeholder-gray-400" 
            placeholder="搜索商家、备注..."
          >
        </div>
        <button 
          @click="toggleFilterDrawer" 
          class="w-11 h-11 rounded-2xl bg-[#1f2937] text-white flex items-center justify-center shadow-lg shadow-gray-300 hover:scale-105 transition active:scale-95"
        >
          <i class="fa-solid fa-sliders text-sm"></i>
        </button>
      </div>
      
      <!-- Quick Filter Chips -->
      <div class="flex gap-2 overflow-x-auto hide-scrollbar pb-2 -mx-6 px-6">
        <button 
          @click="setQuickFilter('all')"
          class="filter-chip whitespace-nowrap px-4 py-1.5 rounded-full text-xs font-bold bg-white text-gray-600 shadow-sm"
          :class="{ 'active': activeQuickFilter === 'all' }"
        >
          全部
        </button>
        <button 
          @click="setQuickFilter('month')"
          class="filter-chip whitespace-nowrap px-4 py-1.5 rounded-full text-xs font-bold bg-white text-gray-600 shadow-sm hover:bg-gray-50"
          :class="{ 'active': activeQuickFilter === 'month' }"
        >
          本月
        </button>
        <button 
          @click="setQuickFilter('expense')"
          class="filter-chip whitespace-nowrap px-4 py-1.5 rounded-full text-xs font-bold bg-white text-gray-600 shadow-sm hover:bg-gray-50"
          :class="{ 'active': activeQuickFilter === 'expense' }"
        >
          支出
        </button>
        <button 
          @click="setQuickFilter('income')"
          class="filter-chip whitespace-nowrap px-4 py-1.5 rounded-full text-xs font-bold bg-white text-gray-600 shadow-sm hover:bg-gray-50"
          :class="{ 'active': activeQuickFilter === 'income' }"
        >
          收入
        </button>
        <button 
          @click="setQuickFilter('large')"
          class="filter-chip whitespace-nowrap px-4 py-1.5 rounded-full text-xs font-bold bg-white text-gray-600 shadow-sm hover:bg-gray-50"
          :class="{ 'active': activeQuickFilter === 'large' }"
        >
          大额交易
        </button>
      </div>
    </div>

    <!-- Content -->
    <div class="flex-1 overflow-y-auto hide-scrollbar px-6 relative z-10 pb-20 pt-4">
      <!-- Statistics -->
      <div class="flex justify-between items-end mb-6 px-2 animate-enter">
        <div>
          <p class="text-[10px] text-gray-400 font-bold uppercase tracking-wider">本期统计</p>
          <div class="flex items-center gap-3 mt-1">
            <span class="text-xs text-gray-500">支出 <b class="text-gray-800">¥ {{ formatAmount(statistics.totalExpense) }}</b></span>
            <span class="text-xs text-gray-500">收入 <b class="text-gray-800">¥ {{ formatAmount(statistics.totalIncome) }}</b></span>
          </div>
        </div>
        <button @click="goToExport" class="text-[10px] font-bold text-indigo-500 flex items-center gap-1 bg-indigo-50 px-2 py-1 rounded-lg">
          <i class="fa-solid fa-file-export"></i> 导出
        </button>
      </div>

      <!-- Loading State -->
      <div v-if="loading" class="py-12 text-center">
        <i class="fa-solid fa-spinner fa-spin text-gray-400 text-2xl"></i>
        <p class="text-gray-400 text-sm mt-2">加载中...</p>
      </div>

      <!-- Empty State -->
      <div v-else-if="groupedTransactions.length === 0" class="py-12 text-center">
        <i class="fa-solid fa-inbox text-gray-300 text-4xl"></i>
        <p class="text-gray-400 text-sm mt-2">暂无交易记录</p>
      </div>

      <!-- Transaction Groups -->
      <div v-else class="space-y-6">
        <div 
          v-for="([date, dayTransactions], index) in groupedTransactions" 
          :key="date"
          class="mb-6 animate-enter"
          :style="{ animationDelay: `${index * 0.1}s` }"
        >
          <h3 class="text-xs font-bold text-gray-400 mb-3 px-1 flex justify-between">
            <span>{{ formatDate(date) }}</span>
            <span>{{ getDailySummary(dayTransactions) >= 0 ? '+' : '' }}{{ formatAmount(getDailySummary(dayTransactions)) }}</span>
          </h3>
          <div class="space-y-3">
            <GlassCard 
              v-for="transaction in dayTransactions" 
              :key="transaction.id"
              class="p-4 flex items-center justify-between shadow-sm hover:shadow-md transition cursor-pointer active:scale-[0.98]"
            >
              <div class="flex items-center gap-4">
                <div 
                  class="w-11 h-11 rounded-[1rem] flex items-center justify-center text-lg shadow-inner"
                  :class="`bg-${getCategoryColor(transaction.category)}-50 text-${getCategoryColor(transaction.category)}-500`"
                >
                  <i class="fa-solid" :class="getCategoryIcon(transaction.category)"></i>
                </div>
                <div>
                  <h4 class="font-bold text-gray-800 text-sm">{{ transaction.description || transaction.title || transaction.category?.name || '未分类' }}</h4>
                  <div class="flex items-center gap-1.5 mt-0.5">
                    <span class="text-[10px] text-gray-400">{{ formatTime(transaction) }} · {{ transaction.category?.name || '未分类' }}</span>
                    <span v-if="transaction.tags && transaction.tags.length > 0" class="w-1 h-1 rounded-full bg-gray-300"></span>
                    <span v-if="transaction.tags && transaction.tags.length > 0" class="text-[10px] text-gray-400">#{{ transaction.tags[0] }}</span>
                  </div>
                </div>
              </div>
              <div class="text-right">
                <span 
                  class="block font-bold"
                  :class="transaction.type === 'income' ? 'text-green-600' : 'text-gray-800'"
                >
                  {{ transaction.type === 'income' ? '+' : '-' }} {{ formatAmount(transaction.amount) }}
                </span>
                <span class="text-[10px] text-gray-400">{{ transaction.account?.account_name || '未知账户' }}</span>
              </div>
            </GlassCard>
          </div>
        </div>
      </div>

      <p class="text-center text-[10px] text-gray-300 py-4">到底啦 ~</p>
    </div>

    <!-- Filter Drawer Overlay -->
    <div 
      v-show="showFilterDrawer"
      @click="toggleFilterDrawer" 
      class="drawer-overlay absolute inset-0 bg-black/20 backdrop-blur-sm z-[100]"
      :class="{ 'open': showFilterDrawer }"
    ></div>

    <!-- Filter Drawer -->
    <div 
      class="bottom-drawer absolute bottom-0 left-0 right-0 bg-white rounded-t-[2.5rem] p-6 pb-20 z-[101] shadow-[0_-10px_40px_rgba(0,0,0,0.1)] max-h-[75vh]"
      :class="{ 'open': showFilterDrawer }"
    >
      <div class="w-12 h-1 bg-gray-200 rounded-full mx-auto mb-6"></div>
      
      <div class="flex justify-between items-center mb-6">
        <h3 class="text-lg font-bold text-gray-800">高级筛选</h3>
        <span @click="resetFilters" class="text-xs text-gray-400 cursor-pointer">重置</span>
      </div>

      <div class="space-y-6 max-h-[45vh] overflow-y-auto pb-4">
        <!-- Date Range -->
        <div>
          <p class="text-[10px] text-gray-400 font-bold uppercase tracking-widest mb-3">日期范围</p>
          <div class="flex gap-3">
            <input 
              v-model="advancedFilters.startDate"
              type="date"
              class="glass-card flex-1 p-3 rounded-xl text-center text-sm font-bold text-gray-700 bg-gray-50 border-transparent outline-none"
            >
            <span class="text-gray-300 self-center">-</span>
            <input 
              v-model="advancedFilters.endDate"
              type="date"
              class="glass-card flex-1 p-3 rounded-xl text-center text-sm font-bold text-gray-700 bg-gray-50 border-transparent outline-none"
            >
          </div>
        </div>

        <!-- Transaction Type -->
        <div>
          <p class="text-[10px] text-gray-400 font-bold uppercase tracking-widest mb-3">交易类型</p>
          <div class="grid grid-cols-3 gap-3">
            <button 
              @click="setAdvancedFilterType('all')"
              class="py-2.5 rounded-xl text-xs font-bold transition"
              :class="advancedFilters.type === 'all' ? 'bg-gray-900 text-white shadow-md' : 'bg-white text-gray-600 border border-gray-100'"
            >
              全部
            </button>
            <button 
              @click="setAdvancedFilterType('expense')"
              class="py-2.5 rounded-xl text-xs font-bold transition"
              :class="advancedFilters.type === 'expense' ? 'bg-gray-900 text-white shadow-md' : 'bg-white text-gray-600 border border-gray-100'"
            >
              支出
            </button>
            <button 
              @click="setAdvancedFilterType('income')"
              class="py-2.5 rounded-xl text-xs font-bold transition"
              :class="advancedFilters.type === 'income' ? 'bg-gray-900 text-white shadow-md' : 'bg-white text-gray-600 border border-gray-100'"
            >
              收入
            </button>
          </div>
        </div>

        <!-- Accounts -->
        <div>
          <p class="text-[10px] text-gray-400 font-bold uppercase tracking-widest mb-3">账户</p>
          <div class="flex gap-3 overflow-x-auto hide-scrollbar pb-1">
            <button 
              @click="setAdvancedFilterAccount(null)"
              class="flex items-center gap-2 px-4 py-2 rounded-xl text-xs font-bold whitespace-nowrap transition"
              :class="advancedFilters.accountId === null ? 'bg-gray-900 text-white shadow-md' : 'bg-white border border-gray-100 text-gray-600'"
            >
              全部账户
            </button>
            <button 
              v-for="acc in accounts"
              :key="acc.id"
              @click="setAdvancedFilterAccount(acc.id)"
              class="flex items-center gap-2 px-4 py-2 rounded-xl text-xs font-bold whitespace-nowrap transition"
              :class="advancedFilters.accountId === acc.id ? 'bg-gray-900 text-white shadow-md' : 'bg-white border border-gray-100 text-gray-600'"
            >
              <i :class="getIconClasses(acc)" :style="{ color: advancedFilters.accountId === acc.id ? 'white' : undefined }"></i>
              {{ acc.account_name }}
            </button>
          </div>
        </div>
      </div>

      <button 
        @click="applyAdvancedFilters" 
        class="w-full bg-gray-900 text-white font-bold py-4 rounded-2xl shadow-xl shadow-gray-200 mt-8 active:scale-[0.98] transition"
      >
        确认筛选 ({{ transactions.length }})
      </button>
    </div>
  </div>
</template>

<style scoped>
.hide-scrollbar::-webkit-scrollbar {
  display: none;
}

/* 核心玻璃拟态样式 */
.glass-card {
  background: rgba(255, 255, 255, 0.65);
  backdrop-filter: blur(20px);
  -webkit-backdrop-filter: blur(20px);
  border: 1px solid rgba(255, 255, 255, 0.8);
  box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.02);
}

/* 筛选胶囊样式 */
.filter-chip {
  transition: all 0.2s;
  border: 1px solid transparent;
}
.filter-chip.active {
  background-color: #1f2937;
  color: white;
  box-shadow: 0 4px 10px rgba(31, 41, 55, 0.2);
}

/* 底部抽屉动画 */
.drawer-overlay { 
  transition: opacity 0.3s; 
  pointer-events: none; 
  opacity: 0; 
}
.drawer-overlay.open { 
  pointer-events: auto; 
  opacity: 1; 
}

.bottom-drawer { 
  transition: transform 0.3s cubic-bezier(0.16, 1, 0.3, 1); 
  transform: translateY(100%); 
}
.bottom-drawer.open { 
  transform: translateY(0); 
}

/* 进场动画 */
@keyframes fadeInUp { 
  from { 
    opacity: 0; 
    transform: translateY(20px); 
  } 
  to { 
    opacity: 1; 
    transform: translateY(0); 
  } 
}
.animate-enter { 
  animation: fadeInUp 0.5s cubic-bezier(0.16, 1, 0.3, 1) forwards; 
  opacity: 0; 
}

.blob { 
  position: absolute; 
  filter: blur(50px); 
  z-index: 0; 
  opacity: 0.6; 
}
</style>
