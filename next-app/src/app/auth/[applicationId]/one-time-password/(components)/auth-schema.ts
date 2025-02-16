import { z } from "zod";

export const formSchema = z.object({
  email: z.string().email("Invalid e-mail address"),
});
