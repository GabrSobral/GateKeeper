import useSWR from "swr";

import { api } from "../base/gatekeeper-api";

import type { IServiceOptions } from "@/types/service-options";

type Request = {
  organizationId?: string;
};

export type ApplicationsResponse = {
  id: string;
  name: string;
  description: string;
  createdAt: Date;
  updatedAt: Date;
  deactivatedAt?: Date;
  badges: string[];
}[];

const fetcher = (url: string, options: IServiceOptions) =>
  api
    .get<ApplicationsResponse>(url, {
      headers: {
        Authorization: `Bearer ${options.accessToken}`,
      },
    })
    .then((res) => res.data);

export function useApplicationsSWR(request: Request, options: IServiceOptions) {
  return useSWR(
    request?.organizationId
      ? `/v1/organizations/${request?.organizationId}/applications`
      : null,
    (url) => fetcher(url, options)
  );
}
