import apiClient from './client'

// 用户认证 API
export const authAPI = {
  // 注册
  register(data) {
    return apiClient.post('/auth/register', data)
  },

  // 登录
  login(data) {
    return apiClient.post('/auth/login', data)
  },

  // 刷新 Token
  refreshToken(refreshToken) {
    return apiClient.post('/auth/refresh', { refresh_token: refreshToken })
  },

  // 登出
  logout(refreshToken) {
    return apiClient.post('/auth/logout', { refresh_token: refreshToken })
  }
}

// 用户管理 API
export const userAPI = {
  // 获取当前用户信息
  getCurrentUser() {
    return apiClient.get('/users/me')
  },

  // 更新用户信息
  updateCurrentUser(data) {
    return apiClient.put('/users/me', data)
  },

  // 修改密码
  updatePassword(data) {
    return apiClient.put('/users/me/password', data)
  },

  // 获取用户统计
  getUserStats() {
    return apiClient.get('/users/me/stats')
  }
}
