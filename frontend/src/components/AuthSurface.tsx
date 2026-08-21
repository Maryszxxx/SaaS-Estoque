import type { ReactNode } from 'react'

interface AuthSurfaceProps {
  documentCode: string
  children: ReactNode
  footer: ReactNode
}

export function AuthSurface({ documentCode, children, footer }: AuthSurfaceProps) {
  return <main className="kraft-surface min-h-screen px-4 py-4 sm:p-8 lg:grid lg:place-items-center lg:p-10">
    <div className="mx-auto grid w-full max-w-6xl gap-6 lg:grid-cols-[.9fr_1.1fr] lg:items-center">
      <section className="relative hidden min-h-145 overflow-hidden rounded-2xl border border-[#3D332B]/25 bg-[#3D332B]/18 p-10 text-[#F5F0E8] lg:flex lg:flex-col lg:justify-between">
        <div><div className="inline-flex items-center gap-3 border border-[#F5F0E8]/35 px-3 py-2 text-xs font-semibold tracking-[.18em]"><span className="size-2 bg-[#173F3A]" /> STOC</div><p className="mt-7 text-xs font-medium tracking-[.2em] text-[#F5F0E8]/75">GESTÃO DE ESTOQUE</p></div>
        <div className="relative max-w-sm"><div className="mb-7 h-px w-16 bg-[#F5F0E8]/55" /><h1 className="text-5xl font-semibold leading-[1.08] tracking-[-.05em]">Visibilidade clara para decisões melhores.</h1><p className="mt-7 text-lg leading-8 text-[#F5F0E8]/85">Centralize seus produtos e mantenha sua operação em movimento.</p></div>
        <div className="flex items-end justify-between border-t border-[#F5F0E8]/25 pt-5 text-xs text-[#F5F0E8]/70"><span>CONTROLE · ORGANIZAÇÃO · FLUXO</span><svg viewBox="0 0 80 56" aria-hidden="true" className="h-12 w-16 fill-none stroke-current stroke-[1.2]"><path d="M10 18 40 4l30 14v28L40 52 10 46Z" /><path d="m10 18 30 15 30-15M40 33v19" /><path d="M28 10 58 25" /></svg></div>
      </section>
      <section className="document-sheet relative overflow-hidden rounded-2xl border border-[#735A45]/30 px-6 py-8 sm:px-10 sm:py-10 lg:px-14 lg:py-12">
        <div className="flex items-center justify-between text-[10px] font-semibold tracking-[.18em] text-[#735A45]"><span>STOC · DOCUMENTO DE ACESSO</span><span>{documentCode}</span></div>
        <div className="document-rule mt-5 pt-8">{children}</div>
        <div className="document-rule mt-8 pt-5 text-center text-sm text-[#735A45]">{footer}</div>
      </section>
    </div>
  </main>
}
