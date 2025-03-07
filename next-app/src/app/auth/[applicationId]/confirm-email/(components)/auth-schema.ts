import { z } from "zod";

export const formSchema = z.object({
  code: z.string({ message: "Code must contain 6 characters" }).min(6).max(6),
});
