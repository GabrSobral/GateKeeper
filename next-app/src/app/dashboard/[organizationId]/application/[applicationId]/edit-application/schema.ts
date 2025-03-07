import { z } from "zod";

export const formSchema = z.object({
  name: z
    .string()
    .min(2, "Name must hast at least 2 characters.")
    .max(75, "Name must have at most 75 characters."),

  description: z.string().optional(),

  passwordHashSecret: z
    .string()
    .min(32, "Password hash secret must have at least 32 characters.")
    .max(128, "Password hash secret must have at most 128 characters."),

  badges: z.array(z.string()),
  hasMfaAuthApp: z.boolean(),
  hasMfaEmail: z.boolean(),
  canSelfSignUp: z.boolean(),
  canSelfForgotPass: z.boolean(),
});
