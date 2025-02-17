import { Checkbox } from "@/components/ui/checkbox";
import { Label } from "@/components/ui/label";
import {
  Tooltip,
  TooltipContent,
  TooltipTrigger,
} from "@/components/ui/tooltip";

type Props = {
  isEditEnabled: boolean;
};

export function ApplicationRolesSection({ isEditEnabled }: Props) {
  console.log(isEditEnabled);

  return (
    <div className="flex flex-col gap-1">
      <span className="text-sm font-medium"> Application Roles </span>

      <span className="text-muted-foreground my-2 text-sm">
        Choose the roles that will be assigned to this user.
      </span>

      <div className="flex items-center space-x-2">
        <Checkbox id="asd1" />

        <Tooltip>
          <TooltipTrigger>
            <Label
              htmlFor="asd1"
              className="text-sm leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
            >
              User
            </Label>
          </TooltipTrigger>

          <TooltipContent>User Description</TooltipContent>
        </Tooltip>
      </div>

      <div className="flex items-center space-x-2">
        <Checkbox id="asd2" />

        <Tooltip>
          <TooltipTrigger>
            <Label
              htmlFor="asd2"
              className="text-sm leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
            >
              Admin
            </Label>
          </TooltipTrigger>

          <TooltipContent>Admin Description</TooltipContent>
        </Tooltip>
      </div>
    </div>
  );
}
