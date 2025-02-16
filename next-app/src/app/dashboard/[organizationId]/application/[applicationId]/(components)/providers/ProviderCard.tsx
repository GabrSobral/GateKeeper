"use client";

import { Badge } from "@/components/ui/badge";
import { buttonVariants } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import {
  Sheet,
  SheetClose,
  SheetContent,
  SheetDescription,
  SheetFooter,
  SheetHeader,
  SheetTitle,
  SheetTrigger,
} from "@/components/ui/sheet";
import { Switch } from "@/components/ui/switch";
import { useState } from "react";

type Props = {
  title: string;
  description: string;
  isEnabled: boolean;
  clientId: string;
  clientSecret: string;
};

export function ProviderCard({
  clientId,
  clientSecret,
  description,
  isEnabled,
  title,
}: Props) {
  const [draftClientId, setDraftClientId] = useState(clientId);
  const [draftClientSecret, setDraftClientSecret] = useState(clientSecret);
  const [draftIsEnabled, setDraftIsEnabled] = useState(isEnabled);

  return (
    <Sheet>
      <SheetTrigger>
        <Card className="transition-all hover:scale-[1.01] hover:cursor-pointer hover:shadow-lg">
          <CardHeader>
            <CardTitle className="flex flex-wrap justify-between gap-4">
              {title}
            </CardTitle>

            <CardDescription className="text-left">
              {description}
            </CardDescription>
          </CardHeader>

          <CardContent className="flex gap-1">
            <Badge variant="outline" className="w-fit">
              {draftClientId && draftClientSecret
                ? "Configured"
                : "Not configured"}
            </Badge>

            <Badge
              variant={draftIsEnabled ? "default" : "secondary"}
              className="w-fit"
            >
              {draftIsEnabled ? "Enabled" : "Disabled"}
            </Badge>
          </CardContent>
        </Card>
      </SheetTrigger>

      <SheetContent side="right">
        <SheetHeader>
          <SheetTitle>Configure Provider</SheetTitle>
          <SheetDescription>
            Make changes to your Google authentication provider configuration.
            Then click &quot;Save changes&quot; to apply them.
          </SheetDescription>
        </SheetHeader>

        <div className="my-4 flex flex-col gap-4">
          <div className="flex flex-col gap-2">
            <Label htmlFor="client-id">Client ID</Label>
            <Input
              id="client-id"
              value={draftClientId}
              type="text"
              placeholder="Type the client ID"
              className="col-span-3"
              onChange={(e) => setDraftClientId(e.target.value)}
            />
          </div>

          <div className="flex flex-col gap-2">
            <Label htmlFor="client-secret">Client Secret</Label>
            <Input
              id="client-secret"
              type="password"
              placeholder="Type the client secret"
              value={draftClientSecret}
              className="col-span-3"
              onChange={(e) => setDraftClientSecret(e.target.value)}
            />
          </div>

          <div className="flex items-center space-x-2">
            <Switch
              id="provider-enabled"
              checked={draftIsEnabled}
              onChange={() => setDraftIsEnabled(!draftIsEnabled)}
            />

            <Label htmlFor="provider-enabled">
              {draftIsEnabled ? "Enabled" : "Disabled"}
            </Label>
          </div>
        </div>

        <SheetFooter>
          <SheetClose className={buttonVariants({ variant: "default" })}>
            Save changes
          </SheetClose>
        </SheetFooter>
      </SheetContent>
    </Sheet>
  );
}
