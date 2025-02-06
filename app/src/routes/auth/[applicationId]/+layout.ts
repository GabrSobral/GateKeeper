import { getApplicationByIdService } from '$lib/services/use-application-auth-query';
import type { LayoutLoad } from './$types';

export const load: LayoutLoad = async ({ params }) => {
	const applicationId = params.applicationId;
	
	return {
		applicationData: await getApplicationByIdService({ applicationId }),
	};
};