"use client";

import { useRef } from "react";

import { generateMfaToken } from "@/lib/utils/generate-mfa-token";
import { generateQrCode } from "@/lib/utils/generate-qrcode";

export function GenerateMfaSecretButton() {
  const qrCodeCanvas = useRef<HTMLCanvasElement | null>(null);

  function generateMfaTokenHandler() {
    if (qrCodeCanvas.current) {
      generateMfaToken()
        .then((data) => {
          const qrCode = generateQrCode(
            data.otpUrl,
            qrCodeCanvas.current as HTMLCanvasElement
          );

          console.log("QR Code generated:", qrCode);
        })
        .catch((error: unknown) => {
          console.error("Error generating MFA token:", error);
        });
    }
  }

  return (
    <div className="flex flex-col items-center justify-center p-4 bg-white rounded-md shadow-md">
      <button
        onClick={generateMfaTokenHandler}
        type="button"
        className="px-4 py-2 text-white bg-blue-500 rounded-md shadow-md hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 focus:ring-offset-white"
      >
        Generate MFA Secret
      </button>

      <div className="mt-4 bg-red-50">
        <canvas
          ref={qrCodeCanvas}
          id="qrcode"
          width={256}
          height={256}
          className="border border-gray-300 rounded-md"
        ></canvas>
      </div>
    </div>
  );
}
