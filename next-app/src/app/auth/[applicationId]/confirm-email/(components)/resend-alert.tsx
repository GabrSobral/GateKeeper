"use client";

import { toast } from "sonner";
import { useState } from "react";
import { useParams, useSearchParams } from "next/navigation";

import { resendConfirmEmailApi } from "@/services/auth/resend-confirm-email";

export function ResendAlert() {
  const applicationId = useParams().applicationId as string;
  const searchParams = useSearchParams();

  const [isLoading, setIsLoading] = useState(false);

  const email = searchParams.get("email") || "";

  async function resendConfirmationToken() {
    setIsLoading(true);

    const [err] = await resendConfirmEmailApi({ applicationId, email });

    if (err) {
      toast.error(err.response?.data?.message || err.message);
      console.error(err);
      setIsLoading(false);
      return;
    }

    setIsLoading(false);

    toast.success("Confirmation code sent successfully");
  }

  return (
    <div className="flex flex-col  items-center justify-center p-3 rounded-sm bg-gray-50 dark:bg-gray-900 shadow-sm">
      <p className="text-muted-foreground text-sm">
        It seems you are trying to sign in. Please confirm your e-mail before
        signing in.
      </p>

      <button
        type="button"
        onClick={resendConfirmationToken}
        disabled={isLoading}
        data-isloading={isLoading}
        className="hover:underline data-[isloading=true]:pointer-events-none text-primary text-sm font-semibold mt-2"
      >
        {isLoading
          ? "Sending..."
          : "Click here to resend the confirmation code"}
      </button>
    </div>
  );
}
