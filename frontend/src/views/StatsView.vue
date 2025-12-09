<script setup>
import { ref, computed, onMounted } from 'vue'
import { useToast } from '../composables/useToast'
import { transactionAPI } from '@/api/transaction'
import GlassCard from '../components/GlassCard.vue'

const { showToast } = useToast()

// 数据
const startDate = ref(new Date(new Date().getFullYear(), new Date().getMonth(), 1))
const endDate = ref(new Date())
const monthlyExpense = ref(0)
const previousMonthExpense = ref(0)
const categoryStats = ref([])
const weeklyData = ref([])
const loading = ref(false)
const showMonthRangePicker = ref(false)
const hoveredPointIndex = ref(null)
const tooltipPos = ref({ x: 0, y: 0 })

// 计算属性
const monthRangeDisplay = computed(() => {
    const months = ['1月', '2月', '3月', '4月', '5月', '6月', '7月', '8月', '9月', '10月', '11月', '12月']
    const startMonth = months[startDate.value.getMonth()]
    const endMonth = months[endDate.value.getMonth()]

    if (startDate.value.getFullYear() === endDate.value.getFullYear()) {
        return `${startMonth}-${endMonth}`
    }
    return `${startDate.value.getFullYear()}年${startMonth}-${endDate.value.getFullYear()}年${endMonth}`
})

const expenseChangePercent = computed(() => {
    if (previousMonthExpense.value === 0) return 0
    return Math.round(((monthlyExpense.value - previousMonthExpense.value) / previousMonthExpense.value) * 100)
})

const expenseChangeSign = computed(() => {
    return expenseChangePercent.value >= 0 ? '+' : ''
})

const maxCategoryAmount = computed(() => {
    if (categoryStats.value.length === 0) return 1
    return Math.max(...categoryStats.value.map(c => c.total_amount))
})

const getCategoryColor = (categoryName) => {
    const colorMap = {
        '购物消费': 'pink',
        '餐饮美食': 'orange',
        '交通出行': 'blue',
        '娱乐休闲': 'purple',
        '生活服务': 'green',
        '教育培训': 'cyan',
        '其他': 'gray'
    }
    return colorMap[categoryName] || 'indigo'
}

const getCategoryIcon = (icon) => {
    return icon || 'fa-receipt'
}

// 计算当月的四个周
const getMonthWeeks = (start, end) => {
    const weeks = []
    const monthStart = new Date(start)
    monthStart.setHours(0, 0, 0, 0)
    const monthEnd = new Date(end)
    monthEnd.setHours(23, 59, 59, 999)
    
    // 生成4周的数据
    for (let i = 0; i < 4; i++) {
        const weekStart = new Date(monthStart)
        weekStart.setDate(weekStart.getDate() + i * 7)
        weekStart.setHours(0, 0, 0, 0)
        
        const weekEnd = new Date(weekStart)
        weekEnd.setDate(weekEnd.getDate() + 6)
        weekEnd.setHours(23, 59, 59, 999)
        
        // 确保周数据不超出月份范围
        if (weekStart.toDateString() > monthEnd.toDateString()) break
        
        weeks.push({
            weekNum: i + 1,
            start: weekStart,
            end: weekEnd < monthEnd ? weekEnd : monthEnd
        })
    }
    
    return weeks
}

// 获取本月统计
const loadMonthlyTrendData = async () => {
    try {
        const start = startDate.value.toISOString().split('T')[0]
        const end = endDate.value.toISOString().split('T')[0]

        const response = await transactionAPI.getTransactions({
            type: 'expense',
            start_date: start,
            end_date: end,
            page: 1,
            page_size: 1000
        })

        // 按日期分组
        const dailyMap = {}

        if (response.data?.items) {
            response.data.items.forEach(item => {
                // transaction_date 可能是 YYYY-MM-DD 格式或 ISO 格式，两种都处理
                let date = item.transaction_date
                if (date.includes('T')) {
                    date = date.split('T')[0]
                }
                if (!dailyMap[date]) {
                    dailyMap[date] = 0
                }
                dailyMap[date] += item.amount
            })
        }

        // 计算每周的消费总额
        const weeks = getMonthWeeks(startDate.value, endDate.value)
        const weeklyTotals = weeks.map(week => {
            let total = 0
            let current = new Date(week.start)
            
            while (current <= week.end) {
                const dateStr = current.toISOString().split('T')[0]
                total += dailyMap[dateStr] || 0
                current.setDate(current.getDate() + 1)
            }
            
            return {
                weekNum: week.weekNum,
                amount: total
            }
        })

        weeklyData.value = weeklyTotals
    } catch (error) {
        console.error('Failed to load monthly trend data:', error)
    }
}

