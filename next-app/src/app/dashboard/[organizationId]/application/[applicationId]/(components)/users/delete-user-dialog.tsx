"use client";

import { toast } from "sonner";
import { useState } from "react";
import { useParams } from "next/navigation";

import { Button, buttonVariants } from "@/components/ui/button";
import {
  Dialog,
  DialogClose,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";

import { ApplicationUser } from ".";
import { deleteApplicationUserApi } from "@/services/dashboard/delete-application-user";

type Props = {
  isOpened: boolean;
  onOpenChange: (value: boolean) => void;
  user: ApplicationUser | null;
  removeUser: (user: ApplicationUser) => void;
};

export function DeleteUserDialog({
  isOpened,
  onOpenChange,
  user,
  removeUser,
}: Props) {
  const [isLoading, setIsLoading] = useState(false);
  const { organizationId, applicationId } = useParams() as {
    organizationId: string;
    applicationId: string;
  };

  async function handler() {
    if (!user) {
      console.error("User is not defined");
      toast.error("User is not defined");
      return;
    }

    setIsLoading(true);

    const [err] = await deleteApplicationUserApi(
      { applicationId, organizationId, userId: user?.id },
      { accessToken: "" }
    );

    if (err) {
      console.error(err);
      toast.error("Failed to delete role");
      setIsLoading(false);
      return;
    }

    removeUser(user);

    // Logic here
    setIsLoading(false);

    onOpenChange(false);
  }

  return (
    <Dialog open={isOpened} onOpenChange={(value) => onOpenChange(value)}>
      <DialogContent className="sm:max-w-[450px]">
        <DialogHeader>
          <DialogTitle>Delete User</DialogTitle>

          <DialogDescription>
            On deleting this user, it will be permanently removed from the
            application. Are you sure?
          </DialogDescription>
        </DialogHeader>

        <DialogFooter>
          <DialogClose className={buttonVariants({ variant: "outline" })}>
            Cancel
          </DialogClose>

          <Button type="submit" onClick={handler} variant="destructive">
            {isLoading ? "Deleting..." : "Delete"}
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
}
