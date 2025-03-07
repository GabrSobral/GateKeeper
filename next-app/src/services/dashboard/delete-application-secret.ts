import {
  APIError,
  IServiceOptions,
  ResultWithoutResponse,
} from "@/types/service-options";
import { api } from "../base/gatekeeper-api";

type Request = {
  secretId: string;
  applicationId: string;
  organizationId?: string;
};

export async function deleteApplicationSecretApi(
  { secretId, applicationId, organizationId }: Request,
  { accessToken }: IServiceOptions
): Promise<ResultWithoutResponse<APIError>> {
  try {
    await api.delete(
      `/v1/organizations/${organizationId}/applications/${applicationId}/secrets/${secretId}`,
      {
        headers: {
          Authorization: `Bearer ${accessToken}`,
        },
      }
    );
    return [null];
  } catch (error: unknown) {
    return [error as APIError];
  }
}
