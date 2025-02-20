"use client";

import { z } from "zod";
import { useState } from "react";
import { useForm, UseFormReturn } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { useParams, useRouter } from "next/navigation";

import { Button } from "@/components/ui/button";
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
import { Separator } from "@/components/ui/separator";

import { formSchema } from "../schema";
import { MultiFactorAuthForm } from "./multi-factor-auth-form";
import { ApplicationRolesSection } from "./application-roles-section";

export type FormType = UseFormReturn<z.infer<typeof formSchema>>;

const roles = [
  { id: "id1", name: "User", description: "User role description" },
  { id: "id2", name: "Admin", description: "Admin role description" },
];

export function UserDetailForm() {
  const [isLoading, setIsLoading] = useState(false);

  const router = useRouter();
  const { applicationId, organizationId } = useParams() as {
    organizationId: string;
    applicationId: string;
  };

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      displayName: "",
      email: "",
      firstName: "",
      lastName: "",
      hasMfaAuthAppEnabled: false,
      hasMfaEmailEnabled: false,
      roles: [],
    },
  });

  function onSubmit() {
    setIsLoading(true);

    // Logic here
    setIsLoading(false);

    console.log(organizationId, applicationId, router);

    // const userId = "123";

    console.log(form.getValues());

    // router.push(
    //   `/dashboard/${organizationId}/application/${applicationId}/user/${userId}`
    // );
  }

  return (
    <Form {...form}>
      <div className="flex items-center justify-between gap-4">
        <FormField
          control={form.control}
          name="displayName"
          render={({ field }) => (
            <FormItem className="w-full">
              <FormLabel className="flex gap-1 sr-only">
                Display Name
                <span className="text-red-500">*</span>
              </FormLabel>

              <div className="w-full flex gap-2">
                <FormControl>
                  <Input
                    placeholder="Type the user display name"
                    autoComplete="name"
                    type="text"
                    style={{
                      fontSize: "1.75rem",
                      fontWeight: 700,
                      height: "3.5rem",
                      lineHeight: "3.5rem",
                    }}
                    className="max-w-[700px]"
                    {...field}
                  />
                </FormControl>
              </div>

              <FormDescription>
                The name that will be displayed to the user.
              </FormDescription>
              <FormMessage></FormMessage>
            </FormItem>
          )}
        />
      </div>

      <form
        onSubmit={form.handleSubmit(onSubmit)}
        className="mt-4 max-w-[700px]"
      >
        <div className="grid gap-4">
          <fieldset className="flex gap-2">
            <FormField
              control={form.control}
              name="firstName"
              render={({ field }) => (
                <FormItem className="w-full">
                  <FormLabel className="flex gap-1">
                    First Name
                    <span className="text-red-500">*</span>
                  </FormLabel>

                  <div className="w-full flex gap-2">
                    <FormControl>
                      <Input
                        className="w-full"
                        placeholder="Type the user first name"
                        autoComplete="given-name"
                        type="text"
                        {...field}
                      />
                    </FormControl>
                  </div>

                  <FormMessage></FormMessage>
                </FormItem>
              )}
            />

            <FormField
              control={form.control}
              name="lastName"
              render={({ field }) => (
                <FormItem className="w-full">
                  <FormLabel className="flex gap-1">
                    Last Name
                    <span className="text-red-500">*</span>
                  </FormLabel>

                  <div className="w-full flex gap-2">
                    <FormControl>
                      <Input
                        className="w-full"
                        placeholder="Type the user last name"
                        autoComplete="family-name"
                        type="text"
                        {...field}
                      />
                    </FormControl>
                  </div>

                  <FormMessage></FormMessage>
                </FormItem>
              )}
            />
          </fieldset>

          <Separator className="my-2" />

          <FormField
            control={form.control}
            name="email"
            render={({ field }) => (
              <FormItem>
                <FormLabel className="flex gap-1">
                  E-mail
                  <span className="text-red-500">*</span>
                </FormLabel>

                <div className="w-full flex gap-2">
                  <FormControl>
                    <Input
                      className="w-full"
                      placeholder="Type the user e-mail"
                      autoComplete="email"
                      type="text"
                      {...field}
                    />
                  </FormControl>
                </div>
                <FormMessage></FormMessage>
              </FormItem>
            )}
          />

          <FormField
            control={form.control}
            name="temporaryPassword"
            render={({ field }) => (
              <FormItem className="w-full">
                <FormLabel className="flex gap-1">
                  Temporary Password
                  <span className="text-red-500">*</span>
                </FormLabel>

                <div className="w-full flex gap-2">
                  <FormControl>
                    <Input
                      className="w-full"
                      placeholder="Type the user temporary password"
                      autoComplete="password"
                      type="text"
                      {...field}
                    />
                  </FormControl>
                </div>

                <FormDescription>
                  The temporary password that will be sent to the user at the
                  first access.
                </FormDescription>
                <FormMessage></FormMessage>
              </FormItem>
            )}
          />

          <Separator className="my-2" />

          <MultiFactorAuthForm form={form} />

          <Separator className="my-2" />

          <ApplicationRolesSection form={form} roles={roles} />
        </div>

        <Button type="submit" className="float-right mt-4" disabled={isLoading}>
          Create User
        </Button>
      </form>
    </Form>
  );
}
