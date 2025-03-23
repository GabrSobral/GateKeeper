import { z } from "zod";

export const formSchema = z.object({
  code: z.string().min(6).max(6),
});
