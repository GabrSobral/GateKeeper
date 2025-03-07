import { AlertCircle } from "lucide-react";
import { Alert, AlertDescription, AlertTitle } from "./ui/alert";

type Props = {
  title: string;
  message: string;
};

export function ErrorAlert({ message, title }: Props) {
  return (
    <Alert variant="destructive" className="bg-red-500/10">
      <AlertCircle className="h-4 w-4" />
      <AlertTitle>{title}</AlertTitle>
      <AlertDescription>{message}</AlertDescription>
    </Alert>
  );
}
