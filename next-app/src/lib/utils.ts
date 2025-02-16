import { clsx, type ClassValue } from "clsx";
import { toast } from "sonner";
import { twMerge } from "tailwind-merge";

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
}

export function copy(value: string) {
  navigator.clipboard.writeText("Hello, world!");
  toast.success(`"${value}" was copied to clipboard!`);
}

export function formatDate(date: Date) {
  // Format to "MM/DD/YYYY at HH:MM"
  return new Intl.DateTimeFormat("en-US", {
    month: "2-digit",
    day: "2-digit",
    year: "numeric",
    hour: "numeric",
    minute: "numeric",
  }).format(date);
}
