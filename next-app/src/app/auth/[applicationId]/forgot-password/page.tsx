import Link from "next/link";
import { redirect } from "next/navigation";
import { ArrowLeftIcon } from "lucide-react";

import { AuthForm } from "./(components)/auth-form";
import { Background } from "../(components)/background";
import { ErrorAlert } from "@/components/error-alert";

import { getApplicationAuthDataService } from "@/services/auth/get-application-auth-data";

type Props = {
  params: Promise<{ applicationId: string }>;
  searchParams: Promise<{
    trying_to_sign_in: string;
    email: string;
    redirect_uri: string;
    response_type: string;
    scope: string;
    code_challenge_method: string;
    code_challenge: string;
    state: string;
  }>;
};

export default async function ForgotPasswordPage({
  params,
  searchParams,
}: Props) {
  const { applicationId } = await params;
  const searchParamsData = await searchParams;

  const [application, err] = await getApplicationAuthDataService({
    applicationId,
  });

  const urlParams = new URLSearchParams(searchParamsData);

  if (!application?.canSelfForgotPass) {
    return redirect(`/auth/${applicationId}/sign-in`);
  }

  return (
    <Background application={application} page="forgot-password">
      <Link
        href={`/auth/${applicationId}/sign-in?${urlParams.toString()}`}
        className="text-primary flex items-center gap-3 hover:underline absolute top-4 left-4"
      >
        <div className="flex items-center justify-center rounded-md bg-secondary text-primary p-2 hover:brightness-90 transition-all">
          <ArrowLeftIcon />
        </div>
        Back to sign in
      </Link>

      <div className="flex flex-col space-y-2 text-center">
        <h1 className="text-2xl font-semibold tracking-tight">
          Forgot your password?
        </h1>
        <p className="text-muted-foreground text-sm">
          Enter your e-mail below to continue
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
