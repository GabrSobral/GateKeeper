import { NextRequest, NextResponse } from "next/server";
import { createHash, randomBytes } from "node:crypto";

interface RequestBody {
  clientId: string;
  redirectUri: string;
}

export async function POST(request: NextRequest) {
  const data = (await request.json()) as RequestBody;

  const { clientId, redirectUri } = data;

  // Faça o processamento necessário com "data"

  const state = randomBytes(16).toString("hex");

  // Se estiver utilizando PKCE, gere o code_verifier e o code_challenge aqui
  const codeVerifier = randomBytes(32).toString("hex");
  const codeChallenge = generateCodeChallenge(codeVerifier);

  // Armazene o state (e o codeVerifier, se usar PKCE) na sessão ou cookie do usuário
  // para posterior validação no callback

  const params = new URLSearchParams({
    client_id: clientId,
    redirect_uri: redirectUri,
    response_type: "code",
    scope: "openid profile email",
    state,
    // Inclua code_challenge e code_challenge_method se usar PKCE
    code_challenge: codeChallenge,
    code_challenge_method: "S256",
  });

  const response = NextResponse.redirect(
    `http://localhost:3000/auth/01952b62-2e18-7f0a-9595-8b30b27a183d/sign-in?${params.toString()}`
  );

  response.cookies.set("code_verifier", codeVerifier, {
    httpOnly: true,
    secure: process.env.NODE_ENV === "production",
    path: "/",
    sameSite: "strict",
  });

  response.cookies.set("state", state, {
    httpOnly: true,
    secure: process.env.NODE_ENV === "production",
    path: "/",
    sameSite: "strict",
  });

  // Redireciona o usuário para o endpoint de autorização
  return response;
}

function generateCodeChallenge(codeVerifier: string): string {
  const hash = createHash("sha256");
  hash.update(codeVerifier);

  function base64UrlEncode(buffer: Buffer): string {
    return buffer
      .toString("base64")
      .replace(/\+/g, "-")
      .replace(/\//g, "_")
      .replace(/=/g, "");
  }

  return base64UrlEncode(hash.digest());
}
