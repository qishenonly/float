<template>
  <div class="fixed inset-0 overflow-hidden bg-gradient-to-br from-cyan-50 via-white to-violet-50">
    <div class="h-full flex items-center justify-center px-4 py-8">
      <div class="w-full max-w-md -mt-16">
        <!-- Logo & Title - 居中对齐，移除返回按钮 -->
        <div class="mb-7 animate-enter text-center">
          <div class="w-16 h-16 rounded-2xl bg-gradient-to-tr from-violet-600 to-indigo-500 flex items-center justify-center text-white text-2xl shadow-lg shadow-indigo-200/50 mb-4 mx-auto">
            <i class="fa-solid fa-cloud"></i>
          </div>
          <h1 class="text-3xl font-extrabold text-gray-800 mb-2">入住浮岛</h1>
          <p class="text-gray-400 text-sm">开始记录您的生活足迹 ☁️</p>
        </div>

        <!-- Register Form -->
        <form @submit.prevent="handleRegister" class="space-y-4 animate-enter delay-100">
          <!-- Username -->
          <div class="glass-input rounded-2xl px-4 py-3.5 flex items-center gap-3">
            <i class="fa-regular fa-user text-gray-400 w-5 text-center"></i>
            <input
              v-model="formData.username"
              type="text"
              placeholder="岛民ID（用户名）"
              class="bg-transparent w-full outline-none text-gray-700 font-medium placeholder-gray-400"
              required
              minlength="3"
            >
          </div>

          <!-- Email -->
          <div class="glass-input rounded-2xl px-4 py-3.5 flex items-center gap-3">
            <i class="fa-regular fa-envelope text-gray-400 w-5 text-center"></i>
            <input
              v-model="formData.email"
              type="email"
              placeholder="电子邮箱"
              class="bg-transparent w-full outline-none text-gray-700 font-medium placeholder-gray-400"
              required
            >
          </div>

          <!-- Password -->
          <div class="glass-input rounded-2xl px-4 py-3.5 flex items-center gap-3">
            <i class="fa-solid fa-lock text-gray-400 w-5 text-center"></i>
            <input
              v-model="formData.password"
              :type="showPassword ? 'text' : 'password'"
              placeholder="设置密钥"
              class="bg-transparent w-full outline-none text-gray-700 font-medium placeholder-gray-400"
              required
              minlength="6"
            >
            <i 
              @click="showPassword = !showPassword"
              :class="showPassword ? 'fa-eye' : 'fa-eye-slash'"
              class="far text-gray-400 cursor-pointer hover:text-indigo-600 transition"
            ></i>
          </div>

          <!-- Confirm Password -->
          <div class="glass-input rounded-2xl px-4 py-3.5 flex items-center gap-3">
            <i class="fa-solid fa-check-double text-gray-400 w-5 text-center"></i>
            <input
              v-model="confirmPassword"
              :type="showPassword ? 'text' : 'password'"
              placeholder="确认密钥"
              class="bg-transparent w-full outline-none text-gray-700 font-medium placeholder-gray-400"
              required
            >
          </div>

          <!-- Terms -->
          <div class="flex items-start gap-3 mt-2 px-1">
            <div class="relative flex items-center">
              <input
                v-model="agreedToTerms"
                type="checkbox"
                id="tos"
                class="peer h-4 w-4 cursor-pointer appearance-none rounded border border-gray-300 shadow-sm checked:border-indigo-500 checked:bg-indigo-500 transition-all"
              >
              <i class="fa-solid fa-check text-white absolute left-1/2 top-1/2 -translate-x-1/2 -translate-y-1/2 text-[10px] opacity-0 peer-checked:opacity-100 pointer-events-none"></i>
            </div>
            <label for="tos" class="text-xs text-gray-500 leading-tight cursor-pointer">
              我已阅读并同意 <a @click.prevent="showTerms" class="text-indigo-600 font-bold cursor-pointer hover:underline">浮岛居民协议</a>
            </label>
          </div>

          <!-- Error Message -->
          <div v-if="errorMessage" class="bg-red-50 border border-red-200 text-red-600 px-4 py-3 rounded-xl text-sm">
            {{ errorMessage }}
          </div>

          <!-- Submit Button -->
          <button
            type="submit"
            :disabled="loading || !agreedToTerms"
            class="w-full bg-gradient-to-r from-violet-600 to-indigo-500 text-white font-bold py-4 rounded-2xl shadow-lg shadow-indigo-200 hover:shadow-xl hover:scale-[1.02] active:scale-[0.98] transition duration-200 mt-6 disabled:opacity-50"
          >
            <span v-if="!loading">完成入住</span>
            <span v-else class="flex items-center justify-center gap-2">
              <i class="fa-solid fa-spinner fa-spin"></i>
              注册中...
            </span>
          </button>
        </form>

        <!-- Login Link -->
        <div class="mt-7 text-center animate-enter delay-200">
          <p class="text-sm text-gray-500">
            已有居民证? 
            <RouterLink to="/login" class="font-bold text-indigo-600 hover:text-indigo-700">
              回家
            </RouterLink>
          </p>
        </div>
      </div>
    </div>

    <!-- Terms Dialog - 避开底部导航栏 -->
    <div v-if="showTermsDialog" class="fixed inset-0 bg-black/50 flex items-center justify-center z-50 px-4" @click="showTermsDialog = false">
      <div class="bg-white rounded-3xl p-6 max-w-md w-full flex flex-col" style="max-height: calc(100vh - 180px); margin-bottom: 80px;" @click.stop>
        <div class="flex justify-between items-center mb-4 flex-shrink-0">
          <h2 class="text-xl font-bold text-gray-800">浮岛居民协议</h2>
          <button @click="showTermsDialog = false" class="w-8 h-8 rounded-full bg-gray-100 flex items-center justify-center text-gray-600 hover:bg-gray-200 flex-shrink-0">
            <i class="fa-solid fa-xmark"></i>
          </button>
        </div>
        
        <!-- 可滚动内容区域 -->
        <div class="flex-1 overflow-y-auto space-y-4 text-sm text-gray-600 leading-relaxed mb-4">
          <p class="font-bold text-gray-800">欢迎来到浮岛 Float Island！</p>
          
          <p>感谢您选择浮岛作为您的个人财务管理工具。在开始使用我们的服务之前，请仔细阅读以下条款：</p>
          
          <div>
            <h3 class="font-bold text-gray-800 mb-2">1. 服务说明</h3>
            <p>浮岛是一款个人财务管理应用，帮助您记录和分析日常收支，管理资产和预算。</p>
          </div>
          
          <div>
            <h3 class="font-bold text-gray-800 mb-2">2. 隐私保护</h3>
            <p>我们重视您的隐私。您的所有财务数据将被加密存储，仅供您个人使用，不会与第三方共享。</p>
          </div>
          
          <div>
            <h3 class="font-bold text-gray-800 mb-2">3. 数据安全</h3>
            <p>我们采用业界标准的安全措施保护您的数据，包括但不限于数据加密、安全传输协议等。</p>
          </div>
          
          <div>
            <h3 class="font-bold text-gray-800 mb-2">4. 用户责任</h3>
            <p>您需要妥善保管账号密码，对账号下的所有活动负责。如发现账号异常，请立即联系我们。</p>
          </div>
          
          <div>
            <h3 class="font-bold text-gray-800 mb-2">5. 服务变更</h3>
            <p>我们保留随时修改或终止服务的权利，但会提前通知用户重大变更。</p>
          </div>
          
          <p class="text-xs text-gray-400 mt-4">最后更新时间：2025年12月5日</p>
        </div>
        
        <!-- 固定底部按钮 -->
        <button
          @click="acceptTerms"
          class="w-full bg-gradient-to-r from-violet-600 to-indigo-500 text-white font-bold py-3 rounded-2xl shadow-lg hover:shadow-xl transition flex-shrink-0"
        >
          我已阅读并同意
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const router = useRouter()
const authStore = useAuthStore()

