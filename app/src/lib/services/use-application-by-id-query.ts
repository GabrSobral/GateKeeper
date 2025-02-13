import type { IServiceOptions } from '$lib/types/service-options';
import { useQuery } from '@sveltestack/svelte-query';
import { api } from './base/gatekeeper-api';

export interface IApplication {
	id: string;
	name: string;
	description: string;
	badges: string[];
	createdAt: Date;
	updatedAt: Date;
	deactivatedAt?: Date;
	multiFactorAuthEnabled: boolean;
	passwordHashingSecret: string;
	secrets: {
		id: string;
		name: string;
		value: string;
		expirationDate: Date;
	}[];
	users: {
		totalCount: number;
		data: {
			id: string;
			displayName: string;
			email: string;
			roles: {
				id: string;
				name: string;
				description: string;
			}[];
			deactivatedAt?: Date;
		}[];
	};
	roles: {
		totalCount: number;
		data: {
			id: string;
			name: string;
			description: string;
		}[];
	};
	oauthProviders: {
		id: string;
		name: string;
		description: string;
		clientId: string;
		clientSecret: string;
		isEnabled: boolean;
	}[];
}

type Request = {
	applicationId: string;
	organizationId?: string;
};

type Response = IApplication;

async function getApplicationByIdService(
	{ applicationId, organizationId }: Request,
	{ accessToken }: IServiceOptions
): Promise<Response | null> {
	const { data } = await api.get<Response>(
		`/v1/organizations/${organizationId}/applications/${applicationId}`,
		{
			headers: {
				Authorization: `Bearer ${accessToken}`
			}
		}
	);

	return data;
}

export function useApplicationByIdQuery(request: Request, options: IServiceOptions) {
	return useQuery(
		['get-application-by-id', request.applicationId, request.organizationId],
		() => getApplicationByIdService(request, options),
		{
			enabled: !!request.applicationId && !!request.organizationId
		}
	);
}
