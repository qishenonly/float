import apiClient from './client'

// 账户管理 API
export const accountAPI = {
  // 获取账户列表
  getAccounts() {
    return apiClient.get('/accounts')
  },

  // 获取账户详情
  getAccountById(id) {
    return apiClient.get(`/accounts/${id}`)
  },

  // 创建账户
  createAccount(data) {
    return apiClient.post('/accounts', data)
  },

  // 更新账户
  updateAccount(id, data) {
    return apiClient.put(`/accounts/${id}`, data)
  },

  // 删除账户
  deleteAccount(id) {
    return apiClient.delete(`/accounts/${id}`)
  },

  // 获取账户余额汇总
  getAccountBalance() {
    return apiClient.get('/accounts/balance')
  }
}