// 获取月度及上月统计
const loadMonthlyStats = async () => {
    try {
        const start = startDate.value.toISOString().split('T')[0]
        const end = endDate.value.toISOString().split('T')[0]

        const response = await transactionAPI.getStatistics({
            type: 'expense',
            start_date: start,
            end_date: end
        })

        monthlyExpense.value = response.data?.total_expense || 0

        // 获取上个月的开始和结束日期
        const prevStart = new Date(startDate.value)
        prevStart.setMonth(prevStart.getMonth() - 1)
        const prevEnd = new Date(endDate.value)
        prevEnd.setMonth(prevEnd.getMonth() - 1)

        const prevStartStr = prevStart.toISOString().split('T')[0]
        const prevEndStr = prevEnd.toISOString().split('T')[0]

        const prevResponse = await transactionAPI.getStatistics({
            type: 'expense',
            start_date: prevStartStr,
            end_date: prevEndStr
        })

        previousMonthExpense.value = prevResponse.data?.total_expense || 0
    } catch (error) {
        console.error('Failed to load monthly stats:', error)
    }
}

// 获取分类统计
const loadCategoryStats = async () => {
    try {
        const start = startDate.value.toISOString().split('T')[0]
        const end = endDate.value.toISOString().split('T')[0]

        const response = await transactionAPI.getCategoryStatistics({
            start_date: start,
            end_date: end
        })

        categoryStats.value = (response.data || [])
            .filter(item => item.category?.type === 'expense')
            .sort((a, b) => b.total_amount - a.total_amount)
    } catch (error) {
        console.error('Failed to load category stats:', error)
    }
}

// 加载所有数据
const loadData = async () => {
  loading.value = true
  try {
    await Promise.all([
      loadMonthlyStats(),
      loadCategoryStats(),
      loadMonthlyTrendData()
    ])
    } catch (error) {
        showToast('加载统计数据失败', 'error')
    } finally {
        loading.value = false
    }
}

// 设置日期范围
const setDateRange = (start, end) => {
    startDate.value = start
    endDate.value = end
    showMonthRangePicker.value = false
    loadData()
}

// 快速选择按钮
const selectThisMonth = () => {
    const now = new Date()
    const start = new Date(now.getFullYear(), now.getMonth(), 1)
    const end = new Date(now.getFullYear(), now.getMonth() + 1, 0)
    setDateRange(start, end)
}

const selectLastMonth = () => {
    const now = new Date()
    const start = new Date(now.getFullYear(), now.getMonth() - 1, 1)
    const end = new Date(now.getFullYear(), now.getMonth(), 0)
    setDateRange(start, end)
}

const selectLast3Months = () => {
    const now = new Date()
    const start = new Date(now.getFullYear(), now.getMonth() - 2, 1)
    const end = new Date(now.getFullYear(), now.getMonth() + 1, 0)
    setDateRange(start, end)
}

// 格式化金额
const formatAmount = (amount) => {
    return amount.toFixed(2)
}

// 计算进度条宽度
const getProgressWidth = (amount) => {
    if (maxCategoryAmount.value === 0) return 0
    return (amount / maxCategoryAmount.value) * 100
}

// 计算最大周消费
const maxWeeklyAmount = computed(() => {
    if (weeklyData.value.length === 0) return 1
    return Math.max(...weeklyData.value.map(d => d.amount))
})

