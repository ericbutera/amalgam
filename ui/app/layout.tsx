import type { Metadata } from "next";
import "./globals.css";

import Nav from "./nav";

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
        <div>
          {children}
        </div>
      </body>
    </html>
  );
}
