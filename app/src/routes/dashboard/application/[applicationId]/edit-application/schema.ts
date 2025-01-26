import { z } from 'zod';

export const formSchema = z.object({
	name: z.string().min(2).max(50),
	Description: z.string().optional(),
	passwordHashSecret: z.string(),
	badges: z.array(z.string()),
	hasMFA: z.boolean()
});

export type FormSchema = typeof formSchema;
