import { api } from "../base/gatekeeper-api";
import { APIError, IServiceOptions, Result } from "@/types/service-options";

type Request = {
  organizationId: string;
  applicationId: string;
  name: string;
  clientId: string;
  clientSecret: string;
  redirectUri: string;
  enabled: boolean;
};

export async function configureApplicationOauthProviderApi(
  {
    applicationId,
    clientId,
    clientSecret,
    enabled,
    name,
    organizationId,
    redirectUri,
  }: Request,
  { accessToken }: IServiceOptions
): Promise<Result<Response, APIError>> {
  try {
    const { data } = await api.put<Response>(
      `/v1/organizations/${organizationId}/applications/${applicationId}/oauth-provider`,
      {
        name,
        clientId,
        clientSecret,
        redirectUri,
        enabled,
      },
      {
        headers: {
          Authorization: `Bearer ${accessToken}`,
        },
      }
    );
    return [data, null];
  } catch (error: unknown) {
    return [null, error as APIError];
  }
}
