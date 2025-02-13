import type { AxiosError } from 'axios';
import { api } from './base/gatekeeper-api';
import type { IServiceOptions } from '$lib/types/service-options';

type Request = {
	name: string;
	description: string | null;
	applicationId: string;
	organizationId: string;
};

type Response = {
	id: string;
	name: string;
	description: string | null;
	createdAt: Date;
	updatedAt: Date | null;
};

export async function createApplicationRoleApi(
	{ name, description, applicationId, organizationId }: Request,
	{ accessToken }: IServiceOptions
): Promise<[Response | null, AxiosError<{ message: string }> | null]> {
	try {
		const { data } = await api.post<Response>(
			`/v1/organizations/${organizationId}/applications/${applicationId}/roles`,
			{
				name,
				description
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
