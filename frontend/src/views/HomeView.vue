<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { useToast } from '../composables/useToast'
import { transactionAPI } from '@/api/transaction'
import { accountAPI } from '@/api/account'
import GlassCard from '../components/GlassCard.vue'

const router = useRouter()
const authStore = useAuthStore()
const { showToast } = useToast()

const user = computed(() => authStore.currentUser)
const displayName = computed(() => user.value?.display_name || user.value?.username || 'User')
const avatarSeed = computed(() => user.value?.username || 'User')

// 今日交易数据
const todayTransactions = ref([])
const statistics = ref({
  totalIncome: 0,
  totalExpense: 0,
  netAmount: 0
})
const totalBalance = ref(0)

// 获取今天的日期
const getTodayDate = () => {
  const today = new Date()
  return today.toISOString().split('T')[0]
}

// 加载数据
const loadData = async () => {
  try {
    const today = getTodayDate()
    
    // 获取今日交易列表
    const transRes = await transactionAPI.getTransactions({
      start_date: today,
      end_date: today,
      page: 1,
      page_size: 20
    })
    todayTransactions.value = transRes.data?.items || []
    
    // 获取本月统计（用于大屏展示）
    const firstDay = new Date()
    firstDay.setDate(1)
    const startDate = firstDay.toISOString().split('T')[0]
    const statsRes = await transactionAPI.getStatistics({
      start_date: startDate,
      end_date: today
    })
    if (statsRes.data) {
      const income = statsRes.data?.total_income || 0
      const expense = statsRes.data?.total_expense || 0
      statistics.value = {
        totalIncome: income,
        totalExpense: expense,
        netAmount: income - expense
      }
    }
    
    // 获取账户总余额
    const accountsRes = await accountAPI.getAccounts()
    if (accountsRes.data) {
      totalBalance.value = accountsRes.data.reduce((sum, acc) => sum + (acc.balance || 0), 0)
    }
  } catch (error) {
    console.error('Failed to load data:', error)
    showToast('数据加载失败', 'error')
  }
}

onMounted(() => {
  loadData()
})

const formatAmount = (amount) => {
  return amount.toFixed(2)
}

const getCategoryIcon = (category) => {
  if (!category) return 'fa-receipt'
  return category.icon || 'fa-receipt'
}

const getCategoryColor = (category) => {
  if (!category) return 'gray'
  const colorMap = {
    'orange': 'orange',
    'purple': 'purple',
    'blue': 'blue',
    'green': 'green',
    'red': 'red',
    'pink': 'pink'
  }
  return colorMap[category.color] || 'gray'
}

const getAccountName = (transaction) => {
  return transaction.account?.account_name || '未知账户'
}

const getBudgetPercentage = () => {
  if (statistics.value.totalIncome === 0) return 0
  return Math.min((statistics.value.totalExpense / statistics.value.totalIncome * 100), 100)
}

const formatDateTime = (createdAt) => {
  if (!createdAt) return ''
  const date = new Date(createdAt)
  return date.toLocaleTimeString('zh-CN', { hour: '2-digit', minute: '2-digit', hour12: false })
}

const openTransactionHistory = () => {
  router.push('/transactions')
}
</script>

