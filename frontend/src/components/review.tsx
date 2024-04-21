import { AvatarImage, AvatarFallback, Avatar } from "@/components/ui/avatar";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { ReplyIcon, ShareIcon } from "lucide-react";
import { useState } from "react";
import { ReviewData } from "@/lib/use-reviews";
import { ReviewDataFull } from "@/lib/use-review";
import { Skeleton } from "@/components/ui/skeleton";
import Link from "next/link";

export function ReviewSkeleton() {
  return (
    <div className="max-w-4xl mx-auto p-6 rounded-lg shadow-md w-full ">
      <div className="flex items-start space-x-4">
        <Skeleton className="w-10 h-10 rounded-full" />{" "}
        {/* Placeholder for Avatar */}
        <div className="flex-1">
          <div className="flex items-center justify-between">
            <Skeleton className="w-24 h-6 rounded-md" />{" "}
            {/* Placeholder for Badge */}
            <Skeleton className="w-36 h-4 rounded-md" />{" "}
            {/* Placeholder for Date */}
          </div>
          <Skeleton className="mt-4 h-20 w-full rounded-md" />{" "}
          {/* Placeholder for Review Text */}
          <Skeleton className="mt-4 h-6 w-36 rounded-md" />{" "}
          {/* Placeholder for Toggle Button */}
          <div className="flex items-center justify-between mt-6">
            <Skeleton className="w-24 h-6 rounded-md" />{" "}
            {/* Placeholder for Replies Button */}
            <Skeleton className="w-6 h-6 rounded-full" />{" "}
            {/* Placeholder for Share Icon */}
          </div>
          <Skeleton className="mt-2 h-8 w-full rounded-md" />{" "}
          {/* Placeholder for Disclaimer */}
        </div>
      </div>
    </div>
  );
}

function truncateStr(str: string, num: number) {
  if (str.length <= num) return str;
  return str.slice(0, num) + "...";
}

// IProps interface using conditional types
interface IProps {
  review: ReviewData | ReviewDataFull;
  type?: "full" | "small";
}

export function Review({ review, type = "small" }: IProps) {
  const fullReviewText = review.text;
  const isLongReview = fullReviewText.length > 300;

  const reviewClass =
    review.sentiment === "positive"
      ? "bg-green-100"
      : review.sentiment === "negative"
      ? "bg-red-100"
      : "";

  const [showFullReview, setShowFullReview] = useState(false);

  return (
    <div
      className={`max-w-4xl mx-auto p-6 rounded-lg shadow-md w-full ${reviewClass}`}
    >
      <div className="flex items-start space-x-4">
        <Avatar>
          <AvatarImage
            alt="User Avatar"
            src="/placeholder.svg?height=50&width=50"
          />
          <AvatarFallback>CN</AvatarFallback>
        </Avatar>
        <div className="flex-1">
          <div className="flex items-center justify-between">
            <div className="flex items-center space-x-2">
              <Badge variant="secondary">{review.authorStatus}</Badge>
              <span className="text-sm text-muted-foreground">
                {review.date}
              </span>
            </div>
          </div>
          <div className="mt-4">
            <p className="mt-2">
              {showFullReview || type === "full"
                ? fullReviewText
                : truncateStr(fullReviewText, 300)}
            </p>
            {type === "small" && isLongReview && (
              <>
                <Button
                  className="mt-4"
                  variant="ghost"
                  onClick={() => setShowFullReview(!showFullReview)}
                >
                  {showFullReview ? "Свернуть" : "Показать полностью..."}
                </Button>
              </>
            )}
          </div>
          {type === "small" && (
            <div className="flex items-center justify-between mt-6">
              <div className="flex items-center space-x-2">
                <Link href={`/review/${review.reviewId}`} legacyBehavior passHref>
                  <Button className="text-sm" variant="ghost">
                    Ответов
                    <ReplyIcon className="w-4 h-4 ml-1" />
                    <span className="font-semibold">{review.repliesCount}</span>
                  </Button>
                </Link>
              </div>
              <Link href={`/review/${review.reviewId}`} legacyBehavior passHref>
                <Button variant="ghost">
                  <ShareIcon className="w-5 h-5" />
                </Button>
              </Link>
            </div>
          )}
          <div className="text-xs text-muted-foreground mt-2">
            Отзыв является личным мнением автора и может не совпадать с
            действительностью.
          </div>
        </div>
      </div>
    </div>
  );
}
