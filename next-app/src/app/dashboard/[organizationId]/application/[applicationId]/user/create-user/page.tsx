import Link from "next/link";
import { Metadata } from "next";
import { ChevronLeft } from "lucide-react";

import { Breadcrumbs } from "@/components/bread-crumbs";
import { UserDetailForm } from "./(components)/user-detail-form";

type Props = {
  params: Promise<{
    organizationId: string;
    applicationId: string;
    userId: string;
  }>;
};

export const metadata: Metadata = {
  title: "Add User - GateKeeper",
};

export default async function CreateUserPage({ params }: Props) {
  const { organizationId, applicationId, userId } = await params;

  return (
    <>
      <Breadcrumbs
        items={[
          { name: "Dashboard", path: "/dashboard" },
          { name: "Applications", path: "/dashboard/application" },
          {
            name: applicationId,
            path: `/dashboard/${organizationId}/application/${applicationId}?tab=users`,
          },
          {
            name: userId,
            path: `/dashboard/${organizationId}/application/${applicationId}/user/${userId}`,
          },
        ]}
      />

      <main className="flex flex-col p-4">
        <Link
          href={`/dashboard/${organizationId}/application/${applicationId}?tab=users`}
          className="text-md mb-4 flex items-center gap-2 text-gray-600 dark:text-gray-300 hover:text-gray-800  hover:underline"
        >
          <ChevronLeft size={24} />
          Go back to application detail
        </Link>

        <UserDetailForm />
      </main>
    </>
  );
}
