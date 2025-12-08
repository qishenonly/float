<template>
  <div class="min-h-screen bg-gradient-to-br from-violet-50 via-pink-50 to-cyan-50 pb-8">
    <!-- Header -->
    <div class="flex items-center justify-between px-6 pt-12 pb-4 animate-enter">
      <button @click="handleCancel" class="w-10 h-10 rounded-full flex items-center justify-center text-gray-700 hover:bg-white/50 transition">
        <i class="fa-solid fa-arrow-left"></i>
      </button>
      <h1 class="text-lg font-bold text-gray-800">编辑资料</h1>
      <button 
        @click="handleSave" 
        :disabled="loading"
        class="px-5 py-2 bg-indigo-600 text-white text-sm font-bold rounded-full hover:bg-indigo-700 transition disabled:opacity-50 shadow-md"
      >
        {{ loading ? '保存中...' : '保存' }}
      </button>
    </div>

    <!-- Content -->
    <div class="px-6 py-6 space-y-6 animate-enter delay-100">
      <!-- Avatar Section -->
      <div class="text-center">
        <div class="relative inline-block">
          <img 
            :src="user?.avatar_url || `https://api.dicebear.com/7.x/avataaars/svg?seed=${user?.username || 'User'}&backgroundColor=ffdfbf`" 
            class="w-28 h-28 rounded-full shadow-xl"
            alt="avatar"
          >
          <button 
            type="button"
            @click="showToast('头像上传功能开发中...', 'info')"
            class="absolute bottom-0 right-0 w-10 h-10 bg-indigo-600 rounded-full flex items-center justify-center text-white shadow-lg hover:bg-indigo-700 transition"
          >
            <i class="fa-solid fa-camera"></i>
          </button>
        </div>
        <p class="text-xs text-gray-400 mt-3">点击更换头像</p>
      </div>

      <!-- Form Fields -->
      <div class="space-y-4 animate-enter delay-200">
        <!-- Display Name -->
        <div>
          <label class="block text-sm font-medium text-gray-600 mb-2 px-1">昵称</label>
          <input
            v-model="formData.display_name"
            type="text"
            placeholder="请输入昵称"
            maxlength="50"
            class="w-full px-4 py-3 bg-white rounded-2xl border-none outline-none text-gray-800 placeholder-gray-400 shadow-sm"
          >
        </div>

        <!-- Phone -->
        <div>
          <label class="block text-sm font-medium text-gray-600 mb-2 px-1">手机号</label>
          <input
            v-model="formData.phone"
            type="tel"
            placeholder="请输入手机号"
            maxlength="20"
            class="w-full px-4 py-3 bg-white rounded-2xl border-none outline-none text-gray-800 placeholder-gray-400 shadow-sm"
          >
        </div>

        <!-- Email (Read-only) -->
        <div>
          <label class="block text-sm font-medium text-gray-400 mb-2 px-1">邮箱（不可修改）</label>
          <input
            :value="user?.email"
            type="email"
            disabled
            class="w-full px-4 py-3 bg-gray-100 rounded-2xl border-none outline-none text-gray-500 shadow-sm cursor-not-allowed"
          >
        </div>
      </div>

      <!-- Security Section -->
      <div class="pt-4 animate-enter delay-300">
        <h2 class="text-sm font-medium text-gray-600 mb-3 px-1">账号安全</h2>
        
        <RouterLink
          to="/profile/change-password"
          class="w-full flex items-center justify-between px-4 py-4 bg-white rounded-2xl shadow-sm hover:shadow-md transition"
        >
          <div class="flex items-center gap-3">
            <div class="w-10 h-10 rounded-full bg-indigo-100 flex items-center justify-center">
              <i class="fa-solid fa-lock text-indigo-600"></i>
            </div>
            <div class="text-left">
              <p class="text-sm font-medium text-gray-800">修改密码</p>
              <p class="text-xs text-gray-400">定期更换密码保护账户安全</p>
            </div>
          </div>
          <i class="fa-solid fa-chevron-right text-gray-400"></i>
        </RouterLink>
      </div>
    </div>
    <!-- Confirm Modal -->
    <ConfirmModal
      :show="showConfirmModal"
      title="放弃修改"
      content="确定要放弃当前的修改吗？未保存的内容将会丢失。"
      confirm-text="确定放弃"
      @close="showConfirmModal = false"
      @confirm="confirmCancel"
    />
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { userAPI } from '@/api/user'
import { useToast } from '@/composables/useToast'
import ConfirmModal from '@/components/ConfirmModal.vue'

const router = useRouter()
const authStore = useAuthStore()
const { showToast } = useToast()

const user = computed(() => authStore.currentUser)
const loading = ref(false)
const showConfirmModal = ref(false)

const formData = ref({
  display_name: '',
  phone: ''
})

onMounted(() => {
  // 初始化表单数据
  if (user.value) {
    formData.value = {
      display_name: user.value.display_name || '',
      phone: user.value.phone || ''
    }
  }
})

const handleSave = async () => {
  loading.value = true

  try {
    // 准备更新数据（只发送有值的字段）
    const updateData = {}
    
    if (formData.value.display_name && formData.value.display_name !== user.value?.display_name) {
      updateData.display_name = formData.value.display_name
    }
    
    if (formData.value.phone && formData.value.phone !== user.value?.phone) {
      updateData.phone = formData.value.phone
    }

    // 如果没有修改，直接返回
    if (Object.keys(updateData).length === 0) {
      showToast('没有修改内容', 'info')
      setTimeout(() => {
        router.back()
      }, 1000)
      return
    }

    // 调用 API 更新
    await userAPI.updateCurrentUser(updateData)
    
    // 刷新用户信息
    await authStore.fetchCurrentUser()
    
    showToast('保存成功！', 'success')
    
    // 1.5秒后返回
    setTimeout(() => {
      router.back()
    }, 1500)
  } catch (error) {
    console.error('Update profile error:', error)
    showToast(error.response?.data?.message || '保存失败，请重试', 'error')
  } finally {
    loading.value = false
  }
}

const handleCancel = () => {
  // Check if form changed
  const hasChanges = 
    formData.value.display_name !== (user.value?.display_name || '') ||
    formData.value.phone !== (user.value?.phone || '')

  if (hasChanges) {
    showConfirmModal.value = true
  } else {
    router.back()
  }
}

const confirmCancel = () => {
  showConfirmModal.value = false
  router.back()
}
</script>

<style scoped>
/* No additional styles needed - using Tailwind utilities */
</style>
