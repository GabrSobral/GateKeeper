import { cookies } from "next/headers";
import { decryptSessionObjectWithPassword } from "./hash-session";

export async function getServerSession(): Promise<
  [GateKeeperSession | null, GateKeeperSessionError | null]
> {
  const cookieStore = await cookies();
  const encryptedSession = cookieStore.get("gk_session");

  if (!encryptedSession) {
    return [null, { message: "Session Invalid" }];
  }

  const sessionSecret = process.env.SESSION_SECRET;

  if (!sessionSecret) {
    return [null, { message: "SESSION_SECRET not set" }];
  }

  const [session, err] = await decryptSessionObjectWithPassword(
    encryptedSession.value,
    sessionSecret
  );

  if (err) {
    return [null, { message: err.message }];
  }

  return [session, null];
}
