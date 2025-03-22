import { Breadcrumbs } from "@/components/bread-crumbs";
import { Metadata } from "next";

type Props = {
  params: Promise<{
    organizationId: string;
  }>;
};

export const metadata: Metadata = {
  title: "Organizations - GateKeeper",
};

export default async function OrganizationPage({ params }: Props) {
  const { organizationId } = await params;
  return (
    <>
      <Breadcrumbs
        items={[
          { name: "Dashboard", path: "/dashboard" },
          {
            name: organizationId,
            path: `/dashboard/${organizationId}`,
          },
          {
            name: "Applications",
            path: `/dashboard/${organizationId}/application`,
          },
        ]}
      />

      <main className="flex flex-col p-4">
        <h2 className="text-3xl font-bold tracking-tight">{organizationId}</h2>

        <span className="mt-3 text-sm tracking-tight text-gray-600 dark:text-gray-300">
          Applications are the projects you have created. You can add, edit, and
          delete them here.
        </span>
      </main>
    </>
  );
}
