<script setup>
defineProps({
  show: Boolean,
  status: {
    type: String,
    enum: ['loading', 'success', 'error'],
    default: 'loading'
  },
  title: String,
  data: Object,
  error: String,
  url: String
})

defineEmits(['close'])
</script>

<template>
  <Transition name="fade">
    <div v-if="show" class="fixed inset-0 z-[100] flex items-center justify-center px-4" @click.self="$emit('close')">
      <!-- Backdrop -->
      <div class="absolute inset-0 bg-black/40 backdrop-blur-sm"></div>

      <!-- Modal -->
      <div class="bg-white rounded-3xl w-full max-w-sm p-6 relative z-10 shadow-2xl animate-scale">
        <!-- Loading State -->
        <div v-if="status === 'loading'" class="text-center">
          <div class="flex justify-center mb-4">
            <i class="fa-solid fa-spinner fa-spin text-3xl text-indigo-600"></i>
          </div>
          <h3 class="text-lg font-bold text-gray-800 mb-2">健康检查中...</h3>
          <p class="text-gray-500 text-sm">请求地址: {{ url }}</p>
        </div>

        <!-- Success State -->
        <div v-else-if="status === 'success'" class="text-center">
          <div class="flex justify-center mb-4">
            <div class="w-12 h-12 rounded-full bg-green-100 flex items-center justify-center">
              <i class="fa-solid fa-check text-xl text-green-600"></i>
            </div>
          </div>
          <h3 class="text-lg font-bold text-gray-800 mb-2">{{ title }}</h3>
          <p class="text-gray-500 text-sm mb-4">请求地址: {{ url }}</p>
          
          <!-- Data Display -->
          <div class="bg-gray-50 rounded-xl p-4 text-left max-h-64 overflow-y-auto">
            <pre class="text-xs text-gray-700 font-mono whitespace-pre-wrap break-words">{{ JSON.stringify(data, null, 2) }}</pre>
          </div>
        </div>

        <!-- Error State -->
        <div v-else-if="status === 'error'" class="text-center">
          <div class="flex justify-center mb-4">
            <div class="w-12 h-12 rounded-full bg-red-100 flex items-center justify-center">
              <i class="fa-solid fa-exclamation text-xl text-red-600"></i>
            </div>
          </div>
          <h3 class="text-lg font-bold text-gray-800 mb-2">{{ title }}</h3>
          <p class="text-gray-500 text-sm mb-4">请求地址: {{ url }}</p>
          
          <!-- Error Display -->
          <div class="bg-red-50 rounded-xl p-4 text-left max-h-64 overflow-y-auto border border-red-200">
            <p class="text-xs text-gray-600 font-bold mb-2">原因:</p>
            <p class="text-xs text-red-700 font-mono whitespace-pre-wrap break-words leading-relaxed">{{ error }}</p>
          </div>
        </div>

        <!-- Close Button -->
        <button 
          @click="$emit('close')" 
          class="w-full mt-6 py-3 bg-indigo-600 text-white font-bold rounded-xl hover:bg-indigo-700 transition shadow-lg shadow-indigo-200"
        >
          {{ status === 'loading' ? '等待中...' : '确定' }}
        </button>
      </div>
    </div>
  </Transition>
</template>

<style scoped>
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.2s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}

@keyframes scaleIn {
  from { transform: scale(0.95); opacity: 0; }
  to { transform: scale(1); opacity: 1; }
}

.animate-scale {
  animation: scaleIn 0.2s cubic-bezier(0.16, 1, 0.3, 1);
}
</style>
