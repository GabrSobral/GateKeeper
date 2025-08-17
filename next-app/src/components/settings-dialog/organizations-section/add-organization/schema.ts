import { z } from "zod";

export const formSchema = z.object({
  name: z
    .string()
    .min(2, "Name must hast at least 2 characters.")
    .max(75, "Name must have at most 75 characters."),

  description: z
    .string()
    .min(2, "Description must hast at least 2 characters.")
    .max(250, "Description must have at most 250 characters."),
});
