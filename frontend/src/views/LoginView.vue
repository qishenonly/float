<template>
  <div class="fixed inset-0 overflow-hidden bg-gradient-to-br from-violet-50 via-white to-cyan-50">
    <div class="h-full flex items-center justify-center px-4 py-8">
      <div class="w-full max-w-md -mt-16">
        <!-- Logo & Title -->
         <div class="text-center mb-8 animate-enter">
           <div class="inline-block animate-float">
             <div 
               @click="handleLogoClick"
               class="w-20 h-20 rounded-[2rem] bg-gradient-to-tr from-violet-600 to-indigo-500 flex items-center justify-center text-white text-3xl shadow-xl shadow-indigo-200/50 mb-6 mx-auto border-4 border-white/30 backdrop-blur-md cursor-pointer hover:shadow-2xl hover:shadow-indigo-300/50 transition"
             >
               <i class="fa-solid fa-mountain-sun"></i>
             </div>
           </div>
          <h1 class="text-3xl font-extrabold text-gray-800 mb-2 tracking-tight">
            浮岛 <span class="text-indigo-600">Float Island</span>
          </h1>
          <p class="text-gray-400 text-sm font-medium">构建你的财富岛屿 🏝️</p>
        </div>

        <!-- Login Form -->
         <form @submit.prevent="handleLogin" class="space-y-5 animate-enter delay-100">
           <!-- Email -->
           <div class="glass-input rounded-2xl px-4 py-3.5 flex items-center gap-3">
             <i class="fa-solid fa-envelope text-gray-400 w-5 text-center"></i>
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
            <i class="fa-solid fa-key text-gray-400 w-5 text-center"></i>
            <input
              v-model="formData.password"
              :type="showPassword ? 'text' : 'password'"
              placeholder="通行密钥"
              class="bg-transparent w-full outline-none text-gray-700 font-medium placeholder-gray-400"
              required
            >
            <i 
              @click="showPassword = !showPassword"
              :class="showPassword ? 'fa-eye' : 'fa-eye-slash'"
              class="far text-gray-400 cursor-pointer hover:text-indigo-600 transition"
            ></i>
          </div>

          <!-- Forgot Password -->
          <div class="flex justify-end">
            <a href="#" class="text-xs font-bold text-indigo-600 hover:text-indigo-700">忘记密钥?</a>
          </div>

          <!-- Error Message -->
          <div v-if="errorMessage" class="bg-red-50 border border-red-200 text-red-600 px-4 py-3 rounded-xl text-sm">
            {{ errorMessage }}
          </div>

          <!-- Submit Button -->
          <button
            type="submit"
            :disabled="loading"
            class="w-full bg-gradient-to-r from-violet-600 to-indigo-500 text-white font-bold py-4 rounded-2xl shadow-xl shadow-indigo-100 hover:scale-[1.02] active:scale-[0.98] transition duration-200 disabled:opacity-50"
          >
            <span v-if="!loading">登岛</span>
            <span v-else class="flex items-center justify-center gap-2">
              <i class="fa-solid fa-spinner fa-spin"></i>
              登录中...
            </span>
          </button>
        </form>

        <!-- Social Login (Optional) -->
        <div class="mt-10 animate-enter delay-200">
          <div class="flex items-center gap-4 mb-6">
            <div class="h-px bg-gray-200 flex-1"></div>
            <span class="text-[10px] text-gray-400 font-medium uppercase tracking-widest">快速通行</span>
            <div class="h-px bg-gray-200 flex-1"></div>
          </div>
          
          <div class="flex justify-center gap-6">
            <button type="button" class="w-14 h-14 rounded-full bg-white/70 border border-gray-200 flex items-center justify-center text-green-500 text-xl shadow-sm hover:scale-110 transition">
              <i class="fa-brands fa-weixin"></i>
            </button>
            <button type="button" class="w-14 h-14 rounded-full bg-white/70 border border-gray-200 flex items-center justify-center text-gray-900 text-xl shadow-sm hover:scale-110 transition">
              <i class="fa-brands fa-apple"></i>
            </button>
          </div>
        </div>

        <!-- Register Link -->
        <div class="mt-7 text-center animate-enter delay-300">
          <p class="text-sm text-gray-500">
            新来的岛民? 
            <RouterLink to="/register" class="font-bold text-indigo-600 hover:text-indigo-700">
              注册居民证
            </RouterLink>
          </p>
        </div>
      </div>
    </div>

    <!-- Health Check Modal -->
    <HealthCheckModal 
      :show="healthCheckShow"
      :status="healthCheckStatus"
      :title="healthCheckTitle"
      :data="healthCheckData"
      :error="healthCheckError"
      :url="healthCheckUrl"
      @close="healthCheckShow = false"
    />
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import HealthCheckModal from '@/components/HealthCheckModal.vue'

const router = useRouter()
const authStore = useAuthStore()
const apiBaseUrl = import.meta.env.VITE_API_BASE_URL

const formData = ref({
  email: '',
  password: ''
})

const showPassword = ref(false)
const loading = ref(false)
const errorMessage = ref('')

// Health Check Modal
const healthCheckShow = ref(false)
const healthCheckStatus = ref('loading')
const healthCheckTitle = ref('健康检查')
const healthCheckData = ref(null)
const healthCheckError = ref('')
const healthCheckUrl = ref('')

// Click Counter for Health Check (5 clicks in 3 seconds)
const clickCount = ref(0)
const clickTimeout = ref(null)
const clickResetTimeout = ref(null)

// 错误信息翻译
const translateError = (error) => {
  const errorMap = {
    'Network Error': '网络连接失败，请检查网络设置',
    'timeout': '请求超时，请稍后重试',
    'Request failed with status code 400': '邮箱或密码错误',
    'Request failed with status code 401': '邮箱或密码错误',
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
  return '登录失败，请检查邮箱和密码'
}

const handleLogin = async () => {
  loading.value = true
  errorMessage.value = ''

  try {
    await authStore.login(formData.value)
    router.push('/')
  } catch (error) {
    console.error('Login error:', error)
    errorMessage.value = translateError(error)
  } finally {
    loading.value = false
  }
}

const handleLogoClick = () => {
  clickCount.value++
  
  // Clear existing reset timeout
  if (clickResetTimeout.value) {
    clearTimeout(clickResetTimeout.value)
  }
  
  // Check if 5 clicks reached
  if (clickCount.value === 5) {
    checkHealth()
    clickCount.value = 0
    return
  }
  
  // Reset click count after 3 seconds of no clicks
  clickResetTimeout.value = setTimeout(() => {
    clickCount.value = 0
  }, 3000)
}

const checkHealth = async () => {
  healthCheckShow.value = true
  healthCheckStatus.value = 'loading'
  healthCheckTitle.value = '健康检查'
  
  // Get base URL without /api/v1 suffix
  const baseUrl = apiBaseUrl.replace('/api/v1', '')
  healthCheckUrl.value = `${baseUrl}/health`
  
  try {
    const response = await fetch(healthCheckUrl.value)
    const data = await response.json()
    
    healthCheckStatus.value = 'success'
    healthCheckTitle.value = '健康检查成功'
    healthCheckData.value = data
  } catch (error) {
    healthCheckStatus.value = 'error'
    healthCheckTitle.value = '健康检查失败'
    healthCheckError.value = error.message
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

@keyframes float {
  0% { transform: translateY(0px); }
  50% { transform: translateY(-10px); }
  100% { transform: translateY(0px); }
}

.animate-float {
  animation: float 6s ease-in-out infinite;
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
.delay-300 { animation-delay: 0.3s; }
</style>
