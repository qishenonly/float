<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useToast } from '../composables/useToast'

const router = useRouter()
const { showToast } = useToast()

const searchQuery = ref('')
const feedback = ref('')
const feedbackLength = ref(0)
const activeFaq = ref(null)

const faqs = [
  {
    id: 1,
    icon: 'fa-cloud-arrow-up',
    color: 'blue',
    question: '数据如何同步？',
    answer: '极光记账开启了云端自动同步功能。只要您登录了账号，数据会在 WiFi 环境下自动上传至云端，确保数据永不丢失。'
  },
  {
    id: 2,
    icon: 'fa-file-export',
    color: 'green',
    question: '支持导出 Excel 吗？',
    answer: '支持。您可以在"我的"页面点击"账单导出"，选择您需要的时间范围，我们将生成 Excel 或 PDF 文件发送至您的邮箱。'
  },
  {
    id: 3,
    icon: 'fa-unlock-keyhole',
    color: 'orange',
    question: '忘记手势密码怎么办？',
    answer: '别担心，您可以卸载重装 App，然后通过手机号验证码重新登录，您的账单数据不会丢失。'
  }
]

const toggleFaq = (id) => {
  activeFaq.value = activeFaq.value === id ? null : id
}

const updateFeedbackLength = () => {
  feedbackLength.value = feedback.value.length
}

const handleSubmit = () => {
  if (!feedback.value.trim()) {
    showToast('请输入反馈内容', 'warning')
    return
  }
  showToast('感谢您的反馈！我们会尽快处理', 'success')
  feedback.value = ''
  feedbackLength.value = 0
}

const handleScreenshot = () => {
  showToast('截图功能开发中...', 'info')
}

const handleSocial = (platform) => {
  showToast(`${platform}联系功能开发中...`, 'info')
}
</script>

<template>
  <div>
    <!-- Header -->
    <div class="px-6 pt-12 pb-4 relative z-10 flex items-center justify-between animate-enter">
      <button @click="router.back()" class="w-10 h-10 rounded-full bg-white/50 flex items-center justify-center backdrop-blur-md text-gray-600 hover:bg-white transition active-press">
        <i class="fa-solid fa-arrow-left"></i>
      </button>
      <h1 class="text-lg font-bold text-gray-800">帮助与反馈</h1>
      <div class="w-10"></div>
    </div>

    <div class="flex-1 overflow-y-auto hide-scrollbar px-6 relative z-10 pb-8">
      
      <!-- Search -->
      <div class="glass-card px-4 py-3 rounded-2xl flex items-center gap-3 mb-8 text-gray-500 shadow-sm focus-within:ring-2 ring-indigo-200 transition animate-enter delay-100">
        <i class="fa-solid fa-magnifying-glass"></i>
        <input 
          v-model="searchQuery"
          type="text" 
          class="bg-transparent w-full outline-none text-gray-700 placeholder-gray-400" 
          placeholder="搜索常见问题..."
        >
      </div>

      <!-- FAQ Section -->
      <h3 class="font-bold text-gray-800 mb-4 px-1 animate-enter delay-200">常见问题</h3>
      <div class="space-y-3 mb-8 animate-enter delay-300">
        
        <div 
          v-for="faq in faqs" 
          :key="faq.id"
          @click="toggleFaq(faq.id)"
          class="glass-card p-4 rounded-2xl cursor-pointer group active-press"
        >
          <div class="flex justify-between items-center">
            <div class="flex items-center gap-3">
              <div :class="`w-8 h-8 rounded-full bg-${faq.color}-100 text-${faq.color}-500 flex items-center justify-center`">
                <i :class="`fa-solid ${faq.icon} text-xs`"></i>
              </div>
              <span class="text-sm font-bold text-gray-700">{{ faq.question }}</span>
            </div>
            <i 
              :class="activeFaq === faq.id ? 'rotate-90' : ''" 
              class="fa-solid fa-chevron-right text-gray-300 text-xs transition-transform duration-300"
            ></i>
          </div>
          <Transition
            enter-active-class="transition-all duration-300 ease-out"
            enter-from-class="max-h-0 opacity-0"
            enter-to-class="max-h-48 opacity-100"
            leave-active-class="transition-all duration-300 ease-in"
            leave-from-class="max-h-48 opacity-100"
            leave-to-class="max-h-0 opacity-0"
          >
            <div v-show="activeFaq === faq.id" class="overflow-hidden">
              <p class="text-xs text-gray-500 leading-relaxed px-1 mt-2">
                {{ faq.answer }}
              </p>
            </div>
          </Transition>
        </div>

      </div>

      <!-- Feedback Section -->
      <h3 class="font-bold text-gray-800 mb-4 px-1 animate-enter delay-400">意见反馈</h3>
      <div class="glass-card p-4 rounded-2xl mb-8 animate-enter delay-500">
        <textarea 
          v-model="feedback"
          @input="updateFeedbackLength"
          maxlength="200"
          class="w-full bg-transparent outline-none text-sm text-gray-700 placeholder-gray-400 resize-none h-24 mb-2" 
          placeholder="请详细描述您遇到的问题或建议，我们会认真倾听..."
        ></textarea>
        
        <div class="flex justify-between items-center mt-2">
          <button @click="handleScreenshot" class="w-16 h-16 rounded-xl border border-dashed border-gray-300 flex flex-col items-center justify-center text-gray-400 hover:bg-white/50 transition active-press">
            <i class="fa-solid fa-camera text-sm mb-1"></i>
            <span class="text-[10px]">截图</span>
          </button>
          <span class="text-xs text-gray-300">{{ feedbackLength }}/200</span>
        </div>
      </div>
      
      <!-- Submit Button -->
      <button @click="handleSubmit" class="w-full bg-gradient-to-r from-indigo-500 to-blue-500 text-white font-bold py-4 rounded-2xl shadow-lg shadow-indigo-200 hover:shadow-xl hover:scale-[1.02] transition active:scale-[0.98] animate-enter delay-600">
        提交反馈
      </button>
      
      <!-- Social Media -->
      <div class="mt-8 text-center animate-enter delay-700">
        <p class="text-[10px] text-gray-400 mb-2">或者通过以下方式联系我们</p>
        <div class="flex justify-center gap-6 text-gray-300">
          <i @click="handleSocial('微信')" class="fa-brands fa-weixin text-xl hover:text-green-500 transition cursor-pointer active-press"></i>
          <i @click="handleSocial('微博')" class="fa-brands fa-weibo text-xl hover:text-red-500 transition cursor-pointer active-press"></i>
          <i @click="handleSocial('邮箱')" class="fa-solid fa-envelope text-xl hover:text-blue-500 transition cursor-pointer active-press"></i>
        </div>
      </div>

    </div>
  </div>
</template>
