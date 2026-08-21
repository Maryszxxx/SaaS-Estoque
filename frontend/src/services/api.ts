import axios, { AxiosError, type InternalAxiosRequestConfig } from 'axios'
import { tokenStorage } from './tokenStorage'
const api = axios.create({ baseURL: import.meta.env.VITE_API_URL ?? '/api', headers: { 'Content-Type': 'application/json' } })
api.interceptors.request.use((config: InternalAxiosRequestConfig) => { const token = tokenStorage.getAccessToken(); if (token) config.headers.Authorization = `Bearer ${token}`; return config })
let refreshPromise: Promise<string> | null = null
api.interceptors.response.use((response) => response, async (error: AxiosError) => {
  const request = error.config
  if (!request || error.response?.status !== 401 || request.url === '/login' || request.url === '/refresh') return Promise.reject(error)
  const retryableRequest = request as InternalAxiosRequestConfig & { _retried?: boolean }
  if (retryableRequest._retried) return Promise.reject(error); retryableRequest._retried = true
  const refreshToken = tokenStorage.getRefreshToken(); if (!refreshToken) return Promise.reject(error)
  try { refreshPromise ??= axios.post<{ access_token: string }>(`${api.defaults.baseURL}/refresh`, { refresh_token: refreshToken }).then(({ data }) => data.access_token).finally(() => { refreshPromise = null }); const accessToken = await refreshPromise; tokenStorage.save(accessToken, refreshToken); retryableRequest.headers.Authorization = `Bearer ${accessToken}`; return api(retryableRequest) } catch (refreshError) { tokenStorage.clear(); return Promise.reject(refreshError) }
})
export default api
