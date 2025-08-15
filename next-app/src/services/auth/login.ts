import { api } from "../base/gatekeeper-api";
import { APIError, Result } from "@/types/service-options";

type Request = {
  email: string;
  password: string;
  applicationId: string;
  redirectUri: string;
  codeChallengeMethod: string;
  responseType: string;
  scope: string;
  state: string;
  codeChallenge: string;
};

export enum EMfaType {
  MfaApp = "totp",
  MfaEmail = "email",
}

type Response = {
  mfaType: EMfaType;
  mfaId?: string;
  message: string;
  sessionCode: string;
  userId: string;
  changePasswordCode: string;
};

export async function loginApi({
  email,
  password,
  applicationId,
  redirectUri,
  codeChallengeMethod,
  responseType,
  scope,
  state,
  codeChallenge,
}: Request): Promise<Result<Response, APIError>> {
  try {
    const { data } = await api.post<Response>(`/v1/auth/login`, {
      email,
      password,
      applicationId,
      redirectUri,
      codeChallengeMethod,
      responseType,
      scope,
      codeChallenge,
      state,
    });
    return [data, null];
  } catch (error: unknown) {
    return [null, error as APIError];
  }
}
