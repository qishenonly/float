<script setup>
import { ref, onMounted, computed } from 'vue'
import { useRouter } from 'vue-router'
import { useToast } from '../composables/useToast'
import { accountAPI } from '@/api/account'
import ConfirmModal from '@/components/ConfirmModal.vue'

const router = useRouter()
const { showToast } = useToast()

const accounts = ref([])
const totalBalance = ref(0)
const loading = ref(false)
const showModal = ref(false)
const editingAccount = ref(null)
const showDeleteConfirm = ref(false)
const pendingDeleteAccount = ref(null)

const formData = ref({
  account_type: 'bank',
  account_name: '',
  account_number: '',
  initial_balance: '',
  icon: 'fa-credit-card',
  color: 'blue',
  include_in_total: true
})

const accountTypes = [
  { value: 'bank', label: '银行卡', icon: 'fa-credit-card' },
  { value: 'alipay', label: '支付宝', icon: 'fa-alipay' },
  { value: 'wechat', label: '微信', icon: 'fa-weixin' },
  { value: 'cash', label: '现金', icon: 'fa-money-bill' },
  { value: 'other', label: '其他', icon: 'fa-wallet' }
]

const availableColors = [
  'blue', 'green', 'red', 'orange', 'purple', 'pink', 
  'indigo', 'yellow', 'teal', 'cyan', 'slate', 'amber'
]

const availableIcons = [
  'fa-credit-card', 'fa-alipay', 'fa-weixin', 'fa-money-bill', 'fa-wallet',
  'fa-piggy-bank', 'fa-landmark', 'fa-money-check', 'fa-coins', 'fa-sack-dollar'
]

onMounted(() => {
  loadData()
})

const loadData = async () => {
  loading.value = true
  try {
    const [accountsRes, balanceRes] = await Promise.all([
      accountAPI.getAccounts(),
      accountAPI.getAccountBalance()
    ])
    accounts.value = accountsRes.data || []
    totalBalance.value = balanceRes.data?.total_balance || 0
  } catch (error) {
    console.error('Failed to load accounts:', error)
    showToast('加载数据失败', 'error')
  } finally {
    loading.value = false
  }
}

const openAddModal = () => {
  editingAccount.value = null
  formData.value = {
    account_type: 'bank',
    account_name: '',
    account_number: '',
    initial_balance: '',
    icon: 'fa-credit-card',
    color: 'blue',
    include_in_total: true
  }
  showModal.value = true
}

const openEditModal = (account) => {
  editingAccount.value = account
  formData.value = {
    account_type: account.account_type,
    account_name: account.account_name,
    account_number: account.account_number,
    initial_balance: account.initial_balance,
    icon: account.icon,
    color: account.color,
    include_in_total: account.include_in_total
  }
  showModal.value = true
}

const saveAccount = async () => {
  if (!formData.value.account_name.trim()) {
    showToast('请输入账户名称', 'error')
    return
  }

  loading.value = true
  try {
    const data = {
      ...formData.value,
      initial_balance: Number(formData.value.initial_balance) || 0
    }

    if (editingAccount.value) {
      await accountAPI.updateAccount(editingAccount.value.id, data)
      showToast('账户更新成功', 'success')
    } else {
      await accountAPI.createAccount(data)
      showToast('账户创建成功', 'success')
    }
    
    showModal.value = false
    await loadData()
  } catch (error) {
    console.error('Failed to save account:', error)
    showToast(error.response?.data?.message || '保存失败', 'error')
  } finally {
    loading.value = false
  }
}

const deleteAccount = (account) => {
  pendingDeleteAccount.value = account
  showDeleteConfirm.value = true
}

