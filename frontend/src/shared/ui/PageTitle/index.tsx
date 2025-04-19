import { ReactNode } from 'react'

export function PageTitle({ children }: { children: ReactNode }) {
  return <h2 className="text-xl font-semibold mb-4">{children}</h2>
}
