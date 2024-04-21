import UniversityBlock from "@/components/university-block";
import { Metadata } from "next";

export const metadata: Metadata = {
  title: "Сравнение ВУЗов",
};

export default function Page({
  searchParams,
}: {
  searchParams: { [key: string]: string };
}) {

  return (
    <main className="flex min-h-screen justify-around">
      <UniversityBlock universityId={searchParams["first"]} />
      <UniversityBlock universityId={searchParams["second"]} />
    </main>
  );
}
