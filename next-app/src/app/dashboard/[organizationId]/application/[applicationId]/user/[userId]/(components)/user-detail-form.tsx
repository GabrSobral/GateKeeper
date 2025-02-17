"use client";

import { z } from "zod";
import { useState } from "react";
import { useForm } from "react-hook-form";
import { Copy, Pencil } from "lucide-react";
import { zodResolver } from "@hookform/resolvers/zod";
import { useParams, useRouter } from "next/navigation";

import { Badge } from "@/components/ui/badge";
import { Button, buttonVariants } from "@/components/ui/button";
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
import { Label } from "@/components/ui/label";
import { Switch } from "@/components/ui/switch";
import { Separator } from "@/components/ui/separator";
import {
  Tooltip,
  TooltipContent,
  TooltipTrigger,
} from "@/components/ui/tooltip";

import { cn, copy } from "@/lib/utils";

import { formSchema } from "../schema";
import { DeleteUserDialog } from "./delete-user-dialog";
import { MultiFactorAuthForm } from "./multi-factor-auth-form";
import { ApplicationRolesSection } from "./application-roles-section";

export function UserDetailForm() {
  const [isLoading, setIsLoading] = useState(false);
  const [isEditEnabled, setIsEditEnabled] = useState(false);

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
      multiFactorAuth: [],
      roles: [],
    },
  });

  function onSubmit() {
    setIsLoading(true);

    // Logic here
    setIsLoading(false);
  }

  return (
    <>
      {isEditEnabled && (
        <Badge className="mb-4 w-fit text-sm" title="Edit is enabled">
          Editing
        </Badge>
      )}

      <Form {...form}>
        <div className="flex items-center justify-between gap-4">
          {isEditEnabled ? (
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
          ) : (
            <div className="flex gap-4">
              <h2 className="text-3xl font-bold tracking-tight">Ken OConner</h2>

              <Tooltip delayDuration={0}>
                <TooltipTrigger
                  className={buttonVariants({ variant: "outline" })}
                  onClick={() => copy(form.getValues("displayName"))}
                >
                  <Copy />
                </TooltipTrigger>

                <TooltipContent>Copy display name</TooltipContent>
              </Tooltip>
            </div>
          )}

          <div className="flex gap-1">
            <DeleteUserDialog />

            <Tooltip delayDuration={0}>
              <TooltipTrigger
                className={cn(
                  buttonVariants({ variant: "outline" }),
                  "mb-[6px]"
                )}
                onClick={() => {
                  setIsEditEnabled((state) => !state);

                  // if (isEditEnabled) {
                  //   router.push(
                  //     `/dashboard/${organizationId}/application/${applicationId}/user/${applicationId}?edit=${isEditEnabled}`
                  //   );
                  // } else {
                  //   router.push(
                  //     `/dashboard/${organizationId}/application/${applicationId}/user/${applicationId}`
                  //   );
                  // }
                }}
              >
                <Pencil />
              </TooltipTrigger>

              <TooltipContent>Enable Changes</TooltipContent>
            </Tooltip>
          </div>
        </div>

        <div className="mt-4 flex flex-col gap-1">
          <Label
            className="text-foreground text-sm font-semibold"
            htmlFor="user-status-switch"
          >
            Status
          </Label>

          <div className="flex items-center gap-2">
            <Switch
              checked={true}
              aria-labelledby="status-label"
              id="user-status-switch"
              disabled={!isEditEnabled}
            />
            <span className="text-muted-foreground text-xs">Enabled</span>
          </div>
        </div>

        <form onSubmit={onSubmit} className="mt-4 max-w-[700px]">
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
                          data-iseditenabled={isEditEnabled}
                          readOnly={!isEditEnabled}
                          className="data-[iseditenabled=true]:outline-none w-full"
                          placeholder="Type the user first name"
                          autoComplete="given-name"
                          type="text"
                          {...field}
                        />
                      </FormControl>

                      {!isEditEnabled && (
                        <Tooltip delayDuration={0}>
                          <TooltipTrigger
                            className={buttonVariants({ variant: "outline" })}
                            onClick={() => copy(field.value)}
                          >
                            <Copy />
                          </TooltipTrigger>

                          <TooltipContent>Copy first name</TooltipContent>
                        </Tooltip>
                      )}
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
                          data-iseditenabled={isEditEnabled}
                          readOnly={!isEditEnabled}
                          className="data-[iseditenabled=true]:outline-none w-full"
                          placeholder="Type the user last name"
                          autoComplete="family-name"
                          type="text"
                          {...field}
                        />
                      </FormControl>

                      {!isEditEnabled && (
                        <Tooltip delayDuration={0}>
                          <TooltipTrigger
                            className={buttonVariants({ variant: "outline" })}
                            onClick={() => copy(field.value || "")}
                          >
                            <Copy />
                          </TooltipTrigger>

                          <TooltipContent>Copy last name</TooltipContent>
                        </Tooltip>
                      )}
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
                        data-iseditenabled={isEditEnabled}
                        readOnly={!isEditEnabled}
                        className="data-[iseditenabled=true]:outline-none w-full"
                        placeholder="Type the user e-mail"
                        autoComplete="email"
                        type="text"
                        {...field}
                      />
                    </FormControl>

                    {!isEditEnabled && (
                      <Tooltip delayDuration={0}>
                        <TooltipTrigger
                          className={buttonVariants({ variant: "outline" })}
                          onClick={() => copy(field.value || "")}
                        >
                          <Copy />
                        </TooltipTrigger>

                        <TooltipContent>Copy e-mail</TooltipContent>
                      </Tooltip>
                    )}
                  </div>
                  <FormMessage></FormMessage>
                </FormItem>
              )}
            />

            <div className="flex flex-col gap-1">
              <span className="text-sm font-medium"> Reset User Password</span>

              <span className="text-muted-foreground my-2 text-sm">
                Reset the user password. On click, the user will receive an
                e-mail with the new password, and the user will be required to
                change it on the next login.
              </span>

              {/* <ResetPasswordDialog /> */}
            </div>

            <Separator className="my-2" />

            <MultiFactorAuthForm isEditEnabled={isEditEnabled} />

            <Separator className="my-2" />

            <ApplicationRolesSection isEditEnabled={isEditEnabled} />
          </div>

          <Button
            type="submit"
            className="float-right mt-4"
            disabled={isLoading}
          >
            Apply Changes
          </Button>
        </form>
      </Form>
    </>
  );
}
