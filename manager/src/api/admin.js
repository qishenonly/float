import request from './request'

export function login(data) {
  return request({
    url: '/auth/login',
    method: 'post',
    data
  })
}

export function getSystemOverview() {
  return request({
    url: '/admin/overview',
    method: 'get'
  })
}

export function getUsers(params) {
  return request({
    url: '/admin/users',
    method: 'get',
    params
  })
}

export function getAppUpdateHistory(platform) {
  return request({
    url: '/app-updates/history',
    method: 'get',
    params: { platform }
  })
}

export function uploadAppUpdate(formData) {
  return request({
    url: '/app-updates',
    method: 'post',
    data: formData,
    headers: {
      'Content-Type': 'multipart/form-data'
    }
  })
}

export function getTransactionStats() {
    return request({
        url: '/transactions/statistics',
        method: 'get'
    })
}
