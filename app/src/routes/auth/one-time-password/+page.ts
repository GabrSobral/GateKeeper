import { zod } from 'sveltekit-superforms/adapters';
import { superValidate } from 'sveltekit-superforms';

import type { PageLoad } from './$types';
import { formSchema } from './schema';

export const load: PageLoad = async () => {
	return {
		form: await superValidate(zod(formSchema))
	};
};
