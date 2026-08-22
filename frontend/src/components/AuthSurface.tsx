import type { ReactNode } from 'react'

interface AuthSurfaceProps { documentCode: string; children: ReactNode; footer: ReactNode }

function BoxMark() { return <svg viewBox="0 0 32 32" aria-hidden="true" className="size-9 fill-none stroke-current stroke-[1.55]"><path d="m4 10 12-6 12 6v13l-12 6-12-6Z" /><path d="m4 10 12 6 12-6M16 16v13" /></svg> }

export function AuthSurface({ documentCode, children, footer }: AuthSurfaceProps) {
  return <main className="kraft-surface relative min-h-screen overflow-hidden px-4 py-8 sm:p-10 lg:grid lg:place-items-center lg:px-16 lg:py-14">
    <div className="auth-paper-grid relative mx-auto grid w-full max-w-6xl gap-8 lg:grid-cols-[.92fr_1.08fr] lg:items-center">
      <section className="paper-note paper-note-left relative hidden min-h-155 p-12 text-[#302A25] lg:flex lg:flex-col lg:justify-between"><span className="paper-tape paper-tape-top" aria-hidden="true" /><span className="paper-tape paper-tape-bottom" aria-hidden="true" /><div><div className="flex items-center gap-3 text-[#173F3A]"><BoxMark /><span className="text-2xl font-semibold tracking-[-.035em]">STOC</span></div><p className="mt-10 text-xs font-semibold tracking-[.22em] text-[#735A45]">GESTÃO DE ESTOQUE</p></div><div><h1 className="max-w-sm text-5xl font-semibold leading-[1.18] tracking-[-.055em]">Visibilidade clara para decisões melhores.</h1><div className="mt-7 h-px w-15 bg-[#BCA68F]" /><p className="mt-6 max-w-sm text-lg leading-8 text-[#3D332B]">Centralize seus produtos e mantenha sua operação em movimento.</p></div><div className="flex items-end justify-between border-t border-[#BCA68F] pt-6 text-xs font-medium tracking-wide text-[#735A45]"><span>CONTROLE · ORGANIZAÇÃO · FLUXO</span><BoxMark /></div></section>
      <section className="paper-note paper-note-form relative px-7 py-9 sm:px-12 sm:py-12 lg:px-14"><span className="paper-tape paper-tape-center" aria-hidden="true" /><div className="flex items-center justify-between text-[10px] font-semibold tracking-[.18em] text-[#735A45]"><span>STOC · DOCUMENTO DE ACESSO</span><span>{documentCode}</span></div><div className="document-rule mt-5 pt-9">{children}</div><div className="document-rule mt-8 pt-5 text-center text-sm text-[#735A45]">{footer}</div></section>
    </div>
    <footer className="absolute inset-x-0 bottom-0 hidden border-t border-[#BCA68F] bg-[#F5F0E8] py-4 text-center text-xs text-[#735A45] lg:block">© 2025 Stoc. Todos os direitos reservados.</footer>
  </main>
}
