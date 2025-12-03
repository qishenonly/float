import { ref } from 'vue'

const toastState = ref({
  show: false,
  message: '',
  type: 'info'
})

export function useToast() {
  const showToast = (message, type = 'info') => {
    toastState.value = {
      show: true,
      message,
      type
    }
  }

  const hideToast = () => {
    toastState.value.show = false
  }

  return {
    toastState,
    showToast,
    hideToast
  }
}
