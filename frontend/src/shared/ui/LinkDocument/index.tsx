interface LinkDocumentProps {
  href: string
  children: React.ReactNode
  className?: string
}
export function LinkDocument({ href, children, className }: LinkDocumentProps) {
  return (
    <a
      href={href}
      target="_blank"
      rel="noopener noreferrer"
      className={`text-blue-600 hover:underline whitespace-nowrap ${className}`}
    >
      {children}
    </a>
  )
}
