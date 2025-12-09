<script setup>
import { ref, computed, onMounted } from 'vue'
import { transactionAPI } from '@/api/transaction'
import { useToast } from '../composables/useToast'
import GlassCard from '../components/GlassCard.vue'

const { showToast } = useToast()

const transactions = ref([])
const loading = ref(false)
const page = ref(1)
const pageSize = ref(30)
const total = ref(0)

const filters = ref({
  type: 'all',
  search_keyword: ''
})

onMounted(() => {
  loadTransactions()
})

const loadTransactions = async () => {
  loading.value = true
  try {
    const params = {
      page: page.value,
      page_size: pageSize.value
    }

    if (filters.value.type !== 'all') {
      params.type = filters.value.type
    }
    if (filters.value.search_keyword) {
      params.search_keyword = filters.value.search_keyword
    }

    const res = await transactionAPI.getTransactions(params)
    transactions.value = res.data?.items || []
    total.value = res.data?.total || 0
  } catch (error) {
    console.error('Failed to load transactions:', error)
    showToast('加载交易记录失败', 'error')
  } finally {
    loading.value = false
  }
}

const formatCurrency = (amount, type) => {
  const sign = type === 'expense' ? '-' : type === 'income' ? '+' : ''
  return `${sign}¥${amount.toFixed(2)}`
}

const getTypeColor = (type) => {
  switch (type) {
    case 'expense':
      return 'text-red-500'
    case 'income':
      return 'text-green-500'
    case 'transfer':
      return 'text-blue-500'
    default:
      return 'text-gray-500'
  }
}

const handleDelete = async (id) => {
  if (!confirm('确定删除此交易吗？')) return
  
  try {
    await transactionAPI.deleteTransaction(id)
    showToast('交易已删除', 'success')
    await loadTransactions()
  } catch (error) {
    console.error('Failed to delete transaction:', error)
    showToast('删除失败', 'error')
  }
}

const handleFilterChange = () => {
  page.value = 1
  loadTransactions()
}

const totalPages = computed(() => Math.ceil(total.value / pageSize.value))

const prevPage = () => {
  if (page.value > 1) {
    page.value--
    loadTransactions()
  }
}

const nextPage = () => {
  if (page.value < totalPages.value) {
    page.value++
    loadTransactions()
  }
}

