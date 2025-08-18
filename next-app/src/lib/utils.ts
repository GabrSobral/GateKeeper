import { clsx, type ClassValue } from "clsx";
import { toast } from "sonner";
import { twMerge } from "tailwind-merge";

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
}

export function copy(value: string) {
  navigator.clipboard.writeText(value);
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

/**
 * Return the value of a cookie by its name.
 * @param name Nome do cookie
 * @returns Cookie value if it exists
 */
export function getCookieValue(name: string): string | null {
  if (typeof document === "undefined") return null;

  const cookies = document.cookie.split(";").map((cookie) => cookie.trim());
  if (!cookies.length) return null;
  for (const cookie of cookies) {
    const [key, ...rest] = cookie.split("=");
    if (key === name) {
      return decodeURIComponent(rest.join("="));
    }
  }

  return null;
}
