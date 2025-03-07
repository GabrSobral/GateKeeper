import { Skeleton } from "@/components/ui/skeleton";
import { Background } from "../(components)/background";

export default function LoadingSignUpPage() {
  return (
    <Background application={null} page="sign-up">
      <div className="flex flex-col space-y-2 text-center">
        <h1 className="text-2xl font-semibold tracking-tight">
          Create an account
        </h1>

        <p className="text-muted-foreground text-sm">
          Enter your data below to create your account
        </p>
      </div>

      <div className="grid gap-4">
        <div className="flex flex-col gap-2">
          <Skeleton className="w-[80px] h-[1rem]" />
          <Skeleton className="w-full h-[2rem]" />
        </div>

        <div className="flex flex-col gap-2">
          <Skeleton className="w-[80px] h-[1rem]" />
          <Skeleton className="w-full h-[2rem]" />
        </div>

        <div className="flex flex-col gap-2">
          <Skeleton className="w-[80px] h-[1rem]" />
          <Skeleton className="w-full h-[2rem]" />
        </div>

        <div className="flex flex-col gap-2">
          <Skeleton className="w-[80px] h-[1rem]" />
          <Skeleton className="w-full h-[2rem]" />
        </div>

        <Skeleton className="w-full h-[2rem]" />
      </div>
    </Background>
  );
}
