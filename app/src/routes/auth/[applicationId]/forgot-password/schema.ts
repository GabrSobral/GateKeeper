import { z } from "zod";
 
export const formSchema = z.object({
 firstName: z.string().min(2).max(50),
 lastName: z.string().min(2).max(50),
 email: z.string().email(),
 password: z.string()
});
 
export type FormSchema = typeof formSchema;