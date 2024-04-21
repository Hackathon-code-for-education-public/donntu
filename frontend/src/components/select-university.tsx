"use client";
import { Button } from "@/components/ui/button";
import { useState } from "react";
import { useUniversities } from "@/lib/use-universities";
import { Dialog, DialogContent, DialogTrigger } from "@/components/ui/dialog";

import { SearchInput } from "@/components/university/search-input";
import { UniversitySkeleton } from "./university/university-skeleton";
import { UniversityCard } from "./university/university-card";

/** @ts-ignore */
export function SelectUniversity({ onSelect }) {
  const [open, setOpen] = useState(false);
  const [searchQuery, setSearchQuery] = useState("");
  const { data, error } = useUniversities(searchQuery);

  /** @ts-ignore */
  const handleSelectUniversity = (university) => {
    setOpen(false);
    onSelect(university);
  };

  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogTrigger asChild>
        <Button>Выберите университет</Button>
      </DialogTrigger>
      <DialogContent className="sm:max-w-[425px]">
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
                onClick={() => handleSelectUniversity(uni)}
              />
            ))
          )}
          {error && <p>Ошибка загрузки данных.</p>}
        </div>
      </DialogContent>
    </Dialog>
  );
}
