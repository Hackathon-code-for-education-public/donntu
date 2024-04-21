"use client";
import { Reply, ReplySkeleton } from "@/components/reply";
import { Review, ReviewSkeleton } from "@/components/review";
import { useReview } from "@/lib/use-review";

interface Params {
  id: string;
}

export default function Page({ params }: { params: Params }) {
  const { data, isLoading } = useReview(params.id);

  return (
    <main className="min-h-screen mx-auto max-w-4xl">
      <div className="flex flex-col">
        <h1 className="text-2xl font-bold mb-4">Отзыв:</h1>
        {isLoading && <ReviewSkeleton />}
        {data && (
          <>
            <Review review={data.review} type="full" />
          </>
        )}
        <h2 className="text-xl font-bold mb-4 py-4">Ответы: </h2>
        {isLoading && <>
            <ReplySkeleton />
            <ReplySkeleton />
        </>}
        {data && (
          <>
            {data.replies.map((reply) => (
              <Reply key={reply.id} reply={reply} />
            ))}
          </>
        )}
      </div>
    </main>
  );
}
