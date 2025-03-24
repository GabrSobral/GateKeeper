"use client";

import { z } from "zod";
import { toast } from "sonner";
import { useState } from "react";
import { useForm } from "react-hook-form";
import { useParams, useSearchParams } from "next/navigation";

import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { LoadingSpinner } from "@/components/ui/loading-spinner";

import { formSchema } from "./auth-schema";
import { zodResolver } from "@hookform/resolvers/zod";
import { changePasswordApi } from "@/services/auth/change-password";

import { ErrorAlert } from "@/components/error-alert";
import { authorizeApi } from "@/services/auth/authorize";

export function AuthForm() {
  const applicationId = useParams().applicationId as string;
  const searchParams = useSearchParams();

  const redirectUri = searchParams.get("redirect_uri") || "/";
  const codeChallengeMethod = searchParams.get("code_challenge_method") || "";
  const responseType = searchParams.get("response_type") || "";
  const scope = searchParams.get("scope") || "";
  const state = searchParams.get("state") || "";
  const email = searchParams.get("email") || "";
  const codeChallenge = searchParams.get("code_challenge") || "";
  const sessionCode = searchParams.get("session_code") || "";

  const changePasswordCode = searchParams.get("change_password_code") || "";
  const userId = searchParams.get("user_id") || "";

  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      password: "",
      confirmPassword: "",
    },
  });

  async function onSubmit(values: z.infer<typeof formSchema>) {
    if (!applicationId) {
      console.error("Application ID is required");
      toast.error("Application ID is required");
      return;
    }

    if (!changePasswordCode) {
      console.error("Change password code is required");
      toast.error("Change password code is required");
      return;
    }

    if (!userId) {
      console.error("User ID is required");
      toast.error("User ID is required");
      return;
    }

    setIsLoading(true);

    const [err] = await changePasswordApi({
      applicationId,
      newPassword: values.password,
      changePasswordCode,
      userId,
    });

    if (err) {
      console.error(err);
      toast.error(err.response?.data?.message || err.message);
      setIsLoading(false);
      return;
    }

    const [authorizeData, authorizeErr] = await authorizeApi({
      email: email.trim(),
      sessionCode: sessionCode,
      applicationId,
      redirectUri,
      responseType,
      scope,
      codeChallengeMethod,
      codeChallenge,
      state,
    });

    if (authorizeErr) {
      console.error(authorizeErr);
      setError(authorizeErr?.response?.data.message || "An error occurred");
      setIsLoading(false);
      setTimeout(() => setError(null), 6000);
      return;
    }

    if (!authorizeData) {
      setError("An error occurred");
      setIsLoading(false);
      setTimeout(() => setError(null), 6000);
      return;
    }

    setIsLoading(false);

    toast.success("Password was changed successfully");

    window.location.href = `${redirectUri}?code=${authorizeData.authorizationCode}&state=${state}&redirect_uri=${redirectUri}&client_id=${applicationId}`;
  }

  return (
    <div className="grid gap-4">
      {error && <ErrorAlert message={error} title="An error occurred..." />}

      <Form {...form}>
        <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-3">
          <FormField
            control={form.control}
            name="password"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Password</FormLabel>
                <FormControl>
                  <Input
                    placeholder="********"
                    autoComplete="new-password"
                    type="password"
                    {...field}
                  />
                </FormControl>

                <FormDescription></FormDescription>
                <FormMessage />
              </FormItem>
            )}
          />

          <FormField
            control={form.control}
            name="confirmPassword"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Confirm Password</FormLabel>
                <FormControl>
                  <Input placeholder="********" type="password" {...field} />
                </FormControl>

                <FormDescription></FormDescription>
                <FormMessage />
              </FormItem>
            )}
          />

          <Button
            type="submit"
            disabled={isLoading}
            className="w-full relative"
          >
            {isLoading && <LoadingSpinner className="absolute left-4" />}
            Save Password
          </Button>
        </form>
      </Form>
    </div>
  );
}
