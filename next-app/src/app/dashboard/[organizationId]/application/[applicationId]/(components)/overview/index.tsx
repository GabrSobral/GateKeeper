import { Badge } from "@/components/ui/badge";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";

import { formatDate } from "@/lib/utils";
import { IApplication } from "@/services/dashboard/get-application-by-id";

import { SecretsSection } from "./secrets-section";

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
            {application?.mfaAuthAppEnabled ? (
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

      <SecretsSection application={application} />
    </section>
  );
}
