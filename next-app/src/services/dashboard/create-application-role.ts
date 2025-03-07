import { api } from "../base/gatekeeper-api";
import { APIError, IServiceOptions, Result } from "@/types/service-options";

type Request = {
  name: string;
  description: string;
  applicationId: string;
  organizationId: string;
};

type Response = {
  id: string;
  name: string;
  description: string;
  createdAt: Date;
  updatedAt: Date | null;
  applicationId: string;
};

export async function createApplicationRoleApi(
  { name, applicationId, description, organizationId }: Request,
  { accessToken }: IServiceOptions
): Promise<Result<Response, APIError>> {
  try {
    const { data } = await api.post<Response>(
      `/v1/organizations/${organizationId}/applications/${applicationId}/roles`,
      {
        name,
        description,
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
