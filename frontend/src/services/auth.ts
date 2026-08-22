import api from './api'
import type { LoginCredentials, LoginResponse, UserRole } from '../types/auth'

export interface RegisterUserPayload { name: string; email: string; password: string; role: UserRole }
export interface RegisterUserResponse extends LoginResponse { message: string }

export async function login(credentials: LoginCredentials) {
  const { data } = await api.post<LoginResponse>('/login', credentials)
  return data
}

export async function registerUser(user: RegisterUserPayload) {
  const { data } = await api.post<RegisterUserResponse>('/users', user)
  return data
}
