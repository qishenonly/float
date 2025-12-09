<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useToast } from '../composables/useToast'
import { categoryAPI } from '@/api/category'
import ConfirmModal from '@/components/ConfirmModal.vue'

const router = useRouter()
const { showToast } = useToast()

const currentTab = ref('expense')
const categories = ref([])
const loading = ref(false)
const showModal = ref(false)
const editingCategory = ref(null)
const showDeleteConfirm = ref(false)
const pendingDeleteCategory = ref(null)

const formData = ref({
  name: '',
  icon: '',
  color: 'orange'
})

const availableColors = [
  'orange', 'purple', 'blue', 'green', 'red', 'pink', 
  'indigo', 'yellow', 'teal', 'amber', 'lime', 'cyan',
  'rose', 'violet', 'fuchsia', 'emerald', 'sky', 'slate'
]
const availableIcons = [
  'fa-utensils', 'fa-bag-shopping', 'fa-bus', 'fa-house', 'fa-heartbeat', 'fa-gamepad',
  'fa-sack-dollar', 'fa-chart-line', 'fa-briefcase', 'fa-gift', 'fa-coffee', 'fa-car',
  'fa-plane', 'fa-book', 'fa-music', 'fa-film', 'fa-shopping-cart', 'fa-credit-card'
]

// 当前tab的分类列表
const currentCategories = computed(() => {
  return categories.value.filter(cat => cat.type === currentTab.value)
})

onMounted(() => {
  loadCategories()
})

// 加载分类
const loadCategories = async () => {
  loading.value = true
  try {
    const response = await categoryAPI.getCategories()
    categories.value = response.data || []
  } catch (error) {
    console.error('Failed to load categories:', error)
    showToast('加载分类失败', 'error')
  } finally {
    loading.value = false
  }
}

// 打开添加分类modal
const openAddModal = () => {
  editingCategory.value = null
  formData.value = {
    name: '',
    icon: availableIcons[0],
    color: availableColors[0]
  }
  showModal.value = true
}

// 打开编辑modal
const openEditModal = (category) => {
  if (category.is_system) {
    showToast('系统分类不可编辑', 'error')
    return
  }
  
  editingCategory.value = category
  formData.value = {
    name: category.name,
    icon: category.icon,
    color: category.color
  }
  showModal.value = true
}

// 保存分类
const saveCategory = async () => {
  if (!formData.value.name.trim()) {
    showToast('请输入分类名称', 'error')
    return
  }

  loading.value = true
  try {
    if (editingCategory.value) {
      // 更新
      await categoryAPI.updateCategory(editingCategory.value.id, formData.value)
      showToast('分类更新成功', 'success')
    } else {
      // 创建
      await categoryAPI.createCategory({
        ...formData.value,
        type: currentTab.value
      })
      showToast('分类创建成功', 'success')
    }
    
    showModal.value = false
    await loadCategories()
  } catch (error) {
    console.error('Failed to save category:', error)
    showToast(error.response?.data?.message || '保存失败', 'error')
  } finally {
    loading.value = false
  }
}

// 删除分类
const deleteCategory = (category) => {
  if (category.is_system) {
    showToast('系统分类不可删除', 'error')
    return
  }

  pendingDeleteCategory.value = category
  showDeleteConfirm.value = true
}

const confirmDeleteCategory = async () => {
  if (!pendingDeleteCategory.value) return
  
  loading.value = true
  try {
    await categoryAPI.deleteCategory(pendingDeleteCategory.value.id)
    showToast('分类删除成功', 'success')
    showDeleteConfirm.value = false
    pendingDeleteCategory.value = null
    await loadCategories()
  } catch (error) {
    console.error('Failed to delete category:', error)
    showToast(error.response?.data?.message || '删除失败', 'error')
  } finally {
    loading.value = false
  }
}

// 关闭modal
const closeModal = () => {
  showModal.value = false
  editingCategory.value = null
}
</script>

