import useSWR from "swr";

import { api } from "../base/gatekeeper-api";

import type { IServiceOptions } from "@/types/service-options";

export type Organization = {
  id: string;
  name: string;
  description: string;
  createdAt: Date;
  updatedAt: Date | null;
};

type Request = {
  id: string;
};
type Response = Organization;

const fetcher = (url: string, options: IServiceOptions) =>
  api
    .get<Response>(url, {
      headers: {
        Authorization: `Bearer ${options.accessToken}`,
      },
    })
    .then((res) => res.data);

export function useOrganizationByIdSWR(
  { id }: Request,
  options: IServiceOptions
) {
  return useSWR(
    id ? `/v1/organizations/${id}/` : null,
    (url) => fetcher(url, options),
    {
      revalidateOnFocus: false,
    }
  );
}
