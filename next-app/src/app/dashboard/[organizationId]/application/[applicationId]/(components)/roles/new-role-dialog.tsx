import { Button, buttonVariants } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Textarea } from "@/components/ui/textarea";
import { cn } from "@/lib/utils";
import { useState } from "react";

export function NewRoleDialog() {
  const [name, setName] = useState("");
  const [description, setDescription] = useState("");
  const [isLoading, setIsLoading] = useState(false);

  function create() {
    setIsLoading(true);

    clear();

    setIsLoading(false);
  }

  function clear() {
    setName("");
    setDescription("");
  }

  return (
    <Dialog onOpenChange={(isOpened) => isOpened && clear()}>
      <DialogTrigger
        className={cn(buttonVariants({ variant: "default" }), "ml-4")}
      >
        Add Role
      </DialogTrigger>

      <DialogContent className="sm:max-w-[450px]">
        <DialogHeader>
          <DialogTitle>New Application Role</DialogTitle>
          <DialogDescription>
            Create a new role for your application. Handle permissions and
            access.
          </DialogDescription>
        </DialogHeader>

        <div className="grid gap-4 py-4">
          <div className="flex flex-col gap-3">
            <Label htmlFor="name">
              Name <span className="text-red-500">*</span>
            </Label>
            <Input
              id="name"
              placeholder="Type the role name"
              value={name}
              onChange={(e) => setName(e.target.value)}
            />
          </div>

          <div className="flex flex-col gap-3">
            <Label htmlFor="description">
              Description <span className="text-red-500">*</span> (
              {120 - description.length})
            </Label>

            <Textarea
              id="description"
              placeholder="Type the role description"
              value={description}
              maxLength={120}
              onChange={(e) => setDescription(e.target.value)}
            />
          </div>
        </div>

        <DialogFooter>
          <Button type="submit" onClick={create}>
            {isLoading ? "Creating..." : "Create"}
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
}
