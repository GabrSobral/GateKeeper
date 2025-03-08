type Props = {
  className?: string;
};

export function LoadingSpinner({ className }: Props) {
  return (
    <div className={className}>
      <div className="spinner"></div>
    </div>
  );
}