<template>
  <div>
    <!-- Header -->
    <div class="px-6 pt-12 pb-4 relative z-10 flex items-center justify-between animate-enter">
      <button @click="router.back()" class="w-10 h-10 rounded-full bg-white/50 flex items-center justify-center backdrop-blur-md text-gray-600 hover:bg-white transition active-press">
        <i class="fa-solid fa-arrow-left"></i>
      </button>
      <h1 class="text-lg font-bold text-gray-800">分类管理</h1>
      <button @click="openAddModal" class="w-10 h-10 rounded-full bg-white/50 flex items-center justify-center backdrop-blur-md text-gray-600 hover:bg-white transition active-press">
        <i class="fa-solid fa-plus"></i>
      </button>
    </div>

    <!-- Tab Switcher -->
    <div class="px-6 mb-6 relative z-10 animate-enter delay-100">
      <div class="glass-card p-1 rounded-2xl flex relative transition-colors">
        <button 
          @click="currentTab = 'expense'" 
          :class="currentTab === 'expense' ? 'bg-white text-indigo-600 shadow-md' : 'text-gray-500'"
          class="flex-1 py-2.5 rounded-xl text-sm font-bold transition-all duration-300"
        >
          支出
        </button>
        <button 
          @click="currentTab = 'income'" 
          :class="currentTab === 'income' ? 'bg-white text-indigo-600 shadow-md' : 'text-gray-500'"
          class="flex-1 py-2.5 rounded-xl text-sm font-bold transition-all duration-300"
        >
          收入
        </button>
      </div>
    </div>

    <!-- Loading State -->
    <div v-if="loading && categories.length === 0" class="px-6 text-center py-12">
      <i class="fa-solid fa-spinner fa-spin text-3xl text-gray-400"></i>
      <p class="text-gray-400 mt-3">加载中...</p>
    </div>

    <!-- Category List -->
    <div v-else class="px-6 pb-32 relative z-10">
      <p class="text-xs text-gray-400 mb-4 px-2">系统分类不可编辑/删除</p>
      
      <div class="space-y-3 animate-enter delay-200">
        <div 
          v-for="category in currentCategories" 
          :key="category.id"
          class="glass-card p-3 rounded-2xl flex items-center justify-between group active:scale-[0.98] transition cursor-pointer"
        >
          <div class="flex items-center gap-4">
            <div :class="`w-10 h-10 rounded-xl bg-${category.color}-100 text-${category.color}-500 flex items-center justify-center shadow-inner`">
              <i :class="`fa-solid ${category.icon}`"></i>
            </div>
            <div>
              <span class="font-bold text-gray-700">{{ category.name }}</span>
              <span v-if="category.is_system" class="ml-2 text-xs text-gray-400">(系统)</span>
            </div>
          </div>
          
          <div v-if="!category.is_system" class="flex gap-2 opacity-0 group-hover:opacity-100 transition">
            <button @click.stop="openEditModal(category)" class="w-8 h-8 rounded-full bg-white text-gray-400 hover:text-blue-500 flex items-center justify-center shadow-sm active-press">
              <i class="fa-solid fa-pen text-xs"></i>
            </button>
            <button @click.stop="deleteCategory(category)" class="w-8 h-8 rounded-full bg-white text-gray-400 hover:text-red-500 flex items-center justify-center shadow-sm active-press">
              <i class="fa-solid fa-trash text-xs"></i>
            </button>
          </div>
        </div>

        <!-- Empty State -->
        <div v-if="currentCategories.length === 0" class="text-center py-12">
          <i class="fa-solid fa-folder-open text-4xl text-gray-300"></i>
          <p class="text-gray-400 mt-3">暂无{{ currentTab === 'expense' ? '支出' : '收入' }}分类</p>
        </div>
      </div>
    </div>

    <!-- Add/Edit Modal -->
    <div v-if="showModal" class="absolute inset-0 bg-black/50 z-50 flex items-center justify-center p-6" @click.self="closeModal">
      <div class="bg-white rounded-3xl w-full max-w-md p-6 space-y-4 animate-enter">
        <div class="flex items-center justify-between">
          <h2 class="text-lg font-bold">{{ editingCategory ? '编辑分类' : '新建分类' }}</h2>
          <button @click="closeModal" class="w-8 h-8 rounded-full hover:bg-gray-100 flex items-center justify-center">
            <i class="fa-solid fa-times text-gray-400"></i>
          </button>
        </div>

        <!-- Form -->
        <div class="space-y-4">
          <!-- Name -->
          <div>
            <label class="block text-sm font-medium text-gray-600 mb-2">分类名称</label>
            <input
              v-model="formData.name"
              type="text"
              placeholder="请输入分类名称"
              maxlength="50"
              class="w-full px-4 py-3 bg-gray-50 rounded-xl border-none outline-none text-gray-800 placeholder-gray-400"
            >
          </div>

          <!-- Icon -->
          <div>
            <label class="block text-sm font-medium text-gray-600 mb-2">图标</label>
            <div class="grid grid-cols-6 gap-2">
              <button
                v-for="icon in availableIcons"
                :key="icon"
                @click="formData.icon = icon"
                :class="formData.icon === icon ? 'bg-indigo-600 text-white shadow-md' : 'bg-gray-50 text-gray-400 hover:bg-gray-100'"
                class="w-full aspect-square rounded-xl flex items-center justify-center transition"
              >
                <i :class="`fa-solid ${icon}`"></i>
              </button>
            </div>
          </div>

          <!-- Color -->
          <div>
            <label class="block text-sm font-medium text-gray-600 mb-2">颜色</label>
            <div class="grid grid-cols-6 gap-2">
              <button
                v-for="color in availableColors"
                :key="color"
                @click="formData.color = color"
                :class="[
                  `bg-${color}-500`,
                  formData.color === color ? 'ring-2 ring-indigo-400' : ''
                ]"
                class="w-10 h-10 rounded-lg transition"
              ></button>
            </div>
          </div>
        </div>

        <!-- Actions -->
        <div class="flex gap-3 pt-4">
          <button @click="closeModal" class="flex-1 py-3 bg-gray-100 text-gray-600 font-bold rounded-xl hover:bg-gray-200 transition">
            取消
          </button>
          <button @click="saveCategory" :disabled="loading" class="flex-1 py-3 bg-indigo-600 text-white font-bold rounded-xl hover:bg-indigo-700 transition disabled:opacity-50">
            {{ loading ? '保存中...' : '保存' }}
          </button>
        </div>
      </div>
    </div>

    <!-- Delete Confirmation Modal -->
    <ConfirmModal 
      :show="showDeleteConfirm"
      title="确认删除"
      :content="`确定要删除分类「${pendingDeleteCategory?.name}」吗？`"
      confirmText="删除"
      cancelText="取消"
      @confirm="confirmDeleteCategory"
      @close="showDeleteConfirm = false"
    />
  </div>
</template>