const formData = ref({
  username: '',
  email: '',
  password: ''
})

const confirmPassword = ref('')
const showPassword = ref(false)
const agreedToTerms = ref(false)
const loading = ref(false)
const errorMessage = ref('')
const showTermsDialog = ref(false)

const showTerms = () => {
  showTermsDialog.value = true
}

const acceptTerms = () => {
  agreedToTerms.value = true
  showTermsDialog.value = false
}

// 错误信息翻译
const translateError = (error) => {
  const errorMap = {
    'Network Error': '网络连接失败，请检查网络设置',
    'timeout': '请求超时，请稍后重试',
    'Request failed with status code 400': '注册失败，用户名或邮箱可能已被使用',
    'Request failed with status code 500': '服务器错误，请稍后重试',
  }
  
  const errorString = error?.message || error?.toString() || '未知错误'
  
  // 检查是否匹配已知错误
  for (const [key, value] of Object.entries(errorMap)) {
    if (errorString.includes(key)) {
      return value
    }
  }
  
  // 检查后端返回的中文错误信息
  if (error?.response?.data?.message) {
    return error.response.data.message
  }
  
  // 默认友好提示
  return '注册失败，请稍后重试'
}

const handleRegister = async () => {
  // 验证密码匹配
  if (formData.value.password !== confirmPassword.value) {
    errorMessage.value = '两次输入的密码不一致'
    return
  }

  if (!agreedToTerms.value) {
    errorMessage.value = '请先同意浮岛居民协议'
    return
  }

  loading.value = true
  errorMessage.value = ''

  try {
    await authStore.register(formData.value)
    router.push('/')
  } catch (error) {
    console.error('Register error:', error)
    errorMessage.value = translateError(error)
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
.glass-input {
  background: rgba(255, 255, 255, 0.6);
  backdrop-filter: blur(12px);
  border: 1px solid rgba(255, 255, 255, 0.8);
  transition: all 0.3s ease;
}

.glass-input:focus-within {
  background: rgba(255, 255, 255, 0.9);
  border-color: #8b5cf6;
  box-shadow: 0 0 0 4px rgba(139, 92, 246, 0.1);
}

@keyframes fadeInUp {
  from { opacity: 0; transform: translateY(20px); }
  to { opacity: 1; transform: translateY(0); }
}

.animate-enter {
  animation: fadeInUp 0.8s cubic-bezier(0.16, 1, 0.3, 1) forwards;
  opacity: 0;
}

.delay-100 { animation-delay: 0.1s; }
.delay-200 { animation-delay: 0.2s; }
</style>
