"use client";

import { z } from "zod";
import { toast } from "sonner";
import { useForm } from "react-hook-form";
import { Fragment, useState } from "react";
import { CheckCircle } from "lucide-react";
import { useParams } from "next/navigation";

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
import { forgotPasswordApi } from "@/services/auth/forgot-password";
import { Alert, AlertDescription, AlertTitle } from "@/components/ui/alert";

export function AuthForm() {
  const [isSent, setIsSent] = useState(false);
  const [isLoading, setIsLoading] = useState(false);

  const applicationId = useParams().applicationId as string;

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      email: "",
    },
  });

  async function onSubmit(values: z.infer<typeof formSchema>) {
    if (!applicationId) {
      toast.error("Application ID is required");
      console.error("Application ID is required");
      return;
    }

    setIsLoading(true);

    const [err] = await forgotPasswordApi({
      applicationId: applicationId,
      email: values.email,
    });

    if (err) {
      toast.error(err.response?.data?.message || err.message);
      console.error(err);
      setIsLoading(false);
      return;
    }

    setIsLoading(false);
    setIsSent(true);
  }

  if (isSent) {
    return (
      <Fragment>
        <div className="grid gap-6">
          <Alert variant="default" className="bg-green-100 dark:bg-green-700">
            <CheckCircle className="w-5 h-5 text-green-500 dark:text-green-400" />
            <AlertTitle>E-mail sent</AlertTitle>

            <AlertDescription>
              Check your email for a link to reset your password. If it
              doesn&apos;t appear within a few minutes, check your spam folder.
            </AlertDescription>
          </Alert>
        </div>

        <Button onClick={() => setIsSent(false)} className="mt-4">
          Go Back
        </Button>
      </Fragment>
    );
  }

  return (
    <div className="grid gap-4">
      <Form {...form}>
        <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-3">
          <FormField
            control={form.control}
            name="email"
            render={({ field }) => (
              <FormItem>
                <FormLabel>E-mail</FormLabel>
                <FormControl>
                  <Input
                    placeholder="example@email.com"
                    autoComplete="email"
                    type="email"
                    {...field}
                  />
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
            Send Mail
          </Button>
        </form>
      </Form>
    </div>
  );
}
