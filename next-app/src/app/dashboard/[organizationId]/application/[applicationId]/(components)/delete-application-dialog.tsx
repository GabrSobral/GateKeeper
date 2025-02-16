"use client";

import { useState } from "react";
import { Trash } from "lucide-react";

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

export function DeleteApplicationDialog() {
  const [isLoading, setIsLoading] = useState(false);

  async function handler() {
    setIsLoading(true);

    // Logic here
    setIsLoading(false);
  }

  return (
    <Dialog>
      <DialogTrigger className={buttonVariants({ variant: "destructive" })}>
        <Trash />
      </DialogTrigger>

      <DialogContent className="sm:max-w-[450px]">
        <DialogHeader>
          <DialogTitle>Delete Application</DialogTitle>
          <DialogDescription>
            On deleting this application, it will be permanently removed from
            the organization. Are you sure?
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
