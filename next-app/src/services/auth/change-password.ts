import { api } from "../base/gatekeeper-api";
import { APIError, ResultWithoutResponse } from "@/types/service-options";

type Request = {
  applicationId: string;
  userId: string;
  changePasswordCode: string;
  newPassword: string;
};

export async function changePasswordApi({
  applicationId,
  changePasswordCode,
  userId,
  newPassword,
}: Request): Promise<ResultWithoutResponse<APIError>> {
  try {
    await api.post<Response>(`/v1/auth/change-password`, {
      applicationId,
      changePasswordCode,
      userId,
      newPassword,
    });

    return [null];
  } catch (error: unknown) {
    return [error as APIError];
  }
}