// 计算周消费百分比
const getWeeklyHeight = (amount) => {
    if (maxWeeklyAmount.value === 0) return 5
    const percent = (amount / maxWeeklyAmount.value) * 100
    return Math.max(5, percent)
}

// 获取周标签
const getWeekDayLabel = (weekNum) => {
    return `第${weekNum}周`
}

// 处理数据点hover
const handlePointHover = (index, event) => {
    hoveredPointIndex.value = index
    const rect = event.currentTarget.getBoundingClientRect()
    tooltipPos.value = {
        x: rect.left + rect.width / 2,
        y: rect.top - 8
    }
}

const handlePointLeave = () => {
    hoveredPointIndex.value = null
}

// 计算每周的x坐标（基于viewBox 100宽度）
const getWeekXPosition = (index) => {
    const count = weeklyData.value.length
    if (count === 1) return 50
    // 均匀分布：间隔 = 总宽度 / (周数 + 1)
    // 点位置 = 间隔 * (索引 + 1)
    const interval = 100 / (count + 1)
    return interval * (index + 1)
}

// 生成折线图SVG路径
const generateLinePath = () => {
    if (weeklyData.value.length === 0) return ''
    
    const viewBoxHeight = 120
    const topPadding = 15
    const bottomPadding = 35
    const chartHeight = viewBoxHeight - topPadding - bottomPadding
    
    const maxAmount = maxWeeklyAmount.value || 1
    const points = weeklyData.value.map((week, index) => {
        const x = getWeekXPosition(index)
        const y = topPadding + (1 - week.amount / maxAmount) * chartHeight
        return { x, y, amount: week.amount }
    })
    
    // 生成路径
    let pathData = `M ${points[0].x} ${points[0].y}`
    for (let i = 1; i < points.length; i++) {
        const p1 = points[i - 1]
        const p2 = points[i]
        const cpx = (p1.x + p2.x) / 2
        const cpy1 = p1.y
        const cpy2 = p2.y
        pathData += ` C ${cpx} ${cpy1}, ${cpx} ${cpy2}, ${p2.x} ${p2.y}`
    }
    
    return { path: pathData, points, viewBoxHeight }
}

const lineChart = computed(() => generateLinePath())

onMounted(() => {
    loadData()
})
</script>

