import { toast } from "sonner";
import { useState } from "react";

import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Button, buttonVariants } from "@/components/ui/button";
import {
  Dialog,
  DialogClose,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";

import { cn } from "@/lib/utils";

export function ResetPasswordDialog() {
  const [isLoading, setIsLoading] = useState(false);

  function handler() {
    setIsLoading(true);

    // Logic here
    setIsLoading(false);

    toast.success("Password reset successfully");
  }

  return (
    <Dialog>
      <DialogTrigger
        className={cn(buttonVariants({ variant: "secondary" }), "w-fit")}
      >
        Reset Password
      </DialogTrigger>

      <DialogContent className="sm:max-w-[550px]">
        <DialogHeader>
          <DialogTitle>Reset Password</DialogTitle>
          <DialogDescription>
            On confirm, the user will receive an e-mail with the new password,
            and the user will be required to change it on the next login.
          </DialogDescription>
        </DialogHeader>

        <div className="flex flex-col gap-3">
          <Label htmlFor="temp-password-input">Temporary Password</Label>
          <Input
            id="temp-password-input"
            type="password"
            placeholder="Type the temporary password"
          />
        </div>

        <DialogFooter>
          <DialogClose className={buttonVariants({ variant: "outline" })}>
            Cancel
          </DialogClose>

          <Button type="submit" onClick={handler}>
            {isLoading ? "Resetting..." : "Confirm"}
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
}
