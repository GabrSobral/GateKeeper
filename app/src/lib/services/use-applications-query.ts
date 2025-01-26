import type { IServiceOptions } from "$lib/types/service-options";
import { useQuery } from "@sveltestack/svelte-query";

type Response = {
    id: string;
    name: string;
    description: string;
    createdAt: Date;
    updatedAt: Date;
    deactivatedAt?: Date;
    badges: string[];
}[]

function getApplicationsService({ accessToken }: IServiceOptions): Promise<Response> {
  // return fetch('http://localhost:3000/v1/applications').then((response) =>
  //   response.json()
  // );

  console.log(accessToken);

  return Promise.resolve([
    {
      id: "4c85936b-be93-4f6f-942e-47f1eae93dc6",
      name: 'ProxyMity',
      description: 'ProxyMity is a proxy service that allows you to access websites that are blocked by your network.',
      createdAt: new Date(),
      updatedAt: new Date(),
      badges: ['C#', 'Javascript', "Chat"]
    },
    {
      id: "890fb73c-53d7-45ef-aa51-a4eb936b2c42",
      name: 'GateKeeper',
      description: 'GateKeeper is a password manager that allows you to store your passwords securely.',
      createdAt: new Date(),
      updatedAt: new Date(),
      badges: ['badge1', 'badge2']
    }
  ])
}

export function useApplicationsQuery(options: IServiceOptions) {
  return useQuery('list-applications', () => getApplicationsService(options))
}