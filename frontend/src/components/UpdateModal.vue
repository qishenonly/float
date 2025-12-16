<script setup>
import { computed } from 'vue'

const props = defineProps({
  show: Boolean,
  progress: Number,
  status: {
    type: String,
    default: 'downloading' // downloading, installing, error
  },
  version: String,
  description: String
})

const emit = defineEmits(['close', 'confirm', 'cancel'])

const statusText = computed(() => {
  switch (props.status) {
    case 'prompt':
       return '新版本已准备就绪'
    case 'downloading':
      return '正在下载更新包...'
    case 'installing':
      return '下载完成，准备安装...'
    case 'error':
      return '下载失败'
    default:
      return '处理中...'
  }
})

const formattedDescription = computed(() => {
  if (!props.description) return '修复了一些已知问题，优化用户体验。'
  
  let text = props.description
  
  // Ensure headers start on new lines
  text = text.replace(/([✨🔧])/g, '\n\n$1 ')
  
  // Format specific "Fixed" items as bullets if they are buried in text
  // Match "修复了" but avoid double replacing if it's already at start of line
  text = text.replace(/([^•\n])(修复了)/g, '$1\n• $2')
  // Also handle cases where "现在" might start a sentence
  text = text.replace(/([^•\n])(现在)/g, '$1\n• $2')
  
  return text.trim()
})
</script>

<template>
  <Transition name="fade">
    <div v-if="show" class="fixed inset-0 z-[100] flex items-center justify-center px-4">
      <!-- Backdrop -->
      <div class="absolute inset-0 bg-black/40 backdrop-blur-sm"></div>

      <!-- Modal -->
      <div class="bg-white rounded-3xl w-full max-w-sm p-6 relative z-10 shadow-2xl animate-scale">
        <div class="text-center mb-6">
          <div class="w-16 h-16 bg-indigo-50 rounded-full flex items-center justify-center mx-auto mb-4 text-indigo-600 text-2xl">
            <i v-if="status === 'prompt'" class="fa-solid fa-gift animate-bounce"></i>
            <i v-else-if="status === 'error'" class="fa-solid fa-circle-exclamation text-red-500"></i>
            <i v-else class="fa-solid fa-cloud-arrow-down animate-bounce"></i>
          </div>
          <h3 class="text-xl font-bold text-gray-800 mb-1">
            {{ status === 'prompt' ? `发现新版本 ${version}` : (status === 'error' ? '更新失败' : `正在更新 ${version}`) }}
          </h3>
          <p class="text-sm text-gray-500">{{ statusText }}</p>
        </div>

        <!-- Prompt Content -->
        <div v-if="status === 'prompt'">
           <div class="bg-gray-50 rounded-xl p-4 mb-6 text-left max-h-40 overflow-y-auto">
              <p class="text-sm text-gray-600 whitespace-pre-wrap leading-relaxed">{{ formattedDescription }}</p>
           </div>
           <div class="flex gap-3">
             <button @click="$emit('cancel')" class="flex-1 py-3 bg-gray-100 text-gray-600 font-bold rounded-xl hover:bg-gray-200 transition">稍后再说</button>
             <button @click="$emit('confirm')" class="flex-1 py-3 bg-indigo-600 text-white font-bold rounded-xl hover:bg-indigo-700 transition shadow-lg shadow-indigo-200">立即更新</button>
           </div>
        </div>

        <!-- Progress Bar -->
        <div v-else-if="status === 'downloading' || status === 'installing'" class="mb-2 relative h-4 bg-gray-100 rounded-full overflow-hidden">
          <div 
            class="absolute top-0 left-0 h-full bg-gradient-to-r from-indigo-500 to-violet-500 transition-all duration-300 ease-out"
            :style="{ width: `${progress}%` }"
          ></div>
        </div>
        <div v-if="status === 'downloading' || status === 'installing'" class="flex justify-between text-xs text-gray-400 font-medium mb-6">
          <span>{{ progress }}%</span>
          <span>100%</span>
        </div>

        <!-- Actions -->
        <div v-if="status === 'error'" class="flex gap-3">
          <button @click="$emit('close')" class="flex-1 py-3 bg-gray-100 text-gray-600 font-bold rounded-xl hover:bg-gray-200 transition">
            关闭
          </button>
        </div>
      </div>
    </div>
  </Transition>
</template>

<style scoped>
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}

@keyframes scaleIn {
  from { transform: scale(0.9); opacity: 0; }
  to { transform: scale(1); opacity: 1; }
}

.animate-scale {
  animation: scaleIn 0.3s cubic-bezier(0.16, 1, 0.3, 1);
}
</style>
