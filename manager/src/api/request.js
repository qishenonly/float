import axios from 'axios'

const service = axios.create({
  baseURL: '/api/v1',
  timeout: 5000
})

// Request Interceptor
service.interceptors.request.use(
  config => {
    const token = localStorage.getItem('token')
    if (token) {
      config.headers['Authorization'] = `Bearer ${token}`
    }
    return config
  },
  error => {
    return Promise.reject(error)
  }
)

// Response Interceptor
service.interceptors.response.use(
  response => {
    const res = response.data
    
    // 1. Direct Array Response (e.g. History API)
    if (Array.isArray(res)) {
        return res
    }

    // 2. Standard Wrapper with 'code' (e.g. Auth API)
    if (typeof res.code === 'number') {
        if (res.code !== 200) {
             console.error('API Error:', res.message)
             return Promise.reject(new Error(res.message || 'Error'))
        }
        return res.data
    }

    // 3. Upload Response or others without 'code'
    // If it has a success message or valid data without error indication
    // The upload API returns { message: "Upload successful", data: {...} }
    if (res.message === 'Upload successful' || (res.data && !res.error)) {
        return res.data || res
    }

    // 4. Fallback: If HTTP status is success (2xx) but structure is unknown, return data or res
    return res.data !== undefined ? res.data : res
  },
  error => {
    console.error('Network Error:', error)
    return Promise.reject(error)
  }
)

export default service
