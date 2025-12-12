import { ref, reactive } from 'vue'
import { appUpdateAPI } from '@/api/appUpdate'
import { Filesystem, Directory } from '@capacitor/filesystem'
import { FileOpener } from '@capacitor-community/file-opener'
import axios from 'axios'
import { Capacitor } from '@capacitor/core'

const updateState = reactive({
  showModal: false,
  progress: 0,
  status: 'downloading', // downloading, installing, error
  latestVersion: null
})

// Current version info
const currentVersion = {
  code: 3,
  name: '0.0.3',
  platform: 'android'
}

export function useAppUpdate() {
  
  const checkUpdate = async (manual = false) => {
    try {
      // Only check on Android or if manual check (for testing)
      if (Capacitor.getPlatform() !== 'android' && !manual) return null

      const response = await appUpdateAPI.checkUpdate(currentVersion.platform, currentVersion.code)
      const data = response.data || response // Handle both raw and intercepted response

      if (data.has_update && data.latest) {
        updateState.latestVersion = data.latest
        return data.latest
      }
      return null
    } catch (error) {
      console.error('Check update failed:', error)
      throw error
    }
  }

  const startUpdate = async (versionInfo) => {
    if (!versionInfo?.download_url) return

    updateState.showModal = true
    updateState.progress = 0
    updateState.status = 'downloading'
    updateState.latestVersion = versionInfo

    try {
      const baseUrl = import.meta.env.VITE_API_BASE_URL || 'http://10.0.2.2:8080/api/v1'
      const downloadUrl = `${baseUrl.replace('/api/v1', '')}${versionInfo.download_url}`
      
      // 1. Download File
      const response = await axios.get(downloadUrl, {
        responseType: 'blob',
        onDownloadProgress: (progressEvent) => {
          if (progressEvent.total) {
            updateState.progress = Math.round((progressEvent.loaded * 100) / progressEvent.total)
          }
        }
      })

      // 2. Convert Blob to Base64
      const base64Data = await blobToBase64(response.data)

      // 3. Save to Filesystem
      const fileName = `update_${versionInfo.version_code}.apk`
      const savedFile = await Filesystem.writeFile({
        path: fileName,
        data: base64Data,
        directory: Directory.Cache,
        recursive: true
      })

      updateState.status = 'installing'
      updateState.progress = 100

      // 4. Open File (Install)
      await FileOpener.open({
        filePath: savedFile.uri,
        contentType: 'application/vnd.android.package-archive'
      })

      // Close modal after a short delay
      setTimeout(() => {
        updateState.showModal = false
      }, 1000)

    } catch (error) {
      console.error('Update failed:', error)
      updateState.status = 'error'
      alert('更新失败: ' + (error.message || '未知错误'))
    }
  }

  const blobToBase64 = (blob) => {
    return new Promise((resolve, reject) => {
      const reader = new FileReader()
      reader.onloadend = () => {
        // remove "data:application/octet-stream;base64," prefix
        const base64String = reader.result.split(',')[1]
        resolve(base64String)
      }
      reader.onerror = reject
      reader.readAsDataURL(blob)
    })
  }

  const openUpdateDialog = (version) => {
    updateState.latestVersion = version
    updateState.showModal = true
    updateState.status = 'prompt'
    updateState.progress = 0
  }

  const confirmUpdate = () => {
    if (updateState.latestVersion) {
      startUpdate(updateState.latestVersion)
    }
  }

  const closeUpdateModal = () => {
    updateState.showModal = false
  }

  return {
    updateState,
    currentVersion,
    checkUpdate,
    startUpdate,
    openUpdateDialog,
    confirmUpdate,
    closeUpdateModal
  }
}
