import { useQuery } from '@sveltestack/svelte-query';

import { api } from './base/gatekeeper-api';
import type { IServiceOptions } from '$lib/types/service-options';

type Request = {
	organizationId?: string;
};

type Response = {
	id: string;
	name: string;
	description: string;
	createdAt: Date;
	updatedAt: Date;
	deactivatedAt?: Date;
	badges: string[];
}[];

export async function getApplicationsService(
	request: Request,
	{ accessToken }: IServiceOptions
): Promise<Response> {
	const { data } = await api.get<Response>(
		`/v1/organizations/${request.organizationId}/applications`,
		{
			headers: {
				Authorization: `Bearer ${accessToken}`
			}
		}
	);

	return data;
}

export function useApplicationsQuery(request: Request, options: IServiceOptions) {
	return useQuery(
		['list-applications', request.organizationId],
		() => getApplicationsService(request, options),
		{
			enabled: !!request.organizationId,
			refetchOnWindowFocus: false
		}
	);
}
