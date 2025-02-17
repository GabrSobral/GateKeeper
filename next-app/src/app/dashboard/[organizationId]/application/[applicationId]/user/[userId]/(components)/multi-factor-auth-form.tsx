import { Label } from "@/components/ui/label";
import { Checkbox } from "@/components/ui/checkbox";
import {
  Tooltip,
  TooltipContent,
  TooltipTrigger,
} from "@/components/ui/tooltip";

type Props = {
  isEditEnabled: boolean;
};

export function MultiFactorAuthForm({ isEditEnabled }: Props) {
  console.log(isEditEnabled);

  return (
    <div className="flex flex-col gap-1">
      <span className="text-sm font-medium"> Multi Factor Authentication </span>

      <span className="text-muted-foreground my-2 text-sm">
        Choose the methods that will be used for multi factor authentication.
      </span>

      <div className="flex items-center space-x-2">
        <Checkbox id="e-mail-mfa-1" />

        <Tooltip>
          <TooltipTrigger>
            <Label
              htmlFor="e-mail-mfa-1"
              className="text-sm leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
            >
              E-mail
            </Label>
          </TooltipTrigger>

          <TooltipContent>
            Send a verification code to the user&apos;s email address.
          </TooltipContent>
        </Tooltip>
      </div>

      <div className="flex items-center space-x-2">
        <Checkbox id="auth-app-mfa-2" />

        <Tooltip>
          <TooltipTrigger>
            <Label
              htmlFor="auth-app-mfa-2"
              className="text-sm leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
            >
              Authenticator App (Microsoft, Google, etc)
            </Label>
          </TooltipTrigger>

          <TooltipContent>
            Use an authenticator app to generate a verification code.
          </TooltipContent>
        </Tooltip>
      </div>
    </div>
  );
}
