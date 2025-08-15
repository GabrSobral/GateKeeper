"use client";

import { signIn } from "./sign-in";

/*
 * This function remove the session cookie and redirects the user to the GateKeeper sign-in page.
 * ONLY ON THE CLIENT SIDE
 */
export async function signOut(
  shouldLoginAgain = false,
  redirectUri: string | undefined = "/"
) {
  await fetch("/api/sign-out", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    mode: "no-cors",
  });

  if (shouldLoginAgain) {
    signIn({ redirectUri });
  }
}
