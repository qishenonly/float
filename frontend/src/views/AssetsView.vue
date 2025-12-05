<script setup>
import { ref, computed, onMounted } from 'vue'
import { accountAPI } from '@/api/account'
import { useToast } from '../composables/useToast'
import GlassCard from '../components/GlassCard.vue'

const { showToast } = useToast()

const accounts = ref([])
const loading = ref(false)
const showModal = ref(false)
const modalType = ref('fund') // 'fund' or 'credit'

const formData = ref({
  account_type: 'bank',
  account_name: '',
  account_number: '',
  initial_balance: '',
  icon: 'fa-credit-card',
  color: 'blue',
  include_in_total: true
})

const fundAccountTypes = [
  { value: 'bank', label: '银行卡', icon: 'fa-building-columns' },
  { value: 'alipay', label: '支付宝', icon: 'fa-alipay' },
  { value: 'wechat', label: '微信', icon: 'fa-weixin' },
  { value: 'cash', label: '现金', icon: 'fa-money-bill' },
  { value: 'other', label: '其他', icon: 'fa-wallet' }
]

const creditAccountTypes = [
  { value: 'credit', label: '信用卡', icon: 'fa-credit-card' },
  { value: 'other', label: '花呗/白条', icon: 'fa-face-smile' } // Using 'other' or maybe we should add more types later. For now map 'credit' to Credit Card.
]

// For the modal, we need to show appropriate types
const currentAccountTypes = computed(() => {
  return modalType.value === 'fund' ? fundAccountTypes : creditAccountTypes
})

const availableColors = [
  'blue', 'green', 'red', 'orange', 'purple', 'pink', 
  'indigo', 'yellow', 'teal', 'cyan', 'slate', 'amber'
]

onMounted(() => {
  loadData()
})

const loadData = async () => {
  loading.value = true
  try {
    const res = await accountAPI.getAccounts()
    accounts.value = res.data || []
  } catch (error) {
    console.error('Failed to load accounts:', error)
    showToast('加载账户数据失败', 'error')
  } finally {
    loading.value = false
  }
}

// Computed Properties for Assets/Liabilities
const fundAccounts = computed(() => {
  return accounts.value.filter(acc => acc.account_type !== 'credit')
})

const creditAccounts = computed(() => {
  return accounts.value.filter(acc => acc.account_type === 'credit')
})

const totalAssets = computed(() => {
  return fundAccounts.value
    .filter(acc => acc.include_in_total)
    .reduce((sum, acc) => sum + acc.balance, 0)
})

const totalLiabilities = computed(() => {
  // Credit accounts usually have negative balance if debt, or positive if we treat it as debt amount?
  // Let's assume negative balance = debt.
  // So we sum the absolute values of negative balances for "Total Liabilities" display.
  return creditAccounts.value
    .filter(acc => acc.include_in_total)
    .reduce((sum, acc) => sum + Math.abs(Math.min(acc.balance, 0)), 0)
})

const netWorth = computed(() => {
  // Net Worth = Assets + Liabilities (where liabilities are negative)
  // Or simply sum of all balances
  return accounts.value
    .filter(acc => acc.include_in_total)
    .reduce((sum, acc) => sum + acc.balance, 0)
})

const formatCurrency = (amount) => {
  return new Intl.NumberFormat('zh-CN', {
    style: 'currency',
    currency: 'CNY'
  }).format(amount)
}

const formatNumber = (num) => {
  if (num >= 10000) {
    return (num / 10000).toFixed(2) + '万'
  }
  return num.toFixed(2)
}

// Modal Logic
const openAddModal = (type) => {
  modalType.value = type
  formData.value = {
    account_type: type === 'fund' ? 'bank' : 'credit',
    account_name: '',
    account_number: '',
    initial_balance: '',
    icon: type === 'fund' ? 'fa-building-columns' : 'fa-credit-card',
    color: 'blue',
    include_in_total: true
  }
  showModal.value = true
}

const closeModal = () => {
  showModal.value = false
}

