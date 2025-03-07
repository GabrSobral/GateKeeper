import { api } from "../base/gatekeeper-api";
import { APIError, ResultWithoutResponse } from "@/types/service-options";

type Request = {
  applicationId: string;
  passwordResetToken: string;
  passwordResetId: string;
  newPassword: string;
};

export async function resetPasswordApi({
  applicationId,
  passwordResetId,
  passwordResetToken,
  newPassword,
}: Request): Promise<ResultWithoutResponse<APIError>> {
  try {
    await api.post<Response>(`/v1/auth/reset-password`, {
      applicationId,
      passwordResetId,
      passwordResetToken,
      newPassword,
    });

    return [null];
  } catch (error: unknown) {
    return [error as APIError];
  }
}
