"use client";

import { SelectUniversity } from "@/components/select-university";
import UniversityBlock from "@/components/university-block";
import { useSearchParams, useRouter, usePathname } from "next/navigation";
import { Suspense } from 'react'

function Page() {
  const searchParams = useSearchParams();
  const pathname = usePathname();
  const { replace } = useRouter();

  /** @ts-ignore */
  const handleSelectUniversity = (university, position) => {
    const params = new URLSearchParams(searchParams);
    params.set(position, university.id);
    replace(`${pathname}?${params.toString()}`);
  };

  const first = searchParams.get("first");
  const second = searchParams.get("second");

  return (
    <main className="flex min-h-screen justify-around">
      <div className="flex items-center mt-10 flex-col">
        <SelectUniversity
          /** @ts-ignore */
          onSelect={(uni) => handleSelectUniversity(uni, "first")}
        />
        {first && <UniversityBlock universityId={first} />}
      </div>
      <div className="flex items-center mt-10 flex-col">
        <SelectUniversity
          /** @ts-ignore */
          onSelect={(uni) => handleSelectUniversity(uni, "second")}
        />
        {second && <UniversityBlock universityId={second} />}
      </div>
    </main>
  );
}

export default function PPage() {
  return (
    <Suspense>
      <Page />
    </Suspense>
  );
}
