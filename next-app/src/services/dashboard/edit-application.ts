import type { AxiosError } from "axios";
import { api } from "../base/gatekeeper-api";
import { IServiceOptions, Result } from "@/types/service-options";

type Request = {
  id: string;
  name: string;
  description?: string;
  passwordHashSecret: string;
  badges: string[];
  hasMfaEmail: boolean;
  hasMfaAuthApp: boolean;
  organizationId: string;
  isActive: boolean;
};

type Response = {
  id: string;
  name: string;
  description?: string;
  passwordHashSecret: string;
  badges: string[];
  hasMfaEmail: boolean;
  hasMfaAuthApp: boolean;
  isActive: boolean;
};

export async function editApplicationApi(
  {
    id,
    name,
    description,
    badges,
    hasMfaAuthApp,
    hasMfaEmail,
    passwordHashSecret,
    organizationId,
    isActive,
  }: Request,
  { accessToken }: IServiceOptions
): Promise<Result<Response, AxiosError<{ message: string }>>> {
  try {
    const { data } = await api.put<Response>(
      `/v1/organizations/${organizationId}/applications/${id}`,
      {
        id,
        name,
        description,
        badges,
        hasMfaAuthApp,
        hasMfaEmail,
        passwordHashSecret,
        organizationId,
        isActive,
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
