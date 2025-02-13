import type { AxiosError } from 'axios';
import { api } from './base/gatekeeper-api';
import type { IServiceOptions } from '$lib/types/service-options';

type Request = {
	name: string;
	expiresAt: Date | null;
	applicationId: string;
	organizationId?: string;
};

type Response = {
	id: string;
	name: string;
	value: string;
	createdAt: Date;
	expiresAt: Date | null;
};

export async function createApplicationSecretApi(
	{ name, expiresAt, applicationId, organizationId }: Request,
	{ accessToken }: IServiceOptions
): Promise<[Response | null, AxiosError<{ message: string }> | null]> {
	try {
		const { data } = await api.post<Response>(
			`/v1/organizations/${organizationId}/applications/${applicationId}/secrets`,
			{
				name,
				expiresAt
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
