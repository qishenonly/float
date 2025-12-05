import { defineStore } from 'pinia'
import { authAPI, userAPI } from '@/api/user'

export const useAuthStore = defineStore('auth', {
  state: () => ({
    user: null,
    accessToken: null,
    refreshToken: null,
    isAuthenticated: false
  }),

  getters: {
    currentUser: (state) => state.user,
    isLoggedIn: (state) => state.isAuthenticated
  },

  actions: {
    // 初始化认证状态（从 localStorage 恢复）
    initAuth() {
      const token = localStorage.getItem('access_token')
      const refreshToken = localStorage.getItem('refresh_token')
      const user = localStorage.getItem('user')

      if (token && user) {
        this.accessToken = token
        this.refreshToken = refreshToken
        this.user = JSON.parse(user)
        this.isAuthenticated = true
      }
    },

    // 注册
    async register(userData) {
      try {
        const response = await authAPI.register(userData)
        const { user_id, username, email, display_name, avatar_url, access_token, refresh_token } = response.data

        // 保存用户信息和 token
        this.user = { id: user_id, username, email, display_name, avatar_url }
        this.accessToken = access_token
        this.refreshToken = refresh_token
        this.isAuthenticated = true

        // 持久化存储
        localStorage.setItem('access_token', access_token)
        localStorage.setItem('refresh_token', refresh_token)
        localStorage.setItem('user', JSON.stringify(this.user))

        return response
      } catch (error) {
        throw error.response?.data || error
      }
    },

    // 登录
    async login(credentials) {
      try {
        const response = await authAPI.login(credentials)
        const { user_id, username, email, display_name, avatar_url, access_token, refresh_token } = response.data

        // 保存用户信息和 token
        this.user = { id: user_id, username, email, display_name, avatar_url }
        this.accessToken = access_token
        this.refreshToken = refresh_token
        this.isAuthenticated = true

        // 持久化存储
        localStorage.setItem('access_token', access_token)
        localStorage.setItem('refresh_token', refresh_token)
        localStorage.setItem('user', JSON.stringify(this.user))

        return response
      } catch (error) {
        throw error.response?.data || error
      }
    },

    // 登出
    async logout() {
      try {
        if (this.refreshToken) {
          await authAPI.logout(this.refreshToken)
        }
      } catch (error) {
        console.error('Logout error:', error)
      } finally {
        // 清除状态
        this.user = null
        this.accessToken = null
        this.refreshToken = null
        this.isAuthenticated = false

        // 清除存储
        localStorage.removeItem('access_token')
        localStorage.removeItem('refresh_token')
        localStorage.removeItem('user')
      }
    },

    // 刷新用户信息
    async fetchCurrentUser() {
      try {
        const response = await userAPI.getCurrentUser()
        this.user = response.data
        localStorage.setItem('user', JSON.stringify(this.user))
        return response
      } catch (error) {
        throw error.response?.data || error
      }
    },

    // 更新用户信息
    async updateUser(userData) {
      try {
        await userAPI.updateCurrentUser(userData)
        await this.fetchCurrentUser()
      } catch (error) {
        throw error.response?.data || error
      }
    }
  }
})