const saveAccount = async () => {
  if (!formData.value.account_name.trim()) {
    showToast('请输入账户名称', 'error')
    return
  }

  loading.value = true
  try {
    let balance = Number(formData.value.initial_balance) || 0
    // If adding a credit account and user enters positive number (e.g. 2000 debt), 
    // we might want to store it as negative? 
    // Usually users enter "Debt: 2000". 
    // Let's assume for 'credit' type, if user enters positive, it means debt, so we make it negative.
    // BUT, sometimes credit card has positive balance (overpaid).
    // Let's stick to: User enters what they see. If they see "-2000", they enter "-2000".
    // OR, we can provide a hint.
    // For now, let's just save as is.
    
    const data = {
      ...formData.value,
      initial_balance: balance
    }

    await accountAPI.createAccount(data)
    showToast('账户创建成功', 'success')
    showModal.value = false
    await loadData()
  } catch (error) {
    console.error('Failed to save account:', error)
    showToast(error.response?.data?.message || '保存失败', 'error')
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div>
    <!-- Header -->
    <div class="px-6 pt-12 pb-4 relative z-10 animate-enter">
      <h1 class="text-2xl font-extrabold text-gray-800">我的资产</h1>
      <p class="text-xs text-gray-400 mt-1">财富积累是一个美好的过程</p>
    </div>

    <div class="px-6 pb-24">
      <!-- Net Worth Card -->
      <div class="relative overflow-hidden rounded-[2rem] p-6 text-white shadow-xl shadow-indigo-200/50 mb-8 animate-enter delay-100 group active-press">
        <div class="absolute inset-0 bg-gradient-to-br from-[#4e54c8] to-[#8f94fb] z-0"></div>
        <div class="absolute right-0 top-0 w-32 h-32 bg-white opacity-10 rounded-full blur-2xl transform translate-x-1/2 -translate-y-1/2"></div>
        
        <div class="relative z-10">
          <p class="text-indigo-100 text-xs font-medium mb-1 tracking-wide opacity-80">当前净资产 (CNY)</p>
          <h2 class="text-3xl font-bold mb-6 tracking-tight">{{ formatCurrency(netWorth) }}</h2>
          
          <div class="flex items-center gap-4">
            <div class="flex-1">
              <div class="w-full bg-black/20 h-1.5 rounded-full mb-2 overflow-hidden backdrop-blur-sm">
                <div class="bg-cyan-300 h-full shadow-[0_0_10px_rgba(103,232,249,0.5)]" :style="{ width: `${totalAssets > 0 ? (totalAssets / (totalAssets + totalLiabilities)) * 100 : 0}%` }"></div>
              </div>
              <div class="flex justify-between text-[10px]">
                <span class="text-indigo-100/70">总资产</span>
                <span class="font-bold text-cyan-300">{{ formatNumber(totalAssets) }}</span>
              </div>
            </div>
            <div class="flex-1">
              <div class="w-full bg-black/20 h-1.5 rounded-full mb-2 overflow-hidden backdrop-blur-sm">
                <div class="bg-pink-400 h-full shadow-[0_0_10px_rgba(244,114,182,0.5)]" :style="{ width: `${totalLiabilities > 0 ? (totalLiabilities / (totalAssets + totalLiabilities)) * 100 : 0}%` }"></div>
              </div>
              <div class="flex justify-between text-[10px]">
                <span class="text-indigo-100/70">总负债</span>
                <span class="font-bold text-pink-300">{{ formatNumber(totalLiabilities) }}</span>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Fund Accounts -->
      <div class="mb-6 animate-enter delay-200">
        <div class="flex justify-between items-center mb-4">
          <h3 class="text-sm font-bold text-gray-800">资金账户</h3>
          <button @click="openAddModal('fund')" class="w-7 h-7 rounded-full bg-white shadow-sm flex items-center justify-center text-gray-500 hover:text-indigo-600 transition active-press">
            <i class="fa-solid fa-plus text-xs"></i>
          </button>
        </div>
        
        <div v-if="fundAccounts.length === 0" class="text-center py-8 text-gray-400 text-xs">
          暂无资金账户
        </div>

        <div class="space-y-3">
          <GlassCard 
            v-for="account in fundAccounts"
            :key="account.id"
            class="p-4 rounded-2xl flex items-center justify-between shadow-sm hover:shadow-md transition cursor-pointer"
          >
            <div class="flex items-center gap-4">
              <div :class="`w-10 h-10 rounded-xl bg-${account.color}-50 flex items-center justify-center shadow-inner`">
                <i :class="`fa-solid ${account.icon} text-${account.color}-500`"></i>
              </div>
              <div>
                <h4 class="font-bold text-gray-800 text-sm">{{ account.account_name }}</h4>
                <p class="text-[10px] text-gray-400 mt-0.5">{{ account.account_number ? `尾号 ${account.account_number}` : '账户' }}</p>
              </div>
            </div>
            <span class="font-bold text-gray-800">{{ formatCurrency(account.balance) }}</span>
          </GlassCard>
        </div>
      </div>

      <!-- Credit Accounts -->
      <div class="mb-6 animate-enter delay-300">
        <div class="flex justify-between items-center mb-4">
          <h3 class="text-sm font-bold text-gray-800">信用账户</h3>
          <button @click="openAddModal('credit')" class="w-7 h-7 rounded-full bg-white shadow-sm flex items-center justify-center text-gray-500 hover:text-indigo-600 transition active-press">
            <i class="fa-solid fa-plus text-xs"></i>
          </button>
        </div>
        
        <div v-if="creditAccounts.length === 0" class="text-center py-8 text-gray-400 text-xs">
          暂无信用账户
        </div>

        <div class="space-y-3">
          <GlassCard 
            v-for="account in creditAccounts"
            :key="account.id"
            class="p-4 rounded-2xl flex items-center justify-between shadow-sm hover:shadow-md transition cursor-pointer"
          >
            <div class="flex items-center gap-4">
              <div :class="`w-10 h-10 rounded-xl bg-${account.color}-50 flex items-center justify-center shadow-inner`">
                <i :class="`fa-solid ${account.icon} text-${account.color}-500`"></i>
              </div>
              <div>
                <h4 class="font-bold text-gray-800 text-sm">{{ account.account_name }}</h4>
                <p class="text-[10px] text-gray-400 mt-0.5">{{ account.account_number ? `尾号 ${account.account_number}` : '信用账户' }}</p>
              </div>
            </div>
            <span class="font-bold text-gray-800">{{ formatCurrency(account.balance) }}</span>
          </GlassCard>
        </div>
      </div>
    </div>

    <!-- Add Account Modal -->
    <div v-if="showModal" class="fixed inset-0 bg-black/50 z-50 flex items-end sm:items-center justify-center sm:p-6" @click.self="closeModal">
      <div class="bg-white rounded-t-3xl sm:rounded-3xl w-full max-w-md p-6 space-y-5 animate-slide-up sm:animate-enter max-h-[90vh] overflow-y-auto">
        <div class="flex items-center justify-between sticky top-0 bg-white z-10 pb-2">
          <h2 class="text-lg font-bold text-gray-800">{{ modalType === 'fund' ? '添加资金账户' : '添加信用账户' }}</h2>
          <button @click="closeModal" class="w-8 h-8 rounded-full bg-gray-100 hover:bg-gray-200 flex items-center justify-center transition">
            <i class="fa-solid fa-times text-gray-500"></i>
          </button>
        </div>

        <!-- Form -->
        <div class="space-y-4">
          <!-- Account Type -->
          <div>
            <label class="block text-sm font-medium text-gray-600 mb-2">账户类型</label>
            <div class="grid grid-cols-3 gap-2">
              <button
                v-for="type in currentAccountTypes"
                :key="type.value"
                @click="formData.account_type = type.value"
                :class="formData.account_type === type.value ? 'bg-indigo-600 text-white shadow-md' : 'bg-gray-50 text-gray-600 hover:bg-gray-100'"
                class="py-2.5 rounded-xl text-sm font-medium transition flex flex-col items-center gap-1"
              >
                <i :class="`fa-solid ${type.icon}`"></i>
                {{ type.label }}
              </button>
            </div>
          </div>

          <!-- Name -->
          <div>
            <label class="block text-sm font-medium text-gray-600 mb-2">账户名称</label>
            <input
              v-model="formData.account_name"
              type="text"
              placeholder="例如：招商银行"
              maxlength="20"
              class="w-full px-4 py-3 bg-gray-50 rounded-xl border-none outline-none text-gray-800 placeholder-gray-400 focus:ring-2 focus:ring-indigo-100 transition"
            >
          </div>

          <!-- Balance & Number -->
          <div class="grid grid-cols-2 gap-4">
            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">
                {{ modalType === 'fund' ? '当前余额' : '当前欠款(负数)' }}
              </label>
              <input
                v-model="formData.initial_balance"
                type="number"
                :placeholder="modalType === 'fund' ? '0.00' : '-0.00'"
                class="w-full px-4 py-3 bg-gray-50 rounded-xl border-none outline-none text-gray-800 placeholder-gray-400 focus:ring-2 focus:ring-indigo-100 transition"
              >
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">卡号后4位</label>
              <input
                v-model="formData.account_number"
                type="text"
                placeholder="选填"
                maxlength="4"
                class="w-full px-4 py-3 bg-gray-50 rounded-xl border-none outline-none text-gray-800 placeholder-gray-400 focus:ring-2 focus:ring-indigo-100 transition"
              >
            </div>
          </div>

          <!-- Color -->
          <div>
            <label class="block text-sm font-medium text-gray-600 mb-2">卡片颜色</label>
            <div class="grid grid-cols-6 gap-2">
              <button
                v-for="color in availableColors"
                :key="color"
                @click="formData.color = color"
                :class="[
                  `bg-${color}-500`,
                  formData.color === color ? 'ring-2 ring-offset-2 ring-gray-400 scale-90' : ''
                ]"
                class="w-full aspect-square rounded-full transition hover:scale-105"
              ></button>
            </div>
          </div>
        </div>

        <!-- Actions -->
        <div class="pt-4">
          <button 
            @click="saveAccount" 
            :disabled="loading" 
            class="w-full py-3 bg-indigo-600 text-white font-bold rounded-xl hover:bg-indigo-700 transition disabled:opacity-50 shadow-lg shadow-indigo-200"
          >
            {{ loading ? '保存中...' : '保存账户' }}
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.animate-slide-up {
  animation: slideUp 0.3s cubic-bezier(0.16, 1, 0.3, 1);
}

@keyframes slideUp {
  from {
    transform: translateY(100%);
  }
  to {
    transform: translateY(0);
  }
}
</style>
