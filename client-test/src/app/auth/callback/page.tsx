import { GenerateMfaSecretButton } from "@/components/generate-mfa-secret-button";
import { LogoutButton } from "@/components/logout-button";
import { getServerSession } from "@/lib/utils/get-server-session";

export default async function CallbackPage() {
  const [session, err] = await getServerSession();

  if (err) {
    return <p>Error: {err.message}</p>;
  }

  return (
    <div>
      <h1> Session</h1>

      {session && (
        <pre>
          <code>{JSON.stringify(session, null, 2)}</code>
        </pre>
      )}

      <LogoutButton />
      <GenerateMfaSecretButton />
    </div>
  );
}
