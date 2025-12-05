import apiClient from './client'

// 分类管理 API
export const categoryAPI = {
  // 获取分类列表
  getCategories(type) {
    const params = type ? { type } : {}
    return apiClient.get('/categories', { params })
  },

  // 获取分类详情
  getCategoryById(id) {
    return apiClient.get(`/categories/${id}`)
  },

  // 创建分类
  createCategory(data) {
    return apiClient.post('/categories', data)
  },

  // 更新分类
  updateCategory(id, data) {
    return apiClient.put(`/categories/${id}`, data)
  },

  // 删除分类
  deleteCategory(id) {
    return apiClient.delete(`/categories/${id}`)
  },

  // 获取系统默认分类
  getSystemCategories(type) {
    const params = type ? { type } : {}
    return apiClient.get('/categories/system', { params })
  }
}
