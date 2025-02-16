import { Skeleton } from "@/components/ui/skeleton";

export default async function Loading() {
  return (
    <>
      <Skeleton className="w-[20rem] h-6 mx-4 mt-1" />

      <main className="flex flex-col p-4">
        <Skeleton className="w-[15rem] h-6 mb-4" />

        <div className="flex items-center justify-between gap-4">
          <Skeleton className="w-[12rem] h-10" />

          <div className="flex gap-1">
            <Skeleton className="w-8 h-8" />
            <Skeleton className="w-8 h-8" />
          </div>
        </div>

        <Skeleton className="w-full h-[2rem] mt-3" />

        <div className="mt-4 flex flex-wrap gap-2">
          <Skeleton className="w-12 h-5 rounded-full" />
          <Skeleton className="w-20 h-5 rounded-full" />
          <Skeleton className="w-15 h-5 rounded-full" />
        </div>
      </main>
    </>
  );
}
