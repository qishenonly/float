import apiClient from './client'

// 交易管理 API
export const transactionAPI = {
  // 获取交易列表
  getTransactions(params = {}) {
    return apiClient.get('/transactions', { params })
  },

  // 获取交易详情
  getTransactionById(id) {
    return apiClient.get(`/transactions/${id}`)
  },

  // 创建交易
  createTransaction(data) {
    return apiClient.post('/transactions', data)
  },

  // 更新交易
  updateTransaction(id, data) {
    return apiClient.put(`/transactions/${id}`, data)
  },

  // 删除交易
  deleteTransaction(id) {
    return apiClient.delete(`/transactions/${id}`)
  },

  // 批量创建交易
  createTransactionBatch(data) {
    return apiClient.post('/transactions/batch', data)
  },

  // 批量删除交易
  deleteTransactionBatch(ids) {
    return apiClient.delete('/transactions/batch', { data: { ids } })
  },

  // 获取交易统计
  getStatistics(params = {}) {
    return apiClient.get('/transactions/statistics', { params })
  },

  // 获取月度统计
  getMonthStatistics(params = {}) {
    return apiClient.get('/transactions/month-statistics', { params })
  },

  // 获取分类统计
  getCategoryStatistics(params = {}) {
    return apiClient.get('/transactions/category-statistics', { params })
  }
}
