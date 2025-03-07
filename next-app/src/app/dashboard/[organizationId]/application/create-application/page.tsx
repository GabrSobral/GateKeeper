import Link from "next/link";
import { ChevronLeft } from "lucide-react";

import { Breadcrumbs } from "@/components/bread-crumbs";
import { CreateApplicationForm } from "./(components)/create-application-form";

type Props = {
  params: Promise<{
    organizationId: string;
  }>;
};
export default async function CreateApplicationPage({ params }: Props) {
  const { organizationId } = await params;

  return (
    <>
      <Breadcrumbs
        items={[
          { name: "Dashboard", path: `/dashboard` },
          {
            name: organizationId,
            path: `/dashboard/${organizationId}`,
          },
          {
            name: "Applications",
            path: `/dashboard/${organizationId}/application`,
          },
          {
            name: "Create Application",
            path: `/dashboard/${organizationId}/application/create-application`,
          },
        ]}
      />

      <main className="flex flex-col p-4">
        <Link
          href={`/dashboard/${organizationId}/application`}
          className="text-md mb-4 flex items-center gap-2 text-gray-600 dark:text-gray-300 hover:dark:text-gray-500 hover:text-gray-800 hover:underline"
        >
          <ChevronLeft size={24} />
          Go back to applications list
        </Link>

        <h2 className="text-3xl font-bold tracking-tight">
          Create Application
        </h2>

        <span className="mt-3 text-sm tracking-tight text-gray-600 dark:text-gray-300">
          Here you can create a new application. Fill in the form below and
          click the &quot;Create Application&quot; button to create a new
          application.
        </span>

        <CreateApplicationForm />
      </main>
    </>
  );
}
