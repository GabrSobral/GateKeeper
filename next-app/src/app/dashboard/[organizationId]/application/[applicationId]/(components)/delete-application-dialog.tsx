"use client";

import { toast } from "sonner";
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
import { IApplication } from "@/services/dashboard/get-application-by-id";
import { deleteApplicationApi } from "@/services/dashboard/delete-application";
import { useParams, useRouter } from "next/navigation";

type Props = {
  application: IApplication | null;
};

export function DeleteApplicationDialog({ application }: Props) {
  const router = useRouter();
  const { organizationId } = useParams() as { organizationId: string };
  const [isLoading, setIsLoading] = useState(false);

  const [isOpened, setIsOpened] = useState(false);

  async function handler() {
    if (!application) {
      console.error("Application not found.");
      toast.error("Application not found.");
      return;
    }

    setIsLoading(true);

    const [err] = await deleteApplicationApi(
      { applicationId: application.id, organizationId },
      {
        accessToken: "",
      }
    );

    if (err) {
      console.error(err);
      toast.error("An error occurred while deleting the application.");
      setIsLoading(false);
      return;
    }

    setIsOpened(false);
    router.push(`/dashboard/${organizationId}/application`);

    // Logic here
    setIsLoading(false);
  }

  return (
    <Dialog open={isOpened} onOpenChange={setIsOpened}>
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
