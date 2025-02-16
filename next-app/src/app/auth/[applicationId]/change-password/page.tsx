import { Fragment } from "react";
import { AuthForm } from "./(components)/AuthForm";

export default function SignInPage() {
  return (
    <Fragment>
      <div className="flex flex-col space-y-2 text-center">
        <h1 className="text-2xl font-semibold tracking-tight">
          Change your password
        </h1>

        <p className="text-muted-foreground text-sm">
          Enter your new password below to continue
        </p>
      </div>

      <AuthForm />
    </Fragment>
  );
}
