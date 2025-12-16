<template>
  <div>
    <!-- Stats Cards -->
    <div class="grid grid-cols-4 gap-6 mb-8">
        <div class="glass-panel p-6 rounded-2xl relative overflow-hidden group" v-for="(stat, index) in stats" :key="index">
            <div class="absolute -right-6 -top-6 w-24 h-24 rounded-full opacity-50 group-hover:scale-125 transition duration-500" :class="stat.bgClass"></div>
            <div class="relative z-10">
                <div class="flex justify-between items-start mb-4">
                    <div class="w-10 h-10 rounded-xl flex items-center justify-center text-lg" :class="stat.iconBgClass">
                        <i :class="stat.icon"></i>
                    </div>
                    <span class="text-xs font-bold px-2 py-1 rounded-lg" :class="stat.trendClass">{{ stat.trend }}</span>
                </div>
                <h3 class="text-3xl font-extrabold text-gray-800">{{ stat.value }}</h3>
                <p class="text-xs text-gray-400 font-medium mt-1">{{ stat.label }}</p>
            </div>
        </div>
    </div>

    <!-- Charts & Feedback -->
    <div class="grid grid-cols-3 gap-6 h-[500px] mb-8">
        <!-- Traffic Chart (CSS only for now as per design) -->
        <div class="col-span-2 glass-panel rounded-3xl p-6 flex flex-col">
            <div class="flex justify-between items-center mb-6">
                <div>
                    <h3 class="font-bold text-gray-800">流量趋势</h3>
                    <p class="text-xs text-gray-400">近 7 天用户活跃度 (DAU)</p>
                </div>
                <div class="flex bg-gray-100/50 p-1 rounded-lg">
                    <button class="px-3 py-1 bg-white rounded-md text-xs font-bold shadow-sm text-gray-800">周</button>
                    <button class="px-3 py-1 text-xs font-medium text-gray-500 hover:text-gray-700">月</button>
                </div>
            </div>
            
            <div class="flex-1 flex items-end justify-between gap-4 px-4 pb-2 border-b border-gray-100/50">
                <div v-for="(day, index) in trafficData" :key="index" class="w-full flex flex-col gap-2 group cursor-pointer">
                    <div class="w-full bg-indigo-500/20 rounded-t-xl relative h-32 group-hover:bg-indigo-500/30 transition">
                        <div class="absolute bottom-0 w-full bg-indigo-500 rounded-t-xl transition-all duration-500" :style="{ height: day.percent + '%' }"></div>
                        <div class="absolute -top-8 left-1/2 -translate-x-1/2 bg-gray-800 text-white text-[10px] px-2 py-1 rounded opacity-0 group-hover:opacity-100 transition">{{ day.value }}</div>
                    </div>
                    <span class="text-center text-xs text-gray-400 font-medium">{{ day.label }}</span>
                </div>
            </div>
        </div>

        <!-- Latest Feedback -->
        <div class="col-span-1 glass-panel rounded-3xl p-6 flex flex-col">
            <h3 class="font-bold text-gray-800 mb-4">最新反馈</h3>
            
            <div class="flex-1 overflow-y-auto pr-2 space-y-3 custom-scroll">
                <div v-for="feedback in feedbacks" :key="feedback.id" class="bg-white/40 p-3 rounded-xl hover:bg-white/60 transition cursor-pointer">
                    <div class="flex justify-between items-start mb-1">
                        <span class="text-xs font-bold text-gray-800">{{ feedback.title }}</span>
                        <span class="text-[10px] text-gray-400">{{ feedback.time }}</span>
                    </div>
                    <p class="text-xs text-gray-500 line-clamp-2">{{ feedback.content }}</p>
                    <div class="mt-2 flex gap-2" v-if="feedback.tags">
                        <span v-for="tag in feedback.tags" :key="tag.text" :class="tag.class" class="text-[10px] px-1.5 py-0.5 rounded">{{ tag.text }}</span>
                    </div>
                </div>
            </div>
            
            <button class="w-full mt-4 py-2 text-xs font-bold text-gray-500 hover:text-indigo-600 hover:bg-indigo-50 rounded-lg transition">
                查看全部反馈 <i class="fa-solid fa-arrow-right ml-1"></i>
            </button>
        </div>
    </div>

    <!-- Recent Users Table -->
    <div class="glass-panel rounded-3xl p-6">
        <div class="flex justify-between items-center mb-6">
            <h3 class="font-bold text-gray-800">最新注册用户</h3>
            <div class="flex gap-2">
                <button class="px-3 py-1.5 rounded-lg bg-white/50 text-xs font-bold text-gray-600 hover:bg-white transition border border-transparent hover:border-gray-200">
                    <i class="fa-solid fa-filter mr-1"></i> 筛选
                </button>
                <button class="px-3 py-1.5 rounded-lg bg-gray-900 text-xs font-bold text-white hover:bg-gray-800 transition">
                    <i class="fa-solid fa-download mr-1"></i> 导出
                </button>
            </div>
        </div>

        <div class="overflow-x-auto">
            <table class="w-full text-left border-collapse">
                <thead>
                    <tr class="text-xs text-gray-400 border-b border-gray-200/50">
                        <th class="pb-3 pl-2 font-medium">用户 ID</th>
                        <th class="pb-3 font-medium">昵称 / 头像</th>
                        <th class="pb-3 font-medium">注册时间</th>
                        <th class="pb-3 font-medium">状态</th>
                        <th class="pb-3 font-medium">会员等级</th>
                        <th class="pb-3 font-medium text-right pr-2">操作</th>
                    </tr>
                </thead>
                <tbody class="text-sm">
                    <tr v-for="user in recentUsers" :key="user.id" class="table-row group border-b border-gray-100/30 hover:bg-white/40 transition">
                        <td class="py-3 pl-2 text-gray-500 font-mono">{{ user.id }}</td>
                        <td class="py-3">
                            <div class="flex items-center gap-3">
                                <img :src="`https://api.dicebear.com/7.x/avataaars/svg?seed=${user.name}`" class="w-8 h-8 rounded-full bg-orange-100">
                                <span class="font-bold text-gray-700">{{ user.name }}</span>
                            </div>
                        </td>
                        <td class="py-3 text-gray-500">{{ user.registeredAt }}</td>
                        <td class="py-3">
                            <span :class="user.statusClass" class="text-[10px] px-2 py-1 rounded-full font-bold">{{ user.status }}</span>
                        </td>
                        <td class="py-3">
                            <span v-if="user.isPro" class="text-indigo-500 font-bold text-xs"><i class="fa-solid fa-crown mr-1"></i>PRO</span>
                            <span v-else class="text-gray-400 text-xs">Free</span>
                        </td>
                        <td class="py-3 text-right pr-2">
                            <button class="text-gray-400 hover:text-indigo-600 px-2 transition"><i class="fa-solid fa-eye"></i></button>
                            <button class="text-gray-400 hover:text-red-500 px-2 transition"><i class="fa-solid fa-ban"></i></button>
                        </td>
                    </tr>
                </tbody>
            </table>
        </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { getSystemOverview, getUsers, getTransactionStats } from '@/api/admin'

