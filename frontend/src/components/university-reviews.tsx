"use client";
import { useReviewsByUniversity } from "@/lib/use-reviews";
import { Review } from "./review";
import { Button } from "./ui/button";

interface IProps {
  universityId: string;
}

export function UniversityReviews({ universityId }: IProps) {
  const { data, isLoading, loadMore, error } = useReviewsByUniversity(universityId);

  return (
    <>
    <div className="flex flex-col gap-10">
      {isLoading && "Loading..."}
      {data &&
        data.map((review) => {
          return <Review key={review.id} review={review} />;
        })}
      <Button onClick={() => loadMore()} disabled={isLoading}>Загрузить больше</Button>
    </div>
    </>
  );
}
