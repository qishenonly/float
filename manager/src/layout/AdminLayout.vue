<template>
  <div class="bg-[#F2F4F8] h-screen overflow-hidden flex relative">
    
    <!-- Background Blobs -->
    <div class="blob w-[500px] h-[500px] bg-blue-200 rounded-full -top-40 -left-20 mix-blend-multiply"></div>
    <div class="blob w-[600px] h-[600px] bg-purple-200 rounded-full bottom-0 right-0 mix-blend-multiply animation-delay-2000"></div>
    <div class="blob w-[300px] h-[300px] bg-pink-200 rounded-full top-1/2 left-1/3 mix-blend-multiply animation-delay-4000"></div>

    <!-- Sidebar -->
    <aside class="w-64 h-full glass-panel z-20 flex flex-col justify-between relative border-r border-white/40">
        <div>
            <div class="h-20 flex items-center gap-3 px-8 border-b border-white/30">
                <div class="w-10 h-10 rounded-xl bg-gradient-to-tr from-indigo-500 to-purple-500 flex items-center justify-center text-white text-xl shadow-lg shadow-indigo-200">
                    <i class="fa-solid fa-mountain-sun"></i>
                </div>
                <div>
                    <h1 class="font-extrabold text-xl text-gray-800 tracking-tight">Float Admin</h1>
                    <p class="text-[10px] text-gray-500 font-medium tracking-widest uppercase">Management</p>
                </div>
            </div>

            <nav class="px-4 py-8 space-y-2">
                <p class="px-4 text-[10px] font-bold text-gray-400 uppercase tracking-wider mb-2">Main</p>
                
                <RouterLink to="/" class="nav-item flex items-center gap-3 px-4 py-3 rounded-xl transition-all" :class="{ 'active': $route.name === 'dashboard', 'inactive': $route.name !== 'dashboard' }">
                    <i class="fa-solid fa-grid-2 text-sm w-5 text-center"></i>
                    <span class="text-sm font-bold">仪表盘 Dashboard</span>
                </RouterLink>
                
                <RouterLink to="/users" class="nav-item flex items-center gap-3 px-4 py-3 rounded-xl transition-all" :class="{ 'active': $route.name === 'users', 'inactive': $route.name !== 'users' }">
                    <i class="fa-solid fa-users text-sm w-5 text-center"></i>
                    <span class="text-sm font-bold">用户管理 Users</span>
                </RouterLink>

                <RouterLink to="/transactions" class="nav-item flex items-center gap-3 px-4 py-3 rounded-xl transition-all" :class="{ 'active': $route.name === 'transactions', 'inactive': $route.name !== 'transactions' }">
                    <i class="fa-solid fa-file-invoice-dollar text-sm w-5 text-center"></i>
                    <span class="text-sm font-bold">交易记录 Data</span>
                </RouterLink>

                <p class="px-4 text-[10px] font-bold text-gray-400 uppercase tracking-wider mb-2 mt-6">System</p>
                
                <RouterLink to="/versions" class="nav-item flex items-center gap-3 px-4 py-3 rounded-xl transition-all" :class="{ 'active': $route.name === 'versions', 'inactive': $route.name !== 'versions' }">
                    <i class="fa-solid fa-cloud-arrow-up text-sm w-5 text-center"></i>
                    <span class="text-sm font-bold">版本发布 Versions</span>
                </RouterLink>

                <a href="#" class="nav-item inactive flex items-center gap-3 px-4 py-3 rounded-xl transition-all">
                    <i class="fa-solid fa-envelope-open-text text-sm w-5 text-center"></i>
                    <span class="text-sm font-bold">反馈中心 Feedback</span>
                    <span class="ml-auto bg-red-500 text-white text-[10px] px-1.5 py-0.5 rounded-md font-bold">3</span>
                </a>
                <RouterLink to="/settings" class="nav-item flex items-center gap-3 px-4 py-3 rounded-xl transition-all" :class="{ 'active': $route.name === 'settings', 'inactive': $route.name !== 'settings' }">
                    <i class="fa-solid fa-gear text-sm w-5 text-center"></i>
                    <span class="text-sm font-bold">系统设置 Settings</span>
                </RouterLink>
            </nav>
        </div>

        <div class="p-4">
            <div class="glass-card p-3 rounded-xl flex items-center gap-3 cursor-pointer">
                <img src="https://api.dicebear.com/7.x/avataaars/svg?seed=Admin" class="w-10 h-10 rounded-full bg-white shadow-sm">
                <div class="flex-1 min-w-0">
                    <h4 class="text-sm font-bold text-gray-800 truncate">Admin User</h4>
                    <p class="text-[10px] text-green-500 flex items-center gap-1">
                        <span class="w-1.5 h-1.5 rounded-full bg-green-500"></span> Online
                    </p>
                </div>
                <i class="fa-solid fa-arrow-right-from-bracket text-gray-400 hover:text-red-500 transition" @click="handleLogout"></i>
            </div>
        </div>
    </aside>

    <main class="flex-1 flex flex-col relative z-10 overflow-hidden">
        <!-- Header -->
        <header class="h-20 px-8 flex items-center justify-between">
            <div>
                <h2 class="text-xl font-extrabold text-gray-800">{{ $route.meta.title || 'Float Admin' }}</h2>
                <p class="text-xs text-gray-500">{{ formatDate(new Date()) }}</p>
            </div>

            <div class="flex items-center gap-4">
                <div class="glass-panel px-4 py-2 rounded-full flex items-center gap-2 text-gray-500 w-64 focus-within:w-80 transition-all duration-300">
                    <i class="fa-solid fa-magnifying-glass text-xs"></i>
                    <input type="text" placeholder="搜索用户 ID、交易号..." class="bg-transparent w-full outline-none text-sm placeholder-gray-400">
                </div>
                
                <button class="w-10 h-10 rounded-full glass-panel flex items-center justify-center text-gray-600 hover:bg-white transition relative">
                    <i class="fa-regular fa-bell"></i>
                    <span class="absolute top-2 right-2.5 w-2 h-2 bg-red-500 rounded-full border border-white"></span>
                </button>
            </div>
        </header>

        <!-- Page Content -->
        <div class="flex-1 overflow-y-auto px-8 pb-8 custom-scroll">
            <RouterView />
        </div>
    </main>

  </div>
</template>

<script setup>
import { useRouter } from 'vue-router'

const router = useRouter()

const handleLogout = () => {
    // Clear token logic here
    router.push('/login')
}

const formatDate = (date) => {
    return new Intl.DateTimeFormat('zh-CN', {
        weekday: 'long',
        year: 'numeric', 
        month: 'long', 
        day: 'numeric'
    }).format(date)
}
</script>

<style scoped>
.nav-item.active {
    background: linear-gradient(90deg, #1f2937 0%, #374151 100%);
    color: white;
    box-shadow: 0 4px 15px rgba(31, 41, 55, 0.3);
}
.nav-item.inactive {
    color: #6b7280;
}
.nav-item.inactive:hover {
    background: rgba(255, 255, 255, 0.5);
    color: #111827;
}
.animation-delay-2000 { animation-delay: 2s; }
.animation-delay-4000 { animation-delay: 4s; }
</style>
