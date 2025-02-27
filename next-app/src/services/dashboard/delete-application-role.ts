import type { AxiosError } from "axios";
import { api } from "../base/gatekeeper-api";
import {
  IServiceOptions,
  ResultWithoutResponse,
} from "@/types/service-options";

type Request = {
  applicationId: string;
  organizationId: string;
  roleId: string;
};

export async function deleteApplicationRoleApi(
  { applicationId, organizationId, roleId }: Request,
  { accessToken }: IServiceOptions
): Promise<ResultWithoutResponse<AxiosError<{ message: string }>>> {
  try {
    await api.delete<Response>(
      `/v1/organizations/${organizationId}/applications/${applicationId}/roles/${roleId}`,
      {
        headers: {
          Authorization: `Bearer ${accessToken}`,
        },
      }
    );
    return [null];
  } catch (error: unknown) {
    return [error as AxiosError<{ message: string }>];
  }
}
