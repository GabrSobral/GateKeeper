"use client";

import { z } from "zod";
import { toast } from "sonner";
import { useState } from "react";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";

import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { Textarea } from "@/components/ui/textarea";
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";

import { formSchema } from "./schema";

import { Organization } from "@/services/dashboard/use-organizations-swr";
import { editOrganizationApi } from "@/services/settings/edit-organization";

import { OrganizationsPages } from "..";

type Props = {
  setPage: (page: OrganizationsPages) => void;
  organization: Organization;
};

export function EditOrganizationForm({ setPage, organization }: Props) {
  const [isLoading, setIsLoading] = useState(false);

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      name: organization?.name ?? "",
      description: organization?.description ?? "",
    },
  });

  async function onSubmit(values: z.infer<typeof formSchema>) {
    if (!organization) {
      console.error("Organization not found.");
      toast.error("Organization not found.");
      return;
    }

    setIsLoading(true);

    const [, err] = await editOrganizationApi(
      { ...values, id: organization.id },
      {
        accessToken: "fake-token",
      }
    );

    if (err) {
      console.error(err);
      toast.error("An error occurred while creating the organization.");
      setIsLoading(false);
      return;
    }

    setPage("default");

    toast.success("Organization created successfully!");

    setIsLoading(false);
  }

  return (
    <Form {...form}>
      <form
        onSubmit={form.handleSubmit(onSubmit)}
        className="mt-4 max-w-[700px]"
      >
        <div className="grid gap-2">
          <FormField
            control={form.control}
            name="name"
            render={({ field }) => (
              <FormItem>
                <FormLabel className="flex gap-1">
                  Name
                  <span className="text-red-500">*</span>
                </FormLabel>

                <FormControl>
                  <Input
                    placeholder="Type the organization name"
                    autoComplete="name"
                    type="text"
                    {...field}
                  />
                </FormControl>

                <FormDescription></FormDescription>
                <FormMessage></FormMessage>
              </FormItem>
            )}
          />

          <FormField
            control={form.control}
            name="description"
            render={({ field }) => (
              <FormItem>
                <FormLabel className="flex gap-1">Description</FormLabel>

                <FormControl>
                  <Textarea
                    placeholder="Type the organization description"
                    autoComplete="description"
                    {...field}
                  />
                </FormControl>

                <FormDescription></FormDescription>
                <FormMessage></FormMessage>
              </FormItem>
            )}
          />

          <Button type="submit" disabled={isLoading} className="ml-auto w-fit">
            {isLoading ? "Editing Organization..." : "Edit Organization"}
          </Button>
        </div>
      </form>
    </Form>
  );
}
