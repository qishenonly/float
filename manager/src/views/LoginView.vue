<template>
  <div class="bg-[#F0F2F5] h-screen flex items-center justify-center relative overflow-hidden">
    <!-- Blobs -->
    <div class="blob w-96 h-96 bg-purple-300 rounded-full top-0 left-0 mix-blend-multiply"></div>
    <div class="blob w-96 h-96 bg-blue-300 rounded-full bottom-0 right-0 mix-blend-multiply animation-delay-2000"></div>

    <div class="glass-card w-[400px] p-10 rounded-3xl relative z-10">
      <div class="text-center mb-10">
        <div class="w-16 h-16 rounded-2xl bg-gradient-to-tr from-indigo-600 to-purple-600 flex items-center justify-center text-white text-3xl shadow-xl shadow-indigo-200 mx-auto mb-4">
          <i class="fa-solid fa-mountain-sun"></i>
        </div>
        <h1 class="text-2xl font-extrabold text-gray-800">Float Admin</h1>
        <p class="text-gray-400 text-sm mt-1">管理系统登录</p>
      </div>

      <form @submit.prevent="handleLogin" class="space-y-5">
        <div class="relative">
          <i class="fa-solid fa-envelope absolute left-4 top-4 text-gray-400"></i>
          <input 
            v-model="form.email"
            type="email" 
            placeholder="管理员账号" 
            required
            class="w-full bg-white/50 border border-gray-100 rounded-xl py-3.5 pl-10 pr-4 text-sm outline-none focus:bg-white focus:border-indigo-500 transition shadow-sm"
          >
        </div>
        
        <div class="relative">
          <i class="fa-solid fa-lock absolute left-4 top-4 text-gray-400"></i>
          <input 
            v-model="form.password"
            type="password" 
            placeholder="密码" 
            required
            class="w-full bg-white/50 border border-gray-100 rounded-xl py-3.5 pl-10 pr-4 text-sm outline-none focus:bg-white focus:border-indigo-500 transition shadow-sm"
          >
        </div>

        <div class="flex justify-between items-center text-xs">
          <label class="flex items-center gap-2 cursor-pointer text-gray-500">
            <input type="checkbox" v-model="form.remember" class="rounded border-gray-300 text-indigo-600 focus:ring-indigo-500"> 记住我
          </label>
          <a href="#" class="text-indigo-600 font-bold hover:underline">忘记密码?</a>
        </div>

        <button 
          type="submit"
          :disabled="loading"
          class="w-full bg-gray-900 text-white font-bold py-4 rounded-xl shadow-lg shadow-gray-300 hover:scale-[1.02] active:scale-[0.98] transition disabled:opacity-50 disabled:cursor-not-allowed"
        >
          {{ loading ? '登录中...' : '登 录' }}
        </button>
      </form>

      <p class="text-center text-xs text-gray-400 mt-8">
        &copy; 2025 Float Island. All rights reserved.
      </p>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
import { login } from '@/api/admin'

const router = useRouter()
const loading = ref(false)

const form = reactive({
  email: '',
  password: '',
  remember: false
})

const handleLogin = async () => {
    loading.value = true
    try {
        const res = await login({
            // Backend expects 'email' key
            email: form.email, 
            password: form.password
        })

        // Save token
        if (res.access_token) {
            localStorage.setItem('token', res.access_token)
            localStorage.setItem('user', JSON.stringify(res)) // Store other info if needed
            router.push('/')
        } else {
            throw new Error('No access token received')
        }
    } catch (error) {
        console.error("Login failed", error)
        alert('登录失败: ' + (error.response?.data?.error || error.message))
    } finally {
        loading.value = false
    }
}
</script>

<style scoped>
.animation-delay-2000 {
  animation-delay: 2s;
}
</style>
