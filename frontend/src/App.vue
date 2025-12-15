<script setup>
import { App } from '@capacitor/app'
import { RouterView, useRoute, useRouter } from 'vue-router'
import { onMounted } from 'vue'
import BottomNav from './components/BottomNav.vue'
import Toast from './components/Toast.vue'
import UpdateModal from './components/UpdateModal.vue'
import { useToast } from './composables/useToast'
import { useAppUpdate } from './composables/useAppUpdate'

const route = useRoute()
const router = useRouter()
const { toastState, hideToast } = useToast()
const { updateState, checkUpdate, openUpdateDialog, confirmUpdate, closeUpdateModal } = useAppUpdate()

onMounted(async () => {
  // Listen for deep links (e.g. from floating window)
  App.addListener('appUrlOpen', data => {
    // Example: float://add?amount=12.50&description=Starbucks
    // Capacitor converts this to a path but we might need to parse it manually if it's custom scheme
    // If scheme is 'float', data.url is full URL.
    
    // Simple parsing
    try {
      const url = new URL(data.url)
      if (url.host === 'add') {
        const amount = url.searchParams.get('amount')
        const description = url.searchParams.get('description')
        
        router.push({
          path: '/add',
          query: {
             amount,
             description,
             auto: 'true'
          }
        })
      }
    } catch (e) {
      console.error('Failed to parse deep link', e)
    }
  })

  try {
    const latest = await checkUpdate()
    if (latest) {
      // 延时一点显示，避免和页面加载冲突
      setTimeout(() => {
        openUpdateDialog(latest)
      }, 1000)
    }
  } catch (error) {
    console.error('Auto check update failed:', error)
  }
})
</script>

<template>
  <div class="bg-[#F2F4F8] sm:flex sm:items-center sm:justify-center min-h-screen">
    <div class="w-full h-screen sm:max-w-[375px] sm:h-[812px] bg-gradient-to-br from-[#eff6ff] via-[#fff1f2] to-[#f0f9ff] sm:shadow-2xl overflow-hidden relative sm:rounded-[40px] flex flex-col sm:border-[8px] sm:border-white">
      
      <!-- Background Blobs -->
      <div class="blob w-64 h-64 bg-blue-200 rounded-full -top-10 -left-10 mix-blend-multiply"></div>
      <div class="blob w-64 h-64 bg-pink-200 rounded-full top-20 -right-20 mix-blend-multiply"></div>

      <!-- Main Content Area -->
      <div class="flex-1 overflow-y-auto hide-scrollbar pb-8 relative z-10">
        <RouterView />
      </div>

      <!-- Bottom Navigation -->
      <BottomNav v-if="route?.name !== 'add'" />
      
      <!-- Toast Notification -->
      <Toast 
        :show="toastState?.show" 
        :message="toastState?.message" 
        :type="toastState?.type"
        @close="hideToast"
      />

      <!-- Update Modal -->
      <UpdateModal
        :show="updateState?.showModal"
        :progress="updateState?.progress"
        :status="updateState?.status"
        :version="updateState?.latestVersion?.version_name"
        :description="updateState?.latestVersion?.description"
        @confirm="confirmUpdate"
        @cancel="closeUpdateModal"
        @close="closeUpdateModal"
      />
    </div>
  </div>
</template>
