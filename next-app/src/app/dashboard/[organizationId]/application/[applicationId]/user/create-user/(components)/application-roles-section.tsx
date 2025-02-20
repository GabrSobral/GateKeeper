"use client";

import { Label } from "@/components/ui/label";
import { Checkbox } from "@/components/ui/checkbox";
import {
  Tooltip,
  TooltipContent,
  TooltipTrigger,
} from "@/components/ui/tooltip";

import { FormType } from "./user-detail-form";
import { IApplication } from "@/services/dashboard/get-application-by-id";

type Props = {
  form: FormType;
  roles: IApplication["roles"]["data"];
};

export function ApplicationRolesSection({ form, roles }: Props) {
  return (
    <div className="flex flex-col gap-1">
      <span className="text-sm font-medium"> Application Roles </span>

      <span className="text-muted-foreground my-2 text-sm">
        Choose the roles that will be assigned to this user.
      </span>

      {roles.map((role) => (
        <div className="flex items-center space-x-2" key={role.id}>
          <Checkbox
            id={role.id}
            checked={form.getValues("roles").includes(role.id)}
            onCheckedChange={(checked) => {
              const currentRoles = form.getValues("roles");

              if (checked) {
                form.setValue("roles", [...currentRoles, role.id]);
              } else {
                form.setValue(
                  "roles",
                  currentRoles.filter((r) => r !== role.id)
                );
              }

              console.log(form.getValues("roles"), { checked });
            }}
          />

          <Tooltip>
            <TooltipTrigger>
              <Label
                htmlFor={role.id}
                className="text-sm leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
              >
                {role.name}
              </Label>
            </TooltipTrigger>

            <TooltipContent>{role.description}</TooltipContent>
          </Tooltip>
        </div>
      ))}
    </div>
  );
}