<template>
    <div>
        <!-- Header -->
        <div class="px-6 pt-12 pb-2 flex justify-between items-center relative z-10 animate-enter">
            <div>
                <h1 class="text-2xl font-extrabold text-gray-800">收支统计</h1>
                <p class="text-xs text-gray-400 mt-1">每一笔花费都值得被记录</p>
            </div>
            <div class="relative">
                <button @click="showMonthRangePicker = !showMonthRangePicker"
                    class="glass-card px-3 py-1.5 rounded-full flex items-center gap-2 text-sm font-bold text-gray-700 shadow-sm cursor-pointer hover:bg-white/80 transition">
                    <span>{{ monthRangeDisplay }}</span>
                    <i :class="showMonthRangePicker ? 'fa-chevron-up' : 'fa-chevron-down'"
                        class="fa-solid text-xs text-gray-400 transition"></i>
                </button>
                <div v-if="showMonthRangePicker"
                    class="absolute right-0 top-full mt-2 bg-white rounded-2xl shadow-lg p-4 z-10 w-64">
                    <div class="space-y-2">
                        <button @click="selectThisMonth"
                            class="w-full text-left px-3 py-2 hover:bg-gray-100 rounded-lg text-sm text-gray-700">本月</button>
                        <button @click="selectLastMonth"
                            class="w-full text-left px-3 py-2 hover:bg-gray-100 rounded-lg text-sm text-gray-700">上个月</button>
                        <button @click="selectLast3Months"
                            class="w-full text-left px-3 py-2 hover:bg-gray-100 rounded-lg text-sm text-gray-700">最近3个月</button>
                        <div class="border-t border-gray-200 pt-2 mt-2">
                            <p class="text-xs text-gray-500 px-3 mb-2">自定义范围</p>
                            <div class="grid grid-cols-2 gap-2 px-1">
                                <select v-model="startDate" class="px-2 py-1 text-xs border border-gray-200 rounded-lg">
                                    <option v-for="m in 12" :key="m"
                                        :value="new Date(new Date().getFullYear(), m - 1, 1)">
                                        {{ m }}月
                                    </option>
                                </select>
                                <select v-model="endDate" class="px-2 py-1 text-xs border border-gray-200 rounded-lg">
                                    <option v-for="m in 12" :key="m" :value="new Date(new Date().getFullYear(), m, 0)">
                                        {{ m }}月
                                    </option>
                                </select>
                            </div>
                            <button @click="loadData; showMonthRangePicker = false"
                                class="w-full mt-2 px-3 py-2 bg-indigo-600 text-white text-xs rounded-lg font-bold hover:bg-indigo-700 transition">
                                确定
                            </button>
                        </div>
                    </div>
                </div>
            </div>
        </div>

        <div class="px-6 pt-6 pb-32">
            <!-- Loading State -->
            <div v-if="loading" class="text-center py-12">
                <i class="fa-solid fa-spinner fa-spin text-3xl text-gray-400"></i>
                <p class="text-gray-400 mt-3">加载中...</p>
            </div>

            <template v-else>
                <!-- Total Expenditure -->
                <div class="text-center mb-8 animate-enter delay-100">
                    <p class="text-xs font-medium text-gray-400 tracking-wider uppercase mb-2">期间总支出</p>
                    <h2 class="text-4xl font-extrabold text-gray-800">¥ {{ formatAmount(monthlyExpense).split('.')[0]
                        }}.<span class="text-2xl text-gray-400 font-bold">{{ formatAmount(monthlyExpense).split('.')[1]
                            }}</span></h2>
                    <div class="flex justify-center items-center gap-2 mt-2">
                        <span
                            :class="expenseChangePercent >= 0 ? 'bg-red-100 text-red-500' : 'bg-green-100 text-green-500'"
                            class="text-[10px] px-2 py-0.5 rounded-full font-bold">
                            {{ expenseChangeSign }}{{ Math.abs(expenseChangePercent) }}%
                        </span>
                        <span class="text-xs text-gray-400">对比前期</span>
                    </div>
                </div>

                <!-- Monthly Trend Chart -->
                <GlassCard class="p-5 rounded-3xl shadow-sm mb-6 animate-enter delay-200">
                    <div class="flex justify-between items-center mb-6">
                        <h3 class="font-bold text-gray-800 text-sm">月消费趋势</h3>
                        <i class="fa-solid fa-arrow-trend-up text-indigo-500 text-sm"></i>
                    </div>

                    <div v-if="weeklyData.length === 0" class="text-center py-12 text-gray-400">
                        <p class="text-sm">本月暂无消费数据</p>
                    </div>

                    <div v-else class="relative w-full">
                        <svg class="w-full" style="height: 180px; display: block;" viewBox="0 0 100 120" preserveAspectRatio="none">
                            <defs>
                                <!-- 主渐变 -->
                                <linearGradient id="lineGradient" x1="0%" y1="0%" x2="0%" y2="100%">
                                    <stop offset="0%" style="stop-color:rgb(139, 92, 246);stop-opacity:0.4" />
                                    <stop offset="50%" style="stop-color:rgb(99, 102, 241);stop-opacity:0.2" />
                                    <stop offset="100%" style="stop-color:rgb(99, 102, 241);stop-opacity:0" />
                                </linearGradient>
                                <!-- 发光效果 -->
                                <filter id="glow">
                                    <feGaussianBlur stdDeviation="1.5" result="coloredBlur"/>
                                    <feMerge>
                                        <feMergeNode in="coloredBlur"/>
                                        <feMergeNode in="SourceGraphic"/>
                                    </feMerge>
                                </filter>
                            </defs>
                            
                            <!-- 背景 -->
                            <rect width="100" height="120" fill="none"/>
                            
                            <!-- 填充面积 -->
                            <path v-if="lineChart.path" :d="lineChart.path + ' L 95 85 L 5 85 Z'" 
                                fill="url(#lineGradient)" filter="url(#glow)"/>
                            
                            <!-- 折线 -->
                            <path v-if="lineChart.path" :d="lineChart.path" 
                                fill="none" stroke="rgb(99, 102, 241)" stroke-width="1.8" stroke-linecap="round" stroke-linejoin="round" filter="url(#glow)"/>
                            
                            <!-- 数据点光晕 -->
                            <circle v-for="(point, index) in lineChart.points" :key="'glow-' + index"
                                :cx="point.x" :cy="point.y" r="3.5" fill="rgb(99, 102, 241)" opacity="0.1"/>
                            
                            <!-- 数据点 - 支持hover显示tooltip -->
                            <circle v-for="(point, index) in lineChart.points" :key="'point-' + index"
                                :cx="point.x" :cy="point.y" r="2.5" fill="white" stroke="rgb(99, 102, 241)" stroke-width="1.5" filter="url(#glow)"
                                class="cursor-pointer transition-all duration-200"
                                @mouseenter="handlePointHover(index, $event)"
                                @mouseleave="handlePointLeave">
                            </circle>
                        </svg>
                        
                        <!-- 自定义Tooltip -->
                        <div v-if="hoveredPointIndex !== null" 
                            class="fixed bg-gray-800 text-white text-sm px-3 py-1 rounded shadow-lg z-50 pointer-events-none"
                            :style="{ 
                                left: tooltipPos.x + 'px', 
                                top: tooltipPos.y + 'px',
                                transform: 'translate(-50%, -100%)'
                            }">
                            ¥{{ formatAmount(weeklyData[hoveredPointIndex].amount) }}
                        </div>
                        
                        <!-- 标签层 - 只显示周标签 -->
                        <div class="grid gap-2 mt-4 px-2" :style="{ gridTemplateColumns: `repeat(${weeklyData.length}, 1fr)` }">
                            <div v-for="(week, index) in weeklyData" :key="'label-' + index"
                                class="flex flex-col items-center justify-center">
                                <div class="text-[10px] font-medium text-gray-500">
                                    {{ getWeekDayLabel(week.weekNum) }}
                                </div>
                            </div>
                        </div>
                    </div>
                </GlassCard>

                <!-- Category Breakdown -->
                <div class="animate-enter delay-300">
                    <h3 class="font-bold text-gray-800 text-sm mb-4">花哪里了?</h3>
                    <div v-if="categoryStats.length === 0" class="text-center py-8">
                        <p class="text-gray-400 text-sm">期间暂无支出记录</p>
                    </div>
                    <div v-else class="space-y-3 pb-6">
                        <GlassCard v-for="(category, index) in categoryStats" :key="index"
                            class="p-4 rounded-2xl flex items-center justify-between shadow-sm">
                            <div class="flex items-center gap-3 w-full">
                                <div :class="`bg-${getCategoryColor(category.category?.name)}-100 text-${getCategoryColor(category.category?.name)}-500`"
                                    class="w-10 h-10 rounded-xl flex items-center justify-center shadow-inner">
                                    <i :class="`fa-solid ${getCategoryIcon(category.category?.icon)}`"></i>
                                </div>
                                <div class="flex-1">
                                    <div class="flex justify-between items-center mb-1">
                                        <span class="text-sm font-bold text-gray-800">{{ category.category?.name ||
                                            '未分类' }}</span>
                                        <span class="text-sm font-bold text-gray-800">¥ {{
                                            formatAmount(category.total_amount) }}</span>
                                    </div>
                                    <div class="w-full bg-gray-200/50 h-1.5 rounded-full overflow-hidden">
                                        <div :class="`bg-${getCategoryColor(category.category?.name)}-500`"
                                            class="h-full rounded-full transition-all duration-300"
                                            :style="{ width: getProgressWidth(category.total_amount) + '%' }"></div>
                                    </div>
                                </div>
                            </div>
                        </GlassCard>
                    </div>
                </div>
            </template>
        </div>
    </div>
</template>

<style scoped>
/* 柱状图生长动画 */
@keyframes growUp {
    from {
        height: 0;
    }
}

.bar-grow {
    animation: growUp 1s ease-out forwards;
}
</style>
