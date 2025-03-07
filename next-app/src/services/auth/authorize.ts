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

type Response = {
  authorizationCode: string;
  state: string;
  codeChallenge: string;
  codeChallengeMethod: string;
  responseType: string;
  redirectUri: string;
};

export async function authorizeApi({
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
    const { data } = await api.post<Response>(`/v1/auth/authorize`, {
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
