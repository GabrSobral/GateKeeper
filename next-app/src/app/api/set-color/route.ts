// app/api/set-color/route.ts
import { cookies } from "next/headers";
import { NextResponse } from "next/server";

export async function POST(req: Request) {
  const { color } = await req.json();
  (await cookies()).set("color-scheme", color, { path: "/" });

  return NextResponse.json({ ok: true });
}
