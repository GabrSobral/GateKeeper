import { z } from "zod";
 
export const formSchema = z.object({
 confirmPassword: z.string(),
 password: z.string()
});
 
export type FormSchema = typeof formSchema;