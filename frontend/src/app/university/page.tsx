"use client";

import { SearchInput } from "@/components/university/search-input";
import { UniversityCard } from "@/components/university/university-card";
import { UniversitySkeleton } from "@/components/university/university-skeleton";
import { useUniversities } from "@/lib/use-universities";
import { useState } from "react";

export default function Page() {
  const [searchQuery, setSearchQuery] = useState("");
  const { data, error } = useUniversities(searchQuery);

  return (
    <main className="flex-1 py-8 px-6 min-h-screen">
      <div className="max-w-xl mx-auto">
        <h2 className="text-2xl font-bold mb-4">Поиск университета</h2>
        <div className="mb-8">
          <SearchInput onChange={setSearchQuery} />
        </div>
        <div className="space-y-4">
          {!data && !error ? (
            <>
              <UniversitySkeleton />
              <UniversitySkeleton />
              <UniversitySkeleton />
            </>
          ) : (
            data?.map((uni) => (
              <UniversityCard
                key={uni.id}
                id={uni.id}
                name={uni.name}
                longName={uni.longName}
                logoUrl={uni.logo}
              />
            ))
          )}
          {error && <p>Ошибка загрузки данных.</p>}
        </div>
      </div>
    </main>
  );
}
