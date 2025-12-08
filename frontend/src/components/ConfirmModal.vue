<script setup>
defineProps({
  show: Boolean,
  title: {
    type: String,
    default: '确认操作'
  },
  content: {
    type: String,
    default: '确定要执行此操作吗？'
  },
  confirmText: {
    type: String,
    default: '确定'
  },
  cancelText: {
    type: String,
    default: '取消'
  }
})

defineEmits(['close', 'confirm'])
</script>

<template>
  <Transition name="fade">
    <div v-if="show" class="fixed inset-0 z-[100] flex items-center justify-center px-4" @click.self="$emit('close')">
      <!-- Backdrop -->
      <div class="absolute inset-0 bg-black/40 backdrop-blur-sm"></div>

      <!-- Modal -->
      <div class="bg-white rounded-3xl w-full max-w-sm p-6 relative z-10 shadow-2xl animate-scale">
        <div class="text-center mb-6">
          <h3 class="text-xl font-bold text-gray-800 mb-2">{{ title }}</h3>
          <p class="text-gray-500 text-sm">{{ content }}</p>
        </div>

        <div class="flex gap-4">
          <button 
            @click="$emit('close')" 
            class="flex-1 py-3 bg-gray-100 text-gray-600 font-bold rounded-xl hover:bg-gray-200 transition"
          >
            {{ cancelText }}
          </button>
          <button 
            @click="$emit('confirm')" 
            class="flex-1 py-3 bg-indigo-600 text-white font-bold rounded-xl hover:bg-indigo-700 transition shadow-lg shadow-indigo-200"
          >
            {{ confirmText }}
          </button>
        </div>
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
