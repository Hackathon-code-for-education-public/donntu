"use client";

import { useUniversity } from "@/lib/use-university";
import { Skeleton } from "@/components/ui/skeleton";

interface UniversityBlockProps {
  universityId: string;
}

const UniversityBlock: React.FC<UniversityBlockProps> = ({ universityId }) => {
  const { data: university, isLoading, error } = useUniversity(universityId);

  if (error) {
    return <div>Error loading university details. Please try again later.</div>;
  }

  const renderField = (value: any, placeholder: string, sizeClass = 'w-3/4', fontSize = 'text-4xl', containerPadding = 'm-2') => (
    isLoading || value == null ? (
      <Skeleton className={`h-6 ${sizeClass} mt-1`} />
    ) : (
      <div className={`${containerPadding}`}>
        <p className={`${fontSize} text-center`}>{value}</p>
        <p>{placeholder}</p>
      </div>
    )
  );

  return (
    <div className="flex flex-col w-2/3 items-center bg-slate-100 p-20 m-20 rounded-xl">
      {isLoading || university?.logo == null ? (
        <Skeleton className="w-24 h-24 rounded-full m-4" />
      ) : (
        <img
          src={university.logo}
          alt="University Logo"
          className="rounded-full bg-slate-300 m-4"
          style={{ width: "120px", height: "120px" }}
        />
      )}
      {renderField(university?.longName, '', 'w-1/2', 'text-2xl')}
      {renderField(university?.name, '', 'w-3/4', 'text-4xl font-bold')}
      {renderField(university?.region, '', 'w-1/4')}
      {renderField(university?.type, '', 'w-1/3')}
      {renderField(university?.rating, 'Рейтинг', 'w-1/4', 'text-5xl', 'm-2')}
      {renderField(university?.studyFields, 'Направлений', 'w-1/2', 'text-5xl', 'm-2')}
      {renderField(university?.budgetPlaces, 'Бюджетные места', 'w-1/4', 'text-5xl', 'm-2')}
    </div>
  );
};

export default UniversityBlock;
