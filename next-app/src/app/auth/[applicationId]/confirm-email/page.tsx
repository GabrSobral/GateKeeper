import Link from "next/link";
import { Fragment } from "react";
import { ArrowLeftIcon } from "lucide-react";

import { AuthForm } from "./(components)/AuthForm";

type Props = {
  params: Promise<{ applicationId: string }>;
};

export default async function SignInPage({ params }: Props) {
  const { applicationId } = await params;

  return (
    <Fragment>
      <Link
        href={`/auth/${applicationId}/sign-in`}
        className="text-primary flex items-center gap-3 hover:underline absolute top-4 left-4"
      >
        <div className="flex items-center justify-center rounded-md bg-secondary text-primary p-2 hover:brightness-90 transition-all">
          <ArrowLeftIcon />
        </div>
        Back to sign in
      </Link>

      <div className="flex flex-col space-y-2 text-center">
        <h1 className="text-2xl font-semibold tracking-tight">
          Confirm E-mail
        </h1>
        <p className="text-muted-foreground text-sm">
          We sent a confirmation code to your e-mail. Enter the code below to
          confirm your e-mail
        </p>
      </div>

      <AuthForm />
    </Fragment>
  );
}
