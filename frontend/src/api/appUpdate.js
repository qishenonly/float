import client from './client'

export const appUpdateAPI = {
  // 检查更新
  checkUpdate: (platform, versionCode) => {
    return client.get('/app-updates/check', {
      params: { platform, version_code: versionCode }
    })
  },

  // 获取最新版本
  getLatest: (platform) => {
    return client.get('/app-updates/latest', {
      params: { platform }
    })
  },

  // 获取更新历史
  getHistory: (platform) => {
    return client.get('/app-updates/history', {
      params: { platform }
    })
  }
}
