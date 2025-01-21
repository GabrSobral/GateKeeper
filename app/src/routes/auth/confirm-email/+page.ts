import { zod } from 'sveltekit-superforms/adapters';
import { superValidate } from 'sveltekit-superforms';

import { formSchema } from './schema';
import type { PageLoad } from '../one-time-password/$types';

export const load: PageLoad = async () => {
	return {
		form: await superValidate(zod(formSchema))
	};
};
