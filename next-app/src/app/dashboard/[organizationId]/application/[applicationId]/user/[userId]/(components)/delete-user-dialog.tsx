"use client";

import { useState } from "react";
import { Trash } from "lucide-react";

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
import { Button, buttonVariants } from "@/components/ui/button";
import {
  Tooltip,
  TooltipContent,
  TooltipTrigger,
} from "@/components/ui/tooltip";

export function DeleteUserDialog() {
  const [isLoading, setIsLoading] = useState(false);

  function handler() {
    setIsLoading(true);

    // Logic here
    setIsLoading(false);
  }

  return (
    <Dialog>
      <Tooltip delayDuration={0}>
        <TooltipTrigger asChild>
          <DialogTrigger className={buttonVariants({ variant: "destructive" })}>
            <Trash />
          </DialogTrigger>
        </TooltipTrigger>

        <TooltipContent>Delete User</TooltipContent>
      </Tooltip>

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
