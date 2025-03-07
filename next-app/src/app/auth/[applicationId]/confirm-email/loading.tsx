import { Skeleton } from "@/components/ui/skeleton";
import { Background } from "../(components)/background";
import { InputOTPSeparator } from "@/components/ui/input-otp";

export default function LoadingConfirmEmailPage() {
  return (
    <Background application={null} page="sign-in">
      <div className="flex flex-col space-y-2 text-center">
        <h1 className="text-2xl font-semibold tracking-tight">
          Confirm E-mail
        </h1>

        <p className="text-muted-foreground text-sm">
          We sent a confirmation code to your e-mail. Enter the code below to
          confirm your e-mail
        </p>
      </div>

      <div className="flex gap-2 items-center mx-auto">
        <div className="flex gap-[2px]">
          <Skeleton className="w-[2.15rem] h-[2.15rem]" />
          <Skeleton className="w-[2.15rem] h-[2.15rem]" />
          <Skeleton className="w-[2.15rem] h-[2.15rem]" />
        </div>

        <InputOTPSeparator />

        <div className="flex gap-[2px]">
          <Skeleton className="w-[2.15rem] h-[2.15rem]" />
          <Skeleton className="w-[2.15rem] h-[2.15rem]" />
          <Skeleton className="w-[2.15rem] h-[2.15rem]" />
        </div>

        <Skeleton className="w-full h-[2rem]" />
      </div>
    </Background>
  );
}
