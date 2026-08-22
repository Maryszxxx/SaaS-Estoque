import { NavLink, Outlet, useNavigate } from 'react-router-dom'
import { useAuth } from '../hooks/useAuth'

function BoxIcon() { return <svg viewBox="0 0 32 32" aria-hidden="true" className="size-10 fill-none stroke-current stroke-[1.55]"><path d="m4 10 12-6 12 6v13l-12 6-12-6Z" /><path d="m4 10 12 6 12-6M16 16v13" /></svg> }

export function AppLayout() {
  const { user, logout } = useAuth()
  const navigate = useNavigate()
  const handleLogout = () => { logout(); navigate('/login') }

  return <div className="kraft-surface min-h-screen p-3 sm:p-5 lg:flex lg:gap-5">
    <aside className="corrugated-panel hidden w-65 shrink-0 flex-col px-6 py-7 text-[#302A25] lg:flex">
      <NavLink to="/dashboard" className="flex items-center gap-3 border-b border-[#735A45]/45 pb-7 text-3xl font-semibold tracking-[-.05em]"><BoxIcon /> Stoc</NavLink>
      <div className="mt-6 border-y border-[#735A45]/45 py-4"><p className="dashboard-hand text-base font-bold tracking-[.12em] text-[#173F3A]">PAINEL DE CONTROLE</p><p className="mt-1 text-xs text-[#735A45]">Navegação principal</p></div>
      <nav className="mt-6 space-y-1" aria-label="Navegação principal"><NavLink to="/dashboard" className={({ isActive }) => `dashboard-hand flex items-center gap-3 px-4 py-3.5 text-base font-medium transition ${isActive ? 'paper-nav-active text-[#173F3A]' : 'text-[#3D332B]'}`}><span className="grid size-5 place-items-center border border-current text-[10px]">▦</span>Dashboard</NavLink><NavLink to="/products" className={({ isActive }) => `dashboard-hand flex items-center gap-3 px-4 py-3.5 text-base font-medium transition ${isActive ? 'paper-nav-active text-[#173F3A]' : 'text-[#3D332B]'}`}><span className="grid size-5 place-items-center border border-current text-[10px]">□</span>Produtos</NavLink><NavLink to="/categories" className={({ isActive }) => `dashboard-hand flex items-center gap-3 px-4 py-3.5 text-base font-medium transition ${isActive ? 'paper-nav-active text-[#173F3A]' : 'text-[#3D332B]'}`}><span className="grid size-5 place-items-center border border-current text-[10px]">▤</span>Categorias</NavLink><NavLink to="/settings" className={({ isActive }) => `dashboard-hand flex items-center gap-3 px-4 py-3.5 text-base font-medium transition ${isActive ? 'paper-nav-active text-[#173F3A]' : 'text-[#3D332B]'}`}><span className="grid size-5 place-items-center border border-current text-[10px]">⚙</span>Configurações</NavLink></nav>
      <div className="dashboard-hand mt-auto border-t border-[#735A45]/45 pt-6 text-lg"><p className="flex items-center gap-3"><span className="grid size-6 place-items-center rounded-full border border-current text-xs">◦</span>{user?.role}</p><button onClick={handleLogout} className="mt-4 text-sm text-[#735A45] hover:text-[#173F3A]">Sair da conta</button></div>
    </aside>
    <div className="min-w-0 flex-1"><header className="paper-strip flex items-center justify-between px-6 py-5 sm:px-10"><NavLink to="/dashboard" className="flex items-center gap-2 text-2xl font-semibold tracking-[-.04em] text-[#173F3A] lg:hidden"><BoxIcon />Stoc</NavLink><div className="hidden lg:block"><p className="dashboard-hand text-xl text-[#173F3A]">DOCUMENTO OPERACIONAL</p><p className="dashboard-hand mt-1 text-base text-[#3D332B]">Visão atual do estoque</p></div><span className="dashboard-hand paper-label text-base text-[#302A25]">{user?.role}</span></header><main className="paper-dashboard mt-4 min-h-[calc(100vh-3rem)] p-6 sm:mt-5 sm:p-9 lg:p-10"><Outlet /></main></div>
  </div>
}
