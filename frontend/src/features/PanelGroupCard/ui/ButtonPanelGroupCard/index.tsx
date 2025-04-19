interface ButtonPanelGroupCardProps {
  onClick: () => void
  children: React.ReactNode
  title?: string
  className?: string
}

export function ButtonPanelGroupCard({ onClick, children, title, className }: ButtonPanelGroupCardProps) {
  return (
    <button onClick={onClick} className={` text-white font-bold py-1 px-2 rounded ${className}`} title={title}>
      {children}
    </button>
  )
}
