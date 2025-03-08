import { Metadata } from "next";
import { redirect } from "next/navigation";

import { AuthForm } from "./(components)/auth-form";
import { Background } from "../(components)/background";
import { ErrorAlert } from "@/components/error-alert";

import { getApplicationAuthDataService } from "@/services/auth/get-application-auth-data";

type Props = {
  params: Promise<{ applicationId: string }>;
};

export const metadata: Metadata = {
  title: "Sign Up - GateKeeper",
  description: "Sign up for an account",
};

export default async function SignInPage({ params }: Props) {
  const { applicationId } = await params;

  const [application, err] = await getApplicationAuthDataService({
    applicationId,
  });

  if (!application?.canSelfSignUp) {
    return redirect(`/auth/${applicationId}/sign-in`);
  }

  return (
    <Background application={application} page="sign-up">
      <div className="flex flex-col space-y-2 text-center">
        <h1 className="text-2xl font-semibold tracking-tight">
          Create an account
        </h1>

        <p className="text-muted-foreground text-sm">
          Enter your data below to create your account
        </p>
      </div>

      {err ? (
        <ErrorAlert message={err.message} title="An error occurred..." />
      ) : (
        <AuthForm application={application} />
      )}
    </Background>
  );
}
