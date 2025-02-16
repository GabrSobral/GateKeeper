"use client";

import { z } from "zod";
import { useForm } from "react-hook-form";
import { Fragment, useState } from "react";

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

import { formSchema } from "./auth-schema";
import { zodResolver } from "@hookform/resolvers/zod";

export function AuthForm() {
  const [isSent, setIsSent] = useState(false);

  // const applicationId = useParams().applicationId;

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      email: "",
    },
  });

  const isLoading = false;

  function onSubmit(values: z.infer<typeof formSchema>) {
    // Do something with the form values.
    // ✅ This will be type-safe and validated.
    console.log(values);
    setIsSent(true);
  }

  if (isSent) {
    return (
      <Fragment>
        <div className="grid gap-6">
          <p className="text-center text-md bg-green-100 p-4 rounded-lg">
            Check your email for a link to reset your password. If it
            doesn&apos;t appear within a few minutes, check your spam folder.
          </p>
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

          <Button type="submit" disabled={isLoading} className="w-full">
            Send Mail
          </Button>
        </form>
      </Form>
    </div>
  );
}