const groupedTransactions = computed(() => {
  const groups = {}
  transactions.value.forEach(t => {
    const date = t.transaction_date
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
    return date.toLocaleDateString('zh-CN', { month: '2-digit', day: '2-digit', weekday: 'short' })
  }
}

const getDailySummary = (transactions) => {
  const summary = { income: 0, expense: 0 }
  transactions.forEach(t => {
    if (t.type === 'income') summary.income += t.amount
    else if (t.type === 'expense') summary.expense += t.amount
  })
  return summary
}
</script>

<template>
  <div>
    <!-- Header -->
    <div class="px-6 pt-10 pb-4 relative z-10 animate-enter">
      <h1 class="text-2xl font-extrabold text-gray-800">交易记录</h1>
      <p class="text-xs text-gray-400 mt-1">共 {{ total }} 条交易</p>
    </div>

    <!-- Filters -->
    <div class="px-6 pb-6 space-y-3 animate-enter delay-100">
      <select
        v-model="filters.type"
        @change="handleFilterChange"
        class="w-full px-4 py-3 text-sm bg-white rounded-2xl border border-gray-100 outline-none focus:ring-2 focus:ring-indigo-100 transition shadow-sm"
      >
        <option value="all">全部类型</option>
        <option value="expense">支出</option>
        <option value="income">收入</option>
        <option value="transfer">转账</option>
      </select>

      <input
        v-model="filters.search_keyword"
        type="text"
        placeholder="搜索交易..."
        @input="handleFilterChange"
        class="w-full px-4 py-3 text-sm bg-white rounded-2xl border border-gray-100 outline-none focus:ring-2 focus:ring-indigo-100 transition shadow-sm"
      >
    </div>

    <!-- Content -->
    <div class="px-6 pb-32 flex-1">
      <div v-if="loading" class="py-12 text-center">
        <i class="fa-solid fa-spinner fa-spin text-gray-400 text-2xl"></i>
        <p class="text-gray-400 text-sm mt-2">加载中...</p>
      </div>

      <div v-else-if="transactions.length === 0" class="py-12 text-center">
        <i class="fa-solid fa-inbox text-gray-300 text-4xl"></i>
        <p class="text-gray-400 text-sm mt-2">暂无交易记录</p>
      </div>

      <div v-else class="space-y-6 animate-enter delay-200">
        <!-- Grouped by Date -->
        <div v-for="[date, dayTransactions] in groupedTransactions" :key="date">
          <!-- Date Header with Summary -->
          <div class="flex justify-between items-center px-1 py-3 border-b border-gray-100">
            <h3 class="text-sm font-bold text-gray-800">{{ formatDate(date) }}</h3>
            <div class="flex gap-6 text-xs">
              <span v-if="getDailySummary(dayTransactions).income > 0" class="text-green-500 font-medium">
                +¥{{ getDailySummary(dayTransactions).income.toFixed(2) }}
              </span>
              <span v-if="getDailySummary(dayTransactions).expense > 0" class="text-red-500 font-medium">
                -¥{{ getDailySummary(dayTransactions).expense.toFixed(2) }}
              </span>
            </div>
          </div>

          <!-- Transaction Items -->
          <div class="space-y-2 mt-3">
            <GlassCard
              v-for="transaction in dayTransactions"
              :key="transaction.id"
              class="p-4 flex items-center justify-between cursor-pointer group hover:shadow-md transition"
            >
              <div class="flex items-center gap-4 flex-1">
                <!-- Icon -->
                <div
                  :class="`w-12 h-12 rounded-[1rem] bg-${transaction.category?.color}-50 flex items-center justify-center text-lg shadow-inner`"
                >
                  <i
                    :class="`fa-solid ${transaction.category?.icon} text-${transaction.category?.color}-500`"
                  ></i>
                </div>

                <!-- Info -->
                <div class="flex-1 min-w-0">
                  <h4 class="font-semibold text-gray-800 text-sm">
                    {{ transaction.title || transaction.category?.name }}
                  </h4>
                  <p class="text-[10px] text-gray-400 mt-0.5 truncate">
                    {{ transaction.account?.account_name }}
                    <span v-if="transaction.description" class="mx-1">·</span>
                    <span v-if="transaction.description">{{ transaction.description }}</span>
                  </p>
                </div>
              </div>

              <!-- Amount -->
              <div class="flex items-center gap-3">
                <p :class="[getTypeColor(transaction.type), 'font-bold text-sm whitespace-nowrap']">
                  {{ formatCurrency(transaction.amount, transaction.type) }}
                </p>
                <button
                  @click.stop="handleDelete(transaction.id)"
                  class="w-8 h-8 rounded-lg bg-red-50 text-red-400 flex items-center justify-center opacity-0 group-hover:opacity-100 transition hover:bg-red-100 hover:text-red-600"
                >
                  <i class="fa-solid fa-trash text-xs"></i>
                </button>
              </div>
            </GlassCard>
          </div>
        </div>
      </div>

      <!-- Pagination -->
      <div v-if="totalPages > 1 && !loading && transactions.length > 0" class="flex items-center justify-between mt-8 pt-6 border-t border-gray-100">
        <button
          @click="prevPage"
          :disabled="page === 1"
          class="px-4 py-2 text-sm font-medium text-gray-600 bg-white border border-gray-200 rounded-xl hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed transition shadow-sm"
        >
          上一页
        </button>

        <span class="text-xs text-gray-500">
          第 {{ page }} / {{ totalPages }} 页
        </span>

        <button
          @click="nextPage"
          :disabled="page === totalPages"
          class="px-4 py-2 text-sm font-medium text-gray-600 bg-white border border-gray-200 rounded-xl hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed transition shadow-sm"
        >
          下一页
        </button>
      </div>
    </div>
  </div>
</template>

<style scoped>
</style>