<template>
  <div>
    <!-- Header -->
    <div class="px-6 pt-10 pb-2 flex justify-between items-center animate-enter">
      <div>
        <div class="flex items-center gap-2 text-gray-500 mb-1">
          <span class="text-xs font-medium tracking-wider uppercase">{{ new Date().toLocaleDateString('zh-CN', { weekday: 'long', month: 'short', day: 'numeric' }) }}</span>
          <i class="fa-solid fa-cloud-sun text-yellow-400"></i>
        </div>
        <h1 class="text-2xl font-extrabold text-gray-800 tracking-tight">Hi, {{ displayName }} 👋</h1>
        <p class="text-xs text-gray-400 mt-1">今天也是充满希望的一天 ✨</p>
      </div>
      <RouterLink to="/profile" class="w-11 h-11 rounded-full p-1 bg-white shadow-sm cursor-pointer hover:scale-105 transition active-press">
        <img :src="`https://api.dicebear.com/7.x/avataaars/svg?seed=${avatarSeed}&backgroundColor=ffdfbf`" class="w-full h-full rounded-full" alt="avatar">
      </RouterLink>
    </div>

    <!-- Balance Card -->
    <div class="px-6 mt-6 animate-enter delay-100">
      <div class="relative overflow-hidden rounded-[2rem] p-6 text-white shadow-xl shadow-indigo-200/50 group transition-transform hover:scale-[1.02] duration-300 active-press">
        <div class="absolute inset-0 bg-gradient-to-tr from-violet-600 to-indigo-500 z-0"></div>
        <div class="absolute -right-10 -top-10 w-40 h-40 bg-white opacity-10 rounded-full blur-2xl"></div>
        <div class="absolute -left-10 -bottom-10 w-32 h-32 bg-purple-400 opacity-20 rounded-full blur-xl"></div>

        <div class="relative z-10">
          <div class="flex justify-between items-start mb-6">
            <div>
              <p class="text-indigo-100 text-xs font-medium mb-1 tracking-wide opacity-80">账户总余额</p>
              <div class="flex items-baseline gap-2">
                <h2 class="text-3xl font-bold tracking-tight">¥ {{ formatAmount(totalBalance) }}</h2>
                <button class="text-indigo-200 hover:text-white transition active-press"><i class="fa-solid fa-eye-slash text-sm"></i></button>
              </div>
            </div>
            <div class="bg-white/20 backdrop-blur-md p-2 rounded-xl">
              <i class="fa-solid fa-chart-simple text-white"></i>
            </div>
          </div>

          <div class="mb-4">
            <div class="flex justify-between text-[10px] text-indigo-100 mb-1.5 px-0.5">
              <span>支出进度</span>
              <span>已用 {{ getBudgetPercentage().toFixed(0) }}%</span>
            </div>
            <div class="w-full bg-black/20 h-1.5 rounded-full overflow-hidden backdrop-blur-sm">
              <div class="bg-white h-full rounded-full shadow-[0_0_10px_rgba(255,255,255,0.5)]" :style="{ width: getBudgetPercentage() + '%' }"></div>
            </div>
          </div>

          <div class="flex gap-3 mt-4">
            <div class="bg-indigo-900/30 backdrop-blur-md rounded-xl px-3 py-2 flex-1 border border-white/10">
              <div class="flex items-center gap-2 mb-1">
                <div class="w-1.5 h-1.5 rounded-full bg-green-400"></div>
                <span class="text-[10px] text-indigo-100">本月收入</span>
              </div>
              <span class="text-sm font-semibold">{{ formatAmount(statistics.totalIncome) }}</span>
            </div>
            <div class="bg-indigo-900/30 backdrop-blur-md rounded-xl px-3 py-2 flex-1 border border-white/10">
              <div class="flex items-center gap-2 mb-1">
                <div class="w-1.5 h-1.5 rounded-full bg-pink-400"></div>
                <span class="text-[10px] text-indigo-100">本月支出</span>
              </div>
              <span class="text-sm font-semibold">{{ formatAmount(statistics.totalExpense) }}</span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Quick Actions -->
    <div class="px-6 mt-6 grid grid-cols-4 gap-4 animate-enter delay-200">
      <div @click="showToast('心愿单功能开发中...', 'info')" class="flex flex-col items-center gap-2 cursor-pointer group active-press">
        <div class="w-14 h-14 rounded-2xl bg-white shadow-sm flex items-center justify-center text-xl text-pink-500 group-hover:shadow-md group-hover:-translate-y-1 transition duration-300">
          <i class="fa-solid fa-gift"></i>
        </div>
        <span class="text-[10px] font-medium text-gray-500">心愿单</span>
      </div>
      <div @click="showToast('日历功能开发中...', 'info')" class="flex flex-col items-center gap-2 cursor-pointer group active-press">
        <div class="w-14 h-14 rounded-2xl bg-white shadow-sm flex items-center justify-center text-xl text-blue-500 group-hover:shadow-md group-hover:-translate-y-1 transition duration-300">
          <i class="fa-solid fa-calendar-days"></i>
        </div>
        <span class="text-[10px] font-medium text-gray-500">日历</span>
      </div>
      <div @click="showToast('存钱功能开发中...', 'info')" class="flex flex-col items-center gap-2 cursor-pointer group active-press">
        <div class="w-14 h-14 rounded-2xl bg-white shadow-sm flex items-center justify-center text-xl text-orange-500 group-hover:shadow-md group-hover:-translate-y-1 transition duration-300">
          <i class="fa-solid fa-piggy-bank"></i>
        </div>
        <span class="text-[10px] font-medium text-gray-500">存钱</span>
      </div>
      <div @click="showToast('账单功能开发中...', 'info')" class="flex flex-col items-center gap-2 cursor-pointer group active-press">
        <div class="w-14 h-14 rounded-2xl bg-white shadow-sm flex items-center justify-center text-xl text-indigo-500 group-hover:shadow-md group-hover:-translate-y-1 transition duration-300">
          <i class="fa-solid fa-file-invoice"></i>
        </div>
        <span class="text-[10px] font-medium text-gray-500">账单</span>
      </div>
    </div>

    <!-- Recent Transactions -->
    <div class="px-6 mt-8 pb-32 animate-enter delay-300">
      <div class="flex justify-between items-center mb-4">
        <h3 class="text-lg font-bold text-gray-800">今日动态</h3>
        <button @click="openTransactionHistory()" class="w-8 h-8 rounded-full bg-white flex items-center justify-center text-gray-400 shadow-sm hover:text-gray-800 active-press">
          <i class="fa-solid fa-ellipsis"></i>
        </button>
      </div>

      <div class="space-y-4">
        <div v-if="todayTransactions.length === 0" class="text-center py-8">
          <p class="text-gray-400 text-sm">今天还没有交易记录</p>
        </div>
        
        <GlassCard 
          v-for="transaction in todayTransactions" 
          :key="transaction.id"
          class="p-4 flex items-center justify-between cursor-pointer"
        >
          <div class="flex items-center gap-4">
            <div 
              class="w-12 h-12 rounded-[1rem] flex items-center justify-center text-lg shadow-inner"
              :class="`bg-${getCategoryColor(transaction.category)}-50 text-${getCategoryColor(transaction.category)}-500`"
            >
              <i class="fa-solid" :class="getCategoryIcon(transaction.category)"></i>
            </div>
            <div>
              <h4 class="font-bold text-gray-800 text-sm">{{ transaction.description || transaction.title }}</h4>
              <p class="text-[10px] text-gray-400 mt-0.5">{{ transaction.category?.name || '未分类' }} • {{ formatDateTime(transaction.created_at) }}</p>
            </div>
          </div>
          <div class="text-right">
            <span class="block font-bold text-gray-800" :class="transaction.type === 'income' ? 'text-green-600' : ''">
              {{ transaction.type === 'income' ? '+' : '-' }} {{ formatAmount(transaction.amount) }}
            </span>
            <span class="text-[10px] text-gray-400">{{ getAccountName(transaction) }}</span>
          </div>
        </GlassCard>
      </div>
    </div>
  </div>
</template>
