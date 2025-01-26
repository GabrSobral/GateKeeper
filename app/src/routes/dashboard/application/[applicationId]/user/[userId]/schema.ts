import { z } from 'zod';

export const formSchema = z.object({
	displayName: z.string().min(2).max(50),
	firstName: z.string().min(2).max(50),
	lastName: z.string().min(2).max(50).optional(),
	email: z.string().email(),
	multiFactorAuth: z.array(z.any()),
	roles: z.array(z.any()),
});

export type FormSchema = typeof formSchema;
