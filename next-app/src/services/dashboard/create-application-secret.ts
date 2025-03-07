import { api } from "../base/gatekeeper-api";
import { APIError, IServiceOptions, Result } from "@/types/service-options";

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
): Promise<Result<Response, APIError>> {
  try {
    const { data } = await api.post<Response>(
      `/v1/organizations/${organizationId}/applications/${applicationId}/secrets`,
      {
        name,
        expiresAt,
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
