<template>
  <div class="flex-1 overflow-y-auto px-8 pb-8 custom-scroll">
    <div class="grid grid-cols-3 gap-6">
        <!-- Release Form -->
        <div class="col-span-2 space-y-6">
            <div class="glass-panel p-8 rounded-3xl">
                <h3 class="font-bold text-gray-800 mb-6">发布新版本</h3>
                <form @submit.prevent="handleUpload">
                    <div class="grid grid-cols-2 gap-6 mb-6">
                        <div>
                            <label class="block text-xs font-bold text-gray-500 mb-2">版本号 (Version Name)</label>
                            <input v-model="form.versionName" type="text" placeholder="e.g. 1.3.0" required class="w-full bg-white/50 border border-gray-200 rounded-xl px-4 py-3 text-sm focus:border-indigo-500 focus:bg-white transition outline-none">
                        </div>
                        <div>
                            <label class="block text-xs font-bold text-gray-500 mb-2">构建号 (Version Code)</label>
                            <input v-model="form.versionCode" type="number" placeholder="e.g. 105" required class="w-full bg-white/50 border border-gray-200 rounded-xl px-4 py-3 text-sm focus:border-indigo-500 focus:bg-white transition outline-none">
                        </div>
                    </div>

                    <div 
                        class="border-2 border-dashed border-indigo-200 bg-indigo-50/50 rounded-2xl p-8 flex flex-col items-center justify-center text-center cursor-pointer hover:bg-indigo-50 transition mb-6 group relative"
                        @click="triggerFileSelect"
                        @dragover.prevent
                        @drop.prevent="handleDrop"
                    >
                        <input type="file" ref="fileInput" class="hidden" accept=".apk,.ipa" @change="handleFileChange">
                        
                        <div v-if="!selectedFile" class="flex flex-col items-center">
                            <div class="w-16 h-16 rounded-full bg-white text-indigo-500 shadow-md flex items-center justify-center text-2xl mb-4 group-hover:scale-110 transition">
                                <i class="fa-solid fa-cloud-arrow-up"></i>
                            </div>
                            <p class="text-sm font-bold text-gray-700">点击或拖拽上传安装包 (.apk / .ipa)</p>
                            <p class="text-xs text-gray-400 mt-1">最大支持 200MB</p>
                        </div>
                        <div v-else class="flex flex-col items-center">
                             <div class="w-16 h-16 rounded-full bg-green-100 text-green-600 shadow-md flex items-center justify-center text-2xl mb-4">
                                <i class="fa-solid fa-check"></i>
                            </div>
                            <p class="text-sm font-bold text-gray-700">{{ selectedFile.name }}</p>
                            <p class="text-xs text-gray-400 mt-1">{{ (selectedFile.size / 1024 / 1024).toFixed(2) }} MB</p>
                        </div>
                    </div>

                    <div class="mb-6">
                        <label class="block text-xs font-bold text-gray-500 mb-2">更新日志 (Changelog)</label>
                        <textarea v-model="form.changelogContent" rows="5" class="w-full bg-white/50 border border-gray-200 rounded-xl px-4 py-3 text-sm focus:border-indigo-500 focus:bg-white transition outline-none resize-none" placeholder="- 修复了已知 Bug&#10;- 优化了启动速度"></textarea>
                    </div>

                    <div class="flex items-center justify-between">
                        <label class="flex items-center gap-2 cursor-pointer">
                            <input type="checkbox" v-model="form.isForceUpdate" class="w-4 h-4 text-indigo-600 rounded border-gray-300 focus:ring-indigo-500">
                            <span class="text-sm font-bold text-gray-600">强制更新</span>
                        </label>
                        <button 
                            type="submit" 
                            :disabled="uploading || !selectedFile"
                            class="px-8 py-3 bg-gray-900 text-white text-sm font-bold rounded-xl shadow-lg hover:scale-105 transition disabled:opacity-50 disabled:cursor-not-allowed"
                        >
                            <span v-if="uploading"><i class="fa-solid fa-spinner fa-spin mr-2"></i> 发布中...</span>
                            <span v-else>立即发布</span>
                        </button>
                    </div>
                </form>
            </div>
        </div>

        <!-- History List -->
        <div class="col-span-1">
            <div class="glass-panel p-6 rounded-3xl h-full flex flex-col">
                <h3 class="font-bold text-gray-800 mb-4">发布历史</h3>
                <div class="flex-1 overflow-y-auto pr-2 space-y-4 custom-scroll">
                    <div v-for="(version, index) in history" :key="version.id" class="relative pl-6 border-l-2 border-gray-200 pb-2">
                        <div class="absolute -left-[9px] top-0 w-4 h-4 rounded-full border-2 border-white" :class="index === 0 ? 'bg-green-500' : 'bg-gray-300'"></div>
                        <div class="p-3 rounded-xl shadow-sm" :class="index === 0 ? 'bg-white/60' : 'bg-white/40'">
                            <div class="flex justify-between items-start mb-1">
                                <span class="font-bold text-gray-800 text-sm">{{ version.version_name }} (Build {{ version.version_code }})</span>
                                <span v-if="index === 0" class="text-[10px] bg-green-100 text-green-600 px-1.5 py-0.5 rounded">Active</span>
                                <span v-else class="text-[10px] text-gray-400 cursor-pointer hover:text-red-500">回滚</span>
                            </div>
                            <p class="text-xs text-gray-400 mb-2">{{ formatDate(version.created_at) }}</p>
                            <p class="text-xs text-gray-600 line-clamp-2 white-space-pre-line">{{ version.description }}</p>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { getAppUpdateHistory, uploadAppUpdate } from '@/api/admin'

