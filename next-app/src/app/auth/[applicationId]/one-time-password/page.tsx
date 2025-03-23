import Link from "next/link";
import { ArrowLeftIcon } from "lucide-react";

import { AuthForm } from "./(components)/auth-form";
import { Background } from "../(components)/background";
import { ErrorAlert } from "@/components/error-alert";

import { getApplicationAuthDataService } from "@/services/auth/get-application-auth-data";

type Props = {
  params: Promise<{ applicationId: string }>;
};

export default async function OneTimePasswordPage({ params }: Props) {
  const { applicationId } = await params;

  const [application, err] = await getApplicationAuthDataService({
    applicationId,
  });

  return (
    <Background application={application} page="one-time-password">
      <div className="flex flex-col space-y-2 text-center">
        <h1 className="text-2xl font-semibold tracking-tight">Multi Factor</h1>
        <p className="text-muted-foreground text-sm">
          Enter your one time password below to sign in
        </p>
      </div>

      {err ? (
        <ErrorAlert message={err.message} title="An error occurred..." />
      ) : (
        <AuthForm />
      )}
    </Background>
  );
}
