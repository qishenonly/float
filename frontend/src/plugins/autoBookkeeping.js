import { registerPlugin } from '@capacitor/core'
import { Capacitor } from '@capacitor/core'

/**
 * 自动记账插件接口定义
 * @typedef {Object} AutoBookkeepingPlugin
 * @property {() => Promise<{overlay: boolean, accessibility: boolean, notification: boolean}>} checkPermissions
 * @property {() => Promise<void>} requestOverlayPermission
 * @property {() => Promise<void>} requestAccessibilityPermission
 * @property {() => Promise<void>} requestNotificationPermission
 */

// Web fallback for when running in browser
const webFallback = {
  async checkPermissions() {
    console.warn('AutoBookkeeping: Running in web mode, returning mock permissions')
    return { overlay: false, accessibility: false, notification: false, battery: false }
  },
  async requestOverlayPermission() {
    console.warn('AutoBookkeeping: Overlay permission not available on web')
  },
  async requestAccessibilityPermission() {
    console.warn('AutoBookkeeping: Accessibility permission not available on web')
  },
  async requestNotificationPermission() {
    console.warn('AutoBookkeeping: Notification permission not available on web')
  },
  async requestIgnoreBatteryOptimizations() {
    console.warn('AutoBookkeeping: Battery optimization not available on web')
  }
}

let AutoBookkeeping

if (Capacitor.isNativePlatform()) {
  AutoBookkeeping = registerPlugin('AutoBookkeeping')
} else {
  AutoBookkeeping = webFallback
}

export default AutoBookkeeping
