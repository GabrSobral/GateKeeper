import { api } from "../base/gatekeeper-api";
import { APIError, Result } from "@/types/service-options";

type Request = {
  email: string;
  code: string;
  mfaId: string;
  applicationId: string;
};

type Response = {
  sessionCode: string;
};

export async function verifyAppMfaApi({
  email,
  code,
  mfaId,
  applicationId,
}: Request): Promise<Result<Response, APIError>> {
  try {
    const { data } = await api.post<Response>(`/v1/auth/verify-mfa/app`, {
      email,
      code,
      mfaId,
      applicationId,
    });
    return [data, null];
  } catch (error: unknown) {
    return [null, error as APIError];
  }
}
