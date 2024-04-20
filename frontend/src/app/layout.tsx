import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";
import { MainNavigation } from "@/components/main-navigation";
import { MainFooter } from "@/components/main-footer";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "ВУЗ-экскурсия",
  description:
    "Сервис по онлайн экскурсиям в университеты с возможностью узнать информацию у студентов и приемной комиссии",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className={inter.className}>
        <MainNavigation />
        {children}
        <MainFooter />
      </body>
    </html>
  );
}
