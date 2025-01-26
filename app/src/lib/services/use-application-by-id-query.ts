import type { IServiceOptions } from "$lib/types/service-options";
import { useQuery } from "@sveltestack/svelte-query";

export interface IApplication {
  id: string;
  name: string;
  description: string;
  badges: string[];
  createdAt: Date;
  updatedAt: Date;
  deactivatedAt?: Date;
  multiFactorAuthEnabled: boolean;
  passwordHashingSecret: string;
  secrets: {
      id: string;
      name: string;
      value: string;
      expirationDate: Date;
  }[];
  users: {
      totalCount: number;
      data: {
          id: string;
          displayName: string;
          email: string;
          roles: {
              id: string;
              name: string;
          }[];
          deactivatedAt?: Date;
      }[]
  },
  roles: {
      totalCount: number;
      data: {
          id: string;
          name: string;
          description: string;
      }[]
  },
  oauthProviders: {
      id: string;
      name: string;
      description: string;
      clientId: string;
      clientSecret: string;
      isEnabled: boolean;
  }[]
}

type Request = {
  applicationId: string;
}

type Response = IApplication

const mock: IApplication[] = [
  {
    id: "890fb73c-53d7-45ef-aa51-a4eb936b2c42",
    name: "GateKeeper",
    description: "GateKeeper is a password manager that allows you to store your passwords securely.",
    badges: ["User-Friendly", "Scalable", "Cloud-Enabled"],
    createdAt: new Date("2024-05-15T09:00:00Z"),
    updatedAt: new Date("2025-01-20T15:30:00Z"),
    deactivatedAt: undefined,
    multiFactorAuthEnabled: false,
    passwordHashingSecret: "hashing-key-2024",
    secrets: [
      {
        id: "3",
        name: "Database Connection String",
        value: "db-connection-string-2024",
        expirationDate: new Date("2025-09-01T00:00:00Z"),
      },
      {
        id: "4",
        name: "JWT Secret",
        value: "jwt-secret-key",
        expirationDate: new Date("2027-01-01T23:59:59Z"),
      },
    ],
    users: {
      totalCount: 3,
      data: [
        {
          id: "user-101",
          displayName: "Alice Johnson",
          email: "alice.johnson@apptracker.com",
          roles: [
            { id: "role-101", name: "Admin" },
            { id: "role-102", name: "Support" },
          ],
          deactivatedAt: undefined,
        },
        {
          id: "user-102",
          displayName: "Bob Carter",
          email: "bob.carter@apptracker.com",
          roles: [{ id: "role-103", name: "Viewer" }],
          deactivatedAt: undefined,
        },
        {
          id: "user-103",
          displayName: "Eve Summers",
          email: "eve.summers@apptracker.com",
          roles: [{ id: "role-104", name: "Moderator" }],
          deactivatedAt: new Date("2024-12-01T08:00:00Z"),
        },
      ],
    },
    roles: {
      totalCount: 4,
      data: [
        {
          id: "role-101",
          name: "Admin",
          description: "Can manage all resources and settings",
        },
        {
          id: "role-102",
          name: "Support",
          description: "Handles user support tickets",
        },
        {
          id: "role-103",
          name: "Viewer",
          description: "Can only view data",
        },
        {
          id: "role-104",
          name: "Moderator",
          description: "Monitors user activity and manages reports",
        },
      ],
    },
    oauthProviders: [
      {
        id: "oauth-201",
        name: "GitHub",
        description: "GitHub OAuth provider",
        clientId: "github-client-id",
        clientSecret: "github-client-secret",
        isEnabled: true,
      },
      {
        id: "oauth-202",
        name: "Microsoft",
        description: "Microsoft OAuth provider",
        clientId: "microsoft-client-id",
        clientSecret: "microsoft-client-secret",
        isEnabled: true,
      },
      {
        id: "oauth-203",
        name: "LinkedIn",
        description: "LinkedIn OAuth provider",
        clientId: "linkedin-client-id",
        clientSecret: "linkedin-client-secret",
        isEnabled: false,
      },
    ],
  },
  {
    id: "4c85936b-be93-4f6f-942e-47f1eae93dc6",
    name: "ProxyMity",
    description: "ProxyMity is a proxy service that allows you to access websites that are blocked by your network.",
    badges: ["Secure", "Fast", "Reliable"],
    createdAt: new Date("2023-01-01T10:00:00Z"),
    updatedAt: new Date("2025-01-01T10:00:00Z"),
    deactivatedAt: new Date("2025-12-31T23:59:59Z"),
    multiFactorAuthEnabled: true,
    passwordHashingSecret: "super-secret-key",
    secrets: [
      {
        id: "1",
        name: "API Key",
        value: "abc123xyz",
        expirationDate: new Date("2025-12-31T23:59:59Z"),
      },
      {
        id: "2",
        name: "Webhook Secret",
        value: "webhook456key",
        expirationDate: new Date("2026-06-30T23:59:59Z"),
      },
    ],
    users: {
      totalCount: 2,
      data: [
        {
          id: "user-1",
          displayName: "John Doe",
          email: "john.doe@example.com",
          roles: [
            { id: "role-1", name: "Admin" },
            { id: "role-2", name: "Editor" },
          ],
          deactivatedAt: undefined,
        },
        {
          id: "user-2",
          displayName: "Jane Smith",
          email: "jane.smith@example.com",
          roles: [{ id: "role-3", name: "Viewer" }],
          deactivatedAt: new Date("2024-01-01T10:00:00Z"),
        },
      ],
    },
    roles: {
      totalCount: 3,
      data: [
        {
          id: "role-1",
          name: "Admin",
          description: "Full access to all resources",
        },
        {
          id: "role-2",
          name: "Editor",
          description: "Can edit content",
        },
        {
          id: "role-3",
          name: "Viewer",
          description: "Can view content only",
        },
      ],
    },
    oauthProviders: [
      {
        id: "oauth-1",
        name: "Google",
        description: "Google OAuth provider",
        clientId: "google-client-id",
        clientSecret: "google-client-secret",
        isEnabled: true,
      },
      {
        id: "oauth-2",
        name: "Facebook",
        description: "Facebook OAuth provider",
        clientId: "",
        clientSecret: "",
        isEnabled: false,
      },
    ],
  }
]

function getApplicationByIdService({ applicationId }: Request, { accessToken }: IServiceOptions): Promise<Response | null> {
  // return fetch('http://localhost:3000/v1/applications').then((response) =>
  //   response.json()
  // );

  console.log(accessToken, applicationId);

  const application = mock.find((app) => app.id === applicationId);

  if (!application) {
    return Promise.reject(null);
  }

  return Promise.resolve(application);
}

export function useApplicationByIdQuery(request: Request, options: IServiceOptions) {
  return useQuery(['get-application-by-id', request.applicationId], () => getApplicationByIdService(request, options))
}