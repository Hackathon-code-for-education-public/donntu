import UniversityBlock from "@/components/univercity-block";
import { Metadata } from "next";

export const metadata: Metadata = {
  title: "Сравнение ВУЗов",
};

export default function Page({
  searchParams,
}: {
  searchParams: { [key: string]: string };
}) {
  //   router.push(`tickets/?page=${Number(page) + 1}&per_page=${perPage}`);

  return (
    <main className="flex min-h-screen justify-around">
      <UniversityBlock universityId={searchParams["first"]} />
      <UniversityBlock universityId={searchParams["second"]} />
    </main>
  );
}
