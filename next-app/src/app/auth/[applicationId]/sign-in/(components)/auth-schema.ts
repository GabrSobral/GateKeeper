import { z } from "zod";

export const formSchema = z.object({
  email: z.string().email("Invalid e-mail address"),
  password: z.string().min(8, "Password must be at least 8 characters"),
});
