import { z } from 'zod';

export const formSchema = z.object({
	name: z.string().min(2).max(50),
	description: z.string().optional(),
	passwordHashSecret: z.string(),
	badges: z.array(z.string()),
	hasMfaAuthApp: z.boolean(),
	hasMfaEmail: z.boolean()
});

export type FormSchema = typeof formSchema;
