import { zodResolver } from '@hookform/resolvers/zod'
import { useState } from 'react'
import { useForm } from 'react-hook-form'
import { Link, Navigate, useLocation, useNavigate } from 'react-router-dom'
import { z } from 'zod'
import { AuthSurface } from '../components/AuthSurface'
import { useAuth } from '../hooks/useAuth'

const loginSchema = z.object({ email: z.email('Informe um e-mail válido.'), password: z.string().min(1, 'Informe sua senha.') })
type LoginForm = z.infer<typeof loginSchema>

export function LoginPage() {
  const { login, isAuthenticated } = useAuth()
  const navigate = useNavigate()
  const location = useLocation()
  const [error, setError] = useState('')
  const { register, handleSubmit, formState: { errors, isSubmitting } } = useForm<LoginForm>({ resolver: zodResolver(loginSchema) })
  const routeState = location.state as { from?: { pathname?: string }; success?: string } | null
  const destination = routeState?.from?.pathname ?? '/dashboard'
  if (isAuthenticated) return <Navigate to="/dashboard" replace />
  const onSubmit = async (data: LoginForm) => { setError(''); try { await login(data); navigate(destination, { replace: true }) } catch { setError('Não foi possível entrar. Verifique seu e-mail, senha e a conexão com a API.') } }

  return <AuthSurface documentCode="AC-01" footer={<><span>Não possui uma conta? </span><Link to="/register" className="font-semibold text-[#173F3A] underline decoration-[#173F3A]/35 underline-offset-4 hover:decoration-[#173F3A]">Crie seu usuário</Link></>}><p className="text-xs font-semibold tracking-[.18em] text-[#735A45]">ACESSO À PLATAFORMA</p><h2 className="mt-4 text-3xl font-semibold tracking-[-.045em] text-[#302A25] sm:text-4xl">Boas-vindas de volta</h2><p className="mt-3 leading-7 text-[#735A45]">Entre para acessar a gestão do seu estoque.</p>{routeState?.success && <div role="status" className="mt-7 border border-[#789685] bg-[#EDF4EE] px-4 py-3 text-sm text-[#173F3A]">{routeState.success}</div>}<form onSubmit={handleSubmit(onSubmit)} className="mt-8 space-y-5" noValidate>{error && <div role="alert" className="border border-[#C99787] bg-[#FFF4EF] px-4 py-3 text-sm text-[#8A3D32]">{error}</div>}<label className="block text-sm font-semibold text-[#3D332B]">E-mail<input type="email" autoComplete="email" {...register('email')} className="mt-2 w-full border border-[#BCA68F] bg-[#FAF8F3] px-4 py-3 text-[#302A25] outline-none transition focus:border-[#173F3A] focus:ring-3 focus:ring-[#173F3A]/15" aria-invalid={Boolean(errors.email)} />{errors.email && <span className="mt-1.5 block text-sm text-[#8A3D32]">{errors.email.message}</span>}</label><label className="block text-sm font-semibold text-[#3D332B]">Senha<input type="password" autoComplete="current-password" {...register('password')} className="mt-2 w-full border border-[#BCA68F] bg-[#FAF8F3] px-4 py-3 text-[#302A25] outline-none transition focus:border-[#173F3A] focus:ring-3 focus:ring-[#173F3A]/15" aria-invalid={Boolean(errors.password)} />{errors.password && <span className="mt-1.5 block text-sm text-[#8A3D32]">{errors.password.message}</span>}</label><button disabled={isSubmitting} className="w-full bg-[#173F3A] px-4 py-3.5 font-semibold text-[#F5F0E8] transition hover:bg-[#102F2B] focus:outline-none focus:ring-3 focus:ring-[#173F3A]/30 disabled:opacity-60">{isSubmitting ? 'Entrando...' : 'Entrar'}</button></form></AuthSurface>
}
