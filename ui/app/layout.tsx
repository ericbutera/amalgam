import type { Metadata } from "next";
import "./globals.css";

import Nav from "./components/nav";
import { Toaster } from "react-hot-toast";

export const metadata: Metadata = {
  title: "amalgam",
  description: "amalgam",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body>
        <Nav />
        <div className="flex w-full">{children}</div>
        <Toaster position="top-right" />
      </body>
    </html>
  );
}
