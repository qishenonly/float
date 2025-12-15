<script setup>
import { ref, onMounted, computed } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { useToast } from '../composables/useToast'
import GlassCard from '../components/GlassCard.vue'

const router = useRouter()
const authStore = useAuthStore()
const { showToast } = useToast()

const user = computed(() => authStore.currentUser)
const loading = ref(false)

onMounted(async () => {
  // 如果需要，从后端刷新最新用户信息
  if (authStore.isAuthenticated) {
    try {
      await authStore.fetchCurrentUser()
    } catch (error) {
      console.error('Failed to fetch user:', error)
    }
  }
})

const handleLogout = async () => {
  if (confirm('确定要退出登录吗？')) {
    loading.value = true
    try {
      await authStore.logout()
      router.push('/login')
    } catch (error) {
      alert('退出登录失败')
    } finally {
      loading.value = false
    }
  }
}

const showComingSoon = (feature) => {
  showToast(`${feature}功能开发中，敬请期待！`, 'info')
}
</script>

<template>
  <div>
    <!-- Header -->
    <div class="px-6 pt-16 pb-8 relative z-10 animate-enter">
      <div class="flex items-center justify-between mb-8">
        <div class="flex items-center gap-5">
          <div class="w-20 h-20 rounded-full p-1 bg-white shadow-lg relative group cursor-pointer">
               <img :src="user?.avatar_url || `https://api.dicebear.com/7.x/avataaars/svg?seed=${user?.username || 'User'}&backgroundColor=ffdfbf`" class="w-full h-full rounded-full group-hover:scale-105 transition duration-300" alt="avatar">
               <div v-if="user?.verified" class="absolute bottom-0 right-0 w-6 h-6 bg-green-400 rounded-full border-2 border-white flex items-center justify-center text-white text-[10px]">
                  <i class="fa-solid fa-check"></i>
               </div>
          </div>
          <div>
              <h1 class="text-2xl font-extrabold text-gray-800">{{ user?.display_name || user?.username || '用户' }}</h1>
              <div class="flex items-center gap-2 mt-1">
                  <span v-if="user?.verified" class="bg-indigo-100 text-indigo-600 text-[10px] px-2 py-0.5 rounded-full font-bold">已实名</span>
                  <p class="text-xs text-gray-400">坚持记账 {{ user?.continuous_days || 0 }} 天</p>
              </div>
          </div>
        </div>
        
        <!-- Edit Button -->
        <RouterLink to="/profile/edit" class="w-10 h-10 rounded-full bg-white shadow-md flex items-center justify-center text-indigo-600 hover:bg-indigo-50 transition active:scale-95">
          <i class="fa-solid fa-pen text-sm"></i>
        </RouterLink>
      </div>

      <GlassCard class="p-4 flex justify-between shadow-sm">
        <div @click="showComingSoon('本月账单详情')" class="text-center flex-1 border-r border-gray-200/50 cursor-pointer hover:bg-white/30 rounded-lg transition active-press">
            <p class="text-lg font-extrabold text-gray-800">{{ user?.total_records || 0 }}</p>
            <p class="text-[10px] text-gray-400 uppercase tracking-wide">本月笔数</p>
        </div>
        <div @click="showComingSoon('徽章系统')" class="text-center flex-1 border-r border-gray-200/50 cursor-pointer hover:bg-white/30 rounded-lg transition active-press">
            <p class="text-lg font-extrabold text-gray-800">{{ user?.total_badges || 0 }}</p>
            <p class="text-[10px] text-gray-400 uppercase tracking-wide">打卡徽章</p>
        </div>
        <div @click="showComingSoon('VIP 会员')" class="text-center flex-1 cursor-pointer hover:bg-white/30 rounded-lg transition active-press">
            <p class="text-lg font-extrabold text-indigo-600">{{ user?.membership_level || 'FREE' }}</p>
            <p class="text-[10px] text-gray-400 uppercase tracking-wide">{{ user?.membership_level === 'VIP' ? '高级版' : '免费版' }}</p>
        </div>
      </GlassCard>
    </div>

    <!-- Menu List -->
    <div class="px-6 py-2 relative z-10 pb-32">
      <GlassCard class="rounded-2xl shadow-sm overflow-hidden mb-6 animate-enter delay-100">
        <div class="divide-y divide-gray-100/50">
            <!-- <div @click="router.push('/accounts')" class="flex items-center justify-between p-4 hover:bg-white/50 cursor-pointer transition group active-press">
                <div class="flex items-center gap-3">
                    <div class="w-9 h-9 rounded-xl bg-indigo-50 text-indigo-500 flex items-center justify-center group-hover:scale-110 transition">
                        <i class="fa-solid fa-wallet text-sm"></i>
                    </div>
                    <span class="text-sm font-bold text-gray-700">账户管理</span>
                </div>
                <i class="fa-solid fa-chevron-right text-gray-300 text-xs group-hover:text-gray-500 transition"></i>
            </div> -->
            <div @click="router.push('/categories')" class="flex items-center justify-between p-4 hover:bg-white/50 cursor-pointer transition group active-press">
                <div class="flex items-center gap-3">
                    <div class="w-9 h-9 rounded-xl bg-blue-50 text-blue-500 flex items-center justify-center group-hover:scale-110 transition">
                        <i class="fa-solid fa-layer-group text-sm"></i>
                    </div>
                    <span class="text-sm font-bold text-gray-700">分类管理</span>
                </div>
                <i class="fa-solid fa-chevron-right text-gray-300 text-xs group-hover:text-gray-500 transition"></i>
            </div>
            <div @click="showToast('账单导出功能开发中...', 'info')" class="flex items-center justify-between p-4 hover:bg-white/50 cursor-pointer transition group active-press">
                <div class="flex items-center gap-3">
                    <div class="w-9 h-9 rounded-xl bg-green-50 text-green-500 flex items-center justify-center group-hover:scale-110 transition">
                        <i class="fa-solid fa-file-export text-sm"></i>
                    </div>
                    <span class="text-sm font-bold text-gray-700">账单导出</span>
                </div>
                <i class="fa-solid fa-chevron-right text-gray-300 text-xs group-hover:text-gray-500 transition"></i>
            </div>
            <div @click="showComingSoon('记账提醒')" class="flex items-center justify-between p-4 hover:bg-white/50 cursor-pointer transition group active-press">
                <div class="flex items-center gap-3">
                    <div class="w-9 h-9 rounded-xl bg-pink-50 text-pink-500 flex items-center justify-center group-hover:scale-110 transition">
                        <i class="fa-solid fa-bell text-sm"></i>
                    </div>
                    <span class="text-sm font-bold text-gray-700">记账提醒</span>
                </div>
                <div class="w-10 h-5 bg-indigo-500 rounded-full relative shadow-inner">
                    <div class="w-4 h-4 bg-white rounded-full absolute top-0.5 right-0.5 shadow-sm"></div>
                </div>
            </div>
        </div>
      </GlassCard>

      <GlassCard class="rounded-2xl shadow-sm overflow-hidden mb-6 animate-enter delay-200">
        <div class="divide-y divide-gray-100/50">
            <div @click="router.push('/settings')" class="flex items-center justify-between p-4 hover:bg-white/50 cursor-pointer transition group active-press">
                <div class="flex items-center gap-3">
                    <div class="w-9 h-9 rounded-xl bg-gray-50 text-gray-500 flex items-center justify-center group-hover:scale-110 transition">
                        <i class="fa-solid fa-gear text-sm"></i>
                    </div>
                    <span class="text-sm font-bold text-gray-700">通用设置</span>
                </div>
                <i class="fa-solid fa-chevron-right text-gray-300 text-xs group-hover:text-gray-500 transition"></i>
            </div>
            <div @click="router.push('/help')" class="flex items-center justify-between p-4 hover:bg-white/50 cursor-pointer transition group active-press">
                <div class="flex items-center gap-3">
                    <div class="w-9 h-9 rounded-xl bg-gray-50 text-gray-500 flex items-center justify-center group-hover:scale-110 transition">
                        <i class="fa-solid fa-circle-question text-sm"></i>
                    </div>
                    <span class="text-sm font-bold text-gray-700">帮助与反馈</span>
                </div>
                <i class="fa-solid fa-chevron-right text-gray-300 text-xs group-hover:text-gray-500 transition"></i>
            </div>
        </div>
      </GlassCard>

      <!-- Logout Button -->
      <button
        @click="handleLogout"
        :disabled="loading"
        class="w-full bg-white/80 backdrop-blur-md text-red-500 font-bold py-4 rounded-2xl shadow-sm hover:shadow-md transition duration-200 mb-4 animate-enter delay-300 disabled:opacity-50"
      >
        <span v-if="!loading">退出登录</span>
        <span v-else class="flex items-center justify-center gap-2">
          <i class="fa-solid fa-spinner fa-spin"></i>
          退出中...
        </span>
      </button>
      
      <p class="text-center text-[10px] text-gray-300 mt-2 mb-4 animate-enter delay-300">Version 0.0.4</p>
    </div>
  </div>
</template>

<style scoped>
@keyframes fadeInUp {
  from { opacity: 0; transform: translateY(20px); }
  to { opacity: 1; transform: translateY(0); }
}

.animate-enter {
  animation: fadeInUp 0.5s ease-out forwards;
  opacity: 0;
}

.delay-100 { animation-delay: 0.1s; }
.delay-200 { animation-delay: 0.2s; }
.delay-300 { animation-delay: 0.3s; }

.active-press:active {
  transform: scale(0.98);
  transition: transform 0.1s;
}
</style>
