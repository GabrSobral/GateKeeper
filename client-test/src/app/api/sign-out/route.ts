import { getServerSession } from "@/lib/utils/get-server-session";
import { NextResponse } from "next/server";

export async function POST() {
  const [session, err] = await getServerSession();

  if (err) {
    return new Response(JSON.stringify({ message: err.message }), {
      status: 401,
    });
  }

  if (!session) {
    return new Response(JSON.stringify({ message: "Unauthorized" }), {
      status: 401,
    });
  }

  const response = NextResponse.json(null);

  response.cookies.delete("gk_session");
  response.cookies.delete("gk_state");
  response.cookies.delete("gk_code_verifier");

  return response;
}
