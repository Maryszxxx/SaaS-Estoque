export type UserRole = 'ADMIN' | 'EMPLOYEE'
export interface AuthUser { id: number; role: UserRole }
export interface LoginCredentials { email: string; password: string }
export interface LoginResponse { token: string; refresh_token: string }
