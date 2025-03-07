import Link from "next/link";
import { ChevronLeft, Pencil } from "lucide-react";

import { Breadcrumbs } from "@/components/bread-crumbs";
import { buttonVariants } from "@/components/ui/button";
import {
  Tooltip,
  TooltipContent,
  TooltipTrigger,
} from "@/components/ui/tooltip";

import { cn } from "@/lib/utils";
import { getApplicationByIdService } from "@/services/dashboard/get-application-by-id";

import { Badge } from "@/components/ui/badge";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";

import { Overview } from "./(components)/overview";
import { Users } from "./(components)/users";
import { Roles } from "./(components)/roles";
import { Providers } from "./(components)/providers";
import { DeleteApplicationDialog } from "./(components)/delete-application-dialog";

type Props = {
  params: Promise<{
    applicationId: string;
    organizationId: string;
  }>;
};

export default async function ApplicationDetailPage({ params }: Props) {
  const { applicationId, organizationId } = await params;

  const [application, err] = await getApplicationByIdService(
    { applicationId, organizationId },
    { accessToken: "" }
  );

  if (err) {
    return <div>Failed to fetch application</div>;
  }

  return (
    <>
      <Breadcrumbs
        items={[
          { name: "Dashboard", path: `/dashboard` },

          { name: organizationId, path: `/dashboard/${organizationId}` },
          {
            name: "Applications",
            path: `/dashboard/${organizationId}/application`,
          },
          {
            name: application?.name || "-",
            path: `/dashboard/${organizationId}/application/${applicationId}`,
          },
        ]}
      />

      <main className="flex flex-col p-4">
        <Link
          href={`/dashboard/${organizationId}/application`}
          className="text-md mb-4 flex items-center gap-2 text-gray-600 dark:text-gray-300 hover:text-gray-800 hover:underline"
        >
          <ChevronLeft size={24} />
          Go back to applications list
        </Link>

        <div className="flex items-center justify-between gap-4">
          <h2 className="text-3xl font-bold tracking-tight">
            {application?.name}
          </h2>

          <div className="flex gap-1">
            <Tooltip>
              <TooltipTrigger asChild>
                <DeleteApplicationDialog application={application} />
              </TooltipTrigger>

              <TooltipContent>
                <p>Delete Application</p>
              </TooltipContent>
            </Tooltip>

            <Tooltip delayDuration={0}>
              <TooltipTrigger
                className={cn(buttonVariants({ variant: "outline" }))}
                asChild
              >
                <Link
                  href={`/dashboard/${organizationId}/application/${applicationId}/edit-application`}
                >
                  <Pencil />
                </Link>
              </TooltipTrigger>

              <TooltipContent>
                <p>Update Application</p>
              </TooltipContent>
            </Tooltip>
          </div>
        </div>

        <span className="mt-3 text-sm tracking-tight text-gray-600 dark:text-gray-300">
          {application?.description}
        </span>

        <div className="mt-4 flex flex-wrap gap-2">
          {application?.badges.map((badge, i) => (
            <Badge variant="outline" key={i}>
              {badge}
            </Badge>
          ))}
        </div>

        <Tabs className="mt-4" defaultValue="overview">
          <TabsList>
            <TabsTrigger value="overview">Overview</TabsTrigger>
            <TabsTrigger value="users">Users</TabsTrigger>
            <TabsTrigger value="roles">Roles</TabsTrigger>
            <TabsTrigger value="providers">Providers</TabsTrigger>
          </TabsList>

          <TabsContent value="overview">
            <Overview application={application} />
          </TabsContent>

          <TabsContent value="users">
            <Users application={application} />
          </TabsContent>

          <TabsContent value="roles">
            <Roles application={application} />
          </TabsContent>

          <TabsContent value="providers">
            <Providers application={application} />
          </TabsContent>
        </Tabs>
      </main>
    </>
  );
}
