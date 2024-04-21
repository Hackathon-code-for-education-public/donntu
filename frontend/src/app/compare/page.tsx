"use client";

import { SelectUniversity } from "@/components/select-university";
import UniversityBlock from "@/components/university-block";
import { useSearchParams, useRouter, usePathname } from "next/navigation";
import { useEffect } from "react";

export default function Page() {
  const searchParams = useSearchParams();
  const pathname = usePathname();
  const { replace } = useRouter();

  const handleSelectUniversity = (university, position) => {
    const params = new URLSearchParams(searchParams);
    params.set(position, university.id);
    replace(`${pathname}?${params.toString()}`);
  };

  const first = searchParams.get("first");
  const second = searchParams.get("second");

  return (
    <main className="flex min-h-screen justify-around">
      <div className="flex items-center flex-col">
        <SelectUniversity
          onSelect={(uni) => handleSelectUniversity(uni, "first")}
        />
        {first && <UniversityBlock universityId={first} />}
      </div>
      <div className="flex items-center flex-col">
        <SelectUniversity
          onSelect={(uni) => handleSelectUniversity(uni, "second")}
        />
        {second && <UniversityBlock universityId={second} />}
      </div>
    </main>
  );
}
