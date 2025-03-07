"use client";

/*
 * This function redirects the user to the GateKeeper sign-in page.
 * ONLY ON THE CLIENT SIDE
 */

type Config = {
  redirectUri: string;
};

export async function signIn({ redirectUri }: Config) {
  const response = await fetch("/api/sign-in", {
    method: "POST",
    body: JSON.stringify({ redirectUri }),
    headers: {
      "Content-Type": "application/json",
    },
    mode: "no-cors",
  });

  const data = await response.json();

  window.location.href = data.url;
}
