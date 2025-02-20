import type { AxiosError } from "axios";
import { api } from "../base/gatekeeper-api";
import { IServiceOptions, Result } from "@/types/service-options";

type Request = {
  name: string;
  description?: string;
  passwordHashSecret: string;
  badges: string[];
  hasMfaEmail: boolean;
  hasMfaAuthApp: boolean;
  organizationId: string;
};

type Response = {
  id: string;
  name: string;
  description?: string;
  passwordHashSecret: string;
  badges: string[];
  hasMfaEmail: boolean;
  hasMfaAuthApp: boolean;
};

export async function createApplicationApi(
  {
    name,
    description,
    badges,
    hasMfaAuthApp,
    hasMfaEmail,
    passwordHashSecret,
    organizationId,
  }: Request,
  { accessToken }: IServiceOptions
): Promise<Result<Response, AxiosError<{ message: string }>>> {
  try {
    const { data } = await api.post<Response>(
      `/v1/organizations/${organizationId}/applications`,
      {
        name,
        description,
        badges,
        hasMfaAuthApp,
        hasMfaEmail,
        passwordHashSecret,
        organizationId,
      },
      {
        headers: {
          Authorization: `Bearer ${accessToken}`,
        },
      }
    );
    return [data, null];
  } catch (error: unknown) {
    return [null, error as AxiosError<{ message: string }>];
  }
}
