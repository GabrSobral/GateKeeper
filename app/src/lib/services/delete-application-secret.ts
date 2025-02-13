import type { AxiosError } from 'axios';
import { api } from './base/gatekeeper-api';
import type { IServiceOptions } from '$lib/types/service-options';

type Request = {
	secretId: string;
	applicationId: string;
	organizationId?: string;
};

export async function deleteApplicationSecretApi(
	{ secretId, applicationId, organizationId }: Request,
	{ accessToken }: IServiceOptions
): Promise<[AxiosError<{ message: string }> | null]> {
	try {
		await api.delete(
			`/v1/organizations/${organizationId}/applications/${applicationId}/secrets/${secretId}`,
			null,
			{
				headers: {
					Authorization: `Bearer ${accessToken}`
				}
			}
		);
		return [null];
	} catch (error: unknown) {
		return [error as AxiosError<{ message: string }>];
	}
}
