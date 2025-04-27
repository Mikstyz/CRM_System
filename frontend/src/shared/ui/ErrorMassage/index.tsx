interface Props {
  errorMassage?: string;
  className?: string;
  error: string | boolean | null | undefined;
}

export function ErrorMassage({ errorMassage, error, className }: Props) {
  if (error) {
    return (
      <div
        className={
          "flex flex-col items-center justify-center h-full " + className
        }
      >
        <p className="text-red-500">{errorMassage || error}</p>
      </div>
    );
  }
}
