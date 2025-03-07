import { Skeleton } from "@/components/ui/skeleton";
import { Background } from "../(components)/background";
import { InputOTPSeparator } from "@/components/ui/input-otp";

export default function LoadingOneTimePasswordPage() {
  return (
    <Background application={null} page="sign-in">
      <div className="flex flex-col space-y-2 text-center">
        <h1 className="text-2xl font-semibold tracking-tight">Multi Factor</h1>

        <p className="text-muted-foreground text-sm">
          Enter your one time password below to sign in
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
