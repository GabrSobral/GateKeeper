"use client";

import { z } from "zod";
import Link from "next/link";
import { useForm } from "react-hook-form";
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

import { formSchema } from "./auth-schema";
import { zodResolver } from "@hookform/resolvers/zod";

export function AuthForm() {
  const applicationId = useParams().applicationId;

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      email: "",
      password: "",
    },
  });

  const isLoading = false;

  function onSubmit(values: z.infer<typeof formSchema>) {
    // Do something with the form values.
    // ✅ This will be type-safe and validated.
    console.log(values);
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

          <div className="flex items-center justify-between gap-2">
            <Link
              href={`/auth/${applicationId}/forgot-password`}
              className="text-md mb-2 text-center hover:underline"
            >
              Forgot password
            </Link>

            <Link
              href={`/auth/${applicationId}/sign-up`}
              className="font-semibold text-md text-center hover:underline"
            >
              Create an account
            </Link>
          </div>

          <Button type="submit" disabled={isLoading} className="w-full">
            Sign In with Email
          </Button>
        </form>
      </Form>

      <div className="relative">
        <div className="absolute inset-0 flex items-center">
          <span className="w-full border-t"></span>
        </div>

        <div className="relative flex justify-center text-xs uppercase">
          <span className="bg-background text-muted-foreground px-2">
            {" "}
            Or continue with{" "}
          </span>
        </div>
      </div>

      <div className="flex flex-col gap-1">
        <Button variant="outline" type="button" disabled={isLoading}>
          GitHub
        </Button>
      </div>
    </div>
  );
}
