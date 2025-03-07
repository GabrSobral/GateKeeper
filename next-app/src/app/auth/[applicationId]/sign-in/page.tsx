import { Metadata } from "next";

import { AuthForm } from "./(components)/auth-form";
import { Background } from "../(components)/background";
import { ErrorAlert } from "@/components/error-alert";

import { getApplicationAuthDataService } from "@/services/auth/get-application-auth-data";

type Props = {
  params: Promise<{ applicationId: string }>;
};

export const metadata: Metadata = {
  title: "Sign In - GateKeeper",
  description: "Sign in to your account",
};

export default async function SignInPage(props: Props) {
  const { applicationId } = await props.params;

  const [application, err] = await getApplicationAuthDataService({
    applicationId,
  });

  return (
    <Background application={application} page="sign-in">
      <div className="flex flex-col space-y-2 text-center">
        <h1 className="text-2xl font-semibold tracking-tight">
          Sign at your account
        </h1>

        <p className="text-muted-foreground text-sm">
          Enter your data below to sign in
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
