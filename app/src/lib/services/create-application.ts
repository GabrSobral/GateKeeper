import type { AxiosError } from 'axios';
import { api } from './base/gatekeeper-api';
import type { IServiceOptions } from '$lib/types/service-options';

type Request = {
	name: string;
	description: string | null;
	passwordHashSecret: string;
	badges: string[];
	hasMfaEmail: boolean;
	hasMfaAuthApp: boolean;
	organizationId: string;
};

type Response = {
	id: string;
	name: string;
	description: string | null;
	passwordHashSecret: string;
	badges: string[];
	hasMfaEmail: boolean;
	hasMfaAuthApp: boolean;
	organizationId: string;
};

export async function createApplicationApi(
	{
		name,
		description,
		badges,
		hasMfaAuthApp,
		hasMfaEmail,
		organizationId,
		passwordHashSecret
	}: Request,
	{ accessToken }: IServiceOptions
): Promise<[Response | null, AxiosError<{ message: string }> | null]> {
	try {
		const { data } = await api.post<Response>(
			`/v1/organizations/${organizationId}/applications`,
			{
				name,
				description,
				passwordHashSecret,
				badges,
				hasMfaEmail,
				hasMfaAuthApp,
				organizationId
			},
			{
				headers: {
					Authorization: `Bearer ${accessToken}`
				}
			}
		);
		return [data, null];
	} catch (error: unknown) {
		return [null, error as AxiosError<{ message: string }>];
	}
}
