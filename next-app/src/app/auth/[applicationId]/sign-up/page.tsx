import { Fragment } from "react";
import { AuthForm } from "./(components)/AuthForm";

export default function SignInPage() {
  return (
    <Fragment>
      <div className="flex flex-col space-y-2 text-center">
        <h1 className="text-2xl font-semibold tracking-tight">
          Create an account
        </h1>

        <p className="text-muted-foreground text-sm">
          Enter your data below to create your account
        </p>
      </div>

      <AuthForm />
    </Fragment>
  );
}
