import { api } from "../base/gatekeeper-api";
import { APIError, IServiceOptions, Result } from "@/types/service-options";

type Request = {
  userId: string;
  applicationId: string;
  organizationId: string;
  displayName: string;
  firstName: string;
  lastName: string;
  email: string;
  isEmailConfirmed: boolean;
  temporaryPasswordHash: string | null;
  isMfaAuthAppEnabled: boolean;
  isMfaEmailEnabled: boolean;
  roles: string[];
  isActive: boolean;
};

type Response = {
  id: string;
  displayName: string;
  firstName: string;
  lastName: string;
  email: string;
  isEmailConfirmed: boolean;
  temporaryPasswordHash: string;
  isMfaAuthAppEnabled: boolean;
  isMfaEmailEnabled: boolean;
  roles: string[];
  applicationId: string;
  organizationId: string;
};

export async function editApplicationUserApi(
  {
    userId,
    applicationId,
    displayName,
    firstName,
    lastName,
    email,
    isEmailConfirmed,
    isMfaAuthAppEnabled,
    isMfaEmailEnabled,
    temporaryPasswordHash,
    roles,
    organizationId,
    isActive,
  }: Request,
  { accessToken }: IServiceOptions
): Promise<Result<Response, APIError>> {
  try {
    const { data } = await api.put<Response>(
      `/v1/organizations/${organizationId}/applications/${applicationId}/users/${userId}`,
      {
        displayName,
        firstName,
        lastName,
        email,
        isEmailConfirmed,
        isMfaAuthAppEnabled,
        isMfaEmailEnabled,
        temporaryPasswordHash,
        roles,
        isActive,
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
