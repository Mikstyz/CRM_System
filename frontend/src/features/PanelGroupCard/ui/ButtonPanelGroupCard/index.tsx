interface ButtonPanelGroupCardProps {
  onClick: () => void;
  children: React.ReactNode;
  title?: string;
  className?: string;
}

export function ButtonPanelGroupCard({
  onClick,
  children,
  title,
  className,
}: ButtonPanelGroupCardProps) {
  return (
    <button
      onClick={onClick}
      className={` w-[2.2em] h-[2.2em] flex justify-center items-center bg-[#2a2a2a] text-white font-bold rounded ${className}`}
      title={title}
    >
      {children}
    </button>
  );
}