const stats = ref([
    { 
        value: '-', label: '总注册用户', trend: '+0%', 
        icon: 'fa-solid fa-users', bgClass: 'bg-blue-100', iconBgClass: 'bg-blue-50 text-blue-500', trendClass: 'text-green-500 bg-green-50' 
    },
    { 
        value: '¥ 0', label: '今日交易流水', trend: '+0%', 
        icon: 'fa-solid fa-sack-dollar', bgClass: 'bg-green-100', iconBgClass: 'bg-green-50 text-green-500', trendClass: 'text-green-500 bg-green-50' 
    },
    { 
        value: '0', label: '今日新增记录 (笔)', trend: '0%', 
        icon: 'fa-solid fa-file-invoice', bgClass: 'bg-purple-100', iconBgClass: 'bg-purple-50 text-purple-500', trendClass: 'text-gray-400 bg-gray-50' 
    },
    { 
        value: '-', label: '系统状态', trend: 'Active', 
        icon: 'fa-solid fa-server', bgClass: 'bg-orange-100', iconBgClass: 'bg-orange-50 text-orange-500', trendClass: 'text-blue-500 bg-blue-50' 
    }
])

const trafficData = ref([
    { label: 'Mon', value: '2,400', percent: 60 },
    { label: 'Tue', value: '3,100', percent: 85 },
    { label: 'Wed', value: '1,800', percent: 45 },
    { label: 'Thu', value: '2,800', percent: 70 },
    { label: 'Fri', value: '3,600', percent: 90 },
    { label: 'Sat', value: '4,000', percent: 100 },
    { label: 'Sun', value: '3,800', percent: 95 },
])

const feedbacks = ref([
    { id: 1, title: '无法导出 PDF', time: '10min ago', content: '我在尝试导出上个月账单时，一直提示生成失败，请问怎么解决...', tags: [{ text: 'Bug', class: 'bg-red-100 text-red-500' }, { text: '待处理', class: 'bg-blue-100 text-blue-500' }] },
    { id: 2, title: '建议增加暗黑模式', time: '2h ago', content: '晚上记账太亮了，希望可以适配系统的深色模式。', tags: [{ text: '建议', class: 'bg-green-100 text-green-500' }] },
])

const recentUsers = ref([])

const loadData = async () => {
    try {
        // 1. System Overview (Total Users)
        const overview = await getSystemOverview()
        if (overview) {
            stats.value[0].value = overview.total_users?.toLocaleString() || '0'
            stats.value[3].value = overview.system_status || 'Normal'
        }

        // 2. Transaction Stats (Total Flow, Count)
        // Not all implementations might be ready, wrapping in try/catch independently or just checking support
        try {
            const transStats = await getTransactionStats()
            if (transStats) {
                // Assuming the structure, if different we handle graceful defaults
                stats.value[1].value = `¥ ${transStats.total_amount?.toLocaleString() || '0'}`
                stats.value[2].value = transStats.total_count?.toLocaleString() || '0'
            }
        } catch (e) {
            console.warn("Failed to load transaction stats", e)
        }

        // 3. Recent Users
        const usersResp = await getUsers({ page: 1, page_size: 5 })
        if (usersResp && usersResp.items) {
            recentUsers.value = usersResp.items.map(u => ({
                id: `#${u.id}`,
                name: u.display_name || u.username,
                registeredAt: new Date(u.created_at).toLocaleString(),
                status: '正常', // No status field in response yet
                statusClass: 'bg-green-100 text-green-600',
                isPro: u.membership_level !== 'FREE'
            }))
        }
    } catch (error) {
        console.error("Failed to load dashboard data", error)
    }
}

onMounted(() => {
    loadData()
})
</script>
