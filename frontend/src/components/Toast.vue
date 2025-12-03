<script setup>
import { ref, watch } from 'vue'

const props = defineProps({
  message: String,
  type: {
    type: String,
    default: 'info' // info, success, warning, error
  },
  show: Boolean
})

const emit = defineEmits(['close'])

const visible = ref(props.show)

watch(() => props.show, (newVal) => {
  visible.value = newVal
  if (newVal) {
    setTimeout(() => {
      visible.value = false
      emit('close')
    }, 2500)
  }
})

const iconMap = {
  info: 'fa-circle-info',
  success: 'fa-circle-check',
  warning: 'fa-triangle-exclamation',
  error: 'fa-circle-xmark'
}

const colorMap = {
  info: 'from-blue-500 to-indigo-500',
  success: 'from-green-500 to-emerald-500',
  warning: 'from-yellow-500 to-orange-500',
  error: 'from-red-500 to-pink-500'
}
</script>

<template>
  <Transition
    enter-active-class="transition duration-300 ease-out"
    enter-from-class="opacity-0 translate-y-4"
    enter-to-class="opacity-100 translate-y-0"
    leave-active-class="transition duration-200 ease-in"
    leave-from-class="opacity-100 translate-y-0"
    leave-to-class="opacity-0 translate-y-4"
  >
    <div v-if="visible" class="fixed top-20 left-1/2 -translate-x-1/2 z-50 w-[90%] max-w-sm">
      <div :class="['relative overflow-hidden rounded-2xl p-4 text-white shadow-2xl backdrop-blur-xl border border-white/20', 'bg-gradient-to-r', colorMap[type]]">
        <div class="absolute -right-10 -top-10 w-32 h-32 bg-white opacity-10 rounded-full blur-2xl"></div>
        <div class="relative z-10 flex items-center gap-3">
          <div class="w-10 h-10 bg-white/20 rounded-xl flex items-center justify-center backdrop-blur-sm">
            <i :class="['fa-solid', iconMap[type], 'text-lg']"></i>
          </div>
          <p class="flex-1 text-sm font-medium">{{ message }}</p>
        </div>
      </div>
    </div>
  </Transition>
</template>
