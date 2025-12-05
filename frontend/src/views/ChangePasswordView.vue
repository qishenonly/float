<template>
  <div class="min-h-screen bg-gradient-to-br from-violet-50 via-pink-50 to-cyan-50 pb-8">
    <!-- Header -->
    <div class="flex items-center justify-between px-6 py-4 animate-enter">
      <button @click="router.back()" class="w-10 h-10 rounded-full flex items-center justify-center text-gray-700 hover:bg-white/50 transition">
        <i class="fa-solid fa-arrow-left"></i>
      </button>
      <h1 class="text-lg font-bold text-gray-800">修改密码</h1>
      <div class="w-10"></div>
    </div>

    <!-- Content -->
    <div class="px-6 py-6 space-y-6 animate-enter delay-100">
      <!-- Tips -->
      <div class="bg-blue-50 border border-blue-200 rounded-2xl px-4 py-3">
        <div class="flex items-start gap-3">
          <i class="fa-solid fa-info-circle text-blue-500 mt-0.5"></i>
          <div class="text-xs text-blue-700 leading-relaxed">
            <p class="font-medium mb-1">密码安全建议：</p>
            <ul class="space-y-1 text-blue-600">
              <li>• 密码长度至少6位</li>
              <li>• 建议包含字母、数字和符号</li>
              <li>• 定期更换密码保护账户安全</li>
            </ul>
          </div>
        </div>
      </div>

      <!-- Form Fields -->
      <div class="space-y-4 animate-enter delay-200">
        <!-- Old Password -->
        <div>
          <label class="block text-sm font-medium text-gray-600 mb-2 px-1">当前密码</label>
          <div class="relative">
            <input
              v-model="formData.old_password"
              :type="showOldPassword ? 'text' : 'password'"
              placeholder="请输入当前密码"
              class="w-full px-4 py-3 pr-12 bg-white rounded-2xl border-none outline-none text-gray-800 placeholder-gray-400 shadow-sm"
            >
            <button
              type="button"
              @click="showOldPassword = !showOldPassword"
              class="absolute right-4 top-1/2 -translate-y-1/2 text-gray-400 hover:text-gray-600"
            >
              <i :class="showOldPassword ? 'fa-eye' : 'fa-eye-slash'" class="far"></i>
            </button>
          </div>
        </div>

        <!-- New Password -->
        <div>
          <label class="block text-sm font-medium text-gray-600 mb-2 px-1">新密码</label>
          <div class="relative">
            <input
              v-model="formData.new_password"
              :type="showNewPassword ? 'text' : 'password'"
              placeholder="请输入新密码（至少6位）"
              minlength="6"
              class="w-full px-4 py-3 pr-12 bg-white rounded-2xl border-none outline-none text-gray-800 placeholder-gray-400 shadow-sm"
            >
            <button
              type="button"
              @click="showNewPassword = !showNewPassword"
              class="absolute right-4 top-1/2 -translate-y-1/2 text-gray-400 hover:text-gray-600"
            >
              <i :class="showNewPassword ? 'fa-eye' : 'fa-eye-slash'" class="far"></i>
            </button>
          </div>
        </div>

        <!-- Confirm Password -->
        <div>
          <label class="block text-sm font-medium text-gray-600 mb-2 px-1">确认新密码</label>
          <div class="relative">
            <input
              v-model="confirmPassword"
              :type="showConfirmPassword ? 'text' : 'password'"
              placeholder="请再次输入新密码"
              class="w-full px-4 py-3 pr-12 bg-white rounded-2xl border-none outline-none text-gray-800 placeholder-gray-400 shadow-sm"
            >
            <button
              type="button"
              @click="showConfirmPassword = !showConfirmPassword"
              class="absolute right-4 top-1/2 -translate-y-1/2 text-gray-400 hover:text-gray-600"
            >
              <i :class="showConfirmPassword ? 'fa-eye' : 'fa-eye-slash'" class="far"></i>
            </button>
          </div>
        </div>
      </div>

      <!-- Submit Button -->
      <div class="pt-4 animate-enter delay-300">
        <button
          @click="handleSubmit"
          :disabled="loading || !isFormValid"
          class="w-full py-4 bg-indigo-600 text-white font-bold rounded-2xl shadow-lg hover:bg-indigo-700 transition disabled:opacity-50 disabled:cursor-not-allowed"
        >
          {{ loading ? '修改中...' : '确认修改' }}
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import { userAPI } from '@/api/user'
import { useToast } from '@/composables/useToast'

const router = useRouter()
const { showToast } = useToast()

const loading = ref(false)
const showOldPassword = ref(false)
const showNewPassword = ref(false)
const showConfirmPassword = ref(false)

const formData = ref({
  old_password: '',
  new_password: ''
})

const confirmPassword = ref('')

const isFormValid = computed(() => {
  return formData.value.old_password.length >= 6 &&
         formData.value.new_password.length >= 6 &&
         confirmPassword.value === formData.value.new_password
})

const handleSubmit = async () => {
  // 验证两次密码输入
  if (formData.value.new_password !== confirmPassword.value) {
    showToast('两次输入的新密码不一致', 'error')
    return
  }

  // 验证新密码长度
  if (formData.value.new_password.length < 6) {
    showToast('新密码至少需要6位', 'error')
    return
  }

  loading.value = true

  try {
    await userAPI.updatePassword({
      old_password: formData.value.old_password,
      new_password: formData.value.new_password
    })

    showToast('密码修改成功！', 'success')

    // 1.5秒后返回
    setTimeout(() => {
      router.back()
    }, 1500)
  } catch (error) {
    console.error('Change password error:', error)
    const errorMsg = error.response?.data?.message || '密码修改失败'
    showToast(errorMsg, 'error')
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
/* No additional styles needed - using Tailwind utilities */
</style>