const fileInput = ref(null)
const selectedFile = ref(null)
const uploading = ref(false)
const history = ref([])

const form = reactive({
    versionName: '',
    versionCode: '',
    changelogContent: '',
    isForceUpdate: false
})

const triggerFileSelect = () => {
    fileInput.value.click()
}

const handleFileChange = (e) => {
    if (e.target.files.length > 0) {
        selectedFile.value = e.target.files[0]
    }
}

const handleDrop = (e) => {
    if (e.dataTransfer.files.length > 0) {
        selectedFile.value = e.dataTransfer.files[0]
    }
}

const fetchHistory = async () => {
    try {
        const data = await getAppUpdateHistory('android')
        history.value = data || []
    } catch (e) {
        console.error("Failed to fetch history", e)
    }
}

const handleUpload = async () => {
    if (!selectedFile.value) return

    uploading.value = true
    const formData = new FormData()
    formData.append('file', selectedFile.value)
    formData.append('version_code', form.versionCode)
    formData.append('version_name', form.versionName)
    formData.append('platform', 'android') // Default
    formData.append('update_type', 'minor')
    formData.append('title', `Release v${form.versionName}`)
    formData.append('description', form.changelogContent) // Use text content for description
    formData.append('changelog', '{}') // Send empty JSON for changelog field as it's not used
    formData.append('is_force_update', form.isForceUpdate)

    try {
        await uploadAppUpdate(formData)
        alert('发布成功！')
        // Reset
        selectedFile.value = null
        form.versionName = ''
        form.versionCode = ''
        form.changelogContent = ''
        form.isForceUpdate = false
        fileInput.value.value = '' // clear input
        
        fetchHistory()
    } catch (e) {
        console.error("Upload failed", e)
        alert('发布失败，请重试')
    } finally {
        uploading.value = false
    }
}

const formatDate = (dateStr) => {
    if (!dateStr) return ''
    return new Date(dateStr).toLocaleString()
}

const getChangelogText = (changelogJson) => {
    try {
        if (typeof changelogJson === 'string') {
             // Handle double encoded json if any
             try {
                 const parsed = JSON.parse(changelogJson)
                 return Array.isArray(parsed) ? parsed.join('\n') : parsed
             } catch {
                 return changelogJson
             }
        }
        return Array.isArray(changelogJson) ? changelogJson.join('\n') : changelogJson
    } catch (e) {
        return changelogJson
    }
}

onMounted(() => {
    fetchHistory()
})
</script>
