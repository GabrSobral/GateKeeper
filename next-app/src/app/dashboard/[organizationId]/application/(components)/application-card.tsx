"use client";

import Link from "next/link";
import { useParams } from "next/navigation";

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
import { APIError } from "@/types/service-options";
import { Alert, AlertDescription, AlertTitle } from "@/components/ui/alert";
import { AlertCircle } from "lucide-react";

export function ApplicationCard() {
  const organizationId = useParams().organizationId as string;
  const { data, isLoading, error } = useApplicationsSWR(
    { organizationId },
    { accessToken: "" }
  );

  const err = error as APIError;

  if (isLoading) {
    <>
      <Skeleton className="h-[133px] max-w-[400px] flex-1" />
      <Skeleton className="h-[133px] max-w-[400px] flex-1" />
    </>;
  }

  if (err) {
    return (
      <Alert variant="destructive" className="bg-red-500/10">
        <AlertCircle className="h-4 w-4" />
        <AlertTitle>An error occurred...</AlertTitle>
        <AlertDescription>{err.response?.data.message}</AlertDescription>
      </Alert>
    );
  }

  if (data?.length === 0) {
    return (
      <Alert variant="default" className="bg-yellow-500/10">
        <AlertCircle className="h-4 w-4" />
        <AlertTitle>No applications found...</AlertTitle>
        <AlertDescription>
          You don&apos;t have any application on this organization. Create one
          to get started.
        </AlertDescription>
      </Alert>
    );
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
