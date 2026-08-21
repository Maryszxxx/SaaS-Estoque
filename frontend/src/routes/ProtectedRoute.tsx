import { Navigate, Outlet, useLocation } from 'react-router-dom'
import { useAuth } from '../hooks/useAuth'
import type { UserRole } from '../types/auth'
export function ProtectedRoute({ roles }: { roles?: UserRole[] }) { const { isAuthenticated, isLoading, user } = useAuth(); const location = useLocation(); if (isLoading) return <div className="grid min-h-screen place-items-center text-sm text-slate-500">Carregando sessão...</div>; if (!isAuthenticated) return <Navigate to="/login" replace state={{ from: location }} />; if (roles && (!user || !roles.includes(user.role))) return <Navigate to="/dashboard" replace />; return <Outlet /> }
