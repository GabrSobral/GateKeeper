import { AuthForm } from "./(components)/auth-form";
import { Background } from "../(components)/background";
import { ErrorAlert } from "@/components/error-alert";

import { getApplicationAuthDataService } from "@/services/auth/get-application-auth-data";

type Props = {
  params: Promise<{ applicationId: string }>;
};

export default async function UpdatePasswordPage({ params }: Props) {
  const { applicationId } = await params;

  const [application, err] = await getApplicationAuthDataService({
    applicationId,
  });

  return (
    <Background
      application={application}
      page="change-password"
      termsAndConditionsEnabled={false}
    >
      <div className="flex flex-col space-y-2 text-center">
        <h1 className="text-2xl font-semibold tracking-tight">
          Change your password
        </h1>

        <p className="text-muted-foreground text-sm">
          Enter your new password below to continue
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
