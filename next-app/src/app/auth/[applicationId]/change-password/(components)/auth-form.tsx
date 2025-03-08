"use client";

import { z } from "zod";
import { toast } from "sonner";
import { useState } from "react";
import { useForm } from "react-hook-form";
import { CheckCircle } from "lucide-react";
import { useParams, useRouter, useSearchParams } from "next/navigation";

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
import { Alert, AlertDescription, AlertTitle } from "@/components/ui/alert";

import { formSchema } from "./auth-schema";
import { zodResolver } from "@hookform/resolvers/zod";
import { resetPasswordApi } from "@/services/auth/reset-password";

export function AuthForm() {
  const applicationId = useParams().applicationId as string;
  const searchParams = useSearchParams();

  const router = useRouter();

  const passwordResetToken = searchParams.get("token");
  const passwordResetId = searchParams.get("id");
  const alreadyChanged = searchParams.get("changed");

  const [isLoading, setIsLoading] = useState(false);

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

    if (!passwordResetToken) {
      console.error("Password reset token is required");
      toast.error("Password reset token is required");
      return;
    }

    if (!passwordResetId) {
      console.error("Password reset ID is required");
      toast.error("Password reset ID is required");
      return;
    }

    setIsLoading(true);

    const [err] = await resetPasswordApi({
      applicationId,
      newPassword: values.password,
      passwordResetToken,
      passwordResetId,
    });

    if (err) {
      console.error(err);
      toast.error(err.response?.data?.message || err.message);
      setIsLoading(false);
      return;
    }

    setIsLoading(false);

    toast.success("Password reset successfully");

    router.push(`/auth/${applicationId}/change-password?changed=true`);
  }

  if (alreadyChanged === "true") {
    return (
      <div className="grid gap-4">
        <Alert variant="default" className="bg-green-100 dark:bg-green-700">
          <CheckCircle className="w-5 h-5 text-green-500 dark:text-green-400" />

          <AlertTitle>Success</AlertTitle>
          <AlertDescription>
            Your password has been changed successfully <br />
          </AlertDescription>
        </Alert>

        <span className="mt-1 font-semibold text-center">
          You can close this page and try logging in with your new password!
        </span>
      </div>
    );
  }

  return (
    <div className="grid gap-4">
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
