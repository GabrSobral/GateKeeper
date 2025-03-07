import { api } from "../base/gatekeeper-api";
import { APIError, ResultWithoutResponse } from "@/types/service-options";

type Request = {
  applicationId: string;
  displayName: string;
  firstName: string;
  lastName: string;
  email: string;
  password: string;
};

export async function signUpApi({
  applicationId,
  displayName,
  firstName,
  email,
  lastName,
  password,
}: Request): Promise<ResultWithoutResponse<APIError>> {
  try {
    await api.post<Response>(`/v1/auth/sign-up`, {
      displayName,
      firstName,
      email,
      lastName,
      password,
      applicationId,
    });

    return [null];
  } catch (error: unknown) {
    return [error as APIError];
  }
}
