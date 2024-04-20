"use client"

import { AvatarImage, AvatarFallback, Avatar } from "@/components/ui/avatar";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";

import { ReplyIcon, ShareIcon } from "lucide-react";
import { useState } from "react";
import { ReviewData } from "@/lib/use-reviews";

function truncateStr(str: string, num: number) {
  if (str.length <= num) return str;
  return str.slice(0, num) + "...";
}

interface IProps {
  review: ReviewData
}

export function Review({ review }: IProps) {
  const fullReviewText = review.text;
  const isLongReview = fullReviewText.length > 300;

  console.log(review.sentiment)

  const reviewClass = review.sentiment === 'positive' ? 'bg-green-100' : 
                      review.sentiment === 'negative' ? 'bg-red-100' : 'bg-gray-100';
  
  const [showFullReview, setShowFullReview] = useState(false);

  return (
    <div className={`max-w-4xl mx-auto p-6 rounded-lg shadow-md w-full ${reviewClass}`}>
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
              {showFullReview
                ? fullReviewText
                : truncateStr(fullReviewText, 300)}
            </p>
            {isLongReview && ( // Render the button only if the review is long
              <Button
                className="mt-4"
                variant="ghost"
                onClick={() => setShowFullReview(!showFullReview)}
              >
                {showFullReview ? "Свернуть" : "Показать полностью..."}
              </Button>
            )}
          </div>
          <div className="flex items-center justify-between mt-6">
            <div className="flex items-center space-x-2">
              <Button className="text-sm" variant="ghost">
                Ответов
                <ReplyIcon className="w-4 h-4 ml-1" />
                <span className="font-semibold">{review.repliesCount}</span>
              </Button>
            </div>
            <Button variant="ghost">
              <ShareIcon className="w-5 h-5" />
            </Button>
          </div>
          <div className="text-xs text-muted-foreground mt-2">
            Отзыв является личным мнением автора и может не совпадать с
            действительностью.
          </div>
        </div>
      </div>
    </div>
  );
}
