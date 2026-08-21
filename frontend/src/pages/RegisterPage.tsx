import { zodResolver } from '@hookform/resolvers/zod'
import { useState } from 'react'
import type { ReactNode } from 'react'
import { useForm } from 'react-hook-form'
import { Link, useNavigate } from 'react-router-dom'
import { z } from 'zod'
import { AuthSurface } from '../components/AuthSurface'
import { registerUser } from '../services/auth'

const registerSchema = z.object({ name: z.string().min(4, 'Informe um nome com pelo menos 4 caracteres.').max(90), email: z.email('Informe um e-mail válido.'), password: z.string().min(6, 'A senha deve ter ao menos 6 caracteres.').regex(/[0-9]/, 'Inclua um número.').regex(/[A-Z]/, 'Inclua uma letra maiúscula.').regex(/[a-z]/, 'Inclua uma letra minúscula.'), confirmPassword: z.string(), role: z.enum(['ADMIN', 'EMPLOYEE'], { error: 'Selecione uma role.' }) }).refine((data) => data.password === data.confirmPassword, { path: ['confirmPassword'], message: 'As senhas não coincidem.' })
type RegisterForm = z.infer<typeof registerSchema>

export function RegisterPage() {
  const navigate = useNavigate(); const [error, setError] = useState('')
  const { register, handleSubmit, formState: { errors, isSubmitting } } = useForm<RegisterForm>({ resolver: zodResolver(registerSchema), defaultValues: { role: 'EMPLOYEE' } })
  const onSubmit = async ({ confirmPassword: _confirmPassword, ...data }: RegisterForm) => { setError(''); try { await registerUser(data); navigate('/login', { replace: true, state: { success: 'Usuário criado com sucesso. Agora você já pode entrar.' } }) } catch { setError('Não foi possível criar o usuário. Verifique os dados e tente novamente.') } }
  return <AuthSurface documentCode="RC-02" footer={<><span>Já possui uma conta? </span><Link to="/login" className="font-semibold text-[#173F3A] underline decoration-[#173F3A]/35 underline-offset-4 hover:decoration-[#173F3A]">Voltar para login</Link></>}><p className="text-xs font-semibold tracking-[.18em] text-[#735A45]">CADASTRO INICIAL</p><h2 className="mt-4 text-3xl font-semibold tracking-[-.045em] text-[#302A25] sm:text-4xl">Crie seu usuário</h2><p className="mt-3 leading-7 text-[#735A45]">Preencha os dados para iniciar o controle do seu estoque.</p><form onSubmit={handleSubmit(onSubmit)} className="mt-8 grid gap-5 sm:grid-cols-2" noValidate>{error && <div role="alert" className="border border-[#C99787] bg-[#FFF4EF] px-4 py-3 text-sm text-[#8A3D32] sm:col-span-2">{error}</div>}<FormField label="Nome" error={errors.name?.message}><input autoComplete="name" {...register('name')} className="document-input" /></FormField><FormField label="E-mail" error={errors.email?.message}><input type="email" autoComplete="email" {...register('email')} className="document-input" /></FormField><FormField label="Senha" error={errors.password?.message}><input type="password" autoComplete="new-password" {...register('password')} className="document-input" /></FormField><FormField label="Confirmar senha" error={errors.confirmPassword?.message}><input type="password" autoComplete="new-password" {...register('confirmPassword')} className="document-input" /></FormField><FormField label="Role" error={errors.role?.message}><select {...register('role')} className="document-input"><option value="EMPLOYEE">EMPLOYEE</option><option value="ADMIN">ADMIN</option></select></FormField><div className="hidden sm:block" /><button disabled={isSubmitting} className="bg-[#173F3A] px-4 py-3.5 font-semibold text-[#F5F0E8] transition hover:bg-[#102F2B] focus:outline-none focus:ring-3 focus:ring-[#173F3A]/30 disabled:opacity-60 sm:col-span-2">{isSubmitting ? 'Criando usuário...' : 'Criar usuário'}</button></form></AuthSurface>
}

function FormField({ label, error, children }: { label: string; error?: string; children: ReactNode }) { return <label className="block text-sm font-semibold text-[#3D332B]">{label}<span className="mt-2 block">{children}</span>{error && <span className="mt-1.5 block text-sm font-normal text-[#8A3D32]">{error}</span>}</label> }
