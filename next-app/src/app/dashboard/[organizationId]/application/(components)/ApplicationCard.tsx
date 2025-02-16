"use client";

import { useParams, useRouter } from "next/navigation";

import {
  Card,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";

import { useApplicationsSWR } from "@/services/dashboard/use-applications-swr";
import { Skeleton } from "@/components/ui/skeleton";
import Link from "next/link";

export function ApplicationCard() {
  const router = useRouter();
  const organizationId = useParams().organizationId as string;
  const { data, isLoading } = useApplicationsSWR(
    { organizationId },
    { accessToken: "" }
  );

  if (isLoading) {
    <>
      <Skeleton className="h-[133px] max-w-[400px] flex-1" />
      <Skeleton className="h-[133px] max-w-[400px] flex-1" />
    </>;
  }

  return (
    data?.map((application) => (
      <Link
        key={application.id}
        href={`/dashboard/${organizationId}/application/${application.id}`}
      >
        <Card className="w-[calc(33.333%-8px)] min-w-[400px] transition-all hover:scale-[1.01] hover:cursor-pointer hover:shadow-lg">
          <CardHeader>
            <CardTitle>{application.name}</CardTitle>
            <CardDescription className="line-clamp-4">
              {application.description}
            </CardDescription>
          </CardHeader>

          <CardFooter className="mt-3">
            {application.badges.map((badge, i) => (
              <Badge key={i} variant="outline">
                {badge}
              </Badge>
            ))}
          </CardFooter>
        </Card>
      </Link>
    )) || []
  );
}
