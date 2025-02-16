"use client";

import { useState } from "react";

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
import { ApplicationRole } from ".";

type Props = {
  isOpened: boolean;
  onOpenChange: (value: boolean) => void;
  role: ApplicationRole | null;
};

export function DeleteRoleDialog({ isOpened, onOpenChange }: Props) {
  const [isLoading, setIsLoading] = useState(false);

  function handler() {
    setIsLoading(true);

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
