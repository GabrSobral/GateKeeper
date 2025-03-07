"use server";

import crypto from "node:crypto";

const ALGORITHM = "aes-256-gcm";
const IV_LENGTH = 12; // GCM works best with 12-byte IVs
const SALT_LENGTH = 16;
const KEY_LENGTH = 32;
const AUTH_TAG_LENGTH = 16;

/**
 * Derives a fast cryptographic key using HKDF (HMAC-based Extract-and-Expand Key Derivation Function).
 */
function deriveKey(password: string, salt: Buffer): Buffer {
  const key = crypto.hkdfSync(
    "sha256",
    Buffer.from(password, "utf8"),
    salt,
    Buffer.alloc(0),
    KEY_LENGTH
  );
  return Buffer.from(key); // Explicit conversion to Buffer
}

/**
 * Encrypts a session object using AES-256-GCM with HKDF key derivation.
 */
export async function hashSessionObjectWithPassword(
  session: Record<string, unknown>,
  password: string
): Promise<[string | null, Error | null]> {
  try {
    const salt = crypto.randomBytes(SALT_LENGTH);
    const iv = crypto.randomBytes(IV_LENGTH);
    const key = deriveKey(password, salt);

    const cipher = crypto.createCipheriv(ALGORITHM, key, iv);
    const jsonBuffer = Buffer.from(JSON.stringify(session), "utf8");

    const encrypted = Buffer.concat([
      cipher.update(jsonBuffer),
      cipher.final(),
    ]);
    const authTag = cipher.getAuthTag();

    return [
      Buffer.concat([salt, iv, authTag, encrypted]).toString("base64"),
      null,
    ];
  } catch (error) {
    console.error("Encryption failed:", error);
    return [null, error as Error];
  }
}

/**
 * Decrypts an encrypted session object using AES-256-GCM with HKDF.
 */
export async function decryptSessionObjectWithPassword(
  encryptedSession: string,
  password: string
): Promise<[GateKeeperSession | null, Error | null]> {
  try {
    const data = Buffer.from(encryptedSession, "base64");

    const salt = data.subarray(0, SALT_LENGTH);
    const iv = data.subarray(SALT_LENGTH, SALT_LENGTH + IV_LENGTH);
    const authTag = data.subarray(
      SALT_LENGTH + IV_LENGTH,
      SALT_LENGTH + IV_LENGTH + AUTH_TAG_LENGTH
    );
    const encrypted = data.subarray(SALT_LENGTH + IV_LENGTH + AUTH_TAG_LENGTH);

    const key = deriveKey(password, salt);
    const decipher = crypto.createDecipheriv(ALGORITHM, key, iv);
    decipher.setAuthTag(authTag);

    const decrypted = Buffer.concat([
      decipher.update(encrypted),
      decipher.final(),
    ]);
    return [JSON.parse(decrypted.toString("utf8")), null];
  } catch (error) {
    console.error("Decryption failed:", error);

    return [null, error as Error];
  }
}