const confirmDeleteAccount = async () => {
  if (!pendingDeleteAccount.value) return
  
  loading.value = true
  try {
    await accountAPI.deleteAccount(pendingDeleteAccount.value.id)
    showToast('账户删除成功', 'success')
    showDeleteConfirm.value = false
    showModal.value = false
    pendingDeleteAccount.value = null
    await loadData()
  } catch (error) {
    console.error('Failed to delete account:', error)
    showToast('删除失败', 'error')
  } finally {
    loading.value = false
  }
}

const closeModal = () => {
  showModal.value = false
  editingAccount.value = null
}

const getAccountTypeLabel = (type) => {
  return accountTypes.find(t => t.value === type)?.label || type
}

const formatCurrency = (amount) => {
  return new Intl.NumberFormat('zh-CN', {
    style: 'currency',
    currency: 'CNY'
  }).format(amount)
}
</script>

<template>
  <div>
    <!-- Header -->
    <div class="px-6 pt-12 pb-4 relative z-10 flex items-center justify-between animate-enter">
      <button @click="router.back()" class="w-10 h-10 rounded-full bg-white/50 flex items-center justify-center backdrop-blur-md text-gray-600 hover:bg-white transition active-press">
        <i class="fa-solid fa-arrow-left"></i>
      </button>
      <h1 class="text-lg font-bold text-gray-800">账户管理</h1>
      <button @click="openAddModal" class="w-10 h-10 rounded-full bg-white/50 flex items-center justify-center backdrop-blur-md text-gray-600 hover:bg-white transition active-press">
        <i class="fa-solid fa-plus"></i>
      </button>
    </div>

    <!-- Total Balance Card -->
    <div class="px-6 py-6 animate-enter delay-100">
      <div class="bg-gradient-to-br from-indigo-600 to-purple-600 rounded-3xl p-6 text-white shadow-lg shadow-indigo-200">
        <p class="text-indigo-100 text-sm mb-1">总资产</p>
        <h2 class="text-3xl font-bold">{{ formatCurrency(totalBalance) }}</h2>
        <div class="mt-4 flex items-center gap-2 text-xs text-indigo-100 bg-white/10 w-fit px-3 py-1 rounded-full">
          <i class="fa-solid fa-shield-halved"></i>
          <span>数据已加密保护</span>
        </div>
      </div>
    </div>

    <!-- Account List -->
    <div class="px-6 pb-24 relative z-10">
      <div v-if="loading && accounts.length === 0" class="text-center py-12">
        <i class="fa-solid fa-spinner fa-spin text-3xl text-gray-400"></i>
        <p class="text-gray-400 mt-3">加载中...</p>
      </div>

      <div v-else class="space-y-4 animate-enter delay-200">
        <div 
          v-for="account in accounts" 
          :key="account.id"
          class="glass-card p-4 rounded-2xl flex items-center justify-between group active:scale-[0.99] transition cursor-pointer"
          @click="openEditModal(account)"
        >
          <div class="flex items-center gap-4">
            <div :class="`w-12 h-12 rounded-xl bg-${account.color}-100 text-${account.color}-600 flex items-center justify-center text-xl`">
              <i :class="`fa-solid ${account.icon}`"></i>
            </div>
            <div>
              <h3 class="font-bold text-gray-800">{{ account.account_name }}</h3>
              <div class="flex items-center gap-2 text-xs text-gray-400 mt-0.5">
                <span class="bg-gray-100 px-1.5 py-0.5 rounded">{{ getAccountTypeLabel(account.account_type) }}</span>
                <span v-if="account.account_number">尾号 {{ account.account_number }}</span>
              </div>
            </div>
          </div>
          
          <div class="text-right">
            <p class="font-bold text-gray-800">{{ formatCurrency(account.balance) }}</p>
            <p v-if="!account.include_in_total" class="text-xs text-gray-400 mt-0.5">
              <i class="fa-solid fa-eye-slash mr-1"></i>不计入总资产
            </p>
          </div>
        </div>

        <!-- Empty State -->
        <div v-if="accounts.length === 0" class="text-center py-12">
          <div class="w-20 h-20 bg-gray-100 rounded-full flex items-center justify-center mx-auto mb-4 text-gray-300 text-3xl">
            <i class="fa-solid fa-wallet"></i>
          </div>
          <h3 class="text-gray-800 font-bold mb-1">暂无账户</h3>
          <p class="text-gray-400 text-sm mb-6">添加您的第一个资产账户</p>
          <button @click="openAddModal" class="px-6 py-2 bg-indigo-600 text-white rounded-full font-bold text-sm shadow-md hover:bg-indigo-700 transition">
            立即添加
          </button>
        </div>
      </div>
    </div>

    <!-- Add/Edit Modal -->
    <div v-if="showModal" class="fixed inset-0 bg-black/50 z-50 flex items-end sm:items-center justify-center sm:p-6" @click.self="closeModal">
      <div class="bg-white rounded-t-3xl sm:rounded-3xl w-full max-w-md p-6 pb-24 space-y-5 animate-slide-up sm:animate-enter max-h-[90vh] overflow-y-auto">
        <div class="flex items-center justify-between sticky top-0 bg-white z-10 pb-2">
          <h2 class="text-lg font-bold text-gray-800">{{ editingAccount ? '编辑账户' : '添加账户' }}</h2>
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
                v-for="type in accountTypes"
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
              placeholder="例如：招商银行工资卡"
              maxlength="20"
              class="w-full px-4 py-3 bg-gray-50 rounded-xl border-none outline-none text-gray-800 placeholder-gray-400 focus:ring-2 focus:ring-indigo-100 transition"
            >
          </div>

          <!-- Balance & Number -->
          <div class="grid grid-cols-2 gap-4">
            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">当前余额</label>
              <input
                v-model="formData.initial_balance"
                type="number"
                placeholder="0.00"
                :disabled="!!editingAccount"
                class="w-full px-4 py-3 bg-gray-50 rounded-xl border-none outline-none text-gray-800 placeholder-gray-400 focus:ring-2 focus:ring-indigo-100 transition disabled:opacity-60 disabled:cursor-not-allowed"
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

          <!-- Options -->
          <div class="flex items-center justify-between py-2">
            <span class="text-sm font-medium text-gray-600">计入总资产</span>
            <button 
              @click="formData.include_in_total = !formData.include_in_total"
              :class="formData.include_in_total ? 'bg-indigo-600' : 'bg-gray-200'"
              class="w-12 h-7 rounded-full relative transition-colors duration-300 focus:outline-none"
            >
              <span 
                :class="formData.include_in_total ? 'translate-x-6' : 'translate-x-1'"
                class="block w-5 h-5 bg-white rounded-full shadow-sm transform transition-transform duration-300"
              ></span>
            </button>
          </div>
        </div>

        <!-- Actions -->
        <div class="flex gap-3 pt-4">
          <button 
            v-if="editingAccount"
            @click="deleteAccount(editingAccount)"
            class="px-4 py-3 bg-red-50 text-red-600 font-bold rounded-xl hover:bg-red-100 transition"
          >
            <i class="fa-solid fa-trash"></i>
          </button>
          <button 
            @click="saveAccount" 
            :disabled="loading" 
            class="flex-1 py-3 bg-indigo-600 text-white font-bold rounded-xl hover:bg-indigo-700 transition disabled:opacity-50 shadow-lg shadow-indigo-200"
          >
            {{ loading ? '保存中...' : '保存账户' }}
          </button>
        </div>
      </div>
    </div>

    <!-- Delete Confirm Modal -->
    <ConfirmModal
      :show="showDeleteConfirm"
      title="删除账户"
      :content="`确定要删除账户「${pendingDeleteAccount?.account_name}」吗？删除后无法恢复。`"
      confirmText="删除"
      @close="showDeleteConfirm = false; pendingDeleteAccount = null"
      @confirm="confirmDeleteAccount"
    />
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
