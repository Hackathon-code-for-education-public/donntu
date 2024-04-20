"use client";
import { OpenDay } from "./open-day";
import { useOpenDaysByUniversity } from "@/lib/use-open-days";

interface IProps {
  universityId: string;
}

export function UniversityOpenDays({ universityId }: IProps) {
  const { data, isLoading, error } = useOpenDaysByUniversity(universityId);

  return (
    <div className="grid grid-cols-1 md:grid-cols-3 gap-4 py-4">
      {isLoading && "Loading..."}
      {data &&
        data.map((day) => {
          return <OpenDay key={day.id} day={day} />;
        })}
    </div>
  );
}
