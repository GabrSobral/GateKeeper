import { api } from "../base/gatekeeper-api";
import { APIError, ResultWithoutResponse } from "@/types/service-options";

type Request = {
  applicationId: string;
  email: string;
};

export async function resendConfirmEmailApi({
  email,
  applicationId,
}: Request): Promise<ResultWithoutResponse<APIError>> {
  try {
    await api.post<Response>(`/v1/auth/confirm-email/resend`, {
      email,
      applicationId,
    });

    return [null];
  } catch (error: unknown) {
    return [error as APIError];
  }
}
