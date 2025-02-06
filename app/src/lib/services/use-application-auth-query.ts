import { useQuery } from "@sveltestack/svelte-query";

export interface IApplicationAuth {
  id: string;
  name: string;
  multiFactorAuthEnabled: boolean;
  canSelfRegister: boolean;
  oauthProviders: {
      id: string;
      name: string;
      clientId: string;
      clientSecret: string;
  }[]
}

type Request = {
  applicationId: string;
}

type Response = IApplicationAuth

const mock: IApplicationAuth[] = [
    {
        id: "890fb73c-53d7-45ef-aa51-a4eb936b2c42",
        name: "GateKeeper",
        multiFactorAuthEnabled: false,
        canSelfRegister: true,
        oauthProviders: [
          {
            id: "oauth-201",
            name: "GitHub",
            clientId: "github-client-id",
            clientSecret: "github-client-secret",
          },
          {
            id: "oauth-202",
            name: "Microsoft",
            clientId: "microsoft-client-id",
            clientSecret: "microsoft-client-secret",
          },
          {
            id: "oauth-203",
            name: "LinkedIn",
            clientId: "linkedin-client-id",
            clientSecret: "linkedin-client-secret",
          },
        ],
      },
      {
        id: "4c85936b-be93-4f6f-942e-47f1eae93dc6",
        name: "ProxyMity",
        multiFactorAuthEnabled: false,
        canSelfRegister: false,
        oauthProviders: [
          {
            id: "oauth-201",
            name: "GitHub",
            clientId: "github-client-id",
            clientSecret: "github-client-secret",
          },
        ],
      },
]

export function getApplicationByIdService({ applicationId }: Request): Promise<Response | null> {
  // return fetch('http://localhost:3000/v1/applications').then((response) =>
  //   response.json()
  // );

  console.log(applicationId);

  const application = mock.find((app) => app.id === applicationId);

  if (!application) {
    return Promise.reject(null);
  }

  return Promise.resolve(application);
}

export function useApplicationAuthQuery(request: Request) {
  return useQuery(['get-application-auth', request.applicationId], () => getApplicationByIdService(request))
}