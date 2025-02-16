import { Badge } from "@/components/ui/badge";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import {
  Tooltip,
  TooltipContent,
  TooltipTrigger,
} from "@/components/ui/tooltip";

// import { formatDate } from "@/lib/utils";
import { NewSecretDialog } from "./new-secret-dialog";
import { DeleteSecretDialog } from "./delete-secret-dialog";

import { IApplication } from "@/services/dashboard/get-application-by-id";
import { formatDate } from "@/lib/utils";

type Props = {
  application: IApplication | null;
};

export function Overview({ application }: Props) {
  return (
    <section className="flex w-full flex-col gap-y-4">
      <Card className="w-full transition-all">
        <CardHeader>
          <CardTitle>Info</CardTitle>
          <CardDescription>Information about the application.</CardDescription>
        </CardHeader>

        <CardContent className="flex flex-wrap gap-x-8 gap-y-4">
          <div className="flex flex-col">
            <span className="text-md font-semibold">Application ID</span>
            <span className="text-sm">{application?.id}</span>
          </div>

          <div className="flex flex-col">
            <span className="text-md font-semibold">Status</span>
            {!application?.deactivatedAt ? (
              <Badge variant="default">Active</Badge>
            ) : (
              <Badge variant="destructive" className="flex gap-1">
                Deactivated at{" "}
                <span className="text-white">
                  {application?.deactivatedAt &&
                    formatDate(application.deactivatedAt)}
                </span>
              </Badge>
            )}
          </div>

          <div className="flex flex-col">
            <span className="text-md font-semibold">Multi Factor Auth</span>
            {application?.multiFactorAuthEnabled ? (
              <Badge variant="outline" className="w-fit">
                Yes
              </Badge>
            ) : (
              <Badge variant="outline" className="w-fit">
                No
              </Badge>
            )}
          </div>
        </CardContent>
      </Card>

      <Card className="w-full transition-all">
        <CardHeader>
          <CardTitle className="flex flex-wrap justify-between gap-4">
            Secrets
            <NewSecretDialog />
          </CardTitle>

          <CardDescription>
            Secrets are used to authenticate your application with the server.
            Keep them safe.
          </CardDescription>
        </CardHeader>

        <CardContent className="flex flex-col gap-y-4">
          {application?.secrets.map((secret) => (
            <div className="flex items-center gap-4" key={secret.id}>
              <div className="space-y-1">
                <p className="text-sm font-medium leading-none">
                  {secret.name}
                </p>
                <p className="text-muted-foreground">{secret.value}</p>
              </div>

              <div className="ml-auto text-sm">
                Expiração:{" "}
                <span className="text-md font-medium">
                  {secret?.expirationDate
                    ? formatDate(new Date(secret.expirationDate))
                    : "Lifetime"}
                </span>
              </div>

              <Tooltip>
                <TooltipTrigger asChild>
                  <DeleteSecretDialog secret={secret} />
                </TooltipTrigger>

                <TooltipContent>
                  <p>Delete secret</p>
                </TooltipContent>
              </Tooltip>
            </div>
          ))}
        </CardContent>
      </Card>
    </section>
  );
}
