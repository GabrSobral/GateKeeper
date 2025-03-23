import { api } from "../base/gatekeeper-api";
import { APIError, Result } from "@/types/service-options";

type Request = {
  email: string;
  code: string;
  applicationId: string;
};

type Response = {
  sessionCode: string;
};

export async function verifyMfaApi({
  email,
  code,
  applicationId,
}: Request): Promise<Result<Response, APIError>> {
  try {
    const { data } = await api.post<Response>(`/v1/auth/verify-mfa`, {
      email,
      code,
      applicationId,
    });
    return [data, null];
  } catch (error: unknown) {
    return [null, error as APIError];
  }
}
