import { useQuery } from '@sveltestack/svelte-query';

import type { IServiceOptions } from '$lib/types/service-options';
import { api } from './base/gatekeeper-api';

type Response = {
	id: string;
	name: string;
	createdAt: Date;
};

async function getOrganizationsService({ accessToken }: IServiceOptions): Promise<Response[]> {
	const { data } = await api.get<Response[]>('/v1/organizations', {
		headers: {
			Authorization: `Bearer ${accessToken}`
		}
	});

	return data;
}

export function useOrganizationsQuery(options: IServiceOptions) {
	return useQuery('list-organizations', () => getOrganizationsService(options), {
		refetchOnWindowFocus: false
	});
}
