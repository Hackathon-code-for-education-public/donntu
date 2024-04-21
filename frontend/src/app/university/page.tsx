"use client";

import { Input } from "@/components/ui/input";
import { useUniversities } from "@/lib/use-universities";
import { UniversityIcon } from "lucide-react";
import Link from "next/link";
import { useState } from "react";

interface UniversityCardProps {
  id: string;
  name: string;
  longName: string;
  logoUrl?: string;
}

function UniversityCard({ id, name, longName, logoUrl }: UniversityCardProps) {
  return (
    <Link href={`/university/${id}`} passHref legacyBehavior>
      <a className="bg-white rounded-md shadow-md p-4 flex items-start">
        {logoUrl ? (
          <img
            alt="University Logo"
            className="mr-4"
            height={60}
            src={logoUrl}
            style={{
              aspectRatio: "60/60",
              objectFit: "contain",
            }}
            width={60}
          />
        ) : (
          <div
            className="mr-4 flex items-center justify-center"
            style={{ width: 60, height: 60 }}
          >
            <UniversityIcon size="60px" />
          </div>
        )}
        <div>
          <h3 className="text-lg font-bold">{name}</h3>
          <p className="text-gray-500 text-sm">{longName}</p>
        </div>
      </a>
    </Link>
  );
}

function UniversitySkeleton() {
  return (
    <div className="bg-white rounded-md shadow-md p-4 flex items-start animate-pulse">
      <div className="mr-4 bg-gray-300" style={{ width: 60, height: 60 }}></div>
      <div className="w-full">
        <div className="h-4 bg-gray-300 rounded w-3/4 mb-2"></div>
        <div className="h-3 bg-gray-300 rounded w-5/6"></div>
      </div>
    </div>
  );
}

function SearchInput({ onChange }: { onChange: (value: string) => void }) {
  return (
    <Input
      className="w-full rounded-md py-2 px-4 border border-gray-300 focus:outline-none focus:ring-2 focus:ring-gray-500"
      placeholder="Название университета"
      type="text"
      onChange={(event) => onChange(event.target.value)}
    />
  );
}

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
                logoUrl={uni.logoUrl}
              />
            ))
          )}
          {error && <p>Ошибка загрузки данных.</p>}
        </div>
      </div>
    </main>
  );
}
