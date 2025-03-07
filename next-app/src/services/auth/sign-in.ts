import { api } from "../base/gatekeeper-api";
import { APIError, Result } from "@/types/service-options";

type Request = {
  authorizationCode: string;
};

type Response = {
  user: {
    id: string;
    displayName: string;
    firstName: string;
    lastName: string;
    email: string;
    createdAt: Date;
    PhotoURL: string | null;
  };
  accessToken: string;
  refreshToken: string;
};

export async function signInApi({
  authorizationCode,
}: Request): Promise<Result<Response, APIError>> {
  try {
    const { data } = await api.post<Response>(`/v1/auth/authorize`, {
      authorizationCode,
    });

    return [data, null];
  } catch (error: unknown) {
    return [null, error as APIError];
  }
}
