import { getServerSession } from "@/lib/utils/get-server-session";

export async function GET() {
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

  return new Response(JSON.stringify(session), {
    headers: {
      "Content-Type": "application/json",
    },
  });
}
