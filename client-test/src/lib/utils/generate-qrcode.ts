import qrCode from "qrcode";

export function generateQrCode(
  otpUrl: string,
  canvas: HTMLCanvasElement
): string {
  const options: qrCode.QRCodeRenderersOptions = {
    color: {
      dark: "#000000", // QR code color
      light: "#FFFFFF", // Background color
    },
    width: 256, // Width of the QR code
  };

  qrCode.toCanvas(canvas, otpUrl, options, (error) => {
    if (error) {
      console.error("Error generating QR code:", error);
      return null;
    }
  });

  return canvas.toDataURL("image/png");
}
