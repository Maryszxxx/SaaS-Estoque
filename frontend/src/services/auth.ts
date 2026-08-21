import api from './api'
import type { LoginCredentials, LoginResponse } from '../types/auth'
export async function login(credentials: LoginCredentials) { const { data } = await api.post<LoginResponse>('/login', credentials); return data }
export async function registerUser(user: { name: string; email: string; password: string; role: 'ADMIN' | 'EMPLOYEE' }) { await api.post('/users', user) }
