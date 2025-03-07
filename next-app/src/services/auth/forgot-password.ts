import { api } from "../base/gatekeeper-api";
import { APIError, ResultWithoutResponse } from "@/types/service-options";

type Request = {
  applicationId: string;
  email: string;
};

export async function forgotPasswordApi({
  applicationId,
  email,
}: Request): Promise<ResultWithoutResponse<APIError>> {
  try {
    await api.post<Response>(`/v1/auth/forgot-password`, {
      applicationId,
      email,
    });

    return [null];
  } catch (error: unknown) {
    return [error as APIError];
  }
}
