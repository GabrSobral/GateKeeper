import useSWR from "swr";

import { api } from "../base/gatekeeper-api";

import type { IServiceOptions } from "@/types/service-options";

type Request = {
  organizationId?: string;
  applicationId: string;
};

type Response = {
  id: string;
  name: string;
  description: string;
}[];

const fetcher = (url: string, options: IServiceOptions) =>
  api
    .get<Response>(url, {
      headers: {
        Authorization: `Bearer ${options.accessToken}`,
      },
    })
    .then((res) => res.data);

export function useApplicationRolesSWR(
  request: Request,
  options: IServiceOptions
) {
  return useSWR(
    request?.organizationId
      ? `/v1/organizations/${request?.organizationId}/applications/${request?.applicationId}/roles`
      : null,
    (url) => fetcher(url, options),
    { revalidateOnFocus: false }
  );
}
