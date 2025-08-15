"use client";

import { FormType } from "./user-detail-form";
import { Button } from "@/components/ui/button";
import {
  Tooltip,
  TooltipContent,
  TooltipTrigger,
} from "@/components/ui/tooltip";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import {
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";

type Props = {
  form: FormType;
  isEditEnabled: boolean;
};

export function MultiFactorAuthForm({ isEditEnabled, form }: Props) {
  return (
    <div className="flex flex-col gap-1">
      <FormField
        control={form.control}
        name="preferred2FAMethod"
        render={({ field }) => (
          <FormItem>
            <FormLabel>Preferred MFA Method</FormLabel>
            <Select
              onValueChange={field.onChange}
              defaultValue={field.value || "Nenhum"}
              disabled={!isEditEnabled}
            >
              <FormControl>
                <SelectTrigger className="min-w-[10rem]">
                  <SelectValue placeholder="Select a MFA method to display" />
                </SelectTrigger>
              </FormControl>

              <SelectContent>
                <SelectItem value="email">E-mail</SelectItem>
                <SelectItem value="totp">Authenticator App</SelectItem>
              </SelectContent>
            </Select>

            <FormDescription>
              Set your preferred method to use for multi-factor authentication
              when signing into the application.
            </FormDescription>
            <FormMessage />
          </FormItem>
        )}
      />

      <ul className="flex flex-col gap-1 mt-3">
        <li className="flex justify-between items-center w-full">
          <div className="flex gap-4 items-center">
            <Tooltip>
              <TooltipTrigger type="button">
                <span className="text-muted-background font-semibold">
                  E-mail
                </span>
              </TooltipTrigger>

              <TooltipContent>
                Send a verification code to the user&apos;s email address.
              </TooltipContent>
            </Tooltip>

            {form.getValues("isMfaEmailConfigured") ? (
              <span className="text-sm text-green bg-green-300/30 px-3 rounded-full border border-green-300 text-green-900 dark:text-green-300">
                Configured
              </span>
            ) : (
              <span className="text-sm bg-red-300/30 px-3 rounded-full border border-red-300 text-red-900 dark:text-red-300">
                Not Configured
              </span>
            )}
          </div>

          {isEditEnabled && (
            <Button variant="outline" className="font-semibold text-red-500">
              Disable
            </Button>
          )}
        </li>

        <li className="flex justify-between items-center w-full">
          <div className="flex gap-4 items-center">
            <Tooltip>
              <TooltipTrigger type="button">
                <span className="text-muted-background font-semibold">
                  Authenticator App
                </span>
              </TooltipTrigger>

              <TooltipContent>
                Use an authenticator app to generate a verification code.
              </TooltipContent>
            </Tooltip>

            {form.getValues("IsMfaAuthAppConfigured") ? (
              <span className="text-sm text-green bg-green-300/30 px-3 rounded-full border border-green-300 text-green-900 dark:text-green-300">
                Configured
              </span>
            ) : (
              <span className="text-sm bg-red-300/30 px-3 rounded-full border border-red-300 text-red-900 dark:text-red-300">
                Not Configured
              </span>
            )}
          </div>

          {isEditEnabled && (
            <Button variant="outline" className="font-semibold text-red-500">
              Disable
            </Button>
          )}
        </li>
      </ul>
    </div>
  );
}
